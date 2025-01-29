package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/a3ylf/plight/db"
	"github.com/a3ylf/plight/flags"
	"github.com/a3ylf/plight/tui"
)

func Start() {
	flags.ParseFlags()

	database, err := db.StartDB()
    xd, err := database.GetSessions()
    var sessions db.Sessions
	json.Unmarshal(xd,&sessions)
	tui.StartTui(sessions)
	os.Exit(1)
	if err != nil {
		log.Println(err)
	}

	if flags.Reset {
		err = database.ResetDB()
		if err != nil {
			log.Println(err)
		}
		return
	}

	args := flags.ParseArgs()
	switch args[0] {
	case "show", "sh":

	case "s", "session":
		l := len(args)
		if l > 1 {
			a1 := args[1]
			if a1 == "show" {
				if l == 2 {
					if flags.Raw {
						sess, err := database.GetSessions()
						if err != nil {
							log.Println(err)
							return
						}
						fmt.Println(string(sess))
						return
					}
				} else if l == 3 {
					if flags.Raw {
						sess, err := database.GetSession(args[2])
						if err != nil {
							log.Println(err)
							return
						}
						fmt.Println(string(sess))
						return
					}
				} else {
					fmt.Println("Too many arguments")
					return
				}
			}

			err = database.SessionAdd(args[1])
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Time added to %v, %v\n", args[1], time.Now().Format(time.TimeOnly))
			}

		} else {
			fmt.Println("Use: plight s (session name)")
		}
	case "h", "hit":
		if len(args) == 2 {
			err = database.HitAdd(args[1])
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
