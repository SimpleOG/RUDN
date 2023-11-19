package db

import (
	"context"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"sync"
)

const (
	//тупо счётчик
	//Образовательная программа
	educationForm = iota // Форма обучения
	lvlOp                // Уровень ОП
	sifer                // Шифр ООП РУДН
	code                 // Код направления
	name                 // Наименование программы
	// Дисциплина или вид учебной  работы
	bloc
	comp
	rupNum
	dop_info
	dicipline_name
	sem_module
	weeks_module
	work_type
	//Ауд. нагр. (час. в нед.)
	lections
	labs
	practice
	typeGia
	crse_works
	crse_projects
	le // затычки
	lee
	leee
	//контингент обучающихся
	group_code
	group_num
	group_count
	group_all
	stud_all
	stud_rus
	stud_inostr
	leeee
	leeeee
	leeeeee
	//Сведения о ППС
	department
	post
	conditions
	FIO

	start = 6
	end   = 1652
)

func ReadExcelColumn(find int) ([]string, error) {

	f, err := excelize.OpenFile("считать.xlsx")
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	rows := make([]string, end)
	cols, err := f.GetCols("УН сводная")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	k := 0
	for i := start; i < end; i++ {
		rows[k] = cols[find][i]
		k++
	}
	return rows, nil
}

//подтянем из экселя все штуки( шучу, не все)

func (q *Queries) FillCourses() (err error) {
	return err
}

func (q *Queries) FillTeachers() (err error) {
	var wg sync.WaitGroup
	var fullName []string
	var departments []string
	var posts []string
	var condition []string
	wg.Add(4)
	go func() {
		fullName, err = ReadExcelColumn(FIO)
		if err != nil {
			log.Fatalln("Не удалось найти фио", err)
		}
		wg.Done()
	}()
	go func() {
		departments, err = ReadExcelColumn(department)
		if err != nil {
			log.Fatalln("Не удалось найти фио", err)
		}
		wg.Done()
	}()
	go func() {
		posts, err = ReadExcelColumn(post)
		if err != nil {
			log.Fatalln("Не удалось найти фио", err)
		}
		wg.Done()
	}()
	go func() {
		condition, err = ReadExcelColumn(conditions)
		if err != nil {
			log.Fatalln("Не удалось найти фио", err)
		}
		wg.Done()
	}()
	wg.Wait()
	arg := make([]CreateTeacherParams, len(fullName))
	for i := range fullName {
		wg.Add(1)
		go func(i int) {
			arg[i] = CreateTeacherParams{
				FullName:   fullName[i],
				Department: departments[i],
				Post:       posts[i],
				Conditions: condition[i],
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	fn := func(args []CreateTeacherParams) []CreateTeacherParams {
		keys := make(map[string]bool)
		var list []CreateTeacherParams
		for _, entry := range args {
			if entry.FullName == "" {
				continue
			}
			if _, value := keys[entry.FullName]; !value {
				keys[entry.FullName] = true
				list = append(list, entry)
			}
		}
		return list
	}
	arg = fn(arg)

	for _, v := range arg {
		wg.Add(1)
		go func(v CreateTeacherParams) {
			_, err = q.CreateTeacher(context.Background(), v)
			if err != nil {
				log.Printf("Ошибка %v в структуре %v \n ", err, v)
			}
			wg.Done()
		}(v)
	}
	wg.Wait()
	fmt.Println("Я всё")
	return
}
func (q *Queries) FillRow() {

}

//func(q *Queries) FillCourses(){}
//func(q *Queries) FillCourses(){}
