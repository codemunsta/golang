package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	sessions2 "github.com/gorilla/sessions"
	"github.com/rs/cors"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

var redisClient *redis.Client
var templates *template.Template
var sessions = sessions2.NewCookieStore([]byte("0skken_lzm_%4s3cr3t"))

func main() {
	fmt.Println("Server Started")
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	templates = template.Must(template.ParseGlob("templates/*.html"))
	router := mux.NewRouter()
	fileServer := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))
	// http.HandleFunc("/", handler)
	router.HandleFunc("/", indexHandler).Methods(http.MethodGet)
	router.HandleFunc("/form", formHandler).Methods(http.MethodGet)
	router.HandleFunc("/form", formHandler).Methods(http.MethodPost)
	router.HandleFunc("/login", loginHandler).Methods(http.MethodGet)
	router.HandleFunc("/login", loginHandler).Methods(http.MethodPost)
	router.HandleFunc("/register", registerHandler).Methods(http.MethodGet)
	router.HandleFunc("/register", registerHandler).Methods(http.MethodPost)
	// router.HandleFunc("/goodbye", goodbyeHandler).Methods(http.MethodGet)
	// router.HandleFunc("/test-session", testHandler).Methods(http.MethodGet)
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
	session, _ := sessions.Get(r, "session")
	name, ok := session.Values["username"]
	if !ok {
		http.Redirect(w, r, "/login", 302)
		return
	} else {
		fmt.Println(name)
		comment, err := redisClient.LRange("comments", 0, 10).Result()
		if err != nil {
			return
		}
		templates.ExecuteTemplate(w, "index.html", comment)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Get(r, "session")
	_, ok := session.Values["username"]
	if !ok {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if r.Method == "POST" {
		r.ParseForm()
		comment := r.PostForm.Get("Comment")
		redisClient.LPush("comments", comment)
		http.Redirect(w, r, "/", 302)
	} else {
		templates.ExecuteTemplate(w, "form.html", nil)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")
		hash, error := redisClient.Get("user:" + username).Bytes()
		if error != nil {
			return
		}
		compare := bcrypt.CompareHashAndPassword(hash, []byte(password))
		if compare != nil {
			return
		}
		session, err := sessions.Get(r, "session")
		if err != nil {
			return
		}
		session.Values["username"] = username
		session.Save(r, w)
		http.Redirect(w, r, "/", 302)
	} else {
		templates.ExecuteTemplate(w, "login.html", nil)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")
		cost := bcrypt.DefaultCost
		hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
		if err != nil {
			return
		}
		redisClient.Set("user:"+username, hash, 0)
		http.Redirect(w, r, "/login", 302)
	} else {
		templates.ExecuteTemplate(w, "register.html", nil)
	}
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
