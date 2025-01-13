package db

import "time"

func (p *Plight) HitAdd(hit string) error {
	data, err := p.ReadDB()
	if err != nil {
		return err
	}

	today := time.Now().Format(time.DateOnly)

	if _, e := data.Hits[hit]; !e {
		data.Hits[hit] = make(map[string]int)

	}
	data.Hits[hit][today]++

	err = p.writeDB(data)

	return err
}
