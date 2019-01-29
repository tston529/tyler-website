package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var phrases = []string{
	"Level 10 Warlock.",
	"World's Best Frycook.*</h2><div style='font-size: .5em;'>*Source Needed</div><h2>",
	"Bob Ross Junkie.",
	"Instant Ramen Connoisseur.",
	"Eclectic Musician.",
	"Casual PC Gamer.",
    "Student of the Sea.",
}

//Index parses and serves the index page
func Index(w http.ResponseWriter, r *http.Request) {
/*
	if r.URL.Path != "/" || r.URL.Path != "/images/" || r.URL.Path != "/styles/" || r.URL.Path != "/scripts/"{
        errorHandler(w, r, http.StatusNotFound)
        return
	}*/
	
	now := time.Now()  // find the time right now
	iVars := PageVars{ //store the date and time in a struct
		PageName:	"Tyler Stoney",
		Year:       now.Format("02-01-2006")[6:],
		HomeNav:    template.HTMLAttr("class='selected'"),
		AboutMeNav: "",
		WorkNav:    "",
		Phrase:     template.HTML(phrases[rand.Intn(len(phrases))]),
	}

	t, err := template.ParseFiles("header.html", "index.html", "footer.html") //parse the html file homepage.html
	if err != nil {                                            // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.ExecuteTemplate(w, "index", iVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                            // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
