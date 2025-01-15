package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/a3ylf/plight/db"
	"github.com/a3ylf/plight/flags"
)

func Start() {
	flags.ParseFlags()

	db, err := db.StartDB()

	if err != nil {
		log.Println(err)
	}

	if flags.Reset {
		err = db.ResetDB()
		if err != nil {
			log.Println(err)
		}
		return
	}

	args := flags.ParseArgs()
	switch args[0] {
	case "s","session":
		l := len(args)
		if l > 1 {
			a1 := args[1]
			if a1 == "show" {
				if l == 2 {
					sess, err := db.GetSessions()
					if err != nil {
						log.Println(err)
						return
					}
					fmt.Println(sess)
					return
				} else if l == 3 {
					sess, err := db.GetSession(args[2])
					if err != nil {
						log.Println(err)
						return
					}
					fmt.Println(sess)
				} else {
					fmt.Println("Too many arguments")
					return
				}
			}

			err = db.SessionAdd(args[1])
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Time added to %v, %v\n", args[1], time.Now().Format(time.TimeOnly))
			}

		} else {
			fmt.Println("Use: plight s (session name)")
		}
	case "h","hit":
		if len(args) == 2 {
			err = db.HitAdd(args[1])
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Hit added to %v\n", args[1])
			}
		} else {
			fmt.Println("Use plight h (hit name)")
		}
	}

}
