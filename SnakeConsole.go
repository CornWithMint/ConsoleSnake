package main

import (
	"fmt"
	"math/rand"
	"os"
	"slices"
	"time"

	"github.com/eiannone/keyboard"
)

//Движение не только по секундам но и по нажатию
//Врезания в себя

type Snake struct {
	points []Point
}

type Point struct {
	PosX int
	PosY int
}

func CreatePoint() *Point {
	return &Point{
		PosX: rand.Intn(FieldWight),
		PosY: rand.Intn(FieldHeight),
	}
}

func (p *Snake) NewPointWithDelete(point *Point) {
	p.points = append([]Point{*point}, p.points...)
	p.points = p.points[:len(p.points)-1]
}
func (p *Snake) NewPointWithApple(point *Point, apple *Point) {
	p.points = append([]Point{*point}, p.points...)
	*apple = *CreatePoint()
}

func NewField(Snake *Snake, apple *Point) {
	for i := range FieldHeight {
		for j := range FieldWight {
			if slices.Contains(Snake.points, Point{j, i}) {
				fmt.Print("🔵")
			} else if apple.PosX == j && apple.PosY == i {
				fmt.Print("🍎")
			} else {
				fmt.Print("⬜️")
			}
		}
		fmt.Print("\n")
	}
}

// Функция устанавливающая в какую сторону повернута змея
func SetFacing(facing *string) {
	keyboard.Open()

	defer keyboard.Close()

	for {
		char, key, _ := keyboard.GetKey()
		switch {
		case key == keyboard.KeyArrowUp && *facing != "Down" && *facing != "Up":
			*facing = "Up"
		case key == keyboard.KeyArrowDown && *facing != "Up" && *facing != "Down":
			*facing = "Down"
		case key == keyboard.KeyArrowLeft && *facing != "Right" && *facing != "Left":
			*facing = "Left"
		case key == keyboard.KeyArrowRight && *facing != "Left" && *facing != "Right":
			*facing = "Right"
		}
		if char == 'q' || char == 'й' {
			os.Exit(0)
		}
	}
}

func (p *Snake) Moving(facing *string, apple *Point) *Point {
	head := p.points[0]
	switch *facing {
	case "Up":
		if head.PosY-1 == apple.PosY && head.PosX == apple.PosX {
			p.NewPointWithApple(&Point{PosX: apple.PosX, PosY: apple.PosY}, apple)
		} else if head.PosY > 0 {
			p.NewPointWithDelete(&Point{PosX: head.PosX, PosY: head.PosY - 1})
		} else {
			p.Wall()
		}
	case "Down":
		if head.PosY+1 == apple.PosY && head.PosX == apple.PosX {
			p.NewPointWithApple(&Point{PosX: apple.PosX, PosY: apple.PosY}, apple)
		} else if head.PosY < FieldHeight-1 {
			p.NewPointWithDelete(&Point{PosX: head.PosX, PosY: head.PosY + 1})
		} else {
			p.Wall()
		}
	case "Left":
		if head.PosY == apple.PosY && head.PosX-1 == apple.PosX {
			p.NewPointWithApple(&Point{PosX: apple.PosX, PosY: apple.PosY}, apple)
		} else if head.PosX > 0 {
			p.NewPointWithDelete(&Point{PosX: head.PosX - 1, PosY: head.PosY})
		} else {
			p.Wall()
		}
	case "Right":
		if head.PosY == apple.PosY && head.PosX+1 == apple.PosX {
			p.NewPointWithApple(&Point{PosX: apple.PosX, PosY: apple.PosY}, apple)
		} else if head.PosX < FieldWight-1 {
			p.NewPointWithDelete(&Point{PosX: head.PosX + 1, PosY: head.PosY})
		} else {
			p.Wall()
		}
	}
	return apple
}

func (p *Snake) Wall() {
	fmt.Println("Для тебя игра окончилась ДРУЖОК, ты в СТЕНЕ")
	if len(p.points) > GetScore(len(p.points)) {
		WriteScore(len(p.points))
	}
	fmt.Println("Твой результат: ", len(p.points), "\nЛучший результат: ", GetScore(len(p.points)))
	os.Exit(0)
}

const (
	FieldWight  = 10
	FieldHeight = 10
)

func main() {

	Snake := &Snake{points: []Point{{5, 5}}}

	facing := ""
	apple := CreatePoint()

	go SetFacing(&facing)

	fmt.Println("Чтобы начать нажми стрелочку")

	for facing == "" {
	}

	for {

		fmt.Println("Счет: ", len(Snake.points))
		NewField(Snake, apple)
		apple = Snake.Moving(&facing, apple)

		time.Sleep(time.Millisecond * 500)
		fmt.Print("\n")
	}
}
