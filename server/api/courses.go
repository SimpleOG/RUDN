package api

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) listAllCourses(ctx *gin.Context) {
	data, err := s.store.ListAllCourses(ctx)
	if err != nil {
		return
	}
	renderTemplate(ctx, "coursesTables.gohtml", gin.H{"data": data})
}
func (s *Server) displayCourse(ctx *gin.Context) {
	name := ctx.Query("name")
	course, err := s.store.GetCourse(ctx, name)
	if err != nil {
		return
	}

	renderTemplate(ctx, "coursesTables.gohtml", gin.H{"data": course})
}
