package main

import (
	"net/http"
	"github.com/siddontang/go/log"
	"sync"
	"text/template"
	"path/filepath"
)

const (
	TEMPLATE_PATH = "C:/Projects/GoDockerClient/src/goblueeprints/ch01/chat/templates"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(
			filepath.Join(TEMPLATE_PATH, t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	// get the room going
	go r.run()
	// start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

/*
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
	<html>
	    <head>
		<title>Chat</title>
	    </head>
	    <body>
		 let' s chat!
	    </chat>
	</html>
	`))
})*/

//http.Handle("/", &templateHandler{filename:"chat.html"})
//if err := http.ListenAndServe(":8080", nil); err != nil {
//	log.Fatal("ListenAndServe:", err)
//}
