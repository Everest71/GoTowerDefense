package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
	"os"
)

const mapPath = "assets/Map/Home.tmx"

type game struct {
	Level    *tiled.Map
	tileHash map[uint32]*ebiten.Image
}

func (game game) Draw(screen *ebiten.Image) {
	drawOptions := ebiten.DrawImageOptions{}
	for tileY := 0; tileY < game.Level.Height; tileY += 1 {
		for tileX := 0; tileX < game.Level.Width; tileX += 1 {
			drawOptions.GeoM.Reset()
			TileXpos := float64(game.Level.TileWidth * tileX)
			TileYpos := float64(game.Level.TileHeight * tileY)
			drawOptions.GeoM.Translate(TileXpos, TileYpos)
			tileToDraw := game.Level.Layers[0].Tiles[tileY*game.Level.Width+tileX]
			ebitenTileToDraw := game.tileHash[tileToDraw.ID]
			fmt.Println(ebitenTileToDraw)
			screen.DrawImage(ebitenTileToDraw, &drawOptions)
		}
	}
}

func (m game) Update() error {
	return nil
}
func (m game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	//TODO implement me
	return outsideWidth, outsideHeight
}

func makeEbiteImagesFromMap(tiledMap tiled.Map) map[uint32]*ebiten.Image {
	idToImage := make(map[uint32]*ebiten.Image)
	for _, tile := range tiledMap.Tilesets[0].Tiles {
		ebitenImageTile, _, err :=
			ebitenutil.NewImageFromFile(tile.Image.Source)
		if err != nil {
			fmt.Println("Error loading tile image:",
				tile.Image.Source, err)
		}
		idToImage[tile.ID] = ebitenImageTile
	}
	return idToImage
}

func main() {
	// Parse .tmx file.
	gameMap, err := tiled.LoadFile(mapPath)
	windowWidth := 800
	windowHeight := 800
	ebiten.SetWindowSize(windowWidth, windowHeight)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}
	ebitenImageMap := makeEbiteImagesFromMap(*gameMap)
	farmer := game{
		Level:    gameMap,
		tileHash: ebitenImageMap,
	}
	fmt.Println("tilesets:", gameMap.Tilesets[0].Tiles)
	//fmt.Println("layers:", gameMap.Layers[0].Tiles)
	fmt.Print("type:", fmt.Sprintf("%T", gameMap.Layers[0].Tiles[0]))
	err = ebiten.RunGame(farmer)
	if err != nil {
		fmt.Println("Couldn't run game:", err)
	}
}
