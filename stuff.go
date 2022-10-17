package main

import (
	"fmt"
)

// func homeHandler (w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
// }

// func contactHandler (w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	fmt.Fprint(w, "<h1>Contact page</h1><p>To get in touch, email me at <a href=\"cbj@gmail.com\">cbj@gmail.com</a>.")
// }

// func faqHandler (w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	fmt.Fprint(w, "<h1>FAQ Page</h1><p> Q: Is there a free version? \n A: Yes")
// }
// // func pathHandler (w http.ResponseWriter, r *http.Request) {
// // 	switch r.URL.Path {
// // 	case "/":
// // 		homeHandler(w, r)
// // 	case "/contact":
// // 		contactHandler(w, r)
// // 	default:
// // 	http.Error(w, "Page not found", http.StatusNotFound)
// // 	}
// // }

// type Router struct {}

// func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	case "/faq":
// 		faqHandler(w, r)
// 	default:
// 	http.Error(w, "Page not found", http.StatusNotFound)
// 	}
// }

// func main() {
// 	r := chi.NewRouter()
// 	fmt.Println("Starting server on port :3000....")
// 	http.ListenAndServe(":3000", r)
// }

 type twins struct {
	playerName string
	playerNumber int
	battingAve float32
	isPlaying bool
 }
func main () {
  aTwins := twins {
	playerName: "Luis Arraez",
	playerNumber: 4,
	battingAve: .317,
	isPlaying: true,
  }

  Twins := twins{
	playerName: "Joe Mauer",
	playerNumber: 7,
	battingAve: .325,
	isPlaying: true,
	}
 
 fmt.Println(aTwins)
}