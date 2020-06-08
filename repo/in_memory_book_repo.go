package repo

import (
. "github.com/sm3saurabh/book_store/models"
)

type InMemoryBookRepository struct {
  books []Book
}


func (repo *InMemoryBookRepository) GetOnlyNonEmptyBooks() (ret []Book) {
  for _, book := range repo.books {
    if book.Id != 0 {
      ret = append(ret, book)
    }
  }
  return
}

func (repo *InMemoryBookRepository) AddBooksToList(newBooks []Book) {
  repo.books = append(repo.books, newBooks...)
}

func (repo *InMemoryBookRepository) UpdateBookInTheList(book Book) bool {
  var bookIndex int = -1

  for index, bookInList := range repo.books {
    if bookInList.Id == book.Id {
      bookIndex = index
      break
    }
  }

  if bookIndex != -1 {
    repo.books[bookIndex] = book
  }
  return bookIndex != -1
}


func (repo *InMemoryBookRepository) DeleteBookInTheList(id int) bool {
  var bookIndex int = -1

  for index, book := range repo.books {
    if (book.Id == id) {
      bookIndex = index
      break
    }
  }

  if (bookIndex == -1) {
    return false
  }

  var length = len(repo.books)
  repo.books[bookIndex] = repo.books[length - 1]
  repo.books[length - 1] = Book{}
  repo.books = repo.books[:length - 1]

  return true
}

var repoSingleton *InMemoryBookRepository = &InMemoryBookRepository {
  books: initializeBooks(),
}

func NewBooksRepository() *InMemoryBookRepository {
  return repoSingleton
}

func initializeBooks() []Book {
  books := make([]Book, 10)

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
  return books
}
