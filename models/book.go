package models


type Book struct {
  Title string `json:"title"`
  Isbn int `json:"isbn"`
  Id int `json:"id"`
  Author *Author `json:"author"`
}



