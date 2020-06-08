package main


import (
  "log"
  "strconv"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  models "github.com/sm3saurabh/book_store/models"
  repo "github.com/sm3saurabh/book_store/repo"
)

var bookRepo repo.BookRepository

func getBooks(w http.ResponseWriter, r *http.Request) {
  nonEmptyBooks := bookRepo.GetOnlyNonEmptyBooks()
  parsingError := json.NewEncoder(w).Encode(nonEmptyBooks)

  if parsingError != nil {
    http.Error(w, parsingError.Error(), http.StatusInternalServerError)
  }
}

// Add the book to the in memory list
// Also returns the new book list
// TODO - Change it with persistent storage like a db in future
func addBooks(w http.ResponseWriter, r *http.Request) {
  var books []models.Book

  parsingError := json.NewDecoder(r.Body).Decode(&books)

  if parsingError != nil {
    http.Error(w, parsingError.Error(), http.StatusBadRequest)
    return
  }

  bookRepo.AddBooksToList(books)

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

  wasPresent := bookRepo.UpdateBookInTheList(book)

  if !wasPresent {
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

    wasPresent := bookRepo.DeleteBookInTheList(id)

    if !wasPresent {
      http.Error(w, "Invalid id", http.StatusBadRequest)
      return
    }

    getBooks(w, r)

  } else {
    http.Error(w, "No id specified", http.StatusBadRequest)
  }
}

func searchBook(w http.ResponseWriter, r *http.Request) {
  bookTitle := r.URL.Query().Get("title")

  book, err := bookRepo.SearchBookByTitle(bookTitle)

  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  json.NewEncoder(w).Encode(book)

}

func getBookServerRouter() *mux.Router {
  router := mux.NewRouter()


  router.Headers("Content-Type", "application/json")

  router.HandleFunc("/books", getBooks).Methods("GET")
  router.HandleFunc("/add", addBooks).Methods("POST")
  router.HandleFunc("/update", updateBook).Methods("PUT")
  router.HandleFunc("/delete/{id}", deleteBook).Methods("DELETE")
  router.HandleFunc("/search", searchBook).Methods("GET")

  return router
}

func main() {
  bookRepo = repo.NewBooksRepository()

  router := getBookServerRouter()


  log.Fatal(http.ListenAndServe(":8000", router))
}

