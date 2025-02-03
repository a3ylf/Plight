package db

type Period struct {
	From string `json:"from"`
	To   string `json:"to"`
}
type Day struct {
    Day_Total string `json:"total"`
	Periods   []Period `json:"periods"`
}

type Days struct {
    Last string
    Days  map[string]Day 
}
type Sessions map[string]Days 

type Data struct {
    Sessions Sessions  `json:"sessions"`
	Hits map[string]map[string]int `json:"hits"`
}
