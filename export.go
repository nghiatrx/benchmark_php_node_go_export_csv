package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

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

	last_id := 0
	limit := 10000

	file, err := os.Create("go_outout.csv")
	if err != nil {
		log.Fatal("Error creating CSV file:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"id", "qty", "price", "total", "column_0", "column_1", "column_2", "column_3", "column_4", "column_5", "column_6", "column_7", "column_8", "column_9", "column_10"})

	for {
		rows, _ := db.Query(fmt.Sprintf("SELECT * FROM products WHERE id > %v ORDER BY id ASC LIMIT %v", last_id, limit))
		defer rows.Close()
		has_row := false

		for rows.Next() {
			has_row = true
			var id, qty, price int
			var column_0, column_1, column_2, column_3, column_4, column_5, column_6, column_7, column_8, column_9, column_10 string

			rows.Scan(&id, &qty, &price, &column_0, &column_1, &column_2, &column_3, &column_4, &column_5, &column_6, &column_7, &column_8, &column_9, &column_10)

			writer.Write([]string{strconv.Itoa(id), strconv.Itoa(qty), strconv.Itoa(price), strconv.Itoa(qty * price), column_0, column_1, column_2, column_3, column_4, column_5, column_6, column_7, column_8, column_9, column_10})

			last_id = id
		}

		if !has_row {
			break
		}
	}

	t2 := time.Now().UnixMilli()

	fmt.Println(t2 - t1)

}
