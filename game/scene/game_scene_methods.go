package scene

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/aiur-adept/sameriver/v5"

	"github.com/aiur-adept/sameriver-null/game/systems"
)

func (s *GameScene) buildWorld() {
	// construct world object
	s.w = sameriver.NewWorld(map[string]any{
		"width":  s.game.WindowSpec.Width,
		"height": s.game.WindowSpec.Height,
	})

	// register components must always be called before AddSystems()
	// since systems might want to create and listen on component bitarray
	// filters

	// add systems
	s.w.RegisterSystems(
		sameriver.NewPhysicsSystem(),
		sameriver.NewSpatialHashSystem(32, 32),
		sameriver.NewCollisionSystem(10*time.Millisecond),
		systems.NewExampleSystem(),
	)
	s.w.SetSystemSchedule("CollisionSystem", 16)

}

func (s *GameScene) spawnInitialEntities() {
	s.player = s.w.Spawn(map[string]any{
		"components": map[sameriver.ComponentID]any{
			sameriver.POSITION_:     sameriver.Vec2D{100, 100},
			sameriver.VELOCITY_:     sameriver.Vec2D{0, 0},
			sameriver.ACCELERATION_: sameriver.Vec2D{0, 0},
			sameriver.BOX_:          sameriver.Vec2D{10, 10},
			sameriver.MASS_:         1.0,
			sameriver.RIGIDBODY_:    true,
		},
		"tags": []string{"player"},
	})
}

func (s *GameScene) SimpleEntityDraw(
	r *sdl.Renderer, e *sameriver.Entity, c sdl.Color) {

	box := e.GetVec2D(sameriver.BOX_)
	pos := e.GetVec2D(sameriver.POSITION_).ShiftedCenterToBottomLeft(*box)
	r.SetDrawColor(c.R, c.G, c.B, c.A)
	s.game.Screen.FillRect(r, &pos, box)
}

func (s *GameScene) playerHandleKeyboardState(kb []uint8) {
	v := s.player.GetVec2D(sameriver.VELOCITY_)
	// get player v1
	v.X = 0.4 * float64(
		int8(kb[sdl.SCANCODE_D]|kb[sdl.SCANCODE_RIGHT])-
			int8(kb[sdl.SCANCODE_A]|kb[sdl.SCANCODE_LEFT]))
	v.Y = 0.4 * float64(
		int8(kb[sdl.SCANCODE_W]|kb[sdl.SCANCODE_UP])-
			int8(kb[sdl.SCANCODE_S]|kb[sdl.SCANCODE_DOWN]))
}
