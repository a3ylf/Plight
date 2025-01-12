package cfg

import (
	"fmt"
	"log"
	"os"
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
		os.Exit(1)
	}

	args := flags.ParseArgs()
	switch args[0] {
	case "a":
		if len(args) == 2 {
			err = db.WriteDB(args[1])
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Tempo adicionado a sessão %v, %v\n", args[1], time.Now().Format(time.TimeOnly))
			}

		} else {
			fmt.Println("Use: plight a (session name)")
		}
	}

}
