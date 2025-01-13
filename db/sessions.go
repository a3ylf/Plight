package db

import (
	"errors"
	"fmt"
	"time"
)

func (p *Plight) GetSession(session string) (Timers, error) {

    data, err := p.ReadDB()
    
    if err != nil {
        return Timers{},err
    }
    if t , e := data.Sessions[session]; e {
        return t,nil

    }
    return Timers{}, errors.New("Unable to find this session")
}

func (p *Plight) GetSessions() (Sessions, error) {

    data, err := p.ReadDB()
    
    if err != nil {
        return Sessions{},err
    }

    return data.Sessions, err
}
func (p *Plight) SessionAdd(session string) error {

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
				{
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

	err = p.writeDB(data)

	return err
}
