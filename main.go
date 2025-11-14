package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// 读取角色的ascii艺术
func readCharacter() (string, error) {
	file, err := os.Open("character.txt") // 打开ascii角色配置文件
	if err != nil {                       // 错误检测
		return "", err
	}
	defer file.Close() // 函数结束时关闭文件

	var character string
	scanner := bufio.NewScanner(file) // scanner 可以逐行读取数据
	for scanner.Scan() {
		character += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return character, nil
}

// 读取嘲讽
func readQuotes() ([]string, error) {
	file, err := os.Open("quotes.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var quotes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		quotes = append(quotes, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return quotes, nil
}

// 输出带有ASCII艺术的文本
func charaSay(text, character string) {
	fmt.Println(character)
	fmt.Printf("  %s  \n", strings.Repeat("=", len(text)+4))
	fmt.Printf("  < %s >\n", text)
	fmt.Printf("  %s  \n", strings.Repeat("=", len(text)+4))
}

// 监听用户输入
func listenForQuit() chan bool {
	ch := make(chan bool)

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("Press 'q' to quit:\n")
			input, _ := reader.ReadString('\n')
			if strings.TrimSpace(input) == "q" {
				ch <- true
				return
			}
		}
	}()
	return ch
}

// 每隔五分钟输出学习进度和嘲讽
func studyOrDie() {
	// 读取角色和话语
	character, err := readCharacter()
	if err != nil {
		fmt.Println("Error reading character: ", err)
		return
	}
	quotes, err := readQuotes()
	if err != nil {
		fmt.Println("Error reading quotes: ", err)
		return
	}

	// 每隔一段时间输出进度
	duration := 5 * time.Second
	ticker := time.NewTicker(duration)

	quitCh := listenForQuit()

	//	for range ticker.C {
	//	quote := quotes[rand.Intn(len(quotes))]
	// charaSay(quote, character)
	// }
	for {
		select {
		case <-ticker.C:
			quote := quotes[rand.Intn(len(quotes))]
			charaSay(quote, character)
		case <-quitCh:
			fmt.Println("Exiting study-or-die... Good luck!")
			return
		}
	}
}

func main() {
	// 初始化随机种子
	// rand.Seed(time.Now().Unix())

	// 开始执行study-or-die函数
	fmt.Println("Starting study-or-die... Stay focused!")
	studyOrDie()
}
