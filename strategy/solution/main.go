package solution

import "fmt"

type Notifier interface {
	SendNotification(message string)
}

type SendEmailNotify struct{}

func (s *SendEmailNotify) SendNotification(message string) {
	fmt.Println("Send via Email: ", message)
}

type SendSmsNotify struct{}

func (s *SendSmsNotify) SendNotification(message string) {
	fmt.Println("Send via SMS: ", message)
}

type NotificationSevervice struct {
	notifier Notifier
}

func main() {
	s := NotificationSevervice{
		notifier: &SendEmailNotify{},
	}
	s.notifier.SendNotification("Hello World")
}
