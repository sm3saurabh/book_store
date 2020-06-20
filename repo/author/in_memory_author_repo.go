package author


import(
  "sort"
  "errors"
  "strings"
  . "github.com/sm3saurabh/book_store/models"
)

type InMemoryAuthorRepository struct {
  authors AuthorList
}

func (repo *InMemoryAuthorRepository) GetOnlyNonEmptyAuthors() (ret AuthorList) {
  for _, author := range repo.authors {
    if author.Id != 0 {
      ret = append(ret, author)
    }
  }

  return
}

func (repo *InMemoryAuthorRepository) GetAuthorById(id uint32) (Author, error) {

  index := sort.Search(len(repo.authors), func (i int) bool {
    return repo.authors[i].Id == id
  })

  if index != -1 {
    return repo.authors[index], nil
  } else {
    return repo.GetEmptyAuthor(), errors.New("Could not find author for the given id")
  }
}

func (repo *InMemoryAuthorRepository) GetAuthorsMatchingName(name string) (ret AuthorList) {

  for _, author := range repo.authors {
    if strings.Contains(strings.ToLower(author.Name), strings.ToLower(name)) {
      ret = append(ret, author)
    }
  }

  return
}

func (repo *InMemoryAuthorRepository) AddAuthor(author Author) error {
  if repo.authors.IsDuplicate(author) {
    return errors.New("This author is already present. Could not add to the list")
  }

  repo.authors = append(repo.authors, author)

  return nil
}

func (repo *InMemoryAuthorRepository) UpdateAuthor(author Author) error {
  var authorIndex int = -1

  for index, listAuthor := range repo.authors {
    if author.Id == listAuthor.Id {
      authorIndex = index
    }
  }

  if authorIndex != -1 {
    repo.authors[authorIndex] = author
  }

  return errors.New("This author does not exist")
}

func (repo *InMemoryAuthorRepository) DeleteAuthor(author Author) error {
  var authorIndex int = -1

  for index, listAuthor := range repo.authors {
    if (author.Id == listAuthor.Id) {
      authorIndex = index
      break
    }
  }

  if (authorIndex != -1) {
    var length = len(repo.authors)
    repo.authors[authorIndex] = repo.authors[length - 1]
    repo.authors[length - 1] = repo.GetEmptyAuthor()
    repo.authors = repo.authors[:length - 1]
  }

  return errors.New("This author does not exist")
}


func (repo *InMemoryAuthorRepository) GetEmptyAuthor() Author {
  return Author{}
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
