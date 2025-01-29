package db

import (
	"errors"
	"fmt"
	"time"
)

func (p *Plight) HitAdd(hit string) error {
	data, err := p.ReadDB()
	if err != nil {
		return err
	}

	today := time.Now().Format(time.DateOnly)

	if _, e := data.Hits[hit]; !e {
        fmt.Printf("You are creating a new hit, rewrite it's name to confirm\n%v: ",hit)
        var check string
        fmt.Scan(&check)
        if check != hit {
            return (errors.New("hit not created"))
        }
		data.Hits[hit] = make(map[string]int)

	}
	data.Hits[hit][today]++

	err = p.writeDB(data)

	return err
}
