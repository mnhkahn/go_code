package main

import (
	"fmt"
	"sync"
	"time"
)

var _self *Singleton

type Singleton struct {
	Name string
	sync.Once
}

func NewInstance(name string) *Singleton {
	fmt.Println("Create instance", name)
	time.Sleep(4 * time.Second)
	_self.Name = name
	return _self
}

func Instance(name string) *Singleton {
	if _self.Name == "" {
		_self.Once.Do(func() { NewInstance(name) })
	}
	return _self
}

func main() {
	_self = new(Singleton)
	_self.Once = sync.Once{}
	go Instance("cyeam")
	go Instance("bryce")
	time.Sleep(10 * time.Second)
	fmt.Println(_self.Name)
}
