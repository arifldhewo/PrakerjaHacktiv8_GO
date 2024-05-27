package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type product struct {
	ID int
	Name string
	Price int
}

var products = []product{}

func main () {
	http.HandleFunc("/products", productsResource)

	fmt.Println("Apps Listening to port 8000")
	http.ListenAndServe(":8000", nil)
}

func productsResource(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	
	case http.MethodGet:
		getProducts(w, r)
	
	case http.MethodPost:
		postProducts(w, r)

	case http.MethodPut:
		putProducts(w, r)

	case http.MethodDelete:
		delProducts(w, r)

	default: 
		http.Error(w, "Invalid Method", http.StatusBadRequest)
	}
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	defaultRes := map[string]any {
		"status" : "success",
		"message" : "Data not found",
	}

	res := map[string]any {
		"status" : "success",
		"data" : products,
	}
	
	if len(products) == 0 {
		json.NewEncoder(w).Encode(defaultRes)
		return
	}

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(res)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func postProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		price := r.FormValue("price")

		priceConv, err := strconv.Atoi(price)

		if err != nil {
			http.Error(w, "Invalid price", http.StatusBadRequest)
		}

		newProduct := product{
			ID: len(products) + 1,
			Name: name,
			Price: priceConv,
		}

		products = append(products, newProduct)

		json.NewEncoder(w).Encode(newProduct)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func putProducts(w http.ResponseWriter, r *http.Request) {

	isNotEdited := true

	w.Header().Set("Content-Type", "application/json")

	pathRes := map[string]any {
		"status" : "failed",
		"message" : "Query Params Only Accepting Int",
	}

	priceRes := map[string]any {
		"status" : "failed",
		"message" : "Price Invalid",
	}

	notFoundRes := map[string]any {
		"status" : "failed",
		"message" : "Data not found",
	}

	if r.Method == "PUT" {
		name := r.FormValue("name")
		price := r.FormValue("price")

		path, errPath := strconv.Atoi(r.URL.RawQuery[3:])
		if errPath != nil {
			json.NewEncoder(w).Encode(pathRes)
			return
		}

		priceConv, errPrice := strconv.Atoi(price)
		if errPrice != nil {
			json.NewEncoder(w).Encode(priceRes)
			return
		}

		for index, value := range products {
			if value.ID == path {
				products[index].Name = name
				products[index].Price = priceConv
				isNotEdited = false
			}
		}

		if isNotEdited == true {
			json.NewEncoder(w).Encode(notFoundRes)
			return
		}

		type newRes struct {
			ID int
			Name string
			Price int
		}

		newRes1 := newRes {
			ID: path,
			Name: name,
			Price: priceConv,
		}

		res := map[string]any {
			"status" : "success",
			"data" : newRes1,
		}

		json.NewEncoder(w).Encode(res)
		return
	}

	http.Error(w, "Invalid Method", http.StatusBadRequest)
}

func delProducts(w http.ResponseWriter, r *http.Request) {
	isNotDeleted := true

	w.Header().Set("Content-Type", "application/json")

	pathRes := map[string]any {
		"status" : "failed",
		"message" : "Query Params Only Accepting Int",
	}

	notFoundRes := map[string]any {
		"status" : "failed",
		"message" : "Data not found",
	}

	if r.Method == "DELETE" {

		path, errPath := strconv.Atoi(r.URL.RawQuery[3:])
		if errPath != nil {
			json.NewEncoder(w).Encode(pathRes)
			return
		}

		for index, value := range products {
			if value.ID == path {
				products = append(products[:index], products[index+1:]...)
				isNotDeleted = false
			}
		}

		if isNotDeleted == true {
			json.NewEncoder(w).Encode(notFoundRes)
			return
		}
		
		res := map[string]any {
			"status" : "success",
			"message" : "Deleted Successfully",
		}

		json.NewEncoder(w).Encode(res)
		return
	}
}