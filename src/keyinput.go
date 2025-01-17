package trisoban

import (
	tl "github.com/JoelOtter/termloop"
)

// Tick listens to keyinput on the titlescreen, so if the enter key is pressed, the game will start.
func (ts *Titlescreen) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyEnter:
			CurrentLevel = 1
			gs = NewGameScreen()
			game.Screen().SetLevel(gs)

		case tl.KeyBackspace:
			CurrentLevel = LoadLevel()
			gs = NewGameScreen()
			game.Screen().SetLevel(gs)
		}
	}
}

// Tick listens to the gamescreen input and handles it accordingly.
func (gs *Gamescreen) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyArrowUp:
			gs.Move("up")
		case tl.KeyArrowDown:
			gs.Move("down")
		case tl.KeyArrowLeft:
			gs.Move("left")
		case tl.KeyArrowRight:
			gs.Move("right")
		case tl.KeyF3:
			gs.RestartLevel()
		case tl.KeyF1:
			gs.SaveConfirmation.SetText("")
			if gs.CheckLevelCompletion() {
				gs.ChangeLevel("next")
			}
		case tl.KeyF2:
			gs.SaveConfirmation.SetText("")
			if gs.CheckLevelCompletion() {
				gs.ChangeLevel("previous")
			}
		case tl.KeyInsert:
			gs.SaveGame()
		}
	}
}
