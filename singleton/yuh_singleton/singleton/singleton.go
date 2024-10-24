package singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

func init() {
	instance = &singleton{count: 100}
}

func GetIntance() Singleton {
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
