package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"html/template"
	"net/http"
)

var redisClient *redis.Client
var templates *template.Template

func main() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	templates = template.Must(template.ParseGlob("templates/*.html"))
	router := mux.NewRouter()
	// http.HandleFunc("/", handler)
	router.HandleFunc("/", indexHandler).Methods(http.MethodGet)
	router.HandleFunc("/form", formHandler).Methods(http.MethodGet)
	router.HandleFunc("/form", formHandler).Methods(http.MethodPost)
	// router.HandleFunc("/goodbye", goodbyeHandler).Methods(http.MethodGet)
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	fmt.Print(handler)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello world")
	//comment := []string{"Hello", "World", "123"}
	comment, err := redisClient.LRange("comments", 0, 10).Result()
	if err != nil {
		return
	}
	templates.ExecuteTemplate(w, "index.html", comment)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		comment := r.PostForm.Get("Comment")
		redisClient.LPush("comments", comment)
		http.Redirect(w, r, "/", 302)
	} else {
		templates.ExecuteTemplate(w, "form.html", nil)
	}
}

//func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "goodbye world")
//}

//import (
//	"fmt"
//	"net/http"
//	"time"
//)
//
//func helloWorldPage(w http.ResponseWriter, r *http.Request) {
//	switch r.URL.Path {
//	case "/":
//		fmt.Fprint(w, "Hello World")
//	case "/ninja":
//		fmt.Fprint(w, "Hello Caleb")
//	default:
//		fmt.Fprint(w, "Big fat error")
//	}
//	fmt.Printf("Handling function with %s request\n", r.Method)
//}
//
//func htmlVsPlain(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "text/html")
//	fmt.Fprint(w, "<h1>Hello World</h1>")
//}
//
//func timeout(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("Timeout Attempted")
//	time.Sleep(2 * time.Second)
//	fmt.Fprint(w, "<h1>Hello World</h1>")
//}
//
//func muskRoute(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "Hello from new musk")
//}
//func main() {
//	http.HandleFunc("/", htmlVsPlain)
//	http.HandleFunc("/timeout", timeout)
//	//http.ListenAndServe("", nil)
//
//	server := http.Server{
//		Addr:         "",
//		Handler:      nil,
//		ReadTimeout:  1000,
//		WriteTimeout: 1000,
//	}
//	var newMux http.ServeMux
//	server.Handler = &newMux
//	newMux.HandleFunc("/", muskRoute)
//	server.ListenAndServe()
//}
