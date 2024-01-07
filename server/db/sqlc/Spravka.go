package db

import (
	"context"
	"fmt"
	docx2 "github.com/gingfrederik/docx"
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

// FillWord передаю сюда имя и нужные поля. Оно возвращает мне путь к файлу
func (q *Queries) FillWord(name string) (string, string, error) {
	f := docx2.NewFile()
	para := f.AddParagraph()
	para.AddText(name).Size(22).Color("808080")
	name += "_данные.docx"
	filepth := "./ForDownload/" + name
	err := f.Save(filepth)
	if err != nil {
		return "", "", err
	}
	return filepth, name, nil
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
	doc, err := docx.Open(Tway + "СправкаПример.docx")
	if err != nil {
		return err
	}
	err = doc.ReplaceAll(m)
	if err != nil {
		return err
	}
	err = MakeDoc(data.TeacherName+"_часы.docx", doc)
	if err != nil {
		return err
	}
	return err
}

// создание папки и word файла
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
