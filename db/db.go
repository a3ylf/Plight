package db

import (
	"encoding/json"
	"fmt"
	"os"
	"plight/flags"
	"sync"
	"time"
)

type Plight struct {
	mux      sync.RWMutex
	filename string
}

func StartDB(filename ...string) (*Plight, error) {
	name := "data.json"
	if len(filename) > 0 {
		name = filename[0]
	}
	if flags.Dev {
		name = "debug.json"
	}
	db := &Plight{
		mux:      sync.RWMutex{},
		filename: name,
	}
	err := db.EnsureDB()
	return db, err
}

func (p *Plight) ResetDB() error {
	err := os.Remove(p.filename)
	if err == nil {
		err = p.EnsureDB()
	}
	return err
}
func (p *Plight) EnsureDB() error {

	_, err := os.Open(p.filename)

	if err != nil {

		err = createDB(p.filename)
	}

	return err
}

func createDB(filename string) error {
	data := &Data{
		Sessions: make(map[string]Timers),
	}
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, b, 0644)
	return err
}

func (p *Plight) ReadDB() (*Data, error) {
	b, err := os.ReadFile(p.filename)

	if err != nil {
		return &Data{}, err
	}

	data := &Data{}

	err = json.Unmarshal(b, data)
	if err != nil {
		return &Data{}, err
	}
	return data, nil
}

// //Format to save on file
// // a := time.Now().Format(time.DateTime)
// // fmt.Println(a)
// // b := time.Now().Format(time.DateTime)
// // fmt.Println(b)
// //
// //Retrieve from file
// fmt.Println()
// c , _ := time.Parse(time.DateTime,a)
//    d , _ := time.Parse(time.DateTime,b)
//    e := d.Sub(c).String()
// fmt.Println(e)
//    x, _ := time.ParseDuration(e)
//    fmt.Println(x)

func (p *Plight) WriteDB(session string) error {

	data, err := p.ReadDB()

	if err != nil {
		return err
	}
	timenow := fmt.Sprint(time.Now().Date())

	if data.Sessions == nil {
		data.Sessions = make(map[string]Timers)
	}
	if _, e := data.Sessions[session]; !e {
		data.Sessions[session] = Timers{
			Days: make(map[string]Day),

			// xd
		}
	}
	last := len(data.Sessions[session].Days[timenow].Periods) - 1

	if last == -1 {
		a := Day{
			Day_Total: "0s",
			Periods: []Period{
				Period{
					From: time.Now().Format(time.TimeOnly)},
			},
		}
		data.Sessions[session].Days[timenow] = a

	} else if data.Sessions[session].Days[timenow].Periods[last].To == "" {
		now := time.Now().Format(time.TimeOnly)
		data.Sessions[session].Days[timenow].Periods[last].To = now
		dur, err := time.ParseDuration(data.Sessions[session].Days[timenow].Day_Total)
		from, err := time.Parse(time.TimeOnly, data.Sessions[session].Days[timenow].Periods[last].From)
		to, err := time.Parse(time.TimeOnly, data.Sessions[session].Days[timenow].Periods[last].To)

		if err != nil {
			return err
		}
		newTotal := to.
			Add(dur).
			Sub(from).
			String()

		s := data.Sessions[session].Days[timenow]
		s.Day_Total = newTotal

		data.Sessions[session].Days[timenow] = s

	} else {
		a := data.Sessions[session].Days[timenow]
		a.Periods = append(a.Periods, Period{
			Id:   last + 1,
			From: time.Now().Format(time.TimeOnly)})
		data.Sessions[session].Days[timenow] = a
	}
	d, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		return err
	}

	p.mux.Lock()
	err = os.WriteFile(p.filename, d, 0644)
	p.mux.Unlock()

	if err != nil {
		return err
	}
	return nil
}
