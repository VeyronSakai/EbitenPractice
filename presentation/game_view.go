package presentation

import (
	"EbitenSample/application"
	"EbitenSample/domain"
	"EbitenSample/domain/player"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var pointerImage = ebiten.NewImage(10, 30)

type Game struct {
	playerApplicationService *application.PlayerApplicationService
	playerId                 player.Id
}

func NewGame(repository player.Repository, factory player.Factory) (*Game, error) {

	playerApplicationService := application.NewPlayerApplicationService(repository, factory)

	playerId, err := playerApplicationService.SpawnPlayer()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	game := &Game{
		playerApplicationService: playerApplicationService,
		playerId:                 playerId,
	}

	return game, nil
}

func (game *Game) Update() error {

	moveVec := domain.Vector2{}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		moveVec.X += -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		moveVec.X += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		moveVec.Y += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		moveVec.Y += -1
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {

	}

	playerMoveDir := player.MoveDir{Value: moveVec}
	game.playerApplicationService.MovePlayer(game.playerId, playerMoveDir)

	return nil
}

func (game Game) Draw(screen *ebiten.Image) {
	p, err := game.playerApplicationService.GetPlayerData(game.playerId)
	if err != nil {
		log.Fatal(err)
	}

	op := &ebiten.DrawImageOptions{}
	// 画面の中央を原点にする
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	op.GeoM.Translate(p.Pos.Value.X, p.Pos.Value.Y)
	screen.DrawImage(pointerImage, op)
}

func (game Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func InitializeView() {
	pointerImage.Fill(color.RGBA{R: 0xff, A: 0xff})
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Practice")
}
