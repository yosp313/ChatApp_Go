package main

import(
  "fmt"
  "net/http"
)

func main(){
  router := http.NewServeMux()

  router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "templates/index.html")
  })
  
  server := http.Server{
    Addr: ":8080",
    Handler: router,
  }

  server.ListenAndServe()
  fmt.Println("Server Listening on port 8080")
}
