package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// 读取角色的ascii艺术
func readCharacters() ([]string, error) {
	file, err := os.Open("character.txt") // 打开ascii角色配置文件
	if err != nil {                       // 错误检测
		return nil, err
	}
	defer file.Close() // 函数结束时关闭文件

	var characters []string
	var character string
	scanner := bufio.NewScanner(file) // scanner 可以逐行读取数据
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// 空行表示一个角色的结束，保存当前角色
			if character != "" {
				characters = append(characters, character)
				character = "" // 重置当前角色
			}
		} else {
			// 否则将行追加到当前角色
			character += line + "\n"
		}
	}

	// 将最后一个角色添加到列表
	if character != "" {
		characters = append(characters, character)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(characters) == 0 {
		return nil, fmt.Errorf("characters not found")
	}
	return characters, nil
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
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			quotes = append(quotes, line)
		}
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}
	if len(quotes) == 0 {
		return nil, fmt.Errorf("quotes not found")
	}

	return quotes, nil
}

// 输出带有ASCII艺术的文本
func charaSay(text, character string) {
	// fmt.Println(character)
	// fmt.Printf("  %s  \n", strings.Repeat("=", len(text)+4))
	// fmt.Printf("  < %s >\n", text)
	// fmt.Printf("  %s  \n", strings.Repeat("=", len(text)+4))
	// 生成简单对话框
	width := len([]rune(text))
	border := strings.Repeat("-", width+4)

	fmt.Println(border)
	fmt.Printf("| %s\n", text)
	fmt.Println(border)
	fmt.Print(character)
	fmt.Println()
}

// 监听用户输入
func listenForQuit() <-chan struct{} {
	ch := make(chan struct{})

	go func() {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Press 'q' and Enter to quit at any time.")
		for {
			input, err := reader.ReadString('\n')
			if err != nil {
				return
			}
			if strings.TrimSpace(input) == "q" {
				ch <- struct{}{}
				return
			}
		}
	}()
	return ch
}

// 每隔五分钟输出学习进度和嘲讽
func studyOrDie(totalDuration time.Duration, freq time.Duration) {
	// 读取角色和话语
	characters, err := readCharacters()
	if err != nil {
		fmt.Println("Error reading character: ", err)
		return
	}
	quotes, err := readQuotes()
	if err != nil {
		fmt.Println("Error reading quotes: ", err)
		return
	}

	// 定时器：整体运行时间
	timer := time.NewTimer(totalDuration)
	//	for range ticker.C {
	//	quote := quotes[rand.Intn(len(quotes))]
	// charaSay(quote, character)

	ticker := time.NewTicker(freq)
	defer ticker.Stop()
	defer timer.Stop()

	quitCh := listenForQuit()
	// 两个计数器：一个控制角色，一个控制语录
	charIndex := 0
	quoteIndex := 0
	for {
		select {
		case <-ticker.C:
			// 每次输出下一个角色和下一个语录
			currentChar := characters[charIndex%len(characters)]
			currentQuote := quotes[quoteIndex%len(characters)]
			charaSay(currentQuote, currentChar)

			// 递增 （环状）
			charIndex++
			quoteIndex++
		case <-timer.C:
			fmt.Println("Time is up! Exiting study-or-die... Have a good day :)")
			return
		case <-quitCh:
			fmt.Println("Exiting study-or-die... Good luck!")
			return
		}
	}
}

func main() {
	// 初始化随机种子
	// rand.Seed(time.Now().Unix())
	// -t 总时长（分钟）-f 发送频率（秒）
	minutes := flag.Int("t", 5, "total run time in minutes")
	seconds := flag.Int("f", 60, "message frequency in seconds")
	flag.Parse()

	if *minutes <= 0 {
		fmt.Println("Invalid -t value, using default 5 minutes")
	}
	if *seconds <= 0 {
		fmt.Println("Invalid -f value, using default 60 seconds:")
		*seconds = 60
	}
	totalDuration := time.Duration(*minutes) * time.Minute
	freq := time.Duration(*seconds) * time.Second

	fmt.Printf("Running study-or-die for %d minutes, every %d seconds...\n", *minutes, *seconds)

	// 开始执行study-or-die函数
	studyOrDie(totalDuration, freq)
}
