package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"rudnWebApp/pb"
)

// CreateFileForDownload на страничке преподавателя выбираются поля которые надо видеть
// фронт отправляет имя преподавателя и поля
// запрос забирает из бд нужные данные
// в word файл записываются данные
// сформированный документ подаётся в хендлер и отправляется клиенту
func (s *Server) DownloadFile(ctx *gin.Context) {
	if ctx.Request.Method == "OPTIONS" {
		ctx.Writer.WriteHeader(http.StatusOK)
		return
	}
	name := ctx.Param("name")
	var arg Fields
	err := ctx.ShouldBindJSON(&arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	filePath, err := s.FillWord(name, arg.Field)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	filename := filepath.Base(filePath)
	getwd, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Println(getwd)
	data, err := os.ReadFile(filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Header("Content-Disposition", "attachment; filename="+filename)
	ctx.Header("Content-Type", "text/plain")
	ctx.Data(http.StatusOK, "application/msword", data)
}

type Fields struct {
	Field []string `json:"array"`
}

// FillWord передаю сюда имя и нужные поля. Оно возвращает мне путь к файлу
func (s *Server) FillWord(name string, fields []string) (string, error) {

	err := s.store.TakeInfo(fields, name)
	response, err := s.grpcClient.Generate(context.Background(), &pb.GenerateRequest{Name: name, Data: fields})
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	return response.Filepath, nil
}
