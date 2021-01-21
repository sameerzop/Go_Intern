package main

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/google/go-cmp/cmp"
)

func TestCreateCustomer(t *testing.T) {
	// Input: ID, Output: expected complete result from MYSQL server.
	//fmt.Println("dsds")

	testcases := []struct {
		input  Customer
		output Customer
	}{
		{Customer{Id: 1, Name: "CustomerA", DOB: "28/09/1997", Addr: Address{1, "AKJ", "HSR", "U.P.", 1}}, Customer{Id: 1, Name: "CustomerA", DOB: "28/09/1997", Addr: Address{1, "AKJ", "HSR", "U.P.", 1}}},
		{Customer{Id: 2, Name: "CustomerB", DOB: "28/09/1999", Addr: Address{2, "BKJ", "BTM", "U.P.", 2}}, Customer{Id: 2, Name: "CustomerB", DOB: "28/09/1999", Addr: Address{2, "BKJ", "BTM", "U.P.", 2}}},
		//{0, Customer{{1, "CustomerA", "28/09/1997", Address{1, "AKJ", "HSR", "U.P.", 1}}, {2, "CustomerB", "28/09/1999", Address{2, "BKJ", "BTM", "U.P.", 2}}}},
	}

	db, err := sql.Open("mysql", "root:123@/Customer_Service")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, _ = db.Query(`DROP TABLE Address;`)
	_, _ = db.Query(`DROP  TABLE Customer;`)
	_, _ = db.Query(`CREATE TABLE Customer ( id INT NOT NULL, Name VARCHAR(100) NOT NULL, DOB VARCHAR(40) NOT NULL, PRIMARY KEY (id) );`)
	_, _ = db.Query(`CREATE TABLE Address ( id INT NOT NULL , StreetName VARCHAR(100) NOT NULL, City VARCHAR(40) NOT NULL, State VARCHAR(40) NOT NULL, Customer_id INT, PRIMARY KEY (id), FOREIGN KEY (Customer_id)REFERENCES Customer(id));`)
	//fmt.Println("yes")

	for ind := range testcases {

		var c Customer
		c = InsertData(db, testcases[ind].input)
		//fmt.Println(c, testCases[ind].out)
		if !reflect.DeepEqual(c, testcases[ind].output) {
			t.Errorf("FAILED wanted %v got %v", testcases[ind].output, c)
		}
	}
}

//	rows, err := db.Query("TRUNCATE TABLE Address ; TRUNCATE TABLE Customer")
//	for ind := range testcases {
//		ans := getdata(db, testcases[ind].input)
//		if !cmp.Equal(ans, testcases[ind].output) {
//			t.Fatalf(`FAIL: %v Expected ans: %v Got: %v`, testcases[ind].input, testcases[ind].output, ans)
//		}
//	}
//}
