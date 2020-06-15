package author


import(
  . "github.com/sm3saurabh/book_store/models"
)

type InMemoryAuthorRepository struct {
  authors []Author
}

var authorRepoSingleton *InMemoryAuthorRepository = &InMemoryAuthorRepository {
  authors: initializeAuthors(),
}

func initializeAuthors() (ret []Author) {
  ret = append(ret, Author{
    Name: "Saurabh Mishra",
    Id: 1,
  }, Author{
    Name: "John Smith",
    Id: 2,
  }, Author{
    Name: "William Shakespeare",
    Id: 3,
  }, Author{
    Name: "Yo Yo Honey Singh",
    Id: 4,
  }, Author{
    Name: "John Doe",
    Id: 5,
  }, Author{
    Name: "Rainy Day",
    Id: 6,
  }, Author{
    Name: "Munshi Premchand",
    Id: 7,
  })

  return
}
