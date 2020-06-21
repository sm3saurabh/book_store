package models


type Book struct {
  Title string `json:"title"`
  Isbn int `json:"isbn"`
  Id int `json:"id"`
  Price float64 `json:"price"`
  Genre string `json:"genre"`
  PublishedAt uint64 `json:"published_at"`
  Author *Author `json:"author"`
}
