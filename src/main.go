package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/sxiixii/petgo/config"
)

type MyResponse struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func NewMyResponse(name string, age int) MyResponse {
	return MyResponse{
		Name: name,
		Age:  age,
	}
}

func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	myResponse := NewMyResponse("Alex", 33)
	myResponseToByte, err := json.Marshal(&myResponse)
	if err != nil {
		fmt.Printf("Ошибка сериализации данных %v", err)
	}
	_, err = w.Write(myResponseToByte)
	if err != nil {
		fmt.Printf("Ошибка отправки данных %v", err)
	}
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	conf := config.New()
	fmt.Printf("config is %v", conf)

	fmt.Println("Hello world!")
	mux := http.NewServeMux()
	mux.HandleFunc("/", ArticlesCategoryHandler)

	fmt.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
