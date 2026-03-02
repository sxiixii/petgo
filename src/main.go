package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/sxiixii/petgo/config"
	"github.com/sxiixii/petgo/internal/repository"
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

	conn, err := pgx.Connect(context.Background(), conf.PostbresURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	repo := repository.New(conn)
	user, err := repo.Get(context.Background(), "46e0f143-85fd-4502-a067-2652ebd6b424")
	if err != nil {
		fmt.Printf("Пользователь не найден \n")
	}

	var oneUser repository.User
	if len(user) > 0 {
		oneUser = user[0]
	}

	fmt.Printf("пользователь %s, %s\n", oneUser.Name, oneUser.Email)

	mux := http.NewServeMux()
	mux.HandleFunc("/", ArticlesCategoryHandler)

	fmt.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
