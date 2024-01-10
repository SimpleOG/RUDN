package db

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func Do(slice []string, ch chan []string) {
	m := make(map[string]bool)
	for _, v := range slice {
		if m[v] == true {
			continue
		}
		m[v] = true
	}
	ans := make([]string, 0)
	for i := range m {
		ans = append(ans, i)
	}
	ch <- ans
}

func (qur *Queries) FillTables() error {
	start := time.Now()
	excel, err := ReadExcel()
	if err != nil {
		return err
	}
	ch := make(chan []string)
	m := make([][]string, 0)
	for i := c; i <= be; i++ {
		go Do(excel[i], ch)
		m = append(m, <-ch)
	}
	fmt.Println("Считалось")
	var wg sync.WaitGroup
	for i := range m {

		for j := range m[i] {
			wg.Add(1)
			go func(i, j int) {
				query := fmt.Sprintf("INSERT INTO table%v(field1) values($1)", i+1)
				_, err := qur.db.Exec(context.Background(), query, m[i][j])
				if err != nil {
					log.Fatalln(err)
				}
				wg.Done()
			}(i, j)
		}
	}
	wg.Wait()

	fmt.Println(time.Since(start))
	return nil

}
