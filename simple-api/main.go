package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Product struct {
	ProductId string `json:"id"`
	Name      string `json:"Name"`
	Blades    int    `json:"Blades"`
	Warrenty  int    `json:"Warrenty"`
}

var product []Product

func main() {
	fmt.Println("Welcome Api")
	r := mux.NewRouter()
	product = append(product, Product{ProductId: "2", Name: "USHA", Blades: 3, Warrenty: 2})
	product = append(product, Product{ProductId: "3", Name: "	ORIENT", Blades: 4, Warrenty: 3})
	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/product", GetAllProducts).Methods("GET")
	r.HandleFunc("/product/{id}", GetOneProduct).Methods("GET")
	r.HandleFunc("/product", CreateOneProduct).Methods("POST")
	r.HandleFunc("/product/{id}", UpdateOneProduct).Methods("PUT")
	r.HandleFunc("/product/{id}", DeleteOneProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))

}
func ServeHome(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Welcome to AmaFlip")
}
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Product")
	w.Header().Set("Content-Type", "applicatioan/json")
	json.NewEncoder(w).Encode(product)

}
func GetOneProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Product")
	w.Header().Set("Content-Type", "applicatioan/json")
	params := mux.Vars(r)
	for _, pro := range product {
		if pro.ProductId == params["id"] {
			json.NewEncoder(w).Encode(pro)
			return
		}
	}
	json.NewEncoder(w).Encode("Not Found Try again next Time")
	return

}
func CreateOneProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create all Product")
	w.Header().Set("Content-Type", "applicatioan/json")
	rand.Seed(time.Now().UnixNano())
	var products Product
	_ = json.NewDecoder(r.Body).Decode(&products)
	products.ProductId = strconv.Itoa(rand.Intn(100))
	product = append(product, products)
	json.NewEncoder(w).Encode(product)
	return
}
func UpdateOneProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one Product")
	w.Header().Set("Content-Type", "applicatioan/json")
	params := mux.Vars(r)
	for index, pro := range product {
		if pro.ProductId == params["id"] {
			product = append(product[:index], product[index+1:]...)
			var prod Product
			_ = json.NewDecoder(r.Body).Decode(&prod)
			prod.ProductId = params["id"]

			product = append(product, prod)
			json.NewEncoder(w).Encode(prod)
			return

		}
	}
}

func DeleteOneProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one Product")
	w.Header().Set("Content-Type", "applicatioan/json")
	params := mux.Vars(r)
	for index, pro := range product {
		if pro.ProductId == params["id"] {
			product = append(product[:index], product[index+1:]...)
			break
		}
	}
}
