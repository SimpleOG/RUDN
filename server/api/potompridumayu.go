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

func (s *Server) GetCourseInfo(ctx *gin.Context) {
	name := ctx.Query("name")
	info, err := s.store.Course_Info(ctx, name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(200, info)
}
func (s *Server) GetTeachers(ctx *gin.Context) {
	info, err := s.store.Get_information_about_PPS(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	total := make([]db.Teacher_InfoRow, len(info))
	var wg sync.WaitGroup
	for i, v := range info {
		wg.Add(1)
		go func(i int, v db.Get_information_about_PPSRow) {
			total[i], err = s.store.TeacherHours(v.FullName)
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
	type dat struct {
		FullName          string  `json:"full_name"`
		Department        string  `json:"department"`
		Post              string  `json:"post"`
		TermsOfAttraction string  `json:"terms_of_attraction"`
		Total             float64 `json:"total"`
		Lectures          float64 `json:"lectures"`
		Practice          float64 `json:"practice"`
		Labs              float64 `json:"labs"`
	}
	data := make([]dat, len(info))
	for i := range info {
		data[i] = dat{
			FullName:          info[i].FullName,
			Department:        info[i].Department,
			Post:              info[i].Post,
			TermsOfAttraction: info[i].TermsOfAttraction,
			Total:             total[i].Total,
			Lectures:          total[i].Lectures,
			Practice:          total[i].Practice,
			Labs:              total[i].Labs,
		}
	}

	ctx.JSON(200, data)
}
