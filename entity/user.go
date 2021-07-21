package entity

type User struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Address    string `json:"address"`
	Searchable string `json:"searchable"`
}
