package entities

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Food struct {
	imgPath string
	x, y    float64
}

func NewFood(x, y float64) Food {
	return Food{
		imgPath: "assets/asset-pack/Items/Food/Fish.png",
		x:       x,
		y:       y,
	}
}

func (f *Food) Draw(screen *ebiten.Image) {
	img, _, err := ebitenutil.NewImageFromFile(f.imgPath)
	if err != nil {
		log.Fatal(err)
	}
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(f.x, f.y)
	screen.DrawImage(img, opts)
}
