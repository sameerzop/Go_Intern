package main

import (
	"database/sql"
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
type Service struct {
	db *sql.DB
}

func InsertData(db *sql.DB, c Customer) Customer {
	var values []interface{}
	query := `INSERT INTO Customer VALUES (?,?,?);`
	values = append(values, c.Id)
	values = append(values, c.Name)
	values = append(values, c.DOB)
	//values = append(values, c.Addr)
	_, err := db.Query(query, values...)
	if err != nil {
		panic(err.Error())
	}

	//fmt.Println("yes again")
	query = `INSERT INTO Address VALUES (?,?,?,?,?);`
	var addrValues []interface{}
	addrValues = append(addrValues, c.Addr.Id)
	addrValues = append(addrValues, c.Addr.StreetName)
	addrValues = append(addrValues, c.Addr.City)
	addrValues = append(addrValues, c.Addr.State)
	addrValues = append(addrValues, c.Addr.CustomerId)
	_, err1 := db.Query(query, addrValues...)
	if err1 != nil {
		panic(err.Error())
	}

	//fmt.Println("yes")
	cust := getdata(db, c.Id)
	return cust[0]
}
func getdata(db *sql.DB, id int) []Customer {
	query := `select * from Customer INNER JOIN Address ON Customer.id = Address.Customer_id  and Customer_Id=?`
	var ans []Customer
	var ids []interface{}

	ids = append(ids, id)

	//if id != 0 {
	//	query = "select * from Customer INNER JOIN Address ON Customer.id = Address.Customer_id where Customer_Id=?multiStatements=true";"
	//	ids = append(ids, id)
	//}
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
	//db, err := sql.Open("mysql", "root:123@/Customer_Service")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var cust, c Customer
	//cust ={ID:5,Name:}
	//cust = Customer{id: 15, Name: "CustomerA14", DOB: "10/10/2010", Addr: Address{15, "Hyderabad", "Telangana", "1210", 15}}
	//c = InsertData(db, cust)
	//fmt.Println(c)
	//defer db.Close()

	//	var param int = 2
	//var param string = string("1 or 1=1")
	//var param string = string("1 ; Select * FROM Customer")
	//getData(db,param)
	//fmt.Println(getdata(db, param))

}
