package main

import (
	"fmt"
)

func reverseString(s string) string {
	// Преобразуем строку в срез руг, чтобы учесть все символы (включая многобайтовые)
	runes := []rune(s)

	// Используем два указателя: с начала и с конца среза
	left, right := 0, len(runes)-1

	// Переворачиваем срез
	for left < right {
		// Меняем местами символы на позициях left и right
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	// Преобразуем обратно в строку и возвращаем
	return string(runes)
}

func main() {
	// Пример ввода
	input := "wb11.12.2024"
	// Переворачиваем строку
	reversed := reverseString(input)
	// Выводим результат
	fmt.Println(reversed) // "
}
