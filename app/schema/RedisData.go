package schema

import "time"

type Data struct {
	Key 		string			`json:"key"`
	Value 		interface{}		`json:"value"`
	Expiration 	time.Duration	`json:"expiration"`
}
