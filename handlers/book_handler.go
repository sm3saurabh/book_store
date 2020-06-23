package handlers

import(
  "github.com/gorilla/mux"
  "net/http"
  "strconv"
  "encoding/json"
  models "github.com/sm3saurabh/book_store/models"
  repo "github.com/sm3saurabh/book_store/repo/book"
)

var bookRepo repo.BookRepository

func HanldeBookRequests(router *mux.Router) {
  bookRepo = repo.NewBooksRepository()

  router.HandleFunc("/books", getBooks).Methods("GET")
  router.HandleFunc("/add", addBooks).Methods("POST")
  router.HandleFunc("/update", updateBook).Methods("PUT")
  router.HandleFunc("/delete/{id}", deleteBook).Methods("DELETE")
  router.HandleFunc("/search", searchBook).Methods("GET")
  router.HandleFunc("/range", getBooksInPriceRange).Methods("GET")
  router.HandleFunc("/books/{genre}", getBooksInGenre).Methods("GET")
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

func getBooksInGenre(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)

  if genre, ok := params["genre"]; ok {
    genreBooks := bookRepo.GetBooksInGenre(genre)

    if len(genreBooks) == 0 {
      err := json.NewEncoder(w).Encode("Invalid genre or books not found")

      if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
      }
      return
    }

    err := json.NewEncoder(w).Encode(genreBooks)

    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  }
}

func getBooksInPriceRange(w http.ResponseWriter, r *http.Request) {
  u := r.URL.Query().Get("upper")
  l := r.URL.Query().Get("lower")

  var lower, upper float64

  if len(u) == 0 {
    upper = 5000
  } else {
    var err error
    upper, err =  strconv.ParseFloat(u, 64)

    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
  }

  if len(l) == 0 {
    lower = 0
  } else {
    var err error
    lower, err = strconv.ParseFloat(l, 64)

    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
  }

  books := bookRepo.GetBookInPriceRange(lower, upper)

  if len(books) == 0 {
    json.NewEncoder(w).Encode("Could not return books for the given price range")
  } else {
    err := json.NewEncoder(w).Encode(books)

    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  }
}

func getBooks(w http.ResponseWriter, r *http.Request) {
  nonEmptyBooks := bookRepo.GetOnlyNonEmptyBooks()
  parsingError := json.NewEncoder(w).Encode(nonEmptyBooks)

  if parsingError != nil {
    http.Error(w, parsingError.Error(), http.StatusInternalServerError)
  }
}
