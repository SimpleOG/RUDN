package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "rudnWebApp/db/sqlc"
	"rudnWebApp/pb"
	configs "rudnWebApp/util"
)

// Server структура отвечающая за обслуживание вебсервера

type Server struct {
	config     configs.Config
	store      db.Store
	router     *gin.Engine
	grpcClient pb.FileGeneratorClient
}

// NewServer создание сервера и роутинга
func NewServer(config configs.Config, store db.Store, grpcClient pb.FileGeneratorClient) (*Server, error) {
	server := &Server{
		config:     config,
		store:      store,
		grpcClient: grpcClient,
	}

	server.setupRouter()
	return server, nil
}

func (s *Server) setupRouter() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"}                   // Разрешенные источники
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // Разрешенные методы
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"} // Разрешенные заголовки

	router.Use(cors.New(config))
	router.Use(func() gin.HandlerFunc {
		return func(ctx *gin.Context) {
			ctx.Header("Access-Control-Allow-Origin", "*")
		}
	}())
	router.GET("/teacher/:name", s.TeacherHours)
	router.GET("/teachers", s.GetTeachers)
	router.POST("/getWordFile/:name", s.DownloadFile)
	router.GET("/course/:name", s.ListAllTeachersDisciplines)
	s.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
func GoodResponse() gin.H {
	return gin.H{"Статус": "Ура,победа"}
}
