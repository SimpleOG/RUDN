package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// CreateFileForDownload на страничке преподавателя выбираются поля которые надо видеть
// фронт отправляет имя преподавателя и поля
// запрос забирает из бд нужные данные
// в word файл записываются данные
// сформированный документ подаётся в хендлер и отправляется клиенту
func (s *Server) DownloadFile(ctx *gin.Context) {
	if ctx.Request.Method == "OPTIONS" {
		ctx.Writer.WriteHeader(http.StatusOK)
		return
	}
	name := ctx.Param("name")
	var arg Fields
	err := ctx.ShouldBindJSON(&arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	filePath, filename, err := s.store.FillWord(name, arg.Field)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Header("Content-Disposition", "attachment; filename="+filename)
	ctx.Header("Content-Type", "text/plain")
	ctx.Data(http.StatusOK, "application/msword", data)
}

type Fields struct {
	Field []string `json:"array"`
}

//func (s *Server) CreateFileForDownload(ctx *gin.Context) {
//
//	filePath, filename, err := s.store.FillWord(name, arg.Field)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//	request_string := fmt.Sprintf("http://localhost:8080/sendFile?filepath=%v&filename=%v", filePath, filename)
//
//	req, err := http.NewRequest("GET", request_string, nil)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//	_, err = http.DefaultClient.Do(req)
//	if err != nil {
//		ctx.JSON(http.StatusNotFound, errorResponse(err))
//		return
//	}
//
//}

func (s *Server) SendFile(ctx *gin.Context) {
	filepath := ctx.Query("filepath")
	filename := ctx.Query("filename")
	w := ctx.Writer
	data, err := os.ReadFile(filepath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.Header("Content-Disposition", "attachment; filename="+filename)
	ctx.Data(http.StatusOK, "application/msword", data)
	ctx.Writer.Header().Set("Data", "no-cache")

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Пример успешно обработан",
	})
	//Открываем файл
	file, err := os.Open(filepath) // замените на путь к вашему файлу
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Устанавливаем заголовки для отправки файла
	w.Header().Set("Content-Disposition", "attachment; filename=yourfile.pdf") // замените на имя вашего файла
	w.Header().Set("Content-Type", "application/pdf")                          // замените на MIME-тип вашего файла

	// Отправляем содержимое файла в ответ
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
func (s *Server) handlePostGetRequest(ctx *gin.Context) {
	r := ctx.Request
	w := ctx.Writer
	if r.Method == http.MethodPost {
		// Создаем файл при получении POST запроса
		// здесь ваша логика для обработки POST запроса и создания файла
		// Например, создаем временный файл и записываем в него данные из запроса
		tmpFile, err := ioutil.TempFile("", "example")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		defer tmpFile.Close()

		_, err = io.Copy(tmpFile, r.Body)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// После создания файла вызываем GET метод
		s.handleGetRequest(w, r, tmpFile.Name())

	} else if r.Method == http.MethodGet {
		// Вызываем GET метод
		s.handleGetRequest(w, r, "")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleGetRequest(w http.ResponseWriter, r *http.Request, filePath string) {
	if filePath != "" {
		// Отправляем созданный файл в ответ на GET запрос
		file, err := ioutil.ReadFile(filePath)
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Disposition", "attachment; filename=example.txt")
		w.Header().Set("Content-Type", "text/plain")
		w.Write(file)
	} else {
		// Обработка GET запроса без файла, здесь ваша логика
		fmt.Fprintln(w, "GET request processed, no file")
	}
}
