// Package main Notes API server.
//
// @title           Notes API
// @version         1.0
// @description     Учебный REST API для заметок (CRUD).
// @contact.name    Denis Borisov
// @contact.email   denisba1995@gmail.com
// @BasePath        /api
package main

import (
	"log"
	"net/http"

	_ "example.com/goprac11-borisovda/docs"
	httpx "example.com/goprac11-borisovda/internal/http"
	"example.com/goprac11-borisovda/internal/http/handlers"
	"example.com/goprac11-borisovda/internal/repo"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	repo := repo.NewNoteRepoMem()
	h := &handlers.Handler{Repo: repo}
	r := httpx.NewRouter(h)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

	r.Get("/docs/*", httpSwagger.WrapHandler) // для chi: r.Get
}
