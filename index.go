package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var phrases = []string{
	"Level 8 Warlock.",
	"World's Best Frycook.*</h2><p>*Source Needed</p><h2>",
	"Bob Ross Fan.",
	"Instant Ramen Connoisseur.",
	"Eclectic Musician.",
	"Casual Gamer.",
}

//Index parses and serves the index page
func Index(w http.ResponseWriter, r *http.Request) {
	now := time.Now()  // find the time right now
	iVars := PageVars{ //store the date and time in a struct
		Year:       now.Format("02-01-2006")[6:],
		HomeNav:    template.HTMLAttr("class='selected'"),
		AboutMeNav: "",
		WorkNav:    "",
		Phrase:     template.HTML(phrases[rand.Intn(len(phrases))]),
	}

	t, err := template.ParseFiles("index.html", "footer.html") //parse the html file homepage.html
	if err != nil {                                            // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.ExecuteTemplate(w, "index", iVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                            // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
