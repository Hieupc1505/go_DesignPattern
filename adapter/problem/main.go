package problem

import "fmt"

type Handler func() error

func execFunc(f Handler) error {
	return f()
}

type MessageHandler func(message string) error

func main() {
	f := func() error {
		fmt.Println("Hello world")
		return nil
	}

	execFunc(f) //OK

	//But, how I can use MessageHandler with execFunc?
	//wihtout modifying execFunc
}
