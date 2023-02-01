package handlers

import (
	"encoding/json"
	"github.com/aZ4ziL/go_chat/models"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"
)

var (
	chatHTMLBase = filepath.Join("views", "chat", "base.html")
	chatHTMLRoom = filepath.Join("views", "chat", "room.html")
)

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "user")
	if err != nil {
		flash.setFlash("danger", "Please login before accessing this page.")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if len(session.Values) == 0 {
		flash.setFlash("danger", "Please login before accessing this page.")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	query := r.URL.Query()

	if r.Method == http.MethodGet {
		defer flash.delete()

		if query.Get("get_full_name_by_user_id") != "" {
			w.Header().Set("Content-Type", "application/json")
			userId, _ := strconv.Atoi(query.Get("get_full_name_by_user_id"))
			fullName := getUserFullName(uint(userId))
			data := map[string]interface{}{
				"full_name": fullName,
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
			return
		}

		roomID, _ := strconv.Atoi(query.Get("id"))
		group, err := models.GetGroupByID(uint(roomID))
		if err != nil {
			http.Error(w, "Group Not Found", http.StatusNotFound)
			return
		}

		data := map[string]interface{}{
			"id":        session.Values["id"],
			"full_name": session.Values["full_name"],
			"username":  session.Values["username"],
			"email":     session.Values["email"],
			"group":     group,
		}

		tmpl := template.Must(template.New("base.html").Funcs(funcMap).ParseFiles(chatHTMLBase, chatHTMLRoom))
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == http.MethodPost {
		uID, _ := strconv.Atoi(r.PostFormValue("user_id"))
		gID, _ := strconv.Atoi(r.PostFormValue("group_id"))
		text := r.PostFormValue("text")

		chat := models.Chat{
			UserID:  uint(uID),
			GroupID: uint(gID),
			Text:    text,
		}
		err := models.CreateNewChat(&chat)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("done"))
		return
	}
}
