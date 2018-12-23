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

type PageVars struct {
	Phrase     template.HTML
	Year       string
	HomeNav    template.HTMLAttr
	AboutMeNav template.HTMLAttr
	WorkNav    template.HTMLAttr
}

/*func isIn(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}*/

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		FourOhFour(w, r)
	}
}

func main() {
	http.HandleFunc("/", Index)
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("scripts"))))

	for k, v := range dirTree {
		http.HandleFunc("/"+k, v)
	}

	log.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
