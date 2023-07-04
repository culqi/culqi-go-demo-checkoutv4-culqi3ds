package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/culqi/culqi-go"
	"github.com/go-chi/chi"
)

var pk string = "pk_test_90667d0a57d45c48"
var sk string = "sk_test_1573b0e8079863ff"
var puerto string = ":3000"
var encrypt = "1"
var encryptiondData = []byte(`{		
	"rsa_public_key": "-----BEGIN PUBLIC KEY-----
	MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDYR6Oqz+vX2amSnNzPosH1CIMocGnHCnxlr1RuRyYtrAAVv3oxpSx42R9KIbW3yBfWwFxpU9m1us1ZjPmISRmjy64z6q6rv5UZNOWllM5v2A+F2MceWHRIJYOxIwV9oAx36EH89qOEnOekVLqZhkdrAx2LvLfqGprKsDcfX06urwIDAQAB
-----END PUBLIC KEY-----",
	"rsa_id": "f355d27f-e735-46a7-b8bd-9773357ff034"
}`)

func main() {
	r := chi.NewRouter()
	mime.AddExtensionType(".js", "application/javascript; charset=utf-8")

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		workDir, _ := os.Getwd()
		filesDir := filepath.Join(workDir, "js")
		if _, err := os.Stat(filesDir + r.URL.Path); errors.Is(err, os.ErrNotExist) {
			http.ServeFile(w, r, filepath.Join(filesDir, "index-card.html"))
		}
		http.ServeFile(w, r, filesDir+r.URL.Path)
	})
	//http.HandleFunc("/index-card", homePageHandler)
	r.Get("/index-card", homePageHandler)
	r.Get("/", homePage2Handler)
	r.Post("/culqi/generateCards", cardsPageHandler)
	r.Post("/culqi/generateCustomer", customerPageHandler)
	r.Post("/culqi/generateCharge", chargePageHandler)
	r.Get("/admin", adminPageHandler)
	http.ListenAndServe(puerto, r)
}

type Customers struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	AddressCity string `json:"address_city"`
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("index-card.html")
	template.Execute(w, nil)
}

func homePage2Handler(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("index.html")
	template.Execute(w, nil)
}

func adminPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is admin page"))
}
func cardsPageHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	culqi.Key(sk)
	res, err := culqi.CreateCard(reqBody)
	fmt.Println(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
func chargePageHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	log.Printf("error decoding sakura response: %v", reqBody)

	culqi.Key(sk)

	if encrypt == "1" {
		res, err := culqi.CreateCharge(reqBody, encryptiondData)
		fmt.Println(err)
		fmt.Println(res)
		w.Header().Set("Content-Type", "application/json")
		if strings.Contains(res, "REVIEW") {
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusOK)
		}

		json.NewEncoder(w).Encode(res)
	} else {
		res, err := culqi.CreateCharge(reqBody, nil)
		fmt.Println(err)
		fmt.Println(res)
		w.Header().Set("Content-Type", "application/json")
		if strings.Contains(res, "REVIEW") {
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusOK)
		}

		json.NewEncoder(w).Encode(res)
	}

}
func customerPageHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var post Customers
	json.Unmarshal(reqBody, &post)

	culqi.Key(sk)

	res, err := culqi.CreateCustomer(reqBody)
	fmt.Println(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
