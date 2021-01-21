package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	age "github.com/bearbin/go-age"
	"fmt"
)
type Customer struct{
	ID int `json:"id"`
	Name string `json:"name"`
	DOB string `json:"dob"`
	Addr Address `json:"address"`
}
type Address struct{
	ID int `json:"id"`
	Streetname string `json:"streetName"`
	City string `json:"city"`
	State string `json:"state"`
	Customerid int `json:"customerId"`
}
func getDBConnection() *sql.DB{
	db, err := sql.Open("mysql", "root:123@/Customer_Service")
	if err != nil {
		log.Fatal(err)
	}
	return  db
}
func getDOB(year, month, day int) time.Time {
	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return dob
}
func GetName(w http.ResponseWriter,r *http.Request){
	db:=getDBConnection()
	var names []interface{}
	query:=r.URL.Query()
	name:=query.Get("name")
	q:=`SELECT * FROM Customer INNER JOIN Address ON Customer.id=Address.Customer_id`
	if len(name)!=0{
		q=`SELECT * FROM Customer INNER JOIN Address ON Customer.id=Address.Customer_id WHERE Customer.Name=?`
		names=append(names,name)
	}
	rows, err := db.Query(q,names...)
	if err != nil {
		panic(err.Error())
	}

	var ans []Customer
	for rows.Next(){
		var c Customer
		if err := rows.Scan(&c.ID, &c.Name, &c.DOB, &c.Addr.ID, &c.Addr.Streetname, &c.Addr.City, &c.Addr.State,&c.Addr.Customerid); err != nil {
			log.Fatal(err)
		}
		ans = append(ans, c)
	}
	Byte,_:=json.Marshal(ans)
	io.WriteString(w,string(Byte))
}
func GetID(w http.ResponseWriter, r *http.Request) {
	db:=getDBConnection()
	//fmt.Println("ok")
	param:=mux.Vars(r)

	rows, err := db.Query("SELECT * FROM Customer INNER JOIN Address ON Customer.id=Address.Customer_id WHERE Customer_id=?;", param["id"])
	if err != nil {
		panic(err.Error())
	}

	var ans []Customer
	for rows.Next(){
		var c Customer
		if err := rows.Scan(&c.ID, &c.Name, &c.DOB, &c.Addr.ID, &c.Addr.Streetname, &c.Addr.City, &c.Addr.State,&c.Addr.Customerid); err != nil {
			log.Fatal(err)
		}
		ans = append(ans, c)
	}
	Byte,_:=json.Marshal(ans)
	io.WriteString(w,string(Byte))

}
//func Put(w http.ResponseWriter,r* http.Request){
//	param:=mux.Vars(r)
	//var c Customer

//}

func PostCustomer(w http.ResponseWriter, r* http.Request) {
	db := getDBConnection()
	var c Customer
	bytevalue, _ := ioutil.ReadAll(r.Body)
	if len(bytevalue) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	//	json.NewEncoder(w).Encode([]Customer(nil))
	} else {
		err1:=json.Unmarshal(bytevalue, &c)
		if err1!=nil{
			log.Fatal(err1)
		}
		if len(c.DOB) == 0 || len(c.Name) == 0 || len(c.Addr.City) == 0 || len(c.Addr.State) == 0 || len(c.Addr.Streetname) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			//json.NewEncoder(w)
			return

		}
		fmt.Println(c.DOB)
		dob := c.DOB
		dob1 := strings.Split(dob, "/")
		year, _ := strconv.Atoi(dob1[2])
		month, _ := strconv.Atoi(dob1[1])
		d, _ := strconv.Atoi(dob1[0])
		fmt.Println("C :", c)
		getAge := getDOB(year, month, d)

		if age.Age(getAge) >= 18 {
			var value []interface{}

			query := `INSERT INTO Customer VALUES(?,?,?)`
			value = append(value, 0)
			value = append(value, c.Name)
			value = append(value, c.DOB)
			rows, err := db.Exec(query, value...)
			if err != nil {

				panic(err.Error())
			}
			idAddr, _ := rows.LastInsertId()
			var Addrvalue []interface{}
			query = `INSERT INTO Address Values(?,?,?,?,?)`
			Addrvalue = append(Addrvalue, c.Addr.ID)
			Addrvalue = append(Addrvalue, c.Addr.Streetname)
			Addrvalue = append(Addrvalue, c.Addr.City)
			Addrvalue = append(Addrvalue, c.Addr.State)
			Addrvalue = append(Addrvalue, idAddr)
			rows2, err := db.Exec(query, Addrvalue...)
			//GetID(w,r)
			idAddr1, _ := rows2.LastInsertId()
			row, err := db.Query("SELECT * FROM Customer INNER JOIN Address ON Customer.id=Address.Customer_id WHERE Customer_id=?;", idAddr)
			if err != nil {
				panic(err.Error())
			}

			var ans []Customer
			for row.Next() {
				var c Customer
				c.ID = int(idAddr)
				c.Addr.ID = int(idAddr1)
				c.Addr.Customerid = int(idAddr)
				if err := row.Scan(&c.ID, &c.Name, &c.DOB, &c.Addr.ID, &c.Addr.Streetname, &c.Addr.City, &c.Addr.State, &c.Addr.Customerid); err != nil {
					log.Fatal(err)
				}
				ans = append(ans, c)
			}
			w.WriteHeader(http.StatusCreated)
			Byte, _ := json.Marshal(ans)
			io.WriteString(w, string(Byte))
		}

	}
}
func PutCustomer(w http.ResponseWriter,r* http.Request){
	db:=getDBConnection()
	body, _ := ioutil.ReadAll(r.Body)
	var c Customer
	err:=json.Unmarshal(body,&c)
	//fmt.Println(c)
	if err!=nil{
		fmt.Println("in err ")
		log.Fatal(err)
	}
	if c.DOB =="" {

		//fmt.Println(c)
		param := mux.Vars(r)
		id := param["id"]
		if c.Name != "" {
			_, err := db.Exec("update Customer set Name=? where id=?", c.Name, id)
			if err != nil {
				panic(err.Error())
				json.NewEncoder(w).Encode(Customer{})
			}
		}
		var data []interface{}
		query := "update Address set "
		if c.Addr.City != "" {
			query += "City = ? ,"
			data = append(data, c.Addr.City)
		}
		if c.Addr.State != "" {
			query += "State = ? ,"
			data = append(data, c.Addr.State)
		}
		if c.Addr.Streetname != "" {
			query += "StreetName = ? ,"
			data = append(data, c.Addr.Streetname)
		}
		query = query[:len(query)-1]
		query += "where Customer_id = ? and id = ?"
		data = append(data, id)
		data = append(data, c.Addr.ID)
		_, err = db.Exec(query, data...)

		//if err != nil {
		// log.Fatal(err)
		//}
		json.NewEncoder(w).Encode(c)
	}else{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("DOB Can't me updated")
	}
}

