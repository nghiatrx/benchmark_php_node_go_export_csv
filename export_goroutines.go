package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func countDB(db *sql.DB) int {
	rows, _ := db.Query("SELECT COUNT(*) as C FROM products")
	defer rows.Close()
	for rows.Next() {
		var count int
		err := rows.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}
		return count
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return 0
}

func findMidId(db *sql.DB, total int) int {
	rows, _ := db.Query(fmt.Sprintf(`SELECT * FROM (
		SELECT id FROM products ORDER BY id ASC limit %v
   ) as T
   order by T.id desc limit 1
   `, total/2))
	defer rows.Close()
	for rows.Next() {
		var count int
		err := rows.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}
		return count
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return 0
}

func queryAndExport(db *sql.DB, begin_id int, end_id int, file_name string, wg *sync.WaitGroup) {
	defer wg.Done()
	const limit = 10000
	last_id := begin_id

	file, err := os.Create(file_name)
	if err != nil {
		log.Fatal("Error creating CSV file:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for {
		sql := fmt.Sprintf(
			`SELECT * FROM products 
			WHERE id > %v AND id <= %v
			LIMIT %v`, last_id, end_id, limit)

		rows, _ := db.Query(sql)

		has_row := false

		for rows.Next() {
			has_row = true
			var id, qty, price int
			var column_0, column_1, column_2, column_3, column_4, column_5, column_6, column_7, column_8, column_9, column_10 string

			rows.Scan(&id, &qty, &price, &column_0, &column_1, &column_2, &column_3, &column_4, &column_5, &column_6, &column_7, &column_8, &column_9, &column_10)

			writer.Write([]string{strconv.Itoa(id), strconv.Itoa(qty), strconv.Itoa(price), strconv.Itoa(qty * price), column_0, column_1, column_2, column_3, column_4, column_5, column_6, column_7, column_8, column_9, column_10})

			last_id = id
		}
		rows.Close()

		if !has_row {
			break
		}
	}
}

func merge2Files(file_name_1 string, file_name_2 string, output_name string) {
	// Create a new file to write the merged contents to
	var mergedFile *os.File
	var file1 *os.File
	var file2 *os.File
	var wg sync.WaitGroup
	wg.Add(3)

	go func(wg *sync.WaitGroup) {
		mergedFile, _ = os.Create(output_name)
		writer := csv.NewWriter(mergedFile)
		writer.Write([]string{"id", "qty", "price", "total", "column_0", "column_1", "column_2", "column_3", "column_4", "column_5", "column_6", "column_7", "column_8", "column_9", "column_10"})
		writer.Flush()
		wg.Done()
	}(&wg)

	go func(wg *sync.WaitGroup) {
		file1, _ = os.Open(file_name_1)
		wg.Done()
	}(&wg)

	go func(wg *sync.WaitGroup) {
		file2, _ = os.Open(file_name_2)
		wg.Done()
	}(&wg)

	wg.Wait()
	defer mergedFile.Close()
	defer file1.Close()
	defer file2.Close()

	// Write the contents of the first file to the merged file
	io.Copy(mergedFile, file1)
	// Write the contents of the second file to the merged file
	io.Copy(mergedFile, file2)
}

func main() {
	t1 := time.Now().UnixMilli()
	// Open a database connection
	const DB_USER = "root"
	const DB_PASSWORD = "123456"
	const DB_HOST = "localhost"
	const DB_PORT = 3306
	const DB_NAME = "benchmark_test"
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	total := countDB(db)
	mid_id := findMidId(db, total)

	const file_name_1 = "go_out_1.csv"
	const file_name_2 = "go_out_2.csv"
	const output_name = "go_out.csv"
	var wg sync.WaitGroup
	wg.Add(2)
	go queryAndExport(db, 0, mid_id, file_name_1, &wg)
	go queryAndExport(db, mid_id+1, math.MaxInt32, file_name_2, &wg)
	wg.Wait()

	merge2Files(file_name_1, file_name_2, output_name)

	t2 := time.Now().UnixMilli()

	fmt.Println(t2 - t1)

}
