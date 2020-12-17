package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"

	notify "github.com/TheCreeper/go-notify"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	Port     uint16 `short:"p" long:"port" description:"port to listen on" env:"CURST_PORT" default:"4950"`
	Timeout  int32  `short:"t" long:"timeout" description:"Default notification timeout" env:"CURST_TIMEOUT" default:"5000"`
	IconPath string `short:"i" long:"icon_path" description:"Notification icons path" env:"CURST_ICON_PATH" default:"~/.local/share/curst/icons"`
}

type NotifyRequest struct {
	AppIcon string `json:"appicon,omitempty"`
	Summary string `json:"summary,omitempty"`
	Body    string `json:"body,omitempty"`
	Icon    string `json:"icon,omitempty"`
	Timeout int32  `json:"timeout,omitempty"`
}

func defInt(x, y int32) int32 {
	if x > 0 {
		return x
	}
	return y
}

func main() {
	opts := Options{}

	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/notify", func(w http.ResponseWriter, r *http.Request) {
		nr := NotifyRequest{}
		err := json.NewDecoder(r.Body).Decode(&nr)
		if err != nil {
			log.Print("Notification request failed")

			return
		}

		n := notify.Notification{
			AppName: "curst",
			Summary: nr.Summary,
			Body:    nr.Body,
			AppIcon: path.Join(opts.IconPath, nr.Icon),
			Timeout: defInt(nr.Timeout, opts.Timeout),
		}

		if _, err := n.Show(); err != nil {
			log.Print("Error showing notification")
		}
	})

	err = http.ListenAndServe(fmt.Sprintf(":%d", opts.Port), nil)
	if err != nil {
		log.Fatal("Http listen failed")
	}
}
