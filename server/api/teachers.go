package api

import (
	"github.com/gin-gonic/gin"
	"rudnWebApp/server/db/sqlc"
	"sync"
)

func (s *Server) listAllTeachers(ctx *gin.Context) {
	data, err := s.store.ListAllTeachers(ctx)
	if err != nil {
		//if err==pgx.norows
		return
	}
	//renderTemplate(ctx,"techersTables.gohtml",gin.H{"data":data})
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(200, data)

}
func (s *Server) displayProfile(ctx *gin.Context) {
	name := ctx.Request.URL.Query().Get("name")
	teacher, err := s.store.GetTeacher(ctx, name)
	if err != nil {
		return
	}
	data, err := s.store.ListAllTeachersCourses(ctx, teacher.FullName)
	if err != nil {
		return
	}
	var hours int32
	a := make([]db.Course, 0)
	var wg sync.WaitGroup

	for _, v := range data {
		wg.Add(1)
		go func(v db.TeachersCourse) {
			course, err := s.store.GetCourse(ctx, v.CourseName)
			a = append(a, course)
			if err != nil {
				return
			}
			hours += course.LectureHours + course.PractiseHours + course.LaboratoriesHours
			wg.Done()
		}(v)

	}
	wg.Wait()
	renderTemplate(ctx, "profile.gohtml", gin.H{"hours": hours, "teacher": teacher, "data": a})

}
