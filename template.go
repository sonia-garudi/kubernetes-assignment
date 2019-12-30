package main

import (
  "html/template"
  "log"
  "net/http"
  "database/sql"
  "fmt"
  "os"
  _ "github.com/go-sql-driver/mysql"
)

type User struct {
	FirstName         string
	LastName         string
	HelloString		string
}

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/submitted", EnterDB)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request){

    UserVar := User{ 
      FirstName: "Enter FirstName",
      LastName: "Enter LastName",
    }

    t, err := template.ParseFiles("homepage.html") //parse the html file homepage.html
    if err != nil { // if there is an error
  	  log.Print("template parsing error: ", err) // log it
  	}
    err = t.Execute(w, UserVar) //execute the template and pass it the HomePageVars struct to fill in the gaps
    if err != nil { // if there is an error
  	  log.Print("template executing error: ", err) //log it
  	}
}

func EnterDB(w http.ResponseWriter, r *http.Request){

	r.ParseForm()
	var fname = r.FormValue("firstname")
	var lname = r.FormValue("lastname")
	var str = "Hello " + fname + " " + lname
	
	UserVar := User{ 
	  FirstName: "Enter FirstName",
      LastName: "Enter LastName",
	  HelloString: str,
    }
	
	db, e := sql.Open("mysql", "sonia:sonia@tcp(polling-app-mysql:3306)/mysql_test")
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	defer db.Close()
	
	_,useDbErr := db.Exec("USE mysql_test")
	if useDbErr != nil {
		fmt.Println(useDbErr)
	} else {
		fmt.Println("DB selected successfully..")
	}
	
	createTableStmt, createTableErr := db.Prepare("CREATE TABLE IF NOT EXISTS users(id int NOT NULL AUTO_INCREMENT, firstname varchar(50), lastname varchar(50), PRIMARY KEY (id));")
	if createTableErr != nil {
		fmt.Println(createTableErr)
	}
	_, crTbExecErr := createTableStmt.Exec()
	if crTbExecErr != nil {
		fmt.Println(crTbExecErr)
	} else {
		fmt.Println("Table created successfully..")
	}
	
	insertStmt, _ := db.Prepare("INSERT users SET firstname=?, lastname=?")
		
	_, insertErr := insertStmt.Exec(fname,lname)
	if insertErr != nil {
		fmt.Println(insertErr)
	}
	
	t, err := template.ParseFiles("homepage.html") //parse the html file homepage.html
    if err != nil { // if there is an error
  	  log.Print("template parsing error: ", err) // log it
  	}
    err = t.Execute(w, UserVar) //execute the template and pass it the HomePageVars struct to fill in the gaps
    if err != nil { // if there is an error
  	  log.Print("template executing error: ", err) //log it
  	}
	
	fmt.Println("Done")
}
