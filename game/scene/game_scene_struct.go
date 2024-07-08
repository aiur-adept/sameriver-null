package scene

import (
	"github.com/aiur-adept/sameriver/v6"
)

type GameScene struct {

	// Scene "abstract class members"

	// whether the scene is running
	ended bool
	// used to make destroy() idempotent
	destroyed bool
	// the game
	game *sameriver.Game

	// GameScene members
	w      *sameriver.World
	player *sameriver.Entity

	obstacles []*sameriver.Entity
}
