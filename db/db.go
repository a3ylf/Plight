package db

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/a3ylf/plight/flags"
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
		Sessions: make(map[string]Days),
		Hits:     make(map[string]map[string]int),
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
//
//	d , _ := time.Parse(time.DateTime,b)
//	e := d.Sub(c).String()
//
// fmt.Println(e)
//
//	x, _ := time.ParseDuration(e)
//	fmt.Println(x)
func (p *Plight) writeDB(data *Data) error {
	d, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		return err
	}
	p.mux.Lock()
	err = os.WriteFile(p.filename, d, 0644)
	p.mux.Unlock()

	return err

}


