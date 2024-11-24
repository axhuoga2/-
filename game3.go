package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// 预定义单词列表
	words := []string{"apple", "grape", "orange", "banana", "cherry", "mango", "lemon", "peach"}

	// 初始化随机种子并选择一个随机单词
	rand.Seed(time.Now().UnixNano())
	target := words[rand.Intn(len(words))]

	fmt.Println("欢迎来到猜英语单词游戏！")
	fmt.Println("我已经想好了一个单词，你可以试着猜猜看！（提示：全是水果名）")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("请输入你的猜测（小写单词）: ")
		guess, _ := reader.ReadString('\n')
		guess = strings.TrimSpace(guess) // 去掉多余的换行和空格

		// 判断是否猜对
		if guess == target {
			fmt.Println("恭喜你，猜对了！答案是:", target)
			break
		}

		// 提示字母匹配情况
		matching := compareWords(target, guess)
		fmt.Printf("你猜的单词不正确，有 %d 个字母在正确的位置。\n", matching)
	}
}

// compareWords 比较两个单词并返回字母完全匹配的位置数
func compareWords(target, guess string) int {
	count := 0
	for i := 0; i < len(target) && i < len(guess); i++ {
		if target[i] == guess[i] {
			count++
		}
	}
	retu
