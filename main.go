package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TheCreeper/go-notify"
)

type NotifyRequest struct {
	AppIcon string `json:"appicon,omitempty"`
	Summary string `json:"summary"`
	Body    string `json:"body"`
	Timeout int32  `json:"timeout,omitempty"`
}

func main() {
	http.HandleFunc("/notify", func(w http.ResponseWriter, r *http.Request) {
		nr := NotifyRequest{}
		err := json.NewDecoder(r.Body).Decode(&nr)
		if err != nil {
			log.Fatal("Notification request failed")
		}

		n := notify.Notification{
			AppName: "curst",
			AppIcon: "/tmp/img.jpg",
			Summary: nr.Summary,
			Body:    nr.Body,
			Timeout: 5000,
		}

		if _, err := n.Show(); err != nil {
			log.Print("Error showing notification")
		}
	})

	err := http.ListenAndServe(":4590", nil)
	if err != nil {
		log.Fatal("Http listen failed")
	}
}
