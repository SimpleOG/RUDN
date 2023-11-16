package api

import "github.com/gin-gonic/gin"

func (s *Server) listAllGroups(ctx *gin.Context) {
	data, err := s.store.ListAllGroups(ctx)
	if err != nil {
		//if err==pgx.norows
		return
	}

	renderTemplate(ctx, "groupsTables.gohtml", gin.H{"data": data})
}
