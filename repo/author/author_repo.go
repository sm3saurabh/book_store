package author


import(
  . "github.com/sm3saurabh/book_store/models"
)

type AuthorRepository interface {
  GetOnlyNonEmptyAuthors() []Author
  GetAuthorById(id uint32) Author
  GetAuthorsMatchingName(name string) []Author
  AddAuthor(author Author) error
  UpdateAuthor(author Author) error
  DeleteAuthor(author Author) error
}

