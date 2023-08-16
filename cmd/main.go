package main

import (
	"database/sql"
	"log"
	"net/http"
	handler "sha/pkg/http"
	"sha/pkg/repository"
	"sha/pkg/usecase"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/go-sha")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepo := &repository.UserRepository{
		DB: db,
	}

	userUseCase := &usecase.UserUseCase{
		UserRepo: *userRepo,
	}

	userHandler := &handler.UserHandler{
		UserUseCase: *userUseCase,
	}

	http.HandleFunc("/users", userHandler.HandleGetAllUsers)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
