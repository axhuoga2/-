package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	// 随机生成一个目标字母 (A-Z)
	target := 'A' + rune(rand.Intn(26))

	fmt.Println("欢迎来到猜字母游戏！")
	fmt.Println("我已经想好了一个字母（A-Z）。试着猜猜看吧！")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("请输入一个字母: ")
		input, _ := reader.ReadString('\n')

		// 去掉输入末尾的换行符
		input = input[:len(input)-1]

		// 检查输入是否有效
		if len(input) != 1 || input[0] < 'A' || input[0] > 'Z' {
			fmt.Println("请输入一个有效的大写字母 (A-Z)。")
			continue
		}

		guess := rune(input[0])

		if guess < target {
			fmt.Println("你猜的字母太小了！")
		} else if guess > target {
			fmt.Println("你猜的字母太大了！")
		} else {
			fmt.Println("恭喜你，猜对了！")
			break
		}
	}
}
