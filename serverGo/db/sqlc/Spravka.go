package db

import (
	"context"
	"fmt"
	"github.com/lukasjarosch/go-docx"
	"os"
	"path/filepath"
	"time"
	"unicode"
)

const Tway = "./ForDownload/"

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

func (qur *Queries) FillTeacherHours(name string) error {
	data, err := qur.Teacher_Info(context.Background(), name)
	if err != nil {
		return err
	}
	m := docx.PlaceholderMap{
		"FIO":          data.FullName,
		"Lector_Hour":  data.Lectures,
		"Seminar_Hour": data.Practice,
		"Lab_Hour":     data.Labs,
		"Total":        fmt.Sprintf("%.2f", data.Total),
		"year":         time.Now().Year(),
	}
	doc, err := docx.Open(Tway + "СправкаПример.docx")
	if err != nil {
		return err
	}
	err = doc.ReplaceAll(m)
	if err != nil {
		return err
	}
	err = MakeDoc(data.FullName+"_часы.docx", doc)
	if err != nil {
		return err
	}
	return nil
}

// MakeDoc создание папки и word файла
func MakeDoc(name string, doc *docx.Document) error {
	ch := make(chan string)
	go cutter(name, ch)
	d := <-ch
	close(ch)
	dirname := d + "_files"
	err := os.Chdir(Tway)
	err = os.MkdirAll(dirname, os.FileMode(0522))
	if err != nil {
		return err
	}
	err = os.Chdir(dirname)
	if err != nil {
		return err
	}
	err = doc.WriteToFile(name)
	if err != nil {
		return err
	}
	curDir, err := os.Getwd()
	if err != nil {
		return err
	}
	err = os.Chdir(filepath.Join(curDir, "..", ".."))
	if err != nil {
		return err
	}
	return err
}
