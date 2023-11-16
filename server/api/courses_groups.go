package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *Server) listAllCourseGroups(ctx *gin.Context) {
	name := ctx.Query("name")
	data, err := s.store.ListAllCourseGroups(ctx, name)
	if err != nil {
		return
	}
	renderTemplate(ctx, "courses_groups.gohtml", gin.H{"data": data, "name": name})

}
func (s *Server) listAllGroupCourses(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("name"))
	if err != nil {
		return
	}
	data, err := s.store.ListAllGroupCourses(ctx, int32(id))
	if err != nil {
		return
	}
	renderTemplate(ctx, "groups_courses.gohtml", gin.H{"data": data, "name": data[0].Name})

}
