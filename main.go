package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	// Food related variables
	food    []*food
	maxFood int
}

type food struct {
	imgPath string
	x, y    float64
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 100, 255, 255})
	ebitenutil.DebugPrint(screen, "Hello, World!")

	for _, food := range g.food {
		if food == nil {
			continue
		}
		img, _, err := ebitenutil.NewImageFromFile(food.imgPath)
		if err != nil {
			log.Fatal(err)
		}
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(food.x, food.y)
		screen.DrawImage(img, opts)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

// addFood adds a food item to the game along with its collider for convenience
func (g *Game) addFood(f food) {
	g.food = append(g.food, &f)
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{
		food: []*food{
			{
				imgPath: "assets/asset-pack/Items/Food/Fish.png",
				x:       10,
				y:       10,
			},
		},
	}); err != nil {
		log.Fatal(err)
	}
}
