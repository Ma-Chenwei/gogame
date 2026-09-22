package gogame

import (
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

// EventData 是单个事件。
// 对应 Pygame 中的 event 对象。
type EventData struct {
	Type   int
	Key    KeyCode
	Button int
	Pos    [2]int
	Rel    [2]int
}

// EventModule 对应 pygame.event。
type EventModule struct{}

var Event = &EventModule{}

var (
	eventMutex sync.Mutex
	eventQueue []EventData

	previousKeys = make(map[KeyCode]bool)
	currentKeys  = make(map[KeyCode]bool)

	previousMouseButtons [3]bool
	currentMouseButtons  [3]bool

	previousMousePos [2]int
	currentMousePos  [2]int
)

// initEvents 初始化事件系统。
func initEvents() {
	eventMutex.Lock()
	defer eventMutex.Unlock()

	eventQueue = nil

	previousKeys = make(map[KeyCode]bool)
	currentKeys = make(map[KeyCode]bool)

	previousMouseButtons = [3]bool{}
	currentMouseButtons = [3]bool{}

	previousMousePos = [2]int{}
	currentMousePos = [2]int{}
}

// Get 获取当前所有事件。
//
// Pygame:
//
//	for event in pygame.event.get():
//
// Go:
//
//	for _, event := range gogame.Event.Get() {
//		...
//	}
func (e *EventModule) Get() []EventData {
	eventMutex.Lock()
	defer eventMutex.Unlock()

	result := make([]EventData, len(eventQueue))
	copy(result, eventQueue)

	eventQueue = eventQueue[:0]

	return result
}

// Poll 获取一个事件。
//
// 没有事件时返回 NOEVENT。
func (e *EventModule) Poll() EventData {
	eventMutex.Lock()
	defer eventMutex.Unlock()

	if len(eventQueue) == 0 {
		return EventData{
			Type: NOEVENT,
		}
	}

	result := eventQueue[0]

	eventQueue = eventQueue[1:]

	return result
}

// Clear 清空所有事件。
func (e *EventModule) Clear() {
	eventMutex.Lock()
	eventQueue = eventQueue[:0]
	eventMutex.Unlock()
}

// pushEvent 添加事件。
func pushEvent(event EventData) {
	eventMutex.Lock()
	eventQueue = append(eventQueue, event)
	eventMutex.Unlock()
}

// updateEvents 每帧更新输入事件。
func updateEvents() {
	updateKeyboardEvents()
	updateMouseEvents()
}

// updateKeyboardEvents 更新键盘事件。
func updateKeyboardEvents() {
	for _, key := range allKeyCodes() {
		ebitenKey := toEbitenKey(key)

		if ebitenKey == ebiten.Key(0) {
			continue
		}

		pressed := ebiten.IsKeyPressed(ebitenKey)

		currentKeys[key] = pressed

		// KEYDOWN
		if pressed && !previousKeys[key] {
			pushEvent(EventData{
				Type: KEYDOWN,
				Key:  key,
			})
		}

		// KEYUP
		if !pressed && previousKeys[key] {
			pushEvent(EventData{
				Type: KEYUP,
				Key:  key,
			})
		}

		previousKeys[key] = pressed
	}
}

// updateMouseEvents 更新鼠标事件。
func updateMouseEvents() {
	x, y := ebiten.CursorPosition()

	currentMousePos = [2]int{
		x,
		y,
	}

	// 鼠标移动
	if currentMousePos != previousMousePos {
		pushEvent(EventData{
			Type: MOUSEMOTION,

			Pos: currentMousePos,

			Rel: [2]int{
				currentMousePos[0] - previousMousePos[0],
				currentMousePos[1] - previousMousePos[1],
			},
		})

		previousMousePos = currentMousePos
	}

	buttons := []ebiten.MouseButton{
		ebiten.MouseButtonLeft,
		ebiten.MouseButtonMiddle,
		ebiten.MouseButtonRight,
	}

	for i, button := range buttons {
		pressed := ebiten.IsMouseButtonPressed(button)

		currentMouseButtons[i] = pressed

		// 鼠标按下
		if pressed && !previousMouseButtons[i] {
			pushEvent(EventData{
				Type: MOUSEBUTTONDOWN,

				Button: i + 1,

				Pos: currentMousePos,
			})
		}

		// 鼠标释放
		if !pressed && previousMouseButtons[i] {
			pushEvent(EventData{
				Type: MOUSEBUTTONUP,

				Button: i + 1,

				Pos: currentMousePos,
			})
		}

		previousMouseButtons[i] = pressed
	}
}