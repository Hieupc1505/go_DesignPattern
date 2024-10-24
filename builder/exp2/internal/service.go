package internal

import "fmt"

type Logger interface {
	Log(...any)
}

type StdLogger struct{}
type FileLogger struct{}

func (StdLogger) Log(...any) {}

func (FileLogger) Log(...any) {}

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}
type SMSNotifier struct{}

func (EmailNotifier) Send(message string) {
	fmt.Println("Send via Email: ", message)
}

func (SMSNotifier) Send(message string) {
	fmt.Println("Send via SMS: ", message)
}

type DataLayer interface {
	Save()
}

type MySqlLayer struct{}
type MongoDbLayer struct{}

func (MySqlLayer) Save()   {}
func (MongoDbLayer) Save() {}

type Uploader interface {
	Upload()
}

type AWS3Uploader struct{}
type GoogleDriveUploader struct{}

func (AWS3Uploader) Upload()        {}
func (GoogleDriveUploader) Upload() {}

type complexService struct {
	name      string
	logger    Logger
	notifier  Notifier
	datalayer DataLayer
	uploader  Uploader
}

func (s *complexService) setName(n string)         { s.name = n }
func (s *complexService) setLogger(l Logger)       { s.logger = l }
func (s *complexService) setNotifier(n Notifier)   { s.notifier = n }
func (s *complexService) setDataLayer(d DataLayer) { s.datalayer = d }
func (s *complexService) setUploader(u Uploader)   { s.uploader = u }

type Service interface {
	DoService()
}

func (s *complexService) DoService() {
	s.logger.Log()
	s.notifier.Send("Login Success")
	s.datalayer.Save()
	s.uploader.Upload()

	fmt.Println("Service finish")
}
