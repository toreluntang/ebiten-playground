package game

import (
	"eat-and-grow/internal/entities"
	"image/color"
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {

	// Food related variables
	food            []entities.Food
	maxFood         int
	foodTickCounter int
	pause           bool
}

func (g *Game) Update() error {
	err := g.checkKeyPress()
	if err != nil {
		return err
	}

	if g.pause {
		return nil
	}

	g.tickFood()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 100, 255, 255})
	for _, f := range g.food {
		f.Draw(screen)
	}

	if g.pause {
		ebitenutil.DebugPrintAt(screen, "Paused", 320/2, 240/2)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (g *Game) checkKeyPress() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyP) {
		g.Pause()
	}

	return nil
}

func (g *Game) Pause() {
	g.pause = !g.pause
}

func (g *Game) tickFood() {
	g.foodTickCounter += 1
	if g.foodTickCounter == 50 { // TODO: Extract to constants probably
		g.randSpawnFood()
		g.foodTickCounter = 0
	}
}

func (g *Game) randSpawnFood() {
	// Do not add infinite food
	if len(g.food) >= g.maxFood {
		return
	}

	if rand.IntN(100) > 1 {
		size := 16
		randX := rand.IntN(320) - size
		randY := rand.IntN(230) - size
		x := math.Max(0, float64(randX))
		y := math.Max(0, float64(randY))
		g.addFood(entities.NewFood(x, y))
	}
}

// addFood adds a food item to the game along with its collider for convenience
func (g *Game) addFood(f entities.Food) {
	g.food = append(g.food, f)
}

func NewGame() *Game {
	return &Game{
		maxFood: 10,
	}
}
