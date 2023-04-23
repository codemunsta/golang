package handlers

import (
	"demo/models"
	store "demo/sessions"
	"demo/templates"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello world")
	//comment := []string{"Hello", "World", "123"}
	comment, err := models.GetComments()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server error"))
		return
	}
	templates.ExecuteTemplates(w, "index.html", comment)
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		comment := r.PostForm.Get("Comment")
		err := models.PostComment(comment)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server error"))
			return
		}
		http.Redirect(w, r, "/", 302)
	} else {
		templates.ExecuteTemplates(w, "form.html", nil)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")
		err := models.AuthenticateUser(username, password)
		if err != nil {
			switch err {
			case models.UserNotFound:
				templates.ExecuteTemplates(w, "login.html", "unknown user error")
			case models.InvalidPassword:
				templates.ExecuteTemplates(w, "login.html", "invalid password")
			default:
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server error"))
			}
			return
		}
		session, err2 := store.SessionHouse.Get(r, "session")
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server error"))
			return
		}
		session.Values["username"] = username
		session.Save(r, w)
		http.Redirect(w, r, "/", 302)
	} else {
		templates.ExecuteTemplates(w, "login.html", nil)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")
		err := models.RegisterUser(username, password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server error"))
			return
		}
		http.Redirect(w, r, "/login", 302)
	} else {
		templates.ExecuteTemplates(w, "register.html", nil)
	}
}
