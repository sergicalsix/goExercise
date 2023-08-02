package mylib

import (
	"fmt"
)

// エクスポートする場合は大文字スタートのキャメルケース
func GoodEvening() {
	var greeting string = "Hello"
	const World string = "World"
	greeting = "Good evening"
	fmt.Println(greeting, World)
}
