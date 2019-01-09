package main

import (
	"html/template"
	"log"
	"net/http"
)

var dirTree = map[string]func(http.ResponseWriter, *http.Request){
	"index":    Index,
	"about-me": AboutMe,
	"work":     Work,
}

var staticImgs = []string{"EmailMe.png", "LinkedIn.png"}

type workData struct {
        Title   string
        Date    string
        Body    template.HTML
}

type PageVars struct {
	PageName	string
	Phrase     template.HTML
	Year       string
	HomeNav    template.HTMLAttr
	AboutMeNav template.HTMLAttr
	WorkNav    template.HTMLAttr
	WorkSlides	[]workData
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		FourOhFour(w, r)
	}
}

func main(){
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("scripts"))))
	http.HandleFunc("/", Index)
	for k, v := range dirTree {
		http.HandleFunc("/"+k, v)
	}

	log.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
