package handlers

import (
	"fmt"
	"github.com/aZ4ziL/go_chat/models"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

var (
	homeHTMLBase = filepath.Join("views/home/base.html")
	homeHTMLHome = filepath.Join("views/home/index.html")
)

// Home handler
func Home(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "user")
	if err != nil {
		flash.setFlash("info", "Please login before access this page.")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if len(session.Values) == 0 {
		flash.setFlash("info", "Please login before access this page.")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if r.Method == http.MethodGet {
		// if request method is GET
		defer flash.delete()
		groups := models.GetAllGroups()

		tmpl, err := template.ParseFiles(homeHTMLBase, homeHTMLHome)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := map[string]interface{}{
			"id":        session.Values["id"],
			"full_name": session.Values["full_name"],
			"username":  session.Values["username"],
			"email":     session.Values["email"],
			"groups":    groups,
			"flash":     flash,
		}

		w.WriteHeader(http.StatusOK)
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// if request is post
	// add new group from user
	if r.Method == http.MethodPost {
		title := r.PostFormValue("title")
		userId, _ := strconv.Atoi(r.PostFormValue("user_id"))
		description := r.PostFormValue("description")
		f, h, err := r.FormFile("logo")
		if err != nil {
			flash.setFlash("danger", "Please upload a image for group logo")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		defer f.Close()

		filename := h.Filename
		ext := filepath.Ext(filename)
		if !(strings.Contains(ext, "png") || strings.Contains(ext, "jpg")) {
			flash.setFlash("danger", "Please upload file for image PNG|JPG only.")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		var titleSlug string // declare title slug for image path
		if strings.Contains(title, " ") {
			titleSlug = strings.ToLower(strings.ReplaceAll(title, " ", "-"))
		} else {
			titleSlug = strings.ToLower(title)
		}
		dst := filepath.Join("media", "groups", titleSlug)
		_ = os.MkdirAll(dst, 0750) // mkdir in linux
		tempFile, err := os.OpenFile(dst+"/"+filename, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := io.Copy(tempFile, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// save to db
		group := models.Group{
			Title:       title,
			Logo:        "/" + dst + "/" + filename,
			Description: description,
			UserID:      uint(userId),
		}
		if err := models.CreateNewGroup(&group); err != nil {
			flash.setFlash("danger", err.Error())
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		flash.setFlash("success", fmt.Sprintf("Successfully to create new group chat."))
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
}
