package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"

)
type Customer struct {
	Id int
	Name string
	DOB string
	Addr Address
}
type Address struct{
	Id         int
	StreetName string
	City       string
	State      string
	CustomerId int
}

//func GetCustomer(dB *sql, id int){
//
//
//}
func getdata(db *sql.DB,id int)[]Customer{
	if id==0{
		var ans []Customer
		out, err := db.Query("select * from Customer INNER JOIN Address ON Customer.id = Address.Customer_id;")
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		for out.Next(){

			var c Customer
			if err := out.Scan(&c.Id, &c.Name, &c.DOB, &c.Addr.Id, &c.Addr.StreetName, &c.Addr.City, &c.Addr.State,&c.Addr.CustomerId); err != nil {
				log.Fatal(err)
			}
			ans = append(ans, c)

		}
		//fmt.Println(result,"result")
		return  ans
	}else {
		var ans []Customer
		out, err := db.Query(fmt.Sprintf("SELECT * FROM Customer INNER JOIN Address ON Customer.id=Address.Customer_id WHERE Customer.id=%v;", id))
		if err != nil {
			panic(err.Error())
		}
		for out.Next() {
			var c Customer
			if err := out.Scan(&c.Id, &c.Name, &c.DOB, &c.Addr.Id, &c.Addr.StreetName, &c.Addr.City, &c.Addr.State, &c.Addr.CustomerId); err != nil {
				log.Fatal(err)
			}
			ans = append(ans, c)
		}
		return ans
	}
}
func main(){
	db, err := sql.Open("mysql", "root:123@/Customer_Service")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var param int = 2
	//getData(db,param)
	fmt.Println(getdata(db,param))
}
/*func main()  {

	//type struct Customer {
	//	id int "json:id"
	//}
	db, err := sql.Open("mysql", "root:123@/Customer_Service")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	stmtIns, err := db.Prepare("INSERT INTO Customer VALUES( ?, ?, ?)") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()

	for i := 3; i < 25; i++ {
		_, err = stmtIns.Exec(i, "CustomerA", "28-09-1997") // Insert tuples (i, i^2)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

}*/

