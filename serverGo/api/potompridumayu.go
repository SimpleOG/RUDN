package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "rudnWebApp/db/sqlc"
	"sync"
)

func (s *Server) TeacherHours(ctx *gin.Context) {
	var name string
	name = ctx.Param("name")
	err := s.store.FillTeacherHours(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(200, GoodResponse())

}

//func (s *Server) GetCourseInfo(ctx *gin.Context) {
//	name := ctx.Query("name")
//	info, err := s.store.Course_Info(ctx, name)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//	ctx.JSON(200, info)
//}

func (s *Server) GetTeachers(ctx *gin.Context) {
	info, err := s.store.Get_information_about_pps(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	data := make([]db.Teacher_InfoRow, len(info))
	var wg sync.WaitGroup
	for i, v := range info {
		wg.Add(1)
		go func(i int, v db.Get_information_about_ppsRow) {
			data[i], err = s.store.Teacher_Info(ctx, v.FullName)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
			wg.Done()
		}(i, v)
	}
	wg.Wait()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(200, data)
}

type Response struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Programname string `json:"name_of_the_program"`
}

func (s *Server) ListAllTeachersDisciplines(ctx *gin.Context) {
	name := ctx.Param("name")
	semType := ctx.Query("sem_type")
	arg := db.List_All_Teacher_DisciplinesParams{
		FullName:     name,
		SemesterType: semType,
	}
	disciplines, err := s.store.List_All_Teacher_Disciplines(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
	}
	ctx.JSON(200, disciplines)
}
