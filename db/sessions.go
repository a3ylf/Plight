package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func (p *Plight) GetSession(session string) ([]byte, error) {

	data, err := p.ReadDB()

	if err != nil {
		return []byte{}, err
	}
	t, e := data.Sessions[session]

	if !e {
		return []byte{}, errors.New("Unable to find this session")
	}
	send, err := json.MarshalIndent(t, "", "   ")
	return send, nil
}

func (p *Plight) GetData() (*Data, error) {

	data, err := p.ReadDB()

	if err != nil {
		return &Data{}, err
	}
	return data, err
}
func (p *Plight) SessionAdd(session string) error {

	data, err := p.ReadDB()

	if err != nil {
		return err
	}
	daynow := fmt.Sprint(time.Now().Date())
	timenow := time.Now().Format(time.TimeOnly)

	if data.Sessions == nil {
		data.Sessions = make(map[string]Days)
	}
	if _, e := data.Sessions[session]; !e {
		data.Sessions[session] = Days{}
	}
	last := len(data.Sessions[session][daynow].Periods) - 1

	if last == -1 {
		day := Day{
			Day_Total: "0s",
			Periods: []Period{
				{
					From: timenow,
				},
			},
		}
		data.Sessions[session][daynow] = day
		fmt.Printf("Time added to %v, %v\n", session, timenow)
		// forgive me for this
	} else if data.Sessions[session][daynow].Periods[last].To == "" {
		data.Sessions[session][daynow].Periods[last].To = timenow
		dur, err := time.ParseDuration(data.Sessions[session][daynow].Day_Total)
		from, err := time.Parse(time.TimeOnly, data.Sessions[session][daynow].Periods[last].From)
		to, err := time.Parse(time.TimeOnly, data.Sessions[session][daynow].Periods[last].To)

		if err != nil {
			return err
		}
		dursess := to.Sub(from)
		newTotal := to.
			Add(dur).
			Sub(from).
			String()

		s := data.Sessions[session][daynow]
		s.Day_Total = newTotal

		data.Sessions[session][daynow] = s
        fmt.Printf("Session ended\nSession duration: %v\nNew total duration: %v\n", dursess, newTotal)

	} else {
		a := data.Sessions[session][daynow]
		a.Periods = append(a.Periods, Period{
			Id:   last + 1,
			From: timenow,})
		data.Sessions[session][daynow] = a
		fmt.Printf("Time added to %v, %v\n", session, timenow)
	}

	err = p.writeDB(data)

	return err
}
