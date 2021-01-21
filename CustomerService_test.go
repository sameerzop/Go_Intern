package main

import (
	"database/sql"
	"github.com/google/go-cmp/cmp"
	"testing"
)


func TestCustomer(t *testing.T) {
	// Input: ID, Output: expected complete result from MYSQL server.
	//fmt.Println("dsds")

	testcases := []struct {
		input  int
		output []Customer
	}{
		{1, []Customer{{1,"CustomerA", "28/09/1997", Address{1,"AKJ","HSR","U.P.",1}}}},
		{2, []Customer{{2,"CustomerB", "28/09/1999", Address{2,"BKJ","BTM","U.P.",2}}}},
		{0, []Customer{{1,"CustomerA", "28/09/1997", Address{1,"AKJ","HSR","U.P.",1}},{2,"CustomerB", "28/09/1999", Address{2,"BKJ","BTM","U.P.",2}}}},
	}


	// Establishing connection with the database.
	db, err := sql.Open("mysql", "root:123@/Customer_Service")

	if err != nil {
				panic(err)
	}
	defer db.Close()
	for ind := range testcases{
		ans:= getdata(db,testcases[ind].input)
		if !cmp.Equal(ans,testcases[ind].output) {
			t.Fatalf(`FAIL: %v Expected ans: %v Got: %v`, testcases[ind].input,testcases[ind].output,ans)
		}
	}
}

