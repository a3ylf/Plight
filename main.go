package main

import (
	"fmt"
	"log"
	"plight/db"
)

func main() {
    db := db.StartDB()

    db.EnsureDB()
    
    log.Println(db.WriteDB("gaming"))
    
    fmt.Println(db.ReadDB())
   
}
