package server

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
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
func (s *Server) setupRouter() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*") //подключаем папку с html
	router.Static("/css", "./css")     // подключаем визуал
	//настройка роутов
	router.GET("/home", s.homePage)
	router.GET("/all_teachers", s.listAllTeachers)
	router.GET("/all_groups", s.listAllGroups)
	router.GET("/teachers", s.DisplayProfile)

	//router.GET("/groups", s.getProfile)

	s.router = router
}
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("templates/*"))
}

// для обработки html файликов, которые содержат динамические данные
func renderTemplate(c *gin.Context, templateName string, h gin.H) {
	//https://stackoverflow.com/questions/25329647/golang-template-with-multiple-structs
	err := templates.ExecuteTemplate(c.Writer, templateName, h)
	if err != nil {
		// Обработка ошибки
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}
