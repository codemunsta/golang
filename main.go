package main

import (
	"demo/models"
	"demo/routers"
	"demo/templates"
	"fmt"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	fmt.Println("Server Started")
	models.Init()
	templates.LoadTemplates("templates/*.html")
	router := routers.NewRouter()
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

//func testHandler(w http.ResponseWriter, r *http.Request) {
//	session, err := sessions.Get(r, "session")
//	if err != nil {
//		return
//	}
//	name, exist := session.Values["username"]
//	if !exist {
//		return
//	}
//	username, ok := name.(string)
//	if !ok {
//		return
//	}
//	w.Write([]byte(username))
//	fmt.Fprintf(w, "\n %v", username)
//}

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
