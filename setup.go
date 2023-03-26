package main

import (
	"database/sql"
	"fmt"
	"math/rand"

	_ "github.com/go-sql-driver/mysql"
)

const DB_USER = "root"
const DB_PASSWORD = "123456"
const DB_HOST = "localhost"
const DB_PORT = 3306
const DB_NAME = "benchmark_test"

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", DB_NAME))
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec(fmt.Sprintf("USE %v", DB_NAME))
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS products (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
		qty INT UNSIGNED,
		price INT UNSIGNED,
		column_0 VARCHAR(255),
		column_1 VARCHAR(255),
		column_2 VARCHAR(255),
		column_3 VARCHAR(255),
		column_4 VARCHAR(255),
		column_5 VARCHAR(255),
		column_6 VARCHAR(255),
		column_7 VARCHAR(255),
		column_8 VARCHAR(255),
		column_9 VARCHAR(255),
		column_10 VARCHAR(255)
	)`)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 1000000; i++ {
		rand.Seed(int64(i))
		db.Exec(fmt.Sprintf(`INSERT INTO products (qty, price, column_0, column_1, column_2, column_3, column_4, column_5, column_6, column_7, column_8, column_9, column_10) 
			VALUES ('%v', '%v', 'value_0_%v', 'value_1_%v', 'value_2_%v', 'value_3_%v', 'value_4_%v', 'value_5_%v', 'value_6_%v', 'value_7_%v', 'value_8_%v', 'value_9_%v', 'value_10_%v')`,
			rand.Intn(100000), rand.Intn(100000), i, i, i, i, i, i, i, i, i, i, i))
	}

}
