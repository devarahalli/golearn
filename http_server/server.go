package main

//import "fmt"
import "net/http"
import "log"

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":5000", nil)

}
