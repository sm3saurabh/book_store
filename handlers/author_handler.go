package handlers

import (
	repo "github.com/sm3saurabh/book_store/repo/author"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

var authorRepo repo.AuthorRepository

func HandleAuthorRequests(router *mux.Router) {
  authorRepo = repo.NewInMemoryAuthorRepository()

  router.HandleFunc("/authors", getAuthors).Methods("GET")
}

func getAuthors(w http.ResponseWriter, r *http.Request) {
  authors := authorRepo.GetOnlyNonEmptyAuthors()

  err := json.NewEncoder(w).Encode(authors)

  if err != nil {
    http.Error(w, "Could not fetch author list", http.StatusInternalServerError)
  }
}


