package model

type Query struct {
	Time 		int 		`json:"name,omitempty"`
	//Database   	*Database 	`json:"database,omitempty"`
	Query  		string 		`json:"name,omitempty"`
}
