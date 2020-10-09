package main

import(
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once sync.Once
	filename string
	template *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.template = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.template.Execute(w, nil)
}

func main() {
	http.Handle("/", &templateHandler{filename: "index.html"})

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}