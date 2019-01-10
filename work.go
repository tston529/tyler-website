package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"fmt"
	"time"
    "os"
    _ "github.com/lib/pq"
)

var db *sql.DB

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

	t, err := template.New("work").Funcs(fm).ParseFiles("header.html", "work.html", "footer.html") //parse the html file
	if err != nil {                                           // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.ExecuteTemplate(w, "work", WorkVars) //execute the template and pass it the struct
	if err != nil {                              // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func queryWork() ([]WorkData, error) {
	// Set this in app.yaml when running in production.
    datastoreName := os.Getenv("POSTGRES_CONNECTION")
    // datastoreName := "postgres://jhtlhxyr:mOxqh49SFZ5uJi-I6AOSb9Yjdu8UdGne@baasu.db.elephantsql.com:5432/jhtlhxyr?sslmode=disable"
    var err error
    db, err = sql.Open("postgres", datastoreName)
    defer db.Close()

    rows, err := db.Query("SELECT title, work_date, body, disp_order, rowid FROM work ORDER BY disp_order")
    if err != nil {
        return nil, fmt.Errorf("Your table doesn't exist, perchance? %v", err)
    }
    defer rows.Close()

    var work []WorkData
    for rows.Next() {
            var w WorkData
            if err := rows.Scan(&w.Title, &w.Date, &w.Body, &w.DispOrder, &w.RowId); err != nil {
                return nil, fmt.Errorf("Something is funky in one of the rows: %v", err)
            }
            work = append(work, w)
    }

    return work, rows.Err()
}