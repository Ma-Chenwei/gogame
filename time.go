package gogame

import (
	"sync"
	"time"
)

// Clock 对应 pygame.time.Clock。
type ClockObject struct {
	mu sync.Mutex

	lastTick time.Time
	lastFPS  float64
}

// TimeModule 对应 pygame.time。
type TimeModule struct{}

var Time = &TimeModule{}

// Clock 创建一个时钟。
// 对应:
//
//	clock = pygame.time.Clock()
func (t *TimeModule) Clock() *ClockObject {
	return &ClockObject{
		lastTick: time.Now(),
	}
}

// Tick 限制最大帧率。
// 对应:
//
//	clock.tick(60)
//
// 返回本次 Tick 实际经过的毫秒数。
func (c *ClockObject) Tick(fps int) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	if fps <= 0 {
		now := time.Now()

		ms := int(now.Sub(c.lastTick).Milliseconds())

		c.lastTick = now

		if ms <= 0 {
			ms = 1
		}

		c.lastFPS = 1000.0 / float64(ms)

		return ms
	}

	frameDuration := time.Second / time.Duration(fps)

	now := time.Now()
	elapsed := now.Sub(c.lastTick)

	if elapsed < frameDuration {
		time.Sleep(frameDuration - elapsed)
	}

	now = time.Now()

	ms := int(now.Sub(c.lastTick).Milliseconds())

	c.lastTick = now

	if ms <= 0 {
		ms = 1
	}

	c.lastFPS = 1000.0 / float64(ms)

	return ms
}

// GetFPS 获取最近一帧计算出来的 FPS。
// 对应 pygame.time.Clock.get_fps()。
func (c *ClockObject) GetFPS() float64 {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.lastFPS
}

// Delay 延迟指定毫秒。
// 对应:
//
//	pygame.time.delay(100)
func (t *TimeModule) Delay(milliseconds int) {
	if milliseconds <= 0 {
		return
	}

	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}

// Wait 延迟指定毫秒。
// 与 Delay 相同，提供额外兼容名称。
func (t *TimeModule) Wait(milliseconds int) {
	t.Delay(milliseconds)
}

// GetTicks 返回 GoGame 初始化之后经过的毫秒数。
// 对应 pygame.time.get_ticks()。
var startTime = time.Now()

func (t *TimeModule) GetTicks() int {
	return int(time.Since(startTime).Milliseconds())
}