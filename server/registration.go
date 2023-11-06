package server

//func (s *Server) Registration(ctx *gin.Context) {
//	ctx.HTML(http.StatusOK, "login.gohtml", gin.H{"title": "login"})
//}
//func (s *Server) SignIn(ctx *gin.Context) {
//	user := db.User{
//		Username:       ctx.PostForm("login"),
//		HashedPassword: ctx.PostForm("password"),
//	}
//	fmt.Println(user)
//	_, err := s.store.GetUser(ctx, user.Username)
//	if err != nil {
//		return
//	}
//	ctx.HTML(http.StatusOK, "homepage.gohtml", gin.H{"title": "login"})
//
//}
