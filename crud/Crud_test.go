package main

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"reflect"
	"testing"
)

/*type Customer struct {
	ID   int
	Name string
	DOB string
	Addr Address
	//DOB  string
}
type Address struct {
	ID         int
	Streetname string
	City       string
	State      string
	Customerid int
}*/


func TestCustomerGet(t *testing.T) {
	testCases := []struct {
		input  string
		output []Customer
	}{
		{"CustomerA", []Customer{{1, "CustomerA", "28/09/1997", Address{1, "AKJ", "HSR", "U.P.", 1}}}},
		{"CustomerB", []Customer{{2, "CustomerB", "28/09/1999", Address{2, "BKJ", "BTM", "U.P.", 2}}}},
		}
	for i:=range testCases{
		var c []Customer
	//	reqst := "http://192.168.1.219:8082/Customer?name=" + testCases[i].input
	   reqst := "/Customer" + "?name=" + testCases[i].input
		req := httptest.NewRequest("GET",reqst , nil)
		//req = mux.SetURLVars(req, map[string]string{"id": testCases[i].input})
		w := httptest.NewRecorder()
		GetName(w, req)
		//resp := w.Result()
		//val, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(w.Body.Bytes(), &c)
		if err != nil {
			log.Fatal(err)
		}
		if !reflect.DeepEqual(c, testCases[i].output) {
			t.Errorf(`FAIL: %v Expected ans: %v Got: %v`, testCases[i].input, testCases[i].output, c)
		}
	}

}

	func TestCustomerGet2(t *testing.T) {
		testCases := []struct {
			input  string
			output []Customer
		}{
			{"1", []Customer{{1, "CustomerA",  "28/09/1997",Address{1, "AKJ", "HSR", "U.P.", 1}}}},
			{"2", []Customer{{2, "CustomerB",  "28/09/1999",Address{2, "BKJ", "BTM", "U.P.", 2}}}},

		}
		for i:=range testCases{
			var c []Customer
			//reqst := "http://192.168.1.219:8080/Customer?id=" + testCases[i].input
			reqst := "/Customer" + "?id=" + testCases[i].input
			req := httptest.NewRequest("GET", reqst, nil)
			w := httptest.NewRecorder()
			req = mux.SetURLVars(req, map[string]string{"id": testCases[i].input})
			GetID(w, req)
			resp := w.Result()
		    val, _ := ioutil.ReadAll(resp.Body)
			err := json.Unmarshal(val, &c)
			if err != nil {
				log.Fatal(err)
			}
			if !reflect.DeepEqual(c, testCases[i].output) {
				t.Errorf(`FAIL: %v Expected ans: %v Got: %v`, testCases[i].input, testCases[i].output, c)
			}
		}
	}
func TestCustomerPost(t *testing.T) {
	testCases := []struct {
		input  string
		output []Customer
	}{
		{"CustomerA", []Customer{{1, "CustomerA", "28/09/1997", Address{1, "AKJ", "HSR", "U.P.", 1}}}},
		{"CustomerB", []Customer{{2, "CustomerB", "28/09/1999", Address{2, "BKJ", "BTM", "A.P.", 2}}}},
		}
	for i:=range testCases{
		var c Customer
		byte,_:=json.Marshal(testCases[i].input)
		//reqst := "http://192.168.1.223:8080/Customer"
		reqst := "/Customer" + "?Customer=" + testCases[i].input
		req := httptest.NewRequest("POST", reqst, bytes.NewBuffer(byte))
		w := httptest.NewRecorder()
		PostCustomer(w, req)
		resp := w.Result()
		val, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(val, &c)
		if err != nil {
			log.Fatal(err)
		}
		if !reflect.DeepEqual(c, testCases[i].output) {
			t.Errorf(`FAIL: %v Expected ans: %v Got: %v`, testCases[i].input, testCases[i].output, c)
		}
	}
}


func TestCustomerPut(t *testing.T) {
	testCases := []struct {
		input  Customer
		output Customer
	}{
		{Customer{1, "CustomerA", " ", Address{1, "AKJ", "HSR", "U.P.", 1}}, Customer{1, "CustomerA", " ", Address{1, "AKJ", "HSR", "U.P.", 1}}},
		{Customer{2, "CustomerB", " ", Address{2, "BKJ", "BTM", "U.P.", 2}}, Customer{2, "CustomerB", " ", Address{2, "BKJ", "BTM", "U.P.", 2}}},
	}
	/*for i := range testCases {
		b, err := json.Marshal(testCases[i].input)
		if err != nil {
			log.Fatal(err)
		}
		req := httptest.NewRequest("PUT", "/Customer", bytes.NewBuffer(b))
		w := httptest.NewRecorder()
		PutCustomer(w, req)
	//	if !reflect.DeepEqual(testCases[i].output,c){
	//		t.Errorf(`FAIL: %v Expected ans: %v Got: %v`, testCases[i].input, testCases[i].output, c)
		if bytes.Equal(w.Body.Customer(), testCases[i].output) {
			t.Errorf("FAILED , expected is %v got %v", w.Body.Bytes(), testCases[i].output)
		}

	}
}*/
	for i:=range testCases{
		var c Customer
		reqst:="http://192.168.1.219/8080?id="+string(testCases[i].input.ID)
		//reqst := "/Customer" + "?id=" + testCases[i].input
		byte,_:=json.Marshal(testCases[i].input)
		req:=httptest.NewRequest("PUT",reqst,bytes.NewBuffer(byte))
		w:=httptest.NewRecorder()
		PutCustomer(w,req)
		resp:=w.Result()
		val,_:=ioutil.ReadAll(resp.Body)
		err:=json.Unmarshal(val,&c)
		if err!=nil{
			log.Fatal("Error")
		}
		if !reflect.DeepEqual(testCases[i].output,c){
			t.Errorf(`FAIL: %v Expected ans: %v Got: %v`, testCases[i].input, testCases[i].output, c)
		}
	}
}
func TestCustomerDelete(t *testing.T) {
	testCases := []struct {
		input  string
		output []Customer
	}{
		{"1", []Customer{{1, "CustomerA", "28/09/1997", Address{1, "AKJ", "HSR", "U.P.", 1}}}},
		{"2", []Customer{{2, "CustomerB", "28/09/1999", Address{2, "BKJ", "BTM", "U.P.", 2}}}},
	}
	for i:=range testCases{
		var c Customer
		//reqst:="http://192.168.1.223/8080/customer?name="+testCases[i].input
		reqst := "/Customer" + "?id=" + testCases[i].input
		req:=httptest.NewRequest("DELETE",reqst,nil)
		w:=httptest.NewRecorder()
	req = mux.SetURLVars(req, map[string]string{"id": testCases[i].input})
		DeleteCustomerById(w,req)
		resp:=w.Result()
		val,_:=ioutil.ReadAll(resp.Body)
		err:=json.Unmarshal(val,&c)
		if err!=nil{
			log.Fatal(err)
		}
		if !reflect.DeepEqual(c,testCases[i].output){
			t.Errorf(`FAIL: %v Expected ans: %v Got: %v`, testCases[i].input, testCases[i].output, c)
		}
	}
}







