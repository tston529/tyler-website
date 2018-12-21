package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

//Work services my custom 404 page on a request to a bad page.
func Work(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	WorkVars := PageVars{
		Year:       now.Format("02-01-2006")[6:],
		HomeNav:    "",
		AboutMeNav: "",
		WorkNav:    template.HTMLAttr("class='selected'"),
	}

	t, err := template.ParseFiles("work.html", "footer.html") //parse the html file
	if err != nil {                                           // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.ExecuteTemplate(w, "work", WorkVars) //execute the template and pass it the struct
	if err != nil {                              // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
