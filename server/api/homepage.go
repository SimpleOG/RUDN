package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) homePage(ctx *gin.Context) {
	//html := template.Must(template.ParseFiles("./templates/header.gohtml", "./templates/homepage.html"))
	//s.router.SetHTMLTemplate(html)
	ctx.JSON(http.StatusOK, gin.H{
		"title": "home",
	})

}
