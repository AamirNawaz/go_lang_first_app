package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server started listing..")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":9001", nil)
}

//hello server calling
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
