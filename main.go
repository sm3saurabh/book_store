package main


import (
	"log"
  "strconv"
	"net/http"
  "encoding/json"
	"github.com/gorilla/mux"
  models "github.com/sm3saurabh/book_store/models"
)


func getBooks(w http.ResponseWriter, r *http.Request) {
  nonEmptyBooks := models.GetOnlyNonEmptyBooks()
  parsingError := json.NewEncoder(w).Encode(nonEmptyBooks)

  if parsingError != nil {
    http.Error(w, parsingError.Error(), http.StatusInternalServerError)
  }
}

// Add the book to the in memory list
// Also returns the new book list
// TODO - Change it with persistent storage like a db in future
func addBooks(w http.ResponseWriter, r *http.Request) {
  var book []models.Book

  parsingError := json.NewDecoder(r.Body).Decode(&book)

  if parsingError != nil {
    http.Error(w, parsingError.Error(), http.StatusBadRequest)
    return
  }

  models.AddBookToList(book)

  getBooks(w, r)
}

// Updates the given book in the list
// Also returns the newly updated book list
func updateBook(w http.ResponseWriter, r *http.Request) {
  var book models.Book

  parsingError := json.NewDecoder(r.Body).Decode(&book)

  if parsingError != nil {
    http.Error(w, parsingError.Error(), http.StatusBadRequest)
    return
  }

  bookIndex := models.UpdateBookInTheList(book)

  if bookIndex == -1 {
    http.Error(w, "Book id entered is not valid", http.StatusBadRequest)
    return
  }

  getBooks(w, r)
}

// Deletes the book in the list specified by it's id
// Also returns the newly updated list
func deleteBook(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)

  if bookId, ok := params["id"]; ok {
    id, err := strconv.Atoi(bookId)

    if err != nil {
      http.Error(w, "Invalid book id format", http.StatusBadRequest)
      return
    }

    wasPresent := models.DeleteBookInTheList(id)

    if !wasPresent {
      http.Error(w, "Invalid id", http.StatusBadRequest)
      return
    }

    getBooks(w, r)

  } else {
    http.Error(w, "No id specified", http.StatusBadRequest)
  }
}

func main() {
  r := mux.NewRouter()

  r.Headers("Content-Type", "application/json")

  r.HandleFunc("/books", getBooks).Methods("GET")
  r.HandleFunc("/add", addBooks).Methods("POST")
  r.HandleFunc("/update", updateBook).Methods("PUT")
  r.HandleFunc("/delete/{id}", deleteBook).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8000", r))
}

