package model

// Bread is bread entity
type Bread struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

func NewBread(id uint64, name string) *Bread {
	return &Bread{Id: id, Name: name}
}

func (h *Bread) String() string {
	return "Bread: " + h.Name
}
