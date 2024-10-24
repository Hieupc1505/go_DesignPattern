package main

import "fmt"

type Observer interface {
	ReceivedJob(job Job)
}

type Developer struct{}

func (c Developer) ReceivedJob(j Job) {
	fmt.Printf("Woooww!, have a job for you: %s\n", j.Title)
}

type Job struct {
	Title string
}

type ITJobCompany struct {
	observers []Observer
}

func (c *ITJobCompany) Subcriber(o Observer) {
	c.observers = append(c.observers, o)
}

func (c *ITJobCompany) UnSubcriber(o Observer) {

	for i, dev := range c.observers {
		if dev == o {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
			return
		}
	}
}

func (c *ITJobCompany) Notify(j Job) {
	for _, dev := range c.observers {
		dev.ReceivedJob(j)
	}
}

func (c *ITJobCompany) AddNewJob(j Job) {
	c.Notify(j)
}

func main() {
	dev := Developer{}
	dev2 := Developer{}

	comp := ITJobCompany{}

	comp.Subcriber(dev)
	comp.Subcriber(dev2)

	job := Job{
		Title: "Backend NOdejs",
	}
	comp.AddNewJob(job)

	comp.UnSubcriber(dev)

	job2 := Job{
		Title: "Fontend Reactjs",
	}

	comp.AddNewJob(job2)
}
