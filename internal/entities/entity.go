package entities

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// TODO: Consider creating interface for entities.
// It might make sense to create a list of Entities on Game, and loop over them in Update and Draw call entity.Draw() entity.Update
// to handle various behaviour for each sub entity (food, player, enemy, obstacle, whatever)

type Entity interface {
	Draw(screen *ebiten.Image)
	Update() error
}

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

func (f *Food) Update() error {
	return nil
}

type Player struct {
	imagePath string
	x, y      float64
	currImg   image.Image
}

func NewPlayer(x, y float64) Player {
	imagePath := "assets/Snake4.png"
	img, _, err := ebitenutil.NewImageFromFile(imagePath)
	if err != nil {
		log.Fatal(err)
	}

	crop := image.Rect(0, 0, 16, 16)
	croppedImg := img.SubImage(crop)

	return Player{
		imagePath: imagePath,
		x:         x,
		y:         y,
		currImg:   croppedImg,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	// Initially lets just crop a non-animated snake and draw it

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(p.x, p.y)
	screen.DrawImage(ebiten.NewImageFromImage(p.currImg), opts)
}

func (p *Player) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.y += 1
	}
	return nil
}
