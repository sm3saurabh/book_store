package repo

import(
  . "github.com/sm3saurabh/book_store/models"
)

type BookRepository interface {
  // Get all the books which are not empty in the list
  GetOnlyNonEmptyBooks() (ret []Book)
  // Get all the books in the given price range
  GetBookInPriceRange(lower float64, upper float64) []Book
  // Add a list of books to the existing book
  AddBooksToList(books []Book)
  // Update the given book in the list
  UpdateBookInTheList(book Book) bool
  // Delte the book by book id
  DeleteBookInTheList(id int) bool
  // Search book by title in the list of book
  SearchBookByTitle(title string) (Book, error)
}

