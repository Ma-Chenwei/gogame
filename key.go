package gogame

import "github.com/hajimehoshi/ebiten/v2"

// KeyCode 是 GoGame 的键值。
// 用法:
//
//	key := gogame.Key.GetPressed()
//
//	if key[gogame.K_TAB] {
//		...
//	}
type KeyCode int

const (
	K_UNKNOWN KeyCode = iota

	K_A
	K_B
	K_C
	K_D
	K_E
	K_F
	K_G
	K_H
	K_I
	K_J
	K_K
	K_L
	K_M
	K_N
	K_O
	K_P
	K_Q
	K_R
	K_S
	K_T
	K_U
	K_V
	K_W
	K_X
	K_Y
	K_Z

	K_0
	K_1
	K_2
	K_3
	K_4
	K_5
	K_6
	K_7
	K_8
	K_9

	K_UP
	K_DOWN
	K_LEFT
	K_RIGHT

	K_TAB
	K_RETURN
	K_ESCAPE
	K_SPACE
	K_BACKSPACE

	K_LSHIFT
	K_RSHIFT
	K_LCTRL
	K_RCTRL
	K_LALT
	K_RALT

	K_CAPSLOCK
	K_NUMLOCK
	K_SCROLLLOCK

	K_INSERT
	K_DELETE
	K_HOME
	K_END
	K_PAGEUP
	K_PAGEDOWN

	K_F1
	K_F2
	K_F3
	K_F4
	K_F5
	K_F6
	K_F7
	K_F8
	K_F9
	K_F10
	K_F11
	K_F12
)

// KeyModule 对应 pygame.key。
type KeyModule struct{}

var Key = &KeyModule{}

// GetPressed 获取当前按下的所有按键。
// 对应:
//
//	pygame.key.get_pressed()
//
// 使用:
//
//	key := gogame.Key.GetPressed()
//
//	if key[gogame.K_UP] {
//		y -= 5
//	}
func (k *KeyModule) GetPressed() map[KeyCode]bool {
	result := make(map[KeyCode]bool)

	for key, pressed := range currentKeys {
		result[key] = pressed
	}

	return result
}

// Get 获取单个按键状态。
func (k *KeyModule) Get(code KeyCode) bool {
	return currentKeys[code]
}

func toEbitenKey(key KeyCode) ebiten.Key {
	switch key {

	case K_A:
		return ebiten.KeyA
	case K_B:
		return ebiten.KeyB
	case K_C:
		return ebiten.KeyC
	case K_D:
		return ebiten.KeyD
	case K_E:
		return ebiten.KeyE
	case K_F:
		return ebiten.KeyF
	case K_G:
		return ebiten.KeyG
	case K_H:
		return ebiten.KeyH
	case K_I:
		return ebiten.KeyI
	case K_J:
		return ebiten.KeyJ
	case K_K:
		return ebiten.KeyK
	case K_L:
		return ebiten.KeyL
	case K_M:
		return ebiten.KeyM
	case K_N:
		return ebiten.KeyN
	case K_O:
		return ebiten.KeyO
	case K_P:
		return ebiten.KeyP
	case K_Q:
		return ebiten.KeyQ
	case K_R:
		return ebiten.KeyR
	case K_S:
		return ebiten.KeyS
	case K_T:
		return ebiten.KeyT
	case K_U:
		return ebiten.KeyU
	case K_V:
		return ebiten.KeyV
	case K_W:
		return ebiten.KeyW
	case K_X:
		return ebiten.KeyX
	case K_Y:
		return ebiten.KeyY
	case K_Z:
		return ebiten.KeyZ

	case K_0:
		return ebiten.Key0
	case K_1:
		return ebiten.Key1
	case K_2:
		return ebiten.Key2
	case K_3:
		return ebiten.Key3
	case K_4:
		return ebiten.Key4
	case K_5:
		return ebiten.Key5
	case K_6:
		return ebiten.Key6
	case K_7:
		return ebiten.Key7
	case K_8:
		return ebiten.Key8
	case K_9:
		return ebiten.Key9

	case K_UP:
		return ebiten.KeyArrowUp
	case K_DOWN:
		return ebiten.KeyArrowDown
	case K_LEFT:
		return ebiten.KeyArrowLeft
	case K_RIGHT:
		return ebiten.KeyArrowRight

	case K_TAB:
		return ebiten.KeyTab
	case K_RETURN:
		return ebiten.KeyEnter
	case K_ESCAPE:
		return ebiten.KeyEscape
	case K_SPACE:
		return ebiten.KeySpace
	case K_BACKSPACE:
		return ebiten.KeyBackspace

	case K_LSHIFT:
		return ebiten.KeyShiftLeft
	case K_RSHIFT:
		return ebiten.KeyShiftRight
	case K_LCTRL:
		return ebiten.KeyControlLeft
	case K_RCTRL:
		return ebiten.KeyControlRight
	case K_LALT:
		return ebiten.KeyAltLeft
	case K_RALT:
		return ebiten.KeyAltRight

	case K_CAPSLOCK:
		return ebiten.KeyCapsLock

	case K_INSERT:
		return ebiten.KeyInsert
	case K_DELETE:
		return ebiten.KeyDelete
	case K_HOME:
		return ebiten.KeyHome
	case K_END:
		return ebiten.KeyEnd
	case K_PAGEUP:
		return ebiten.KeyPageUp
	case K_PAGEDOWN:
		return ebiten.KeyPageDown

	case K_F1:
		return ebiten.KeyF1
	case K_F2:
		return ebiten.KeyF2
	case K_F3:
		return ebiten.KeyF3
	case K_F4:
		return ebiten.KeyF4
	case K_F5:
		return ebiten.KeyF5
	case K_F6:
		return ebiten.KeyF6
	case K_F7:
		return ebiten.KeyF7
	case K_F8:
		return ebiten.KeyF8
	case K_F9:
		return ebiten.KeyF9
	case K_F10:
		return ebiten.KeyF10
	case K_F11:
		return ebiten.KeyF11
	case K_F12:
		return ebiten.KeyF12
	}

	return ebiten.Key(0)
}

func allKeyCodes() []KeyCode {
	return []KeyCode{
		K_A, K_B, K_C, K_D, K_E, K_F,
		K_G, K_H, K_I, K_J, K_K, K_L,
		K_M, K_N, K_O, K_P, K_Q, K_R,
		K_S, K_T, K_U, K_V, K_W, K_X,
		K_Y, K_Z,

		K_0, K_1, K_2, K_3, K_4,
		K_5, K_6, K_7, K_8, K_9,

		K_UP,
		K_DOWN,
		K_LEFT,
		K_RIGHT,

		K_TAB,
		K_RETURN,
		K_ESCAPE,
		K_SPACE,
		K_BACKSPACE,

		K_LSHIFT,
		K_RSHIFT,
		K_LCTRL,
		K_RCTRL,
		K_LALT,
		K_RALT,

		K_CAPSLOCK,

		K_INSERT,
		K_DELETE,
		K_HOME,
		K_END,
		K_PAGEUP,
		K_PAGEDOWN,

		K_F1,
		K_F2,
		K_F3,
		K_F4,
		K_F5,
		K_F6,
		K_F7,
		K_F8,
		K_F9,
		K_F10,
		K_F11,
		K_F12,
	}
}