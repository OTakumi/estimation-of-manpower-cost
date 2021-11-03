package main

import (
	"fmt"
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
			template.Must(template.ParseFiles(filepath.Join("./web/html/",
				t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	// root
	dir, _ := os.Getwd()
	fmt.Println(http.Dir(dir + "/web/ts"))

	http.Handle("/", &templateHandler{filename: "index.html"})

	// Directory of js files
	http.Handle("/ts/", http.StripPrefix("/ts/",
		http.FileServer(http.Dir(dir+"/ts/"))))

	// Start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
