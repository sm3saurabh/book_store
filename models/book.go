package models


type Book struct {
  Title string `json:"title"`
  Isbn int `json:"isbn"`
  Id int `json:"id"`
  Author *Author `json:"author"`
}

var books []Book

func init() {
  initializeBooks()
}

func initializeBooks() {
  books = append(books, Book {
    Title: "Book 1",
    Isbn: 13982,
    Id: 1,
    Author: &Author {
      FirstName: "Saurabh",
      LastName: "Mishra",
    },
  }, Book {
    Title: "Book 2",
    Isbn: 13982,
    Id: 2,
    Author: &Author {
      FirstName: "Saurabh",
      LastName: "Mishra",
    },
  }, Book {
    Title: "Book 3",
    Isbn: 13982,
    Id: 3,
    Author: &Author {
      FirstName: "Saurabh",
      LastName: "Mishra",
    },
  }, Book {
    Title: "Book 4",
    Isbn: 13982,
    Id: 4,
    Author: &Author {
      FirstName: "Saurabh",
      LastName: "Mishra",
    },
  }, Book {
    Title: "Book 5",
    Isbn: 13982,
    Id: 5,
    Author: &Author {
      FirstName: "Saurabh",
      LastName: "Mishra",
    },
  })
}

// Returns a filtered list of books without empty books
func  GetOnlyNonEmptyBooks() (ret []Book) {
  for _, book := range books {
    if  book.Id != 0 {
      ret = append(ret, book)
    }
  }
  return
}

// Adds a book to the to the private book list
func AddBookToList(book Book) {
  books = append(books, book)
}

func UpdateBookInTheList(book Book) int {
  var bookIndex int = -1

  for index, bookInList := range books {
    if bookInList.Id == book.Id {
      bookIndex = index
      break
    }
  }

  if bookIndex != -1 {
    books[bookIndex] = book
  }
  return bookIndex
}

func DeleteBookInTheList(id int) bool {
  var bookIndex int = -1

  for index, book := range books {
    if (book.Id == id) {
      bookIndex = index
      break
    }
  }

  if (bookIndex == -1) {
    return false
  }

  var length = len(books)
  books[bookIndex] = books[length - 1]
  books[length - 1] = Book{}
  books = books[:length - 1]

  return true
}
