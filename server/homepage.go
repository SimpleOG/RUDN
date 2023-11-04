package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) homePage(ctx *gin.Context) {
	//html := template.Must(template.ParseFiles("./templates/header.gohtml", "./templates/homepage.gohtml"))
	//s.router.SetHTMLTemplate(html)
	ctx.HTML(http.StatusOK, "homepage.gohtml", gin.H{
		"title": "home",
	})

}
