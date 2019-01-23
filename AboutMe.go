package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"time"
    _ "os"
    _ "github.com/lib/pq"
)

var db *sql.DB

//AboutMe services my custom 404 page on a request to a bad page.
func AboutMe(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	ws, _ := querySlides("aboutme")
	AboutMeVars := PageVars{
		PageName:	"Tyler Stoney - The Man, The Myth, The Legend",
		Year:       now.Format("02-01-2006")[6:],
		HomeNav:    "",
		AboutMeNav: template.HTMLAttr("class='selected'"),
		WorkNav:    "",
		WorkSlides:	ws,
	}

	t, err := template.New("aboutMe").Funcs(fm).ParseFiles("header.html", "about-me.html", "footer.html") //parse the html file
	if err != nil {                                               // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.ExecuteTemplate(w, "aboutMe", AboutMeVars) //execute the template and pass it the struct
	if err != nil {                                    // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
