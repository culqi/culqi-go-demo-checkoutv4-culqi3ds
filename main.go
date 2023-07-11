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
	"strconv"
	"text/template"

	"github.com/culqi/culqi-go"
	"github.com/go-chi/chi"
)

var pk string = "INSERTE_SU_LLAVE_PUBLICA"
var sk string = "INSERTE_SU_LLAVE_PRIVADA"
var puerto string = ":3000"
var encrypt = "1"
var encryptiondData = []byte(`{		
	"rsa_public_key": "INSERTE_SU_LLAVE_RSA",
	"rsa_id": "INSERTE_SU_RSA_ID"
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
	r.Post("/culqi/generateOrder", orderPageHandler)
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

	culqi.Key(pk, sk)
	statusCode, res, err := culqi.CreateCard(reqBody, encryptiondData...)
	fmt.Println(statusCode)
	fmt.Println(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
func chargePageHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	log.Printf("error decoding sakura response: %v", reqBody)

	culqi.Key(pk, sk)

	var (
		statusCode string
		res        string
	)

	if encrypt == "1" {
		statusCode, res, _ := culqi.CreateCharge(reqBody, encryptiondData...)
		fmt.Println(statusCode)
		fmt.Println(res)
	} else {
		statusCode, res, _ := culqi.CreateCharge(reqBody)
		fmt.Println(statusCode)
		fmt.Println(res)
	}
	w.Header().Set("Content-Type", "application/json")
	code, _ := strconv.Atoi(statusCode)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(res)

	/*
		if encrypt == "1" {
			statusCode, res, err := culqi.CreateCharge(reqBody, encryptiondData...)
			fmt.Println(err)
			fmt.Println(res)
			w.Header().Set("Content-Type", "application/json")
			if strings.Contains(res, "REVIEW") {
				w.WriteHeader(statusCode)
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
	*/
}
func customerPageHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var post Customers
	json.Unmarshal(reqBody, &post)

	culqi.Key(pk, sk)

	statusCode, res, err := culqi.CreateCustomer(reqBody)
	fmt.Println(statusCode)
	fmt.Println(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func orderPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(reqBody)
	fmt.Println(bodyString)
	log.Printf("error decoding sakura response: %v", reqBody)

	culqi.Key(pk, sk)

	statusCode, res, err := culqi.CreateOrder(reqBody)
	fmt.Println(err)
	fmt.Println("statusCode")
	fmt.Println(statusCode)
	fmt.Println(res)

	w.Header().Set("Content-Type", "application/json")
	//code, _ := strconv.Atoi(statusCode)
	w.WriteHeader(statusCode)
	w.Write([]byte(res)) /*
		jsonData, err := json.Marshal(res)
		w.Write(jsonData)
		fmt.Println(jsonData)*/
	//json.NewDecoder(w).Decode(res)
	//json.NewEncoder(w).Encode(res)

}
