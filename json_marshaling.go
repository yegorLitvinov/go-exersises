package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Email     string `json:"email"`
	ID        int    `json:"id,int"`
	CompanyID int    `json:"company_id,int"`
}

func main5() {
	str := `{"email": "a@a.com", "iD": 1, "company_ID": 2}`
	data := []byte(str)
	user := User{}
	if err := json.Unmarshal(data, &user); err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", user)
}
