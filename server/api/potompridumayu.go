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
	err := s.store.DownloadTeacherHours(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(200, nil)
}

func (s *Server) GetCourseInfo(ctx *gin.Context) {
	name := ctx.Query("name")
	info, err := s.store.Course_Info(ctx, name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(200, info)
}
func (s *Server) GetTeachers(ctx *gin.Context) {
	info, err := s.store.Get_information_about_PPS(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	total := make([]db.Teacher_InfoRow, len(info))
	var wg sync.WaitGroup
	for i, v := range info {
		wg.Add(1)
		go func(i int, v db.Get_information_about_PPSRow) {
			total[i], err = s.store.TeacherHours(v.FullName)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			wg.Done()
		}(i, v)
	}
	wg.Wait()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, gin.H{"info": info, "total": total})
}
