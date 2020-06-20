package models

type Author struct {
  Id uint32
  Name string
}

type AuthorList []Author

func (authors AuthorList) IsDuplicate(author Author) bool {
  for _, a := range authors {
    if a.Id == author.Id || a.Name == author.Name {
      return true
    }
  }

  return false
}
