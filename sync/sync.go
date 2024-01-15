package main

import "sync"

type Contador struct {
	sync.Mutex
	valor int
}

func (c *Contador) Incrementa() {
	c.Lock()
	defer c.Unlock()
	c.valor++
}

func (c *Contador) Valor() int {
	return c.valor
}