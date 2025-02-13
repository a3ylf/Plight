package db

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// func (p *Plight) GetSession(session string) ([]byte, error) {
//
// 	data, err := p.ReadDB()
//
// 	if err != nil {
// 		return []byte{}, err
// 	}
// 	t, e := data.Sessions[session]
//
// 	if !e {
// 		return []byte{}, errors.New("Unable to find this session")
// 	}
// 	send, err := json.MarshalIndent(t, "", "   ")
// 	return send, nil
// }

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
	daynow := time.Now().Format(time.DateOnly)
	timenow := time.Now().Format(time.TimeOnly)

	if data.Sessions == nil {
		data.Sessions = make(map[string]Days)
	}
	if _, e := data.Sessions[session]; !e {
		fmt.Printf("You are creating a new session, rewrite it's name to confirm\n%v: ", session)
		var check string
		fmt.Scan(&check)
		if check != session {
			return (errors.New("session not created"))
		}
		data.Sessions[session] = Days{
			Days: make(map[string]Day),
		}
	}
	last := len(data.Sessions[session].Days[daynow].Periods) - 1

	// Never explaining this code in my life
	if last == -1 {
		did := false
		var day Day
		dlast := data.Sessions[session].Last
		lday, err := time.Parse(time.DateOnly, dlast)
		daydate, err := time.Parse(time.DateOnly, daynow)
		if err != nil {
			log.Println(err)
		}
		if dlast != "" {
			lastlast := len(data.Sessions[session].Days[dlast].Periods) - 1
			if data.Sessions[session].Days[dlast].Periods[lastlast].To == "" {
				if daydate.YearDay()-1 == lday.YearDay() {
					var answer string
					fmt.Println("I think you forgor to close last time bro 😱😱, write bob to save it ")
					fmt.Scan(&answer)
					if answer == "bob" {
						data.Sessions[session].Days[dlast].Periods[lastlast].To = "23:59:59"
						dur, err := time.ParseDuration(data.Sessions[session].Days[dlast].Day_Total)
						from, err := time.Parse(time.TimeOnly, data.Sessions[session].Days[dlast].Periods[lastlast].From)

						to, err := time.Parse(time.TimeOnly, "23:59:59")
						if err != nil {
							return err
						}
						x := to.Sub(from)
						now, err := time.Parse(time.TimeOnly, timenow)
						if err != nil {
							fmt.Println("How?")
						}

						y := now.Add(x).Add(dur).Sub(now)
						s := data.Sessions[session].Days[dlast]
						s.Day_Total = y.String()

						data.Sessions[session].Days[dlast] = s

						midnight, err := time.Parse(time.TimeOnly, "00:00:00")
						nextDay := now.Sub(midnight)
						day = Day{
							Day_Total: nextDay.String(),
							Periods: []Period{
								{
									From: "00:00:00",
									To:   now.Format(time.TimeOnly),
								},
							},
						}
					}
					did = true
				} else {
					fmt.Println("Deleting what you forgot bro")
					kk := data.Sessions[session].Days[dlast].Periods
					newperiods := kk[0:lastlast]
					s := data.Sessions[session].Days[dlast]
					s.Periods = newperiods
					data.Sessions[session].Days[dlast] = s

				}
			} 
			
		}
		if !did {
			day = Day{
				Day_Total: "0s",
				Periods: []Period{
					{
						From: timenow,
					},
				},
			}
		}
		d := data.Sessions[session]
		d.Days[daynow] = day
		d.Last = daynow
		data.Sessions[session] = d
		fmt.Printf("Session %v started\nCurrent time: %v\n", session, timenow)

		// forgive me for this
	} else if data.Sessions[session].Days[daynow].Periods[last].To == "" {
		data.Sessions[session].Days[daynow].Periods[last].To = timenow
		dur, err := time.ParseDuration(data.Sessions[session].Days[daynow].Day_Total)
		from, err := time.Parse(time.TimeOnly, data.Sessions[session].Days[daynow].Periods[last].From)
		to, err := time.Parse(time.TimeOnly, data.Sessions[session].Days[daynow].Periods[last].To)

		if err != nil {
			return err
		}
		dursess := to.Sub(from)
		newTotal := to.
			Add(dur).
			Sub(from).
			String()

		s := data.Sessions[session].Days[daynow]
		s.Day_Total = newTotal

		data.Sessions[session].Days[daynow] = s
		fmt.Printf("Session ended\nSession duration: %v\nNew total duration: %v\n", dursess, newTotal)

	} else {
		a := data.Sessions[session].Days[daynow]
		a.Periods = append(a.Periods, Period{
			From: timenow})
		data.Sessions[session].Days[daynow] = a
		fmt.Printf("Session %v started\nCurrent time: %v\n", session, timenow)
	}

	err = p.writeDB(data)

	return err
}
