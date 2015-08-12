package main

import (
	"fmt"
	"sync"
	"time"
)

var _self *Singleton

type Singleton struct {
	Name string
	sync.Mutex
}

func NewInstance(name string) *Singleton {
	fmt.Println("Create instance", name)
	time.Sleep(4 * time.Second)
	_self.Name = name
	return _self
}

func Instance(name string) *Singleton {
	if _self.Name == "" {
		_self.Mutex.Lock()
		defer _self.Mutex.Unlock()
		if _self.Name == "" {
			return NewInstance(name)
		}
	}
	return _self
}

func main() {
	_self = new(Singleton)
	_self.Mutex = sync.Mutex{}
	go Instance("cyeam")
	go Instance("bryce")
	time.Sleep(10 * time.Second)
	fmt.Println(_self.Name)
}
