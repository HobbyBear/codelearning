package main

import (
	"fmt"
	"time"
)

// 实现一个协程池,设计一个只能有那么多个协程在运行

type pool struct {
	tasks    chan func()
	idleTime time.Duration
}

func (p *pool) Submit(f func()) {
	p.tasks <- f
}

func NewPool(num int) *pool {
	p := &pool{
		tasks:    make(chan func(), num),
		idleTime: time.Second * 10,
	}
	p.start()
	return p
}

func (p *pool) start() {
	go func() {
		for task := range p.tasks {
			go func() {
				task()
				for {
					select {
					case task := <-p.tasks:
						fmt.Println("沿用协程")
						task()
					case <-time.After(p.idleTime):
						fmt.Println("协程销毁")
						return
					}
				}
			}()
		}
	}()
}

func main() {
	po := NewPool(1)
	po.Submit(func() {
		fmt.Println("执行了")
		time.Sleep(time.Second)
	})
	for i := 0; i < 1000; i++ {
		po.Submit(func() {
			fmt.Println("执行了", i)
		})
	}

	time.Sleep(time.Hour)
}
