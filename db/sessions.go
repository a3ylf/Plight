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

func (p *Plight) GetSessions() ([]byte, error) {

	data, err := p.ReadDB()

	if err != nil {
		return []byte{}, err
	}
	d, err := json.MarshalIndent(data.Sessions, "", "   ")
	return d, err
}
func (p *Plight) SessionAdd(session string) error {

	data, err := p.ReadDB()

	if err != nil {
		return err
	}
	timenow := fmt.Sprint(time.Now().Date())

	if data.Sessions == nil {
		data.Sessions = make(map[string]Days)
	}
	if _, e := data.Sessions[session]; !e {
		data.Sessions[session] = Days{}
	}
	last := len(data.Sessions[session][timenow].Periods) - 1

	if last == -1 {
		a := Day{
			Day_Total: "0s",
			Periods: []Period{
				{
					From: time.Now().Format(time.TimeOnly)},
			},
		}
		data.Sessions[session][timenow] = a

	} else if data.Sessions[session][timenow].Periods[last].To == "" {
		now := time.Now().Format(time.TimeOnly)
		data.Sessions[session][timenow].Periods[last].To = now
		dur, err := time.ParseDuration(data.Sessions[session][timenow].Day_Total)
		from, err := time.Parse(time.TimeOnly, data.Sessions[session][timenow].Periods[last].From)
		to, err := time.Parse(time.TimeOnly, data.Sessions[session][timenow].Periods[last].To)

		if err != nil {
			return err
		}
		newTotal := to.
			Add(dur).
			Sub(from).
			String()

		s := data.Sessions[session][timenow]
		s.Day_Total = newTotal

		data.Sessions[session][timenow] = s

	} else {
		a := data.Sessions[session][timenow]
		a.Periods = append(a.Periods, Period{
			Id:   last + 1,
			From: time.Now().Format(time.TimeOnly)})
		data.Sessions[session][timenow] = a
	}

	err = p.writeDB(data)

	return err
}
