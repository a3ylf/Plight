package cfg

import (
	"fmt"
	"log"
	"plight/db"
	"plight/flags"
	"time"
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
	case "a":
		if len(args) == 2 {
			err = db.SessionAdd(args[1])
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Time added to %v, %v\n", args[1], time.Now().Format(time.TimeOnly))
			}

		} else {
			fmt.Println("Use: plight a (session name)")
		}
    case "h":
        if len(args) == 2 {
            err = db.HitAdd(args[1])
            if err != nil {
                fmt.Println(err)
            }else {
                fmt.Printf("Hit added to %v\n", args[1])
            }
        }
	}

}
