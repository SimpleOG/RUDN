package api

import (
	"github.com/gin-gonic/gin"
	"log"
	db "rudnWebApp/db/sqlc"
	configs "rudnWebApp/util"
)

const templatepath = "C:\\Users\\Oleg\\GolandProjects\\wepApp\\templates/"

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
func (s *Server) home(ctx *gin.Context) {
	data, err := s.store.Get_EducationalProgram(ctx, "Прикладная информатика")
	if err != nil {
		log.Fatalln(err)
	}
	ctx.JSON(200, gin.H{"data": data})
}
func (s *Server) setupRouter() {
	router := gin.Default()
	//настройка роутов
	router.Use(func() gin.HandlerFunc {
		return func(ctx *gin.Context) {
			ctx.Header("Access-Control-Allow-Origin", "*")
		}
	}())
	router.GET("/home", s.home)
	s.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
