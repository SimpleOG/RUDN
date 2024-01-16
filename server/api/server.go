package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rudnWebApp/db/sqlc"
	configs "rudnWebApp/util"
)

// Server структура отвечающая за обслуживание вебсервера
type Server struct {
	config configs.Config
	store  db.Store
	router *gin.Engine
}

// NewServer создание сервера и роутинга
func NewServer(config configs.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store}

	server.setupRouter()
	return server, nil
}

func (s *Server) setupRouter() {
	router := gin.Default()

	//настройка роутов
	router.Use(func() gin.HandlerFunc {
		return func(ctx *gin.Context) {
			ctx.Header("Access-Control-Allow-Origin", "*")
		}
	}())
	router.GET("/teacher/:name", s.TeacherHours)

	//router.GET("/course/:name", s.GetCourseInfo)
	router.GET("/teachers", s.GetTeachers)
	router.GET("/fill", s.Fill)
	router.POST("/getWordFile/:name", s.DownloadFile)
	router.GET("/course/:name", s.ListAllTeachersDisciplines)
	router.GET("/groups/:name", s.MockGroupData)

	s.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func (s *Server) Fill(ctx *gin.Context) {
	err := s.store.ReadItAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, GoodResponse())

}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
func GoodResponse() gin.H {
	return gin.H{"Статус": "Ура,победа"}
}
