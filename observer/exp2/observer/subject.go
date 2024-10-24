package observer

type Subject interface {
	Register(o Observer)
	deregister(o Observer)
	notifyAll()
}
