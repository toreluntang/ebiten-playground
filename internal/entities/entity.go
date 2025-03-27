package entities

import (
	"image"
	"log"
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
	sprites     []*ebiten.Image
	tickCounter int
}

func NewPlayer(x, y float64) Player {
	spriteSize := 16
	imagePath := "assets/Snake4.png"
	img, _, err := ebitenutil.NewImageFromFile(imagePath)
	if err != nil {
		log.Fatal(err)
	}

	spriteAtIndex := func(index int) *ebiten.Image {
		row := index % (img.Bounds().Max.X / spriteSize)
		col := index / (img.Bounds().Max.X / spriteSize)
		x0 := row * spriteSize
		y0 := col * spriteSize
		x1 := x0 + 16
		y1 := y0 + 16

		cropSpace := image.Rect(int(x0), int(y0), int(x1), int(y1))
		return ebiten.NewImageFromImage(img.SubImage(cropSpace))
	}

	var allSprites []*ebiten.Image
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

	p.currImg = p.sprites[p.currImgIdx]
}

func (p *Player) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(p.x, p.y)
	screen.DrawImage(ebiten.NewImageFromImage(p.currImg), opts)
}

func (p *Player) move() {

	// Move the player based on the velocity
	p.x += p.dx
	p.y += p.dy

	// Teleport the player when out of bounds
	if p.x < 0 {
		p.x = 320
	}
	if p.x > 320 {
		p.x = 0
	}
	if p.y < 0 {
		p.y = 240
	}
	if p.y > 240 {
		p.y = 0
	}
}

func (p *Player) handleInput() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.dy = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.dy = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.dx = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.dx = 1
	}
}

func (p *Player) Update() error {
	if p.tickCounter >= 100 {
		p.tickCounter = 0
	}
	p.tickCounter++

	p.dx = 0
	p.dy = 0
	p.handleInput()
	p.move()

	if (p.dx != 0 || p.dy != 0) && p.tickCounter%5 == 0 {
		p.updateCurrImage()
	}

	return nil
}
