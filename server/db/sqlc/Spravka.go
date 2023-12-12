package db

import (
	"context"
	"fmt"
	"github.com/lukasjarosch/go-docx"
	"math"
	"os"
	"time"
	"unicode"
)

var w, _ = os.Getwd()

var way = w + "/ForDownload"

func cutter(name string, c chan string) <-chan string {
	r := []rune(name)
	st, en := 0, 0
	for i, v := range r {
		if v == '_' {
			st = i
		}
		if v == '.' && r[i+1] == 'd' {
			en = i
		}
	}
	r = r[st+1 : en]
	r[0] = unicode.ToUpper(r[0])
	c <- string(r)
	return c
}
func (q *Queries) FillDoc(path, name string, m docx.PlaceholderMap) error {
	ch := make(chan string)
	defer close(ch)
	go cutter(name, ch)
	d := <-ch
	doc, err := docx.Open(path)
	if err != nil {
		return err
	}
	err = doc.ReplaceAll(m)
	if err != nil {
		return err
	}
	dirname := d + "_files"
	err = os.Chdir(way)
	err = os.Mkdir(dirname, os.FileMode(0522))
	if err != nil {
		//пока что игнор
	}
	err = os.Chdir(way + dirname)
	if err != nil {
		return err
	}
	err = doc.WriteToFile(name)
	if err != nil {
		return err
	}
	return err
}

func (q *Queries) FillTeacherHours(name string) error {
	data, err := q.Teacher_Info(context.Background(), name)
	if err != nil {
		return err
	}
	m := docx.PlaceholderMap{
		"FIO":          data.TeacherName,
		"Lector_Hour":  data.Lectures,
		"Seminar_Hour": data.Practice,
		"Lab_Hour":     data.Labs,
		"Total":        fmt.Sprintf("%.2f", data.Total),
		"year":         time.Now().Year(),
	}

	err = q.FillDoc(way+"Справка Пример.docx", data.TeacherName+"_часы.docx", m)
	if err != nil {
		return err
	}
	return err
}
func (q *Queries) TeacherHours(name string) (Teacher_InfoRow, error) {
	data, err := q.Teacher_Info(context.Background(), name)
	if err != nil {
		return Teacher_InfoRow{}, nil
	}
	data.Total = math.Round(data.Total)
	return data, nil
}
