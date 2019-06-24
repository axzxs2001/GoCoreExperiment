package main

import _ "github.com/denisenkom/go-mssqldb"
import "database/sql"
import "log"
import "fmt"
import "flag"

func main() {
	flag.Parse() // parse the command line args

	connString := "server=.;user id=sa;password=sa;port=1433;database=testdb;"

	fmt.Printf(" connString:%s\n", connString)

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	stmt, err := conn.Prepare("select 1, 'abc'")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var somenumber int64
	var somechars string
	err = row.Scan(&somenumber, &somechars)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("somenumber:%d\n", somenumber)
	fmt.Printf("somechars:%s\n", somechars)

	fmt.Printf("bye\n")

}
