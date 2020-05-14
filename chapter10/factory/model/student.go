package model

type student struct {
	Name  string
	Score float64
}

func NewStudent(name string, score float64) *student {
	return &student{
		Name:  name,
		Score: score,
	}
}
