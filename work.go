package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"fmt"
	"time"
	_ "google.golang.org/appengine"

    _ "github.com/lib/pq"
)

var db *sql.DB

/*type work struct {
        title   string
        date    string
        body    string
}*/

//Work services the 'Work I've Done' page,
func Work(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	ws, _ := queryWork()
	WorkVars := PageVars{
		PageName:	"Tyler Stoney - Work I've Done",
		Year:       now.Format("02-01-2006")[6:],
		HomeNav:    "",
		AboutMeNav: "",
		WorkNav:    template.HTMLAttr("class='selected'"),
		WorkSlides:	ws,
	}

	t, err := template.ParseFiles("header.html", "work.html", "footer.html") //parse the html file
	if err != nil {                                           // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.ExecuteTemplate(w, "work", WorkVars) //execute the template and pass it the struct
	if err != nil {                              // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func queryWork() ([]workData, error) {
	// Set this in app.yaml when running in production.
    //datastoreName := os.Getenv("POSTGRES_CONNECTION")
    datastoreName := "user=postgres password=rAsIH0MM1C2ka4sq dbname=postgres sslmode=disable"

    var err error
    db, err = sql.Open("postgres", datastoreName)

    rows, err := db.Query("SELECT title, date, body FROM work")
    if err != nil {
            return nil, fmt.Errorf("Your table doesn't exist, perchance? %v", err)
    }
    defer rows.Close()

    var work []workData
    for rows.Next() {
            var w workData
            if err := rows.Scan(&w.Title, &w.Date, &w.Body); err != nil {
                    return nil, fmt.Errorf("Something is funky in one of the rows: %v", err)
            }
            work = append(work, w)
    }

    return work, rows.Err()
}