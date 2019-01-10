package main

import (
	"database/sql"
	"html/template"
	"log"
	"strconv"
	"net/http"
	"os"
	_ "github.com/lib/pq"
)

//Edit allows me to edit slides, preventing the need for me to
//  log in to google cloud, hard-code data, and redeploy
func Edit(w http.ResponseWriter, r *http.Request) {

	var PageNot Notification

	if r.Method == "POST"{

		if r.FormValue("rowid") == "NEW" {
			if !createSlide(r) {
				PageNot.MsgClass = "error"
				PageNot.MsgData  = "Failed to create slide."
			} else {
				PageNot.MsgClass = "success"
				PageNot.MsgData  = "Slide created successfully!"
			}
		} else {
			if !updateSlide(r) {
				PageNot.MsgClass = "error"
				PageNot.MsgData  = "Failed to update slide."
			} else {
				PageNot.MsgClass = "success"
				PageNot.MsgData  = "Slide updated successfully!"
			}
		}
	}

	ws, _ := queryWork()
	EditVars := PageVars{
		PageName:	"Tyler Stoney - Edit",
		WorkSlides:	ws,
		Notif:		PageNot,
	}

	t, err := template.ParseFiles("header.html", "edit.html") //parse the html file
	if err != nil {                                               // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.ExecuteTemplate(w, "edit", EditVars) //execute the template and pass it the struct
	if err != nil {                                    // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func updateSlide(r *http.Request) (bool) {
	datastoreName := os.Getenv("POSTGRES_CONNECTION")
    var err error
    db, err = sql.Open("postgres", datastoreName)
    defer db.Close()
    if r.FormValue("delete") != "" {
    	_, err = db.Exec("DELETE FROM work WHERE rowid = $1", r.FormValue("delete"))    	
    }
    _, err = db.Exec("UPDATE work SET title = $1, work_date = $2, body = $3, disp_order = $4 WHERE rowid = $5", r.FormValue("name"), r.FormValue("date"), r.FormValue("body"), r.FormValue("num"), r.FormValue("rowid"))
    if err != nil {
        return false
    }
    return true;
}

func createSlide(r *http.Request) (bool) {
	datastoreName := os.Getenv("POSTGRES_CONNECTION")
    var err error
    db, err = sql.Open("postgres", datastoreName)
    defer db.Close()
    disp_order, _ := strconv.Atoi(r.FormValue("num"))
    stmt := `INSERT INTO work VALUES ($1, $2, $3, $4)`;
    _, err = db.Exec(stmt, r.FormValue("name"), r.FormValue("date"), r.FormValue("body"), int16(disp_order))
    if err != nil {
        return false
    }
    return true;
}
