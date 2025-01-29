package cmd

import (
	"fmt"
	"log"
	"github.com/a3ylf/plight/db"
	"github.com/a3ylf/plight/flags"
	"github.com/a3ylf/plight/tui"
)

func Start() {
	flags.ParseFlags()

	database, err := db.StartDB()
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
    largs := len(args)
	if largs == 0 {
		fmt.Println("No command used")
		return
	}
	switch args[0] {
	case "sh", "show":
		data, err := database.GetData()
		if err != nil {
			log.Println(err)
			return
		}
		if flags.Raw {
			sess, err := database.GetData()
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println(sess)
			return
		}
		tui.StartTui(data)

	case "s", "session":
		l := len(args)
		if l == 2 {
			err = database.SessionAdd(args[1])
			if err != nil {
				fmt.Println(err)
			} 
			return
			

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
