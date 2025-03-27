package entities

import (
	"image"
	"log"
	"math"
	"slices"

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
	dx, dy    float64

	spriteSheet *ebiten.Image
	spriteSize  int
	currImg     image.Image
	currImgIdx  int
	sprites     []image.Image
	tickCounter int
}

func NewPlayer(x, y float64) Player {
	spriteSize := 16
	imagePath := "assets/Snake4.png"
	img, _, err := ebitenutil.NewImageFromFile(imagePath)
	if err != nil {
		log.Fatal(err)
	}

	spriteAtIndex := func(index int) image.Image {
		row := index % (img.Bounds().Max.X / spriteSize)
		col := index / (img.Bounds().Max.X / spriteSize)
		x0 := row * spriteSize
		y0 := col * spriteSize
		x1 := x0 + 16
		y1 := y0 + 16

		cropSpace := image.Rect(int(x0), int(y0), int(x1), int(y1))
		return img.SubImage(cropSpace)
	}

	var allSprites []image.Image
	for i := range 16 {
		allSprites = append(allSprites, spriteAtIndex(i))
	}

	return Player{
		imagePath:   imagePath,
		spriteSize:  spriteSize,
		x:           x,
		y:           y,
		dx:          0,
		dy:          0,
		spriteSheet: img,
		currImg:     allSprites[0],
		currImgIdx:  0,
		sprites:     allSprites,
	}
}

func (p *Player) updateCurrImage() {
	rightStart := 3
	leftStart := 2
	upStart := 1
	downStart := 0
	if p.dx > 0 { // Moving right
		pickNextInLine := slices.Contains([]int{3, 7, 11, 15}, p.currImgIdx)
		if !pickNextInLine || p.currImgIdx == 15 {
			p.currImgIdx = rightStart
		} else {
			p.currImgIdx += 4
		}

	} else if p.dx < 0 { // Moving left
		pickNextInLine := slices.Contains([]int{2, 6, 10, 14}, p.currImgIdx)
		if !pickNextInLine || p.currImgIdx == 14 {
			p.currImgIdx = leftStart
		} else {
			p.currImgIdx += 4
		}

	} else if p.dy > 0 { // Moving down
		pickNextInLine := slices.Contains([]int{0, 4, 8, 12}, p.currImgIdx)
		if !pickNextInLine || p.currImgIdx == 12 {
			p.currImgIdx = downStart
		} else {
			p.currImgIdx += 4
		}

	} else if p.dy < 0 { // Moving up
		pickNextInLine := slices.Contains([]int{1, 5, 9, 13}, p.currImgIdx)
		if !pickNextInLine || p.currImgIdx == 13 {
			p.currImgIdx = upStart
		} else {
			p.currImgIdx += 4
		}
	} else {
		return
	}

	img := p.sprites[p.currImgIdx]
	p.currImg = ebiten.NewImageFromImage(img)
}

func (p *Player) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(p.x, p.y)
	screen.DrawImage(ebiten.NewImageFromImage(p.currImg), opts)
}

func (p *Player) Update() error {
	if p.tickCounter >= 100 {
		p.tickCounter = 0
	}
	p.tickCounter++
	p.dx = 0
	p.dy = 0
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		loc := p.y - 1
		p.y = math.Max(0, loc)
		p.dy = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		loc := p.y + 1
		p.y = math.Min(240-16, loc) // TODO: Move 240 to contants.WorldSizeHeight move 16 to constants.playerSize
		p.dy = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		loc := p.x - 1
		p.x = math.Max(0, loc)
		p.dx = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		loc := p.x + 1
		p.x = math.Min(320-16, loc)
		p.dx = 1
	}

	if (p.dx != 0 || p.dy != 0) && p.tickCounter%5 == 0 {
		p.updateCurrImage()
	}

	return nil
}
