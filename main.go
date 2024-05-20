package main

// import (
// 	"fmt"
//
// 	"github.com/moverest/mnist"
// )

// func main() {
// 	images, err := mnist.LoadImageFile("./dataset/train-images.idx3-ubyte.gz")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
//
// 	for i := 0; i < 28; i++ {
// 		for j := 0; j < 28; j++ {
// 			if images[1][i*28+j] > 0 {
// 				fmt.Print("1 ")
// 			} else {
// 				fmt.Print("  ")
// 			}
// 		}
// 		fmt.Println()
// 	}
//
// }
//

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/moverest/mnist"
)

type Game struct {
	Images     []image.Image
	imageIdx   int
	frameCount int
}

func (g *Game) Update() error {
	// 5フレームごとに画像を切り替える
	g.frameCount++
	if g.frameCount == 10 {
		g.frameCount = 0
		g.imageIdx = (g.imageIdx + 1) % len(g.Images)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	// ここで画像を描画します

	// 画像を作成
	img := ebiten.NewImageFromImage(g.Images[g.imageIdx])

	// 画像を拡大
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(10, 10)
	screen.DrawImage(img, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 280, 280
}

func main() {
	ebiten.SetWindowSize(280, 280)
	ebiten.SetWindowTitle("Your title here")

	var game Game

	game.imageIdx = 0
	images, err := mnist.LoadImageFile("./dataset/train-images.idx3-ubyte.gz")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []*mnist.Image to []image.Image
	convertedImages := make([]image.Image, len(images))
	for i, img := range images {
		convertedImages[i] = img
	}

	game.Images = convertedImages

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
