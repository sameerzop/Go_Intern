package main

import "C"
import (
	"log"

	"github.com/stretchr/testify/assert"

	//	"github.com/stretchr/testify/assert"
	//"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCustomer(t *testing.T) {
	// Input: ID, Output: expected complete result from MYSQL server.
	//fmt.Println("dsds")

	testcases := []struct {
		input  int
		output []Customer
	}{
		{1, []Customer{{1, "CustomerA", "28/09/1997", Address{1, "AKJ", "HSR", "U.P.", 1}}}},
		{2, []Customer{{2, "CustomerB", "28/09/1999", Address{2, "BKJ", "BTM", "U.P.", 2}}}},
		{0, []Customer{{1, "CustomerA", "28/09/1997", Address{1, "AKJ", "HSR", "U.P.", 1}}, {2, "CustomerB", "28/09/1999", Address{2, "BKJ", "BTM", "U.P.", 2}}}},
	}

//	{
		db, mock, err := sqlmock.New()
		if err != nil {
			log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		//	return db, mock
		//}
		//{
		db, mock := NewMock()
		repo := &repository{db}
		defer func() {
			repo.Close()
		}()
		query := "SELECT Id, Name, DOB, Addr FROM Customer WHERE id = \\?"

		rows := sqlmock.NewRows([]string{"Id", "Name", "DOB", "Addr"}).
			AddRow(C.Id, C.Name, C.DOB, C.Addr)

		mock.ExpectQuery(query).WithArgs(C.Id).WillReturnRows(rows)

		Customer, err := repo.FindByID(C.Id)
		assert.NotNil(t, Customer)
		assert.NoError(t, err)
	}
}

//	// Establishing connection with the database.
//	db, err := sql.Open("mysql", "root:123@/Customer_Service")
//
//	if err != nil {
//		panic(err)
//	}
//	defer db.Close()
//	for ind := range testcases {
//		ans := getdata(db, testcases[ind].input)
//		if !cmp.Equal(ans, testcases[ind].output) {
//			t.Fatalf(`FAIL: %v Expected ans: %v Got: %v`, testcases[ind].input, testcases[ind].output, ans)
//		}
//	}
//}
