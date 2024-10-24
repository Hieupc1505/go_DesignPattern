package singleton_v2

import "sync"

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var (
	instance *singleton
	once     sync.Once
)

func init() {
	instance = &singleton{count: 100}
}

func GetIntance() Singleton {
	if instance == nil {
		once.Do(func() {
			instance = &singleton{count: 100}
		})
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
