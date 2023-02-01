package handlers

import (
	"fmt"
	"github.com/aZ4ziL/go_chat/auth"
	"github.com/aZ4ziL/go_chat/models"
	"net/http"
	"path/filepath"
	"text/template"
)

var (
	userHTMLBase     = filepath.Join("views/user/base.html")
	userHTMLLogin    = filepath.Join("views/user/login.html")
	userHTMLRegister = filepath.Join("views/user/register.html")
)

// Register is handler for registration user
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		defer flash.delete()
		tmpl, err := template.ParseFiles(userHTMLBase, userHTMLRegister)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := map[string]interface{}{
			"flash": flash,
		}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == http.MethodPost {
		firstName := r.PostFormValue("first_name")
		lastName := r.PostFormValue("last_name")
		username := r.PostFormValue("username")
		email := r.PostFormValue("email")
		password1 := r.PostFormValue("password1")
		password2 := r.PostFormValue("password2")

		if (firstName == "") || (lastName == "") || (username == "") || (email == "") || (password1 == "") || (password2 == "") {
			flash.setFlash("danger", "Please fill required field")
			http.Redirect(w, r, "/register", http.StatusFound)
			return
		}

		if password1 != password2 {
			flash.setFlash("danger", "Password doesn't same, please reenter the password.")
			http.Redirect(w, r, "/register", http.StatusFound)
			return
		}

		user := models.User{
			FirstName: firstName,
			LastName:  lastName,
			Username:  username,
			Email:     email,
			Password:  password2,
		}
		err := models.CreateNewUser(&user)
		if err != nil {
			flash.setFlash("danger", err.Error())
			http.Redirect(w, r, "/register", http.StatusFound)
			return
		}

		flash.setFlash("success", fmt.Sprintf("Successfully to create new user with username: %s", username))
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
}

// Login is handler for to get new token for user
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		defer flash.delete()

		tmpl, err := template.ParseFiles(userHTMLBase, userHTMLLogin)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := map[string]interface{}{
			"flash": flash,
		}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	if r.Method == http.MethodPost {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		user, err := models.GetUserByUsername(username)
		if err != nil {
			flash.setFlash("danger", "Username or password is incorrect")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		if !auth.DecryptionPassword(user.Password, password) {
			flash.setFlash("danger", "Username or password is incorrect")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		session, _ := store.Get(r, "user")
		session.Values["id"] = user.ID
		session.Values["full_name"] = user.FirstName + " " + user.LastName
		session.Values["username"] = user.Username
		session.Values["email"] = user.Email
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, "Failed to save your session.", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
}
