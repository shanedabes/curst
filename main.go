package main

import (
	"log"

	notify "github.com/TheCreeper/go-notify"
)

func main() {
	notification := notify.Notification{
		AppName: "curst",
		AppIcon: "/tmp/img.jpg",
		Summary: "Test notification",
		Body:    "Hello I am a notification",
		Timeout: 5000,
	}

	if _, err := notification.Show(); err != nil {
		log.Fatal("Notification failed")
	}
}
