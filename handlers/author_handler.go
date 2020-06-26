package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	repo "github.com/sm3saurabh/book_store/repo/author"
)

var authorRepo repo.AuthorRepository

func HandleAuthorRequests(router *mux.Router) {
  authorRepo = repo.NewInMemoryAuthorRepository()

  router.HandleFunc("/authors", getAuthors).Methods("GET")
  router.HandleFunc("/author/{id}", getAuthorById).Methods("GET")
}

func getAuthorById(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)

  if id, ok := params["id"]; ok {

    actualId, nfe := strconv.ParseUint(id, 10, 32)

    if nfe != nil {
      http.Error(w, "Unsupported id format", http.StatusBadRequest)
      return
    }

    author, err := authorRepo.GetAuthorById(uint32(actualId))

    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    json.NewEncoder(w).Encode(author)
  }
}

func getAuthors(w http.ResponseWriter, r *http.Request) {
  authors := authorRepo.GetOnlyNonEmptyAuthors()

  err := json.NewEncoder(w).Encode(authors)

  if err != nil {
    http.Error(w, "Could not fetch author list", http.StatusInternalServerError)
  }
}


