package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

var pro Proceso
var conexion net.Conn


type Proceso struct {
	Id     uint64
	Tiempo uint64
}

func (p *Proceso) start() {
	for {
		p.Tiempo++
		p.print()
	}
}

func (p *Proceso) print() {
	fmt.Printf("id %d: %d", p.Id, p.Tiempo)
	fmt.Println()
	time.Sleep(time.Millisecond * 800)
}

func cliente() {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	conexion = c
	err = gob.NewDecoder(c).Decode(&pro)
	if err != nil {
		fmt.Println(err)
		return
	} 
	go pro.start()
}


func main() {
	go cliente()

	var input string
	fmt.Scanln(&input)

	err := gob.NewEncoder(conexion).Encode(pro)
	if err != nil {
		fmt.Println(err)
	}
}
