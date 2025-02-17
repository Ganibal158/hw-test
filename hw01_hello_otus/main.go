package main

import (
	"fmt"
)

func main() {
	fmt.Println(String("Hello, OTUS!"))
}

func String(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// пакет "golang.org/x/example/hello/reverse" не хотел подключаться "could not import golang.org/x/example/hello/reverse (no required module provides package "golang.org/x/example/hello/reverse")compilerBrokenImport" И подчёркивался красным, пришлось засунуть функцию сюда. Подскажите, пожалуйста, что можно было сделать?
