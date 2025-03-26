package game

import (
	"image/color"
	"log"
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {

	// Food related variables
	food            []food
	maxFood         int
	foodTickCounter int
}

type food struct {
	imgPath string
	x, y    float64
}

func NewFood(x, y float64) food {
	return food{
		imgPath: "assets/asset-pack/Items/Food/Fish.png",
		x:       x,
		y:       y,
	}
}

func (f *food) draw(screen *ebiten.Image) {
	img, _, err := ebitenutil.NewImageFromFile(f.imgPath)
	if err != nil {
		log.Fatal(err)
	}
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(f.x, f.y)
	screen.DrawImage(img, opts)
}

func (g *Game) Update() error {
	g.tickFood()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 100, 255, 255})
	for _, f := range g.food {
		f.draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
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
		g.addFood(NewFood(x, y))
	}
}

// addFood adds a food item to the game along with its collider for convenience
func (g *Game) addFood(f food) {
	g.food = append(g.food, f)
}

func NewGame() *Game {
	return &Game{
		maxFood: 10,
	}
}
