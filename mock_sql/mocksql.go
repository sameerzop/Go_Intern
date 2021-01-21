package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	Id   int
	Name string
	DOB  string
	Addr Address
}
type Address struct {
	Id         int
	StreetName string
	City       string
	State      string
	CustomerId int
}

func getdata(db *sql.DB, id int) []Customer {
	query := "select * from Customer INNER JOIN Address ON Customer.id = Address.Customer_id ;"
	var ans []Customer
	var ids []interface{}

	if id != 0 {
		query = "select * from Customer INNER JOIN Address ON Customer.id = Address.Customer_id where Customer_Id=?;"
		ids = append(ids, id)
	}
	rows, err := db.Query(query, ids...)
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.Id, &c.Name, &c.DOB, &c.Addr.Id, &c.Addr.StreetName, &c.Addr.City, &c.Addr.State, &c.Addr.CustomerId); err != nil {
			log.Fatal(err)
		}
		ans = append(ans, c)

	}
	return ans
	//	}
	//	else {
	//		var ans []Customer
	//		out, err := db.Query(fmt.Sprintf("SELECT * FROM Customer INNER JOIN Address ON Customer.id=Address.Customer_id WHERE Customer.id=%v;", id))
	//		if err != nil {
	//			panic(err.Error())
	//		}
	//		for out.Next() {
	//			var c Customer
	//			if err := out.Scan(&c.Id, &c.Name, &c.DOB, &c.Addr.Id, &c.Addr.StreetName, &c.Addr.City, &c.Addr.State, &c.Addr.CustomerId); err != nil {
	//				log.Fatal(err)
	//			}
	//			ans = append(ans, c)
	//		}
	//		return ans
	//	}
	//}
}
func main() {
	db, err := sql.Open("mysql", "root:123@/Customer_Service")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var param int = 2
	//getData(db,param)
	fmt.Println(getdata(db, param))

}
