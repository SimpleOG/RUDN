package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) listAllGroups(ctx *gin.Context) {
	data, err := s.store.ListAllGroups(ctx)
	if err != nil {
		//if err==pgx.norows
		return
	}
	ctx.JSON(http.StatusOK, data)
	//renderTemplate(ctx, "groupsTables.gohtml", gin.H{"data": data})
}
