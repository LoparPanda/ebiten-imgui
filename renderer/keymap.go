package renderer

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/inkyblackness/imgui-go/v4"
)

var modKeys = []ebiten.Key{
	ebiten.KeyControlLeft,
	ebiten.KeyControlRight,
	ebiten.KeyAltLeft,
	ebiten.KeyAltRight,
	ebiten.KeyShiftLeft,
	ebiten.KeyShiftRight,
}

var keys = map[int]int{
	imgui.KeyTab:        int(ebiten.KeyTab),
	imgui.KeyLeftArrow:  int(ebiten.KeyLeft),
	imgui.KeyRightArrow: int(ebiten.KeyRight),
	imgui.KeyUpArrow:    int(ebiten.KeyUp),
	imgui.KeyDownArrow:  int(ebiten.KeyDown),
	imgui.KeyPageUp:     int(ebiten.KeyPageUp),
	imgui.KeyPageDown:   int(ebiten.KeyPageDown),
	imgui.KeyHome:       int(ebiten.KeyHome),
	imgui.KeyEnd:        int(ebiten.KeyEnd),
	imgui.KeyInsert:     int(ebiten.KeyInsert),
	imgui.KeyDelete:     int(ebiten.KeyDelete),
	imgui.KeyBackspace:  int(ebiten.KeyBackspace),
	imgui.KeySpace:      int(ebiten.KeySpace),
	imgui.KeyEnter:      int(ebiten.KeyEnter),
	imgui.KeyEscape:     int(ebiten.KeyEscape),
	imgui.KeyA:          int(ebiten.KeyA),
	imgui.KeyC:          int(ebiten.KeyC),
	imgui.KeyV:          int(ebiten.KeyV),
	imgui.KeyX:          int(ebiten.KeyX),
	imgui.KeyY:          int(ebiten.KeyY),
	imgui.KeyZ:          int(ebiten.KeyZ),
}

func sendInput(io *imgui.IO, inputChars []rune) []rune {
	// TODO: get KeySuper somehow (GLFW: KeyLeftSuper    = Key(343), R: 347)
	inputChars = ebiten.AppendInputChars(inputChars)
	if len(inputChars) > 0 {
		io.AddInputCharacters(string(inputChars))
		inputChars = inputChars[:0]
	}
	for _, iv := range keys {
		if inpututil.IsKeyJustPressed(ebiten.Key(iv)) {
			io.KeyPress(iv)
		}
		if inpututil.IsKeyJustReleased(ebiten.Key(iv)) {
			io.KeyRelease(iv)
		}
	}
	for _, iv := range modKeys {
		if inpututil.IsKeyJustPressed(iv) {
			io.KeyPress(int(iv))
		}
		if inpututil.IsKeyJustReleased(iv) {
			io.KeyRelease(int(iv))
		}
	}
	io.KeyAlt(int(ebiten.KeyAltLeft), int(ebiten.KeyAltRight))
	io.KeyShift(int(ebiten.KeyShiftLeft), int(ebiten.KeyShiftRight))
	io.KeyCtrl(int(ebiten.KeyControlLeft), int(ebiten.KeyControlRight))
	return inputChars
}
