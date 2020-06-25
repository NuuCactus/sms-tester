package restapi

import (
	"log"
	"net/http"
)

func PostSMS(w http.ResponseWriter, r *http.Request) {

	log.Println("SMS: send")

	w.Write([]byte("OK"))
}
