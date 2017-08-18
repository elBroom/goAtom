package model

type User struct {
	ID        	string   	`json:"id,omitempty"`
	Login		string   	`json:"login,omitempty"`
	Password  	string   	`json:"password,omitempty"`
	Database	[]*Database	`json:"database,omitempty"`
	Query  	 	[]*Query 	`json:"query,omitempty"`
}
