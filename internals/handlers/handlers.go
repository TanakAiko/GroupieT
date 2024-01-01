package handlers

import (
	"fmt"
	md "groupie/models"
	ts "groupie/tools"
	"net/http"
	"html/template"
)

type Errorr struct {
	Status int
	MesErr string
}

var (
	errs = []Errorr{
		{404, "Not found"},
		{405, "Not allowed"},
		{500, "Internal error server"},
		{400, "Bad Request"},
	}
)

func RenderTemplate(page string, data md.Api_Info) http.HandlerFunc {
	files := []string{"./templates/base.html", "./templates/" + page, "./templates/filterByMember.tmpl"}
	return func(w http.ResponseWriter, r *http.Request) {
		// Test the 404 error
		if r.URL.Path != "/" && r.URL.Path != "/artist" && r.URL.Path != "/search" && r.URL.Path != "/filter" {
			RenderError(w, errs[0])
			return
		}

		var toExec md.Info_OneArtist
		toExec.Info = data

		if r.URL.Path == "/artist" {
			Id := r.URL.Query().Get("Id")
			toPrint, err := ts.GetArtist(Id, data)
			if err != nil {
				RenderError(w, errs[3])
				return
			}
			toExec.ArtistSelected = toPrint
		} else if r.URL.Path == "/search" {
			dataSearch := r.FormValue("inputSearch")
			toExec.Search = ts.GetDataToPrint(dataSearch, data)
		} else if r.URL.Path == "/filter" {
			nbMember := r.FormValue("maCheckbox")
			fmt.Println(nbMember)
		}

		// Parse with the corresponding page
		tmpl, error := template.ParseFiles(files...)
		if error != nil {
			fmt.Println("Parsing")
			RenderError(w, errs[2])
			return
		}

		//execute with the corresponding data
		erro := tmpl.Execute(w, toExec)
		if erro != nil {
			fmt.Println("Executing")
			RenderError(w, errs[2])
			return
		}

	}
}
func RenderError(w http.ResponseWriter, data Errorr) {
	tmpl, err := template.ParseFiles("./templates/errors.html")
	if err != nil {
		http.Error(w, errs[2].MesErr, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(data.Status)
	errorr := tmpl.Execute(w, data)
	if errorr != nil {
		http.Error(w, errs[2].MesErr, http.StatusInternalServerError)
		return
	}
}
