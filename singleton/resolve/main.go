package resolve

import (
	"log"
	"sync"
	"time"
)

type config struct {
	logAllowed bool
}

func (c config) LogAllowed() bool {
	return c.logAllowed
}

func NewConfig(allowed bool) config {
	return config{logAllowed: allowed}
}

// This is all about Singleton Pattern
// the rest of the code is just an example
var singletonApp = &application{
	once: sync.Once{},
}

func GetApplication() *application {
	return singletonApp
}

type application struct {
	once sync.Once
	cfg  *config
}

/*
*

	func (app *application) GetConfig() *config {
		if app.cfg == nil {
			log.Println("It should be run only once")
			app.loadConfig()
		}
		return app.cfg
	}

	Nếu chỉ if đơn giản như trên khi chạy trong goroutine thì rất dễ bị
	data racing, do là khi trong quá trình xử lý thì đã có gorotine vào và vẫn dùng
	app.cfg == nil
*/
func (app *application) GetConfig() *config {
	if app.cfg == nil {
		app.once.Do(func() {
			log.Println("It should be run only once")
			app.loadConfig()
		})
	}
	return app.cfg
}

func (app *application) loadConfig() {
	time.Sleep(100) //sử lý logic ví dụ kết nối db
	app.cfg = &config{
		logAllowed: true,
	}
}

func main() {
	//Demo 1000 request to service at a same time
	//I made this code for simple demo, not a real practice

	rps := 1000
	wg := sync.WaitGroup{}
	wg.Add(rps)

	for i := 0; i < rps; i++ {
		go func(idx int) {
			requestHandler(idx)
			if GetApplication().GetConfig().LogAllowed() {
				log.Printf("Request %d handled successfully\n", idx)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func requestHandler(requestIdx int) {
	//Yes, Now i can get the config

	if GetApplication().GetConfig().LogAllowed() {
		log.Printf("Accessing config %d\n", requestIdx)
	}
}
