package main

import (
	"fmt"
	"net/http"
	"os"

	"luxury.maxg95/internal/response"
	"luxury.maxg95/internal/validator"
)

var recipient = os.Getenv("RECIPIENT")

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	err := response.Page(w, http.StatusOK, data, "pages/home.tmpl")
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) insertRequest(w http.ResponseWriter, r *http.Request) {
	_, err := app.db.Exec(`
        CREATE TABLE IF NOT EXISTS requests (
            number TEXT
        )
    `)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = r.ParseForm()
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	request := r.PostForm.Get("request")

	if !validator.IsValidNumber(request) {
		app.badRequest(w, r, fmt.Errorf("Invalid input."))
		return
	}

	_, err = app.db.Exec("INSERT INTO requests (number) VALUES ($1)", request)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := map[string]interface{}{"Number": request}
	err = app.sendEmail(w, r, recipient, data, "mail.tmpl")
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) sendEmail(w http.ResponseWriter, r *http.Request, recipient string, data map[string]interface{}, templateName string) error {
	err := app.mailer.Send(recipient, data, templateName)
	if err != nil {
		return err
	}
	return nil
}
