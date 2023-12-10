package db

import (
	"context"
	"fmt"
	"github.com/lukasjarosch/go-docx"
	"math"
	"os"
	"time"
)

const tway = "C:\\Users\\Oleg\\GolandProjects\\rudnWebApp\\server\\templates\\"
const fway = "C:\\Users\\Oleg\\GolandProjects\\rudnWebApp\\server\\Files"

func (q *Queries) FillDoc(path, name string, m docx.PlaceholderMap) error {
	doc, err := docx.Open(path)
	if err != nil {
		return err
	}
	err = doc.ReplaceAll(m)
	if err != nil {
		return err
	}
	err = os.Chdir(fway)
	if err != nil {
		return err
	}
	err = doc.WriteToFile(name)
	if err != nil {
		return err
	}
	return err
}

func (q *Queries) DownloadTeacherHours(name string) error {
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
	err = q.FillDoc(tway+"Справка Пример.docx", data.TeacherName+"_часы.docx", m)
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