/*func DeleteCust(w http.ResponseWriter, r* http.Request ){
	db:=getDBConnection()
	params:=mux.Vars(r)
	_,err1:=db.Query("DELETE FROM Address WHERE Customer_id=?;",params["id"])
	if err1 != nil {
		panic(err1.Error())
	}
	//fmt.Printf("%T, %d", params["id"], params["id"])
	//_, err1 = stmt1.Exec(params["id"])
	//db, err = sql.Open("mysql", "sumit:1234@/Cust_Service")
	//defer db.Close()

	if err1 != nil {
		panic(err1)
	}
	if err1 != nil {
		panic(err1.Error())
	}
	_,err:=db.Query("DELETE FROM Customer WHERE id=?;",params["id"])
	if err != nil {
		panic(err.Error())
	}
	//fmt.Printf("%T, %d", params["id"], params["id"])
	//_, err = stmt.Exec(params["id"])
	//db, err = sql.Open("mysql", "sumit:1234@/Cust_Service")
	//defer db.Close()

	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Post with ID = %s was deleted", params["id"])
}*/
func DeleteCustomerById(w http.ResponseWriter, r *http.Request) {
	var ids []interface{}
	param := mux.Vars(r)
	id, err1 := strconv.Atoi(param["id"])
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode([]Customer(nil))
		if err != nil {
			panic(err.Error())
		}
	}
	ids = append(ids, id)
	db := getDBConnection()
	query := `SELECT * FROM Customer INNER JOIN Address ON Customer.id= Address.Customer_id where Customer_id= ?; `
	rows, err := db.Query(query, ids...)
	if err != nil {
		panic(err.Error())
	}
	if !rows.Next() {
		fmt.Println("No rows")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode([]Customer(nil))
	} else {
		query = `DELETE  FROM Customer where id =?; `
		_, err1 = db.Exec(query, ids...)
		if err1 != nil {
			panic(err.Error())
		}
		defer rows.Close()
		var c Customer
		for rows.Next() {
			if err := rows.Scan(&c.ID, &c.Name, &c.DOB,  &c.Addr.ID, &c.Addr.City, &c.Addr.State, &c.Addr.Streetname, &c.Addr.Customerid); err != nil {
				log.Fatal(err)
			}
		}
		w.WriteHeader(http.StatusNoContent)
		// json.NewEncoder(w).Encode(c)
	}
}

func main() {


	router:=mux.NewRouter()
	router.HandleFunc("/Customer",GetName).Methods("GET")
	router.HandleFunc("/Customer/{id:[0-9]+}",GetID).Methods("GET")
	router.HandleFunc("/Customer",PostCustomer).Methods("POST")
	router.HandleFunc("/Customer/{id:[0-9]+}",PutCustomer).Methods("PUT")
	router.HandleFunc("/Customer/{id:[0-9]+}",DeleteCustomerById).Methods("DELETE")
	//router.HandleFunc("/Customer/{id}",Put).Methods("PUT")
	http.ListenAndServe(":8084", router)
}
