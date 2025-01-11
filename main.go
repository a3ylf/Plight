package main

import (
	"plight/db"
)

func main() {
    db := db.StartDB()

    db.EnsureDB()
    
    db.WriteDB("gaming")
    
   
}
