package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

type Request struct {
	ChatID  string `json:"chatId"`
	Message string `json:"message"`
}

type RequestURL struct {
	ChatID   string `json:"chatId"`
	URL      string `json:"urlFile"`
	FileName string `json:"fileName"`
	Caption  string `json:"caption"`
}

var uRLgetSet string
var caption string
var jsonData []byte

func PostMethod(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Ошибка обработки формы", http.StatusBadRequest)
			return
		}
		action := r.FormValue("action")
		fmt.Println(action)

		phoneNumber := r.FormValue("phonenumber")
		phoneNumberforurl := r.FormValue("phonenumberforurl")
		someText := r.FormValue("sometext")
		urlfile := r.FormValue("url")
		filename := filepath.Base(urlfile)
		for i := 0; i < int(len(filename)); i++ {
			if string(filename[i]) == "." {
				break
			} else {
				caption = caption + filename
			}
		}

		phoneNumber = phoneNumber + "@c.us"
		phoneNumberforurl = phoneNumberforurl + "@c.us"

		switch action {
		case "SendMessage":
			uRLgetSet = "https://" + Numb + "." + "api.green-api.com/" + WaInstance + "/sendMessage/" + APItokenInstance
			payload := Request{
				ChatID:  phoneNumber,
				Message: someText,
			}
			jsonData, err = json.Marshal(payload)

		case "SendByURL":
			uRLgetSet = "https://" + Numb + "." + "api.green-api.com/" + WaInstance + "/sendFileByUrl/" + APItokenInstance
			payload := RequestURL{
				ChatID:   phoneNumberforurl,
				URL:      urlfile,
				FileName: filename,
				Caption:  caption,
			}
			jsonData, err = json.Marshal(payload)
		default:
			http.ServeFile(w, r, "index.html")
		}
		fmt.Println(uRLgetSet)

		if err != nil {
			http.Error(w, "Ошибка создания JSON", http.StatusInternalServerError)
			return
		}
		fmt.Println(string(jsonData))
		res, err := http.Post(uRLgetSet, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Printf("Ошибка при отправке запроса: %v", err)
			http.Error(w, "Ошибка при отправке запроса", http.StatusBadGateway)
			return
		}
		defer res.Body.Close()
		fmt.Println(res)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
