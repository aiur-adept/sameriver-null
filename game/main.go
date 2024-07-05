package main

import (
	"fmt"

	"github.com/aiur-adept/sameriver-null/game/scene"
	"github.com/aiur-adept/sameriver/v5"
)

func main() {
	fmt.Println("space cats, motherfucker")
	sameriver.RunGame(sameriver.GameInitSpec{
		WindowSpec: sameriver.WindowSpec{
			Title:      "GAME",
			Width:      800,
			Height:     800,
			Fullscreen: false},
		LoadingScene: &scene.LoadingScene{},
		FirstScene:   &scene.GameScene{},
	})
}
