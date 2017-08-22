package schema

type User struct {
	ID        	int   		`json:"id,omitempty"`
	Login		string   	`json:"login,omitempty"`
	Password  	string   	`json:"password,omitempty"`
	Name	  	string   	`json:"name,omitempty"`
}
