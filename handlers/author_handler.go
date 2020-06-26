package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	repo "github.com/sm3saurabh/book_store/repo/author"
  models "github.com/sm3saurabh/book_store/models"
)

var authorRepo repo.AuthorRepository

func HandleAuthorRequests(router *mux.Router) {
  authorRepo = repo.NewInMemoryAuthorRepository()

  router.HandleFunc("/authors", getAuthors).Methods("GET")
  router.HandleFunc("/authors/{name}", getAuthorsByName).Methods("GET")
  router.HandleFunc("/author", addAuthor).Methods("POST")
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

    parsingError := json.NewEncoder(w).Encode(author)

    if parsingError != nil {
      http.Error(w, parsingError.Error(), http.StatusInternalServerError)
    }
  }
}

func getAuthorsByName(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)

  if name, ok := params["name"]; ok {
    authors := authorRepo.GetAuthorsMatchingName(name)

    if len(authors) == 0 {
      http.Error(w, "No authors matching the name", http.StatusNotFound)
      return
    }

    parsingError := json.NewEncoder(w).Encode(authors)

    if parsingError != nil {
      http.Error(w, parsingError.Error(), http.StatusInternalServerError)
    }
  }
}

func addAuthor(w http.ResponseWriter, r *http.Request) {
  var author models.Author

  parsingError := json.NewDecoder(r.Body).Decode(&author)

  if parsingError != nil {
    http.Error(w, "Unsupported format", http.StatusBadRequest)
    return
  }

  err := authorRepo.AddAuthor(author)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  getAuthors(w, r)
}

func getAuthors(w http.ResponseWriter, r *http.Request) {
  authors := authorRepo.GetOnlyNonEmptyAuthors()

  err := json.NewEncoder(w).Encode(authors)

  if err != nil {
    http.Error(w, "Could not fetch author list", http.StatusInternalServerError)
  }
}


