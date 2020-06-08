package repo

import(
  . "github.com/sm3saurabh/book_store/models"
)

type BookRepository interface {
  GetOnlyNonEmptyBooks() (ret []Book)
  AddBooksToList(books []Book)
  UpdateBookInTheList(book Book) bool
  DeleteBookInTheList(id int) bool
  SearchBookByTitle(title string) (Book, error)
}

