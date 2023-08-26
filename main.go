package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/culqi/culqi-go"
	"github.com/go-chi/chi"
	"github.com/rs/cors"

	config "culqi-go-demo/config"
)

func main() {
	r := chi.NewRouter()
	mime.AddExtensionType(".js", "application/javascript; charset=utf-8")

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		//enableCors(&w)
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
	r.Post("/culqi/generateCard", cardsPageHandler)
	r.Post("/culqi/generateCustomer", customerPageHandler)
	r.Post("/culqi/generateCharge", chargePageHandler)
	r.Post("/culqi/generateOrder", orderPageHandler)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT"},
	})
	handler := c.Handler(r)
	http.ListenAndServe(config.Puerto, handler)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("index-card.html")
	template.Execute(w, nil)
}

func homePage2Handler(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("index.html")
	template.Execute(w, nil)
}

// Consumo de servicios
func cardsPageHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	//culqi.Key(pk, sk)
	if config.Encrypt == "1" {
		statusCode, res, _ := culqi.CreateCard(reqBody, config.EncryptionData...)
		fmt.Println(statusCode)
		fmt.Println(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write([]byte(res))
	} else {
		statusCode, res, _ := culqi.CreateCard(reqBody)
		fmt.Println("Resultados")
		fmt.Println(statusCode)
		fmt.Println(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write([]byte(res))
	}
}
func chargePageHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	log.Printf("error decoding sakura response: %v", reqBody)

	//culqi.Key(pk, sk)

	if config.Encrypt == "1" {
		statusCode, res, _ := culqi.CreateCharge(reqBody, config.EncryptionData...)
		fmt.Println(statusCode)
		fmt.Println(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write([]byte(res))
	} else {
		statusCode, res, _ := culqi.CreateCharge(reqBody)
		fmt.Println("Resultados")
		fmt.Println(statusCode)
		fmt.Println(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write([]byte(res))
	}

}
func customerPageHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Body)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(reqBody)
	fmt.Println(bodyString)
	log.Printf("error decoding sakura response: %v", reqBody)

	//culqi.Key(pk, sk)

	statusCode, res, err := culqi.CreateCustomer(reqBody)
	fmt.Println(err)
	fmt.Println("statusCode")
	fmt.Println(statusCode)
	fmt.Println(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(res))
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

	//culqi.Key(pk, sk)

	statusCode, res, err := culqi.CreateOrder(reqBody, config.EncryptionData...)
	fmt.Println(err)
	fmt.Println("statusCode")
	fmt.Println(statusCode)
	fmt.Println(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(res))
}
