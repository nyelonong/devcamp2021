package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	_ "github.com/lib/pq"
)

var router *httprouter.Router
var db *sql.DB

func GetDB() *sql.DB {
	var err error

	if db == nil {
		connStr := "user=postgres password=docker dbname=devcamp sslmode=disable"
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
	}

	return db
}

func setupRoutes() {
	router = httprouter.New()

	router.PUT("/shop", PutNewShop())
	router.POST("/product", UpdateProduct())
	router.GET("/product", GetProduct())
	router.GET("/productlist", GetProductList())
}

func main() {
	log.Println("Listening on port 8080")
	setupRoutes()

	log.Fatal(http.ListenAndServe(":8080", router))
}

func PutNewShop() httprouter.Handle {
	type Req struct {
		ID       int64  `json:"id"`
		ShopName string `json:"shop_name"`
	}
	type Res struct {
		Error string
	}

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req Req

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			log.Println(err)
			return
		}

		db := GetDB()
		_, err := db.Exec("INSERT INTO shop (id, name) VALUES($1, $2)", req.ID, req.ShopName)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Putting new shop success")

		res := &Res{
			Error: "",
		}
		fmt.Fprintf(w, res.Error)
	}
}

func GetProduct() httprouter.Handle {
	type Product struct {
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		Category string  `json:"category"`
	}

	type Res struct {
		Product Product `json:"product"`
		Error   string  `json:"error"`
	}

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		res := &Res{
			Error: "",
		}
		var product Product

		query := r.URL.Query()
		id, err := strconv.Atoi(query.Get("id"))
		if err != nil {
			log.Println(err)
			res.Error = http.StatusText(http.StatusBadRequest)
		}

		db := GetDB()
		if err := db.QueryRow("SELECT name, price, category FROM PRODUCT WHERE ID = $1", id).Scan(&product.Name, &product.Price, &product.Category); err != nil {
			log.Println(err)
		}

		res.Product = product

		encoder := json.NewEncoder(w)
		err = encoder.Encode(res)
		if err != nil {
			log.Println(err)
		}
	}
}

func UpdateProduct() httprouter.Handle {
	type Req struct {
		ID       int64   `json:"id"`
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		Category string  `json:"category"`
	}
	type Res struct {
		Error string
	}

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req Req

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			log.Println(err)
			return
		}

		db := GetDB()
		_, err := db.Exec("UPDATE product SET name=$1, price=$2, category=$3 WHERE ID=$4", req.Name, req.Price, req.Category, req.ID)
		if err != nil {
			log.Println(err)
			return
		}

		res := &Res{
			Error: "",
		}
		fmt.Fprintf(w, res.Error)
	}
}

func GetProductList() httprouter.Handle {
	type Product struct {
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		Category string  `json:"category"`
	}

	type Res struct {
		Product []Product `json:"product"`
		Error   string    `json:"error"`
	}

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		res := &Res{
			Error: "",
		}

		query := r.URL.Query()
		page, err := strconv.Atoi(query.Get("page"))
		if err != nil || page <= 0 {
			page = 1
		}

		limit, err := strconv.Atoi(query.Get("limit"))
		if err != nil || limit <= 0 {
			limit = 2
		}

		offset := (page - 1) * limit

		db := GetDB()
		rows, err := db.Query(`SELECT name, price, category FROM PRODUCT LIMIT $1 OFFSET $2`, limit, offset)
		if err != nil {
			log.Println(err)
		}

		defer rows.Close()
		for rows.Next() {
			var product Product
			err := rows.Scan(&product.Name, &product.Price, &product.Category)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(product)
			res.Product = append(res.Product, product)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		encoder := json.NewEncoder(w)
		err = encoder.Encode(res)
		if err != nil {
			log.Println(err)
		}
	}
}
