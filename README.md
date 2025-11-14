# study-or-die

一个轻量级、有趣且略带恶意的终端学习计时器。
每隔固定时间输出一条 嘲讽语录 + 角色 ASCII 图，提醒你继续学习。
支持按 q 随时退出。

## 功能特性

- 定时提醒：默认每 5 秒钟输出一次提示，将来会添加自定义间隔时间功能。

- ASCII 角色支持：从 character.txt 加载角色图案，将来会添加多角色轮流输出功能。

- 嘲讽语录：从 quotes.txt 随机选择一句输出。

## 项目结构

study-or-die/
├── main.go          # 主程序
├── character.txt    # ASCII 角色
├── quotes.txt       # 嘲讽语录列表
└── go.mod

## 使用方法

1. 克隆项目
git clone <https://github.com/Michaelwu0905/study-or-die.git>
cd study-or-die

2. 运行程序
go run main.go

或构建可执行文件：

go build -o study
./study

## 自定义角色

编辑 character.txt，可以在各种ascii艺术网站找到好看的角色
里面的内容会作为输出角色展示，例如：

 (\_/)
 ( •_•)
 / >🍪

## 自定义嘲讽语录

编辑 quotes.txt
每行写一句，例如：

五分钟过去了，你干了什么？
别停下来，现在休息会后悔的。
继续卷，别让别人超过你。

## 如何退出？

运行过程中随时按q即可退出程序。
