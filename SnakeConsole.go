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
//Сохранение рекорда по длине
//Врезания в себя

type Player struct {
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

func NewPointWithDelete(player *Player, point *Point) {
	player.points = append([]Point{*point}, player.points...)
	player.points = player.points[:len(player.points)-1]
}
func NewPointWithApple(player *Player, point *Point, apple *Point) {
	player.points = append([]Point{*point}, player.points...)
	*apple = *CreatePoint()
}

func NewField(player *Player, apple *Point) {
	for i := range FieldHeight {
		for j := range FieldWight {
			if slices.Contains(player.points, Point{j, i}) {
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
func SetFacing(player *Player, facing *string) {
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

func Moving(player *Player, facing *string, apple *Point) *Point {
	head := player.points[0]
	switch *facing {
	case "Up":
		if head.PosY-1 == apple.PosY && head.PosX == apple.PosX {
			NewPointWithApple(player, &Point{PosX: apple.PosX, PosY: apple.PosY}, apple)
		} else if head.PosY > 0 {
			NewPointWithDelete(player, &Point{PosX: head.PosX, PosY: head.PosY - 1})
		} else {
			Wall()
		}
	case "Down":
		if head.PosY+1 == apple.PosY && head.PosX == apple.PosX {
			NewPointWithApple(player, &Point{PosX: apple.PosX, PosY: apple.PosY}, apple)
		} else if head.PosY < FieldHeight-1 {
			NewPointWithDelete(player, &Point{PosX: head.PosX, PosY: head.PosY + 1})
		} else {
			Wall()
		}
	case "Left":
		if head.PosY == apple.PosY && head.PosX-1 == apple.PosX {
			NewPointWithApple(player, &Point{PosX: apple.PosX, PosY: apple.PosY}, apple)
		} else if head.PosX > 0 {
			NewPointWithDelete(player, &Point{PosX: head.PosX - 1, PosY: head.PosY})
		} else {
			Wall()
		}
	case "Right":
		if head.PosY == apple.PosY && head.PosX+1 == apple.PosX {
			NewPointWithApple(player, &Point{PosX: apple.PosX, PosY: apple.PosY}, apple)
		} else if head.PosX < FieldWight-1 {
			NewPointWithDelete(player, &Point{PosX: head.PosX + 1, PosY: head.PosY})
		} else {
			Wall()
		}
	}
	return apple
}

func Wall() {
	fmt.Println("Для тебя игра окончилась ДРУЖОК, ты в СТЕНЕ")
	os.Exit(0)
}

const (
	FieldWight  = 10
	FieldHeight = 10
)

func main() {
	player := &Player{points: []Point{{5, 5}}}
	facing := ""
	apple := CreatePoint()

	go SetFacing(player, &facing)

	fmt.Println("Чтобы начать нажми стрелочку")

	for facing == "" {
	}

	for {
		fmt.Print("\033[H\033[2J")

		fmt.Println("Длина змейки: ", len(player.points))
		NewField(player, apple)
		apple = Moving(player, &facing, apple)

		time.Sleep(time.Millisecond * 500)
		fmt.Print("\n")
	}
}
