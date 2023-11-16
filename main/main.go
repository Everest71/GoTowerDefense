package main

//const mapPath = "demoMap.tmx" // Path to your Tiled Map.
//type mapGame struct {
//	Level    *tiled.Map
//	tileHash map[uint32]*ebiten.Image
//}
//
//func (m mapGame) Update() error {
//	return nil
//}
//func (m mapGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
//	//TODO implement me
//	return outsideWidth, outsideHeight
//}
//
//func makeEbiteImagesFromMap(tiledMap tiled.Map) map[uint32]*ebiten.Image {
//	idToImage := make(map[uint32]*ebiten.Image)
//	for _, tile := range tiledMap.Tilesets[0].Tiles {
//		ebitenImageTile, _, err :=
//			ebitenutil.NewImageFromFile(tile.Image.Source)
//		if err != nil {
//			fmt.Println("Error loading tile image:",
//				tile.Image.Source, err)
//		}
//		idToImage[tile.ID] = ebitenImageTile
//	}
//	return idToImage
//}
//
//func main() {
//	// Parse .tmx file.
//	gameMap, err := tiled.LoadFile(mapPath)
//	windowWidth := gameMap.Width * gameMap.TileWidth
//	windowHeight := gameMap.Height * gameMap.TileHeight
//	ebiten.SetWindowSize(windowWidth, windowHeight)
//	if err != nil {
//		fmt.Printf("error parsing map: %s", err.Error())
//		os.Exit(2)
//	}
//	ebitenImageMap := makeEbiteImagesFromMap(*gameMap)
//	oneLevelGame := mapGame{
//		Level:    gameMap,
//		tileHash: ebitenImageMap,
//	}
//	fmt.Println("tilesets:", gameMap.Tilesets[0].Tiles)
//	//fmt.Println("layers:", gameMap.Layers[0].Tiles)
//	fmt.Print("type:", fmt.Sprintf("%T", gameMap.Layers[0].Tiles[0]))
//	err = ebiten.RunGame(&oneLevelGame)
//	if err != nil {
//		fmt.Println("Couldn't run game:", err)
//	}
// }
