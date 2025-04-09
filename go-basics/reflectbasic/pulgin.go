package reflectbasic

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"time"
)

type Plugin interface {
	Run() error
}

type PluginImpl1 struct {
}

type PluginImpl2 struct {
}

type PluginImpl3 struct {
}

func (p *PluginImpl1) String() string {
	fmt.Println(reflect.TypeOf(p).Name())
	return reflect.TypeOf(p).Name()
}

func (p *PluginImpl2) String() string {
	fmt.Println(reflect.TypeOf(p).Name())
	return reflect.TypeOf(p).Name()
}

func (p *PluginImpl3) String() string {
	fmt.Println(reflect.TypeOf(p).Name())
	return reflect.TypeOf(p).Name()
}

func (p *PluginImpl1) Run() error {
	fmt.Println("p1 begin run")
	time.Sleep(time.Second)
	fmt.Println("p1 end run")
	return nil
}

func (p *PluginImpl2) Run() error {
	fmt.Println("p2 begin run")
	time.Sleep(time.Microsecond * 400)
	fmt.Println("p2 end run")
	return nil
}
func (p *PluginImpl3) Run() error {
	fmt.Println("p3 begin run")
	time.Sleep(time.Microsecond * 100)
	fmt.Println("p3 end run")
	return nil
}

var plugins map[string]Plugin = make(map[string]Plugin, 0)

func addAndRemovePlugs() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		for {
			for _, p := range plugins {
				p.Run()
			}
		}
	}()
	p1 := &PluginImpl1{}
	plugins[p1.String()] = p1
	time.Sleep(time.Microsecond * 200)
	p2 := &PluginImpl2{}
	plugins[p2.String()] = p2
	p3 := &PluginImpl3{}
	plugins[p3.String()] = p3

	time.Sleep(time.Second * 10)
}

type PluginManager struct {
	mu      sync.RWMutex
	plugins map[string]Plugin
}

func (pm *PluginManager) RegisterPlugin(p Plugin, name string) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	pm.plugins[name] = p
}

func NewPluginManager() *PluginManager {
	return &PluginManager{
		plugins: make(map[string]Plugin, 0),
	}
}

func doPluginPractice() {
	pm := NewPluginManager()

	p1 := &PluginImpl1{}
	pm.RegisterPlugin(p1, p1.String())

}
