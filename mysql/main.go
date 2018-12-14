package main

import ( "database/sql"
	"fmt" 
 	_ "github.com/go-sql-driver/mysql"
	)

func main () {
	
	db, err := sql.Open("mysql", "rossp:Kin405!405@tcp(ec2-54-66-150-99.ap-southeast-2.compute.amazonaws.com:3306)/test?sql_mode=TRADITIONAL&");
	if err != nil {
    		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
        fmt.Printf("db: %v\n",*db)

        err = db.Ping()
	if err != nil {
    		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT name FROM t WHERE id = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	var Name_  string  //[]byte // we "scan" the result in here

	// Query the square-number of 13
	err = stmtOut.QueryRow(1).Scan(&Name_) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 13 is: %s", Name_)

	var empty []interface{} 
	result,err:=db.Exec("SET GLOBAL read_only = ON",empty);			// doesn't work!
	fmt.Printf("\nresult: %v",result);
}
