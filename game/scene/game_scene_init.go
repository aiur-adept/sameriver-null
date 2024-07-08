package scene

import (
	"time"

	"github.com/aiur-adept/sameriver/v6"
)

func (s *GameScene) Init(game *sameriver.Game, config map[string]string) {
	// set scene "abstract base class" members
	s.destroyed = false
	s.game = game

	// build world and spawn entities
	s.buildWorld()
	s.spawnInitialEntities()

	// just to play a little loading screen fun
	time.Sleep(1 * time.Second)
}
