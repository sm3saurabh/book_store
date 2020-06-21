package author


import(
  . "github.com/sm3saurabh/book_store/models"
)

type AuthorRepository interface {
  GetOnlyNonEmptyAuthors() (ret AuthorList)
  GetAuthorById(id uint32) (Author, error)
  GetAuthorsMatchingName(name string) (ret AuthorList)
  AddAuthor(author Author) error
  UpdateAuthor(author Author) error
  DeleteAuthor(author Author) error
}

