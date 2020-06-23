package main


import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/sm3saurabh/book_store/handlers"
)

func getBookServerRouter() *mux.Router {
  router := mux.NewRouter()

  router.Headers("Content-Type", "application/json")

  // Let the dedicated handler handle define all the end points
  handlers.HanldeBookRequests(router)


  return router
}

func main() {
  router := getBookServerRouter()

  log.Fatal(http.ListenAndServe(":8000", router))
}
