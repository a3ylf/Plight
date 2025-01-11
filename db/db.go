package db

import (
	"os"
	"sync"
	"time"
)

type Plight struct {
    mux sync.RWMutex
    filename string
}

func StartDB(filename ... string) *Plight{
    base := "data.json"
    if len(filename) > 0 {
        base = filename[0]
    }
    return &Plight{
        mux: sync.RWMutex{},
        filename: base,

    }

}

func (p *Plight)EnsureDB () error{
    name := "data.json"
    if p.filename != "" {
        name = p.filename
        
    }
    
    _, err := os.Open(name)
    
    if err != nil {
       err = createDB(name)
    }

    return err
}

func createDB(filename string) error {
    _, err := os.Create(filename)
    return err
}

func(p *Plight) writeDB (to string ,  xd time.Time)  {
}
