package entities

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// TODO: Consider creating interface for entities.
// It might make sense to create a list of Entities on Game, and loop over them in Update and Draw call entity.Draw() entity.Update
// to handle various behaviour for each sub entity (food, player, enemy, obstacle, whatever)

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
