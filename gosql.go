//this is the working code blocks for connecting to CloudSQL using Go from a Standard AppEngine environment
//important reference - https://cloud.google.com/sql/docs/access-control#instanceaccess
//make sure you choose the correct combination
//below code works for Standard AppEngine environment with First Generation CloudSQL instance
//Open issue: Insert query is still not very reliable - found to be inserting the same row multiple times
package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/ziutek/mymysql/godrv"

	//"appengine/cloudsql"
	//_ "github.com/go-sql-driver/mysql"
)

func init() {
	http.HandleFunc("/", connect_sql)
}

func connect_sql(w http.ResponseWriter, r *http.Request) {
	//const dbUserName = "root"
	//const dbPassword = "admin"
	//const dbIP = "tcp(104.197.140.181:3306)"
	//const dbName = "demo"
	//const dbName = "trial"
	//const dbPassword = "pass"
	//const dbIP = "tcp(130.211.117.207:3306)"

	fmt.Fprintf(w, "Attempting to connect to Google Cloud SQL...")
	//db, err := sql.Open("mysql", "root:admin@tcp(104.197.140.181:3306)/demo")
	//db, err := sql.Open("mysql", "root:admin@cloudsql(myappmulti:myappmulti)/demo")
	db, err := sql.Open("mymysql", "cloudsql:myappmulti:myappmulti1*demo/root/admin")
	//db, err := sql.Open("mymysql", "tcp:104.197.140.181:3306*demo/root/admin")
	//db, err := sql.Open("mysql", "root:admin@cloudsql(myappmulti:myappmulti1)/demo") //working code for go-sql-driver
	defer db.Close()
	if err != nil {
		fmt.Fprintf(w, err.Error()+"Connection Error<br>")
	}
	//Use regular queries to get results and output them
	err = db.Ping()
	if err != nil {
		fmt.Fprintf(w, "<br>"+err.Error()+" Ping error<br>")
	}
	query := "insert into stocks (scripid, price) values('ITC',200)"
	_, err = db.Query(query)

	if err != nil {
		fmt.Fprintf(w, err.Error()+"  Insert Error<br>")
	} else {
		fmt.Fprintf(w, "Insert successful")
	}
	_, err = db.Query(`select scripid,price from stocks where price>=0 `)

	if err != nil {
		fmt.Fprintf(w, err.Error()+"  Select Error<br>")
	} else {
		fmt.Fprintf(w, "Select successful")
	}

}
