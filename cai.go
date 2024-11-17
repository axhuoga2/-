package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// 生成一个 1 到 100 的随机数
	target := rand.Intn(100) + 1

	fmt.Println("欢迎来到猜数字游戏！")
	fmt.Println("我已经选择了一个 1 到 100 之间的数字。")

	var guess int
	attempts := 0

	// 游戏循环，直到猜对为止
	for {
		// 提示玩家输入数字
		fmt.Print("请输入你的猜测: ")
		fmt.Scan(&guess)
		attempts++

		// 检查玩家的猜测
		if guess < target {
			fmt.Println("太小了！再试一次。")
		} else if guess > target {
			fmt.Println("太大了！再试一次。")
		} else {
			fmt.Printf("恭喜你，猜对了！你用了 %d 次机会。\n", attempts)
			break
		}
	}
}
