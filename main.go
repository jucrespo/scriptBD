package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hello World!")

	TimeJuan := "2006-01-02"

	connectionString := ""
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	for i := 0; i<=102; i++ {
		fmt.Println("Comenzando archivo numero: ", i)
		path := "/Users/jucrespo/Downloads/installments-dev-part" + strconv.Itoa(i) + ".txt"

		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		ids := string(bytes)

		aORb := regexp.MustCompile("\n")
		matches := aORb.FindAllStringIndex(ids, -1)
		totalFile := len(matches)

		ids = strings.Trim(ids, "\n")
		ids = strings.Trim(ids, ",")

		maxAmount := "700003"
		lastUpdated := time.Now().Format(TimeJuan)

		query := "UPDATE installment " +
			"SET max_amount = " + maxAmount + ", last_updated = STR_TO_DATE(?, '%Y-%m-%d') " +
			"WHERE id in ("+ids+")"

		result, err := db.Exec(query, lastUpdated)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		rowsAffected, _ := result.RowsAffected()

		query = "SELECT COUNT(id) FROM installment WHERE id in ("+ids+")"
		rows, err := db.Query(query)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		var val int
		if rows.Next() {
			rows.Scan(&val)
		}

		fmt.Println("max_amount = ", maxAmount, ", last_updated = ", lastUpdated)
		fmt.Println("Totales archivo: ", totalFile, ". Totales BD: ", val, ". Archivo número ", i, ". Afectados: ", rowsAffected)
		fmt.Println("Finalizado correctamente archivo numero: ", i)
		rows.Close()
	}
}


//SET max_amount = 700000, last_updated = " + time.Now().String() + "