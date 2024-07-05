package scene

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/aiur-adept/sameriver/v5"
)

func (s *GameScene) Name() string {
	return "game-scene"
}

func (s *GameScene) Update(dt_ms float64, allowance_ms float64) {
	s.w.Update(allowance_ms)
}

func (s *GameScene) Draw(w *sdl.Window, r *sdl.Renderer) {
	// draw the player
	s.SimpleEntityDraw(r, s.player, sdl.Color{255, 255, 255, 255})
}

func (s *GameScene) HandleKeyboardState(kb []uint8) {
	s.playerHandleKeyboardState(kb)
}

func (s *GameScene) HandleKeyboardEvent(ke *sdl.KeyboardEvent) {
	if ke.Type == sdl.KEYDOWN {
		if ke.Keysym.Sym == sdl.K_SPACE {
			fmt.Println("you pressed space")
		}
	}
}

func (s *GameScene) IsDone() bool {
	return s.ended
}

func (s *GameScene) NextScene() sameriver.Scene {
	return nil
}

func (s *GameScene) End() {
	fmt.Println(s.w.DumpStatsString())
}

func (s *GameScene) IsTransient() bool {
	return true
}

func (s *GameScene) Destroy() {
	if !s.destroyed {
		s.destroyed = true
	}
}
