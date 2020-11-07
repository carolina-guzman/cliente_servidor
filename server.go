package main

import (
	"container/list"
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

var lista list.List
var posicion uint64

type Proceso struct {
	Id     uint64
	Status bool
	Tiempo uint64
}

func (p *Proceso) start() {
	for {
		if p.Status == true {
			p.Tiempo++
			p.print()
		}
	}
}

func (p *Proceso) print() {
	for {
		fmt.Printf("id %d: %d", p.Id, p.Tiempo)
		fmt.Println()
		time.Sleep(time.Millisecond * 500)
	}
}

func servidor() {
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleClient(c)
	}
}

func handleClient(c net.Conn) { //cuando el cliente arranque
	//enviar
	variable := lista.Remove(lista.Front())
	variable.(*Proceso).Status = false
	err := gob.NewEncoder(c).Encode(variable)
	if err != nil {
		fmt.Println(err)
	}
	//recibir
	var pro Proceso
	err = gob.NewDecoder(c).Decode(&pro)
	if err != nil {
		fmt.Println(err)
		return
	} 
	pro.Status = true
	lista.PushBack(&pro)
	go pro.start()

	c.Close()
}

func crearProcesos() {
	var i uint64
	for i = 0; i < 5; i++ {
		p := Proceso{Id: i+1, Status: true, Tiempo: 0}
		lista.PushBack(&p)
		go p.start()
	}
}

func main() {
	posicion = 0
	crearProcesos()
	go servidor()

	var input string
	fmt.Scanln(&input)
}
