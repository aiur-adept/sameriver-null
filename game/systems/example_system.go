package systems

import (
	"github.com/aiur-adept/sameriver/v6"
)

type ExampleSystem struct {
	w *sameriver.World
}

func NewExampleSystem() *ExampleSystem {
	return &ExampleSystem{}
}

func (s *ExampleSystem) LinkWorld(w *sameriver.World) {
	s.w = w
}

func (s *ExampleSystem) Update(dt_ms float64) {

}

func (s *ExampleSystem) Expand(n int) {

}

func (s *ExampleSystem) GetComponentDeps() []any {
	return []any{}
}
