package entity

type User struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Address    string `json:"address"`
	Searchable string `json:"searchable"`
}

type EsUser struct {
	Name       string `json:"name"`
	Searchable string `json:"searchable"`
}

func (e *User) NewEsUser() *EsUser {
	return &EsUser{
		Name:       e.Name,
		Searchable: e.Searchable,
	}
}
