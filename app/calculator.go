package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Operation struct {
	A  int     `json:"a,omitempty"`
	B  int     `json:"b,omitempty"`
	OP string  `json:"op,omitempty"`
	R  float32 `json:"r,omitempty"`
}

func (this Operation) toString() string {
	return strconv.Itoa(this.A) + this.OP + strconv.Itoa(this.B) + " = " + fmt.Sprint(this.R)
}

var operations []string

func NewOp(a string, b string, operator string) string {
	num1, err := strconv.Atoi(a)
	if err != nil {
		return "Check that the values are entered correctly."
	}
	num2, err := strconv.Atoi(b)
	if err != nil {
		return "Check that the values are entered correctly."
	}
	var op Operation
	op.A = num1
	op.B = num2
	op.OP = operator

	switch operator {
	case "+":
		op.R = float32(num1) + float32(num2)
	case "-":
		op.R = float32(num1) - float32(num2)
	case "*":
		op.R = float32(num1) * float32(num2)
	case "/":
		if num2 == 0.0 {
			return "Cannot divide by 0."
		}
		op.R = float32(num1) / float32(num2)
	}
	operations = append(operations, op.toString())
	return op.toString()
}
func GethHistoric(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(operations)
}
func GetDiv(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(NewOp(params["a"], params["b"], "/"))
}
func GetMul(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(NewOp(params["a"], params["b"], "*"))
}
func GetSub(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(NewOp(params["a"], params["b"], "-"))
}
func GetSum(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(NewOp(params["a"], params["b"], "+"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/calc/sum/{a}/{b}", GetSum).Methods("POST")
	router.HandleFunc("/calc/sub/{a}/{b}", GetSub).Methods("POST")
	router.HandleFunc("/calc/mul/{a}/{b}", GetMul).Methods("POST")
	router.HandleFunc("/calc/div/{a}/{b}", GetDiv).Methods("POST")
	router.HandleFunc("/calc/history", GethHistoric).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
