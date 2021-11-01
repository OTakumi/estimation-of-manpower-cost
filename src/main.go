package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHttp method processes HTTP request
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("../view/html/",
				t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	// root
	dir, _ := os.Getwd()
	log.Print(http.Dir(dir + "../view/dist/js/"))
	http.Handle("/", &templateHandler{filename: "index.html"})
	// http.Handle("../view/dist/js/", http.StripPrefix("../view/dist/js/",
	// 	http.FileServer(http.Dir(dir+"../view/dist/js/"))))

	// Start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
