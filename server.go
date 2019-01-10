package main

import (
	"html/template"
	"log"
    "crypto/subtle"
    "golang.org/x/crypto/bcrypt"
	"net/http"
)

var dirTree = map[string]func(http.ResponseWriter, *http.Request){
	"index":    Index,
	"about-me": AboutMe,
    "work":     Work,
    "edit":     BasicAuth(Edit, "tston529", "$2a$14$6O3jP0wxR.0rLzZsKQmBSOuIMTLphE9dqVJ7WEQKaHcMjYLEbwcgu", "Please enter your username and password:"),
}

var staticImgs = []string{"EmailMe.png", "LinkedIn.png"}

type WorkData struct {
    Title   string
    Date    string
    Body    template.HTML
    DispOrder int
    RowId   int
}

type Notification struct {
    MsgClass template.HTMLAttr
    MsgData string
}

type PageVars struct {
	Funcs		template.FuncMap
	PageName	string
	Phrase     template.HTML
	Year       string
	HomeNav    template.HTMLAttr
	AboutMeNav template.HTMLAttr
	WorkNav    template.HTMLAttr
	WorkSlides	[]WorkData
    Notif      Notification
}

var fm = template.FuncMap{"add": func(a, b int) int {
    return a + b
}}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		FourOhFour(w, r)
	}
}

func BasicAuth(handler http.HandlerFunc, username, pHash, realm string) http.HandlerFunc {

    return func(w http.ResponseWriter, r *http.Request) {

        user, pass, ok := r.BasicAuth()

        pHashErr := bcrypt.CompareHashAndPassword([]byte(pHash), []byte(pass))

        if !ok ||
            subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 ||
            pHashErr != nil {
            w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
            w.WriteHeader(401)
            w.Write([]byte("You shall not pass.\n"))
            return
        }

        handler(w, r)
    }
}

func main(){
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("scripts"))))
    http.HandleFunc("/", Index)
	// http.HandleFunc("/edit", )
	for k, v := range dirTree {
		http.HandleFunc("/"+k, v)
	}

	log.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
