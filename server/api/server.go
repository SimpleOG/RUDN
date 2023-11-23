package api

import (
	"github.com/gin-gonic/gin"

	"rudnWebApp/server/db/sqlc"
	configs "rudnWebApp/server/util"
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
func (s *Server) setupRouter() {
	router := gin.Default()
	//настройка роутов
	router.Use(func() gin.HandlerFunc {
		return func(ctx *gin.Context) {
			ctx.Header("Access-Control-Allow-Origin", "*")
		}
	}())

	s.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
