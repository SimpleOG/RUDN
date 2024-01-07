package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// DownloadFile на страничке преподавателя выбираются поля которые надо видеть
// фронт отправляет имя преподавателя и поля
// запрос забирает из бд нужные данные
// в word файл записываются данные
// сформированный документ подаётся в хендлер и отправляется клиенту
func (s *Server) DownloadFile(ctx *gin.Context) {
	name := ctx.Param("name")
	filePath, filename, err := s.store.FillWord(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.Header("Content-Disposition", "attachment; filename="+filename)
	ctx.Data(http.StatusOK, "application/msword", data)
}
