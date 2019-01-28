package main

import (
	"database/sql"
	"html/template"
	"log"
	"strconv"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	_ "github.com/lib/pq"
)

//Edit allows me to edit slides, preventing the need for me to
//  log in to google cloud, hard-code data, and redeploy
func Edit(w http.ResponseWriter, r *http.Request) {

    r.ParseForm()

	var PageNot Notification

	if r.Method == "POST" {

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

    var ws []WorkData
    var tableNames []string
    var tableErr error
		var fileNames []string
    if r.FormValue("table") != "" {
        ws, tableErr = querySlides(r.FormValue("table"))
        if tableErr != nil {
            PageNot.MsgClass = "error"
            PageNot.MsgData  = "Table not found."
        }
        tableNames, _ = getTables()
    } else {
        tableNames, tableErr = getTables()
				fileNames = getFiles()
        //
    }
	EditVars := PageVars{
		PageName:	"Tyler Stoney - Edit",
		WorkSlides:	ws,
		Notif:		PageNot,
    Table:      r.FormValue("table"),
    TableNames: tableNames,
		FileNames: 	fileNames,
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

// Updates data pertaining to a selected slide
func updateSlide(r *http.Request) (bool) {
	datastoreName := os.Getenv("POSTGRES_CONNECTION")
    var err error
    db, err = sql.Open("postgres", datastoreName)
    defer db.Close()
    if r.FormValue("delete") != "" {
        stmt := "DELETE FROM " + r.FormValue("table") + " WHERE rowid = $1"
    	_, err = db.Exec(stmt, r.FormValue("delete"))
    } else {
        stmt := "UPDATE " + r.FormValue("table") + " SET title = $1, work_date = $2, body = $3, disp_order = $4 WHERE rowid = $5"
        _, err = db.Exec(stmt, r.FormValue("name"), r.FormValue("date"), r.FormValue("body"), r.FormValue("num"), r.FormValue("rowid"))
    }
    if err != nil {
        return false
    }
    return true;
}

// Allows me to create a new slide on the requested page
func createSlide(r *http.Request) (bool) {
	datastoreName := os.Getenv("POSTGRES_CONNECTION")
    var err error
    db, err = sql.Open("postgres", datastoreName)
    defer db.Close()
    disp_order, _ := strconv.Atoi(r.FormValue("num"))
    stmt := `INSERT INTO ` + r.FormValue("table") + ` VALUES ($1, $2, $3, $4)`;
    _, err = db.Exec(stmt, r.FormValue("name"), r.FormValue("date"), r.FormValue("body"), int16(disp_order))
    if err != nil {
        log.Print("error creating slide: ", err)
        return false
    }
    return true;
}

// Will retrieve a list of editable files
func getFiles() (fileNames []string) {
	var fileExtensions = []string {
		".htm",
		".html",
		".css",
		".js",
		".txt",
		".scss",
		".sass",
	}

	err := filepath.Walk(".",
    func(path string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
		for _, ext := range fileExtensions {
			if strings.LastIndex(path, ext) != -1 {
    		fileNames = append(fileNames, path)
				break
			}
		}
		return nil
	})

	if err != nil {
	    log.Println(err)
	}

	return
}

// TODO
// Will react to ajax request saving (and loading?) files to edit
func Submit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		return
	}
	r.ParseForm()

	if r.FormValue("ajax") == "" {
		return
	}


}
