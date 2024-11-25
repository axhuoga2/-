package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

// Game 定义游戏结构体
type Game struct {
	x, y float64 // 玩家位置
}

// Update 更新游戏逻辑
func (g *Game) Update() error {
	// 玩家移动
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.y += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.x -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.x += 2
	}
	return nil
}

// Draw 绘制游戏界面
func (g *Game) Draw(screen *ebiten.Image) {
	// 填充背景
	screen.Fill([4]uint8{240, 240, 240, 255})

	// 画玩家（一个简单的矩形代表角色）
	ebitenutil.DrawRect(screen, g.x, g.y, 50, 50, [4]uint8{255, 0, 0, 255})

	// 显示提示文字
	ebitenutil.DebugPrint(screen, "使用方向键移动玩家")
}

// Layout 设置屏幕尺寸
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func main() {
	// 初始化游戏实例
	game := &Game{x: 375, y: 275}

	// 启动游戏
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("简单蛋仔派对原型")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
