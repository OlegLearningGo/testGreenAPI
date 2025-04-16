package handler

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	URLgetSet, IdInstance, Numb, APItokenInstance, WaInstance string
)

func GetMethod(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing", http.StatusBadRequest)
			return
		}
		IdInstance = r.FormValue("IdInstance")
		Numb = IdInstance[0:4]
		APItokenInstance = r.FormValue("APItokenInstance")
		WaInstance = "waInstance" + IdInstance
		action := r.URL.Query().Get("action")
		switch action {
		case "getSettings":
			URLgetSet = "https://" + Numb + "." + "api.green-api.com/" + WaInstance + "/getSettings/" + APItokenInstance
		case "getStateInstance":
			URLgetSet = "https://" + Numb + "." + "api.green-api.com/" + WaInstance + "/getStateInstance/" + APItokenInstance
		default:
			http.ServeFile(w, r, "index.html")
		}
		fmt.Println(URLgetSet)
		res, err := http.Get(URLgetSet)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
		body, error := ioutil.ReadAll(res.Body)
		if error != nil {
			fmt.Println(error)
		}
		res.Body.Close()
		fmt.Println(string(body))

		tmpl := template.Must(template.ParseFiles("index.html"))

		Data := struct {
			ResponseBody string
		}{
			ResponseBody: string(body),
		}
		tmpl.Execute(w, Data)
		return

	}
}
