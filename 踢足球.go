package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("欢迎来到 Go 语言踢足球游戏！")
	fmt.Println("你需要选择一个方向踢球（左、中、右），然后尝试进球。")

	score := 0
	attempts := 5

	directions := []string{"左", "中", "右"}

	for i := 1; i <= attempts; i++ {
		fmt.Printf("\n第 %d 次射门（输入 0-左, 1-中, 2-右）: ", i)
		var choice int
		fmt.Scan(&choice)

		if choice < 0 || choice > 2 {
			fmt.Println("无效选择，请输入 0, 1 或 2")
			i-- // 保持尝试次数不变
			continue
		}

		// 模拟守门员随机扑救方向
		goalkeeper := rand.Intn(3)

		fmt.Printf("你选择了 %s，守门员扑向了 %s。\n", directions[choice], directions[goalkeeper])

		if choice == goalkeeper {
			fmt.Println("很遗憾，守门员扑到了球！未进球。")
		} else {
			fmt.Println("射门得分！")
			score++
		}
	}

	fmt.Printf("\n比赛结束！你总共进了 %d 球。\n", score)
}
