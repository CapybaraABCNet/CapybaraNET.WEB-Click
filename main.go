package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	ClickCount int
	Cen        int = 45
	Click      int = 1
	Text       string
)

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, struct {
		ClickCount int
		Text       string
		Cen        int
	}{ClickCount, Text, Cen})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func clickHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		ClickCount += Click
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func clickBetter(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if ClickCount >= Cen {
			ClickCount -= Cen
			Click += Cen
			Text = fmt.Sprintf("Вы успешно улучшили капибару, теперь за один клик она будет вам давать %d", Click)
			Cen *= 10
		} else {
			Text = "Вам не хватает капибара коинов для улучшения."
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/click", clickHandler)
	http.HandleFunc("/cl", clickBetter)
	fmt.Println("Сервер запущен на http://localhost:2080")
	if err := http.ListenAndServe(":2081", nil); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
