package problem

import (
	"fmt"
)

type NotificationSevervice struct {
	notifierType string
}

func (s NotificationSevervice) SendNotification(message string) {
	if s.notifierType == "email" {
		fmt.Println("Send via Email: ", message)
	} else if s.notifierType == "sms" {
		fmt.Println("Send via SMS: ", message)
	}
}

func main() {
	s := NotificationSevervice{notifierType: "email"}
	s.SendNotification("Hello World")
}
