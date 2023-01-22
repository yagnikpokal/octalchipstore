package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type RequestBody struct {
	SearchByPartRequest SearchByPartRequest `json:"SearchByPartRequest"`
}

type SearchByPartRequest struct {
	MouserPartNumber string `json:"mouserPartNumber"`
}

func (app *Config) routes() http.Handler {
	// Create the router
	mux := chi.NewRouter()

	// Set up the middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.SessionLoad)
	// Define application routes
	mux.Get("/", app.HomePage)

	mux.Get("/login", app.LoginPage)
	mux.Post("/login", app.PostLoginPage)
	mux.Get("/logout", app.Logout)
	mux.Get("/register", app.RegisterPage)
	mux.Post("/register", app.PostRegisterPage)
	mux.Get("/activate-account", app.RegisterPage)

	mux.Get("/login", app.LoginPage)
	// r := chi.NewRouter()

	// add middleware
	//r.Use(middleware.Logger)

	mux.Get("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		fmt.Println(query)
		fmt.Fprint(w, "Search query: "+query)
		//querytomouser(query)

		switch query {
		case query:
			fmt.Println("Some string is come")
			querytomouser(query)
			mux.Get("/", handleIndex)

		}
	})

	//http.ListenAndServe(":3000", r)
	return mux
}

// This will query to mouser.com for the selected part number
func querytomouser(query string) {
	url := "https://api.mouser.com/api/v1/search/partnumber?apiKey=5ad0a426-d6d8-4eda-acd2-7f9df0b07e50"
	requestBody := RequestBody{SearchByPartRequest: SearchByPartRequest{MouserPartNumber: query}}
	jsonValue, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	url := "https://api.mouser.com/api/v1/search/partnumber?apiKey=5ad0a426-d6d8-4eda-acd2-7f9df0b07e50"
	requestBody := RequestBody{SearchByPartRequest: SearchByPartRequest{MouserPartNumber: "BQ25892"}}
	jsonValue, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	tmpl, _ := template.ParseFiles("index.gohtml")
	tmpl.Execute(w, string(body))
}
