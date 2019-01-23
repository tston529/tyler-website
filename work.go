package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
    _ "os"
    _ "github.com/lib/pq"
)

// var db *sql.DB

//Work services the 'Work I've Done' page,
func Work(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	ws, _ := querySlides("work")
	WorkVars := PageVars{
		PageName:	"Tyler Stoney - Work I've Done",
		Year:       now.Format("02-01-2006")[6:],
		HomeNav:    "",
		AboutMeNav: "",
		WorkNav:    template.HTMLAttr("class='selected'"),
		WorkSlides:	ws,
	}

	t, err := template.New("work").Funcs(fm).ParseFiles("header.html", "work.html", "footer.html") //parse the html file
	if err != nil {                                           // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.ExecuteTemplate(w, "work", WorkVars) //execute the template and pass it the struct
	if err != nil {                              // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
