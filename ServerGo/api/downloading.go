package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"rudnWebApp/pb"
)

type Fields struct {
	Field []string `json:"array"`
}

func (s *Server) DownloadFile(ctx *gin.Context) {
	if ctx.Request.Method == "OPTIONS" {
		ctx.Writer.WriteHeader(http.StatusOK)
		return
	}
	name := ctx.Param("name")
	var arg Fields
	err := ctx.ShouldBindJSON(&arg)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, errorResponse(err))
		return
	}
	filePath, err := s.FillWord(name, arg.Field)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	ctx.Header("Content-Disposition", "attachment; filename="+filepath.Base(filePath))
	ctx.Header("Content-Type", "text/plain")
	ctx.Data(http.StatusOK, "application/msword", data)
}

//func uniqueValues(arr []map[string]string, field string) []map[string]string {
//	uniqueMap := make(map[string]map[string]string)
//
//	for _, m := range arr {
//		if _, ok := uniqueMap[m[field]]; !ok {
//			uniqueMap[m[field]] = m
//		}
//	}
//
//	var uniqueArr []map[string]string
//	for _, v := range uniqueMap {
//		uniqueArr = append(uniqueArr, v)
//	}
//
//	return uniqueArr
//}

// FillWord передаю сюда имя и нужные поля. Оно возвращает мне путь к файлу
func (s *Server) FillWord(name string, fields []string) (string, error) {
	rows, err := s.store.TakeInfo(fields, name)
	if err != nil {
		return "", err
	}
	req := make([]*pb.MyMap, 0)
	for i := range rows {
		req = append(req, &pb.MyMap{Map: rows[i]})
	}
	arg := &pb.GenerateRequest{Name: name, Data: req}
	response, err := s.grpcClient.Generate(context.Background(), arg)
	if err != nil {
		return "", err
	}

	return response.Filepath, nil
}
