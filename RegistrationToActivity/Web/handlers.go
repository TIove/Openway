package main

import (
	"RegistrationToActivity/api/auth"
	"RegistrationToActivity/api/record"
	"context"
	"encoding/json"
	"html/template"
	"net/http"
	"time"
)

type loginData struct {
	Login 	 string
	Password string
	Success  bool
	Data     string
}

type recordData struct {
	Name 		string
	LastName 	string
	Email       string
	Birthday 	time.Time
}

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	data := recordData{
		Name: r.FormValue("name"),
		LastName: r.FormValue("last name"),
		Email: r.FormValue("email"),
		Birthday: time.Now(),
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string {
		"./Web/ui/html/home.page.tmpl",
		"./Web/ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files ...)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	if data.Name != "" && data.LastName != "" && data.Email != "" {
		ctx := context.Background()
		modelJson, _ := json.Marshal(data)
		recordClient.SetRecord(ctx, &record.SetRecordRequest{RecordJson: string(modelJson)})
		infoLog.Printf("Setted " + string(modelJson))
	}
}

func (app *Application) login(w http.ResponseWriter, r *http.Request) {
	data := loginData{
		Login: r.FormValue("login"),
		Password: r.FormValue("password"),
	}

	ctx := context.Background()
	token, err := authClient.Login(ctx, &auth.LoginRequest{Login: data.Login, Password: data.Password})
	if err != nil || token == nil {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	r.Header.Add("token", token.Token)

	files := []string {
		"./Web/ui/html/login.page.tmpl",
		"./Web/ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files ...)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (app *Application) admin(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")

	ctx := context.Background()

	records, err := recordClient.GetRecords(ctx, &record.TokenRequest{Token: token})
	if records == nil || records.IsSuccess == false {
		if err != nil && err.Error() == "forbidden"{
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	//TODO add page
}
