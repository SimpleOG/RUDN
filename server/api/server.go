package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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
	router.GET("/download/:path", s.Download)
	router.GET("/hello", s.SayHello)
	router.GET("/course/:name", s.GetCourseInfo)
	router.GET("/teachers", s.GetTeachers)
	router.GET("/fill", s.Fill)
	s.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func (s *Server) SayHello(ctx *gin.Context) {
	ctx.JSON(200, GoodResponse())
}

func (s *Server) Fill(ctx *gin.Context) {
	err := s.store.ReadItAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, GoodResponse())
}
func (s *Server) Download(ctx *gin.Context) {
	filePath := ctx.Param("path") // Путь к файлу, который вы хотите скачать
	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	defer file.Close()

	// Устанавливаем заголовки для скачивания файла
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Disposition", "attachment; filename=your_file.txt")
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Expires", "0")

	// Копируем содержимое файла в ответ
	ctx.FileAttachment(filePath, "your_file.txt")
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
func GoodResponse() gin.H {
	return gin.H{"Статус": "Ура,победа"}
}
