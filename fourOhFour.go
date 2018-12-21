package main

import (
	"html/template"
	"log"
	"net/http"
)

//FourOhFourVariables will be populated with the name of the page
//  the user tried to access.
type FourOhFourVariables struct {
	Keyword string
}

//FourOhFour services my custom 404 page on a request to a bad page.
func FourOhFour(w http.ResponseWriter, r *http.Request) {

	FourOhFourVars := FourOhFourVariables{ //store the date and time in a struct
		Keyword: r.URL.Path[1:],
	}

	t, err := template.ParseFiles("404.html") //parse the html file
	if err != nil {                           // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, FourOhFourVars) //execute the template and pass it the struct
	if err != nil {                    // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
