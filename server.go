package main

import(
  "fmt"
  "net/http"
)

func main(){
  router := http.NewServeMux()

  router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
  })
  
  server := http.Server{
    Addr: ":8080",
    Handler: router,
  }

  server.ListenAndServe()
  fmt.Println("Server Listening on port 8080")
}
