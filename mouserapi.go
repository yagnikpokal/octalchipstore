/*package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
)

type RequestBody struct {
	SearchByPartRequest SearchByPartRequest `json:"SearchByPartRequest"`
}

type SearchByPartRequest struct {
	MouserPartNumber string `json:"mouserPartNumber"`
}

type ResponseData struct {
	Data json.RawMessage
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":8080", nil)
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
	data := ResponseData{Data: body}

	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, data)

}*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RequestBody struct {
	SearchByPartRequest SearchByPartRequest `json:"SearchByPartRequest"`
}

type SearchByPartRequest struct {
	MouserPartNumber string `json:"mouserPartNumber"`
}

func querytompuser(query string) {
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
func main() {
	querytompuser("BQ25890")
}

/*
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/", handleIndex)

	http.ListenAndServe(":8080", r)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
*/
