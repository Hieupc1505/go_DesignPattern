package solution

import "fmt"

type Handler func() error

func execFunc(f Handler) error {
	return f()
}

type MessageHandler func(message string) error

func handlerAdapter(mhdl MessageHandler, msg string) Handler {
	return func() error {
		return mhdl(msg)
	}
}

func main() {
	f := func() error {
		fmt.Println("Hello world")
		return nil
	}

	execFunc(f) //OK

	msgHdl := func(msg string) error {
		fmt.Println(msg)
		return nil
	}

	f = handlerAdapter(msgHdl, "Hello world with Adapter")

	execFunc(f)

}
