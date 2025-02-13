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
    Last string `json:"last"`
    Days  map[string]Day `json:"days"`
}
type Sessions map[string]Days 

type Data struct {
    Sessions Sessions  `json:"sessions"`
	Hits map[string]map[string]int `json:"hits"`
}
