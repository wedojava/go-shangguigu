package main

import (
	"fmt"
)

// 编写学生考试系统

type Student struct {
	Name  string
	Age   int
	Score int
}
type Graduate struct {
	Student
}
type Pupil struct {
	Student
}

func (s *Student) ShowInfo() {
	fmt.Printf("[%v][%v]:%v", s.Name, s.Age, s.Score)
}

func (s *Student) SetScore(score int) {
	s.Score = score
}

// func (p *Graduate) ShowInfo() {
//         fmt.Printf("[%v][%v]:%v", p.Name, p.Age, p.Score)
// }

// func (p *Graduate) SetScore(score int) {
//         p.Score = score
// }

func (p *Graduate) testing() {
	fmt.Println("\nA graduate testing...")
}

// func (p *Pupil) ShowInfo() {
//         fmt.Printf("[%v][%v]:%v", p.Name, p.Age, p.Score)
// }

// func (p *Pupil) SetScore(score int) {
//         p.Score = score
// }

func (p *Pupil) testing() {
	fmt.Println("\nA pupil testing...")
}

func main() {
	// var g = Graduate{
	//         Name:  "BigTom",
	//         Age:   19,
	//         Score: 0,
	// }
	// g.testing()
	// g.SetScore(100)
	// g.ShowInfo()
	// var pupil = Pupil{
	//         Name:  "Tom",
	//         Age:   18,
	//         Score: 0,
	// }
	// pupil.testing()
	// pupil.SetScore(100)
	// pupil.ShowInfo()
	p := &Pupil{}
	p.Student.Name = "ppp"
	p.Student.Age = 10
	p.testing()
	p.Student.SetScore(100)
	p.Student.ShowInfo()

	g := &Graduate{}
	g.Student.Name = "merry"
	g.Student.Age = 19
	g.testing()
	g.Student.SetScore(90)
	g.Student.ShowInfo()
}
