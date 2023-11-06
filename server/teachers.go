package server

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) listAllTeachers(ctx *gin.Context) {
	data, err := s.store.ListAllTeachers(ctx)
	if err != nil {
		//if err==pgx.norows
		return
	}
	renderTemplate(ctx, "teachersTables.gohtml", gin.H{"data": data})

}
func (s *Server) DisplayProfile(ctx *gin.Context) {
	//id, err := strconv.Atoi(ctx.Request.URL.Query().Get("id"))
	//if err != nil {
	//	return
	//}
	//data, err := s.store.ListAllTeachersGroups(ctx, int32(id))
	//if err != nil {
	//	return
	//}
	//
	//renderTemplate(ctx, "profile.gohtml", gin.H{"data": data})

}
