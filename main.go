package main

import (
	"fmt"
	"log"
	"plight/db"
	"time"
)

func main() {
    //Format to save on file
	a := time.Now().Format(time.DateTime)
	fmt.Println(a)
	b := time.Now().Format(time.DateTime)
	fmt.Println(b)

    //Retrieve from file
	fmt.Println()
	c , _ := time.Parse(time.DateTime,a)
    d , _ := time.Parse(time.DateTime,b)
    e := d.Sub(c).String()
	fmt.Println(e)
    x, _ := time.ParseDuration(e)
    fmt.Println(x)

    log.Println(db.StartDB("bro.xd").EnsureDB())
   
}
