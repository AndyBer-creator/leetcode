// Сравнение строк с использованием регулярных выражений
package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	// Проверка пустой строки и шаблона
	if p == "" {
		return s == ""
	}

	// Определите, совпадает ли первый символ строки с первым символом
	// шаблона
	firstMatch := len(s) > 0 && (p[0] == s[0] || p[0] == '.')

	// Обработка шаблона с '*' (например, 'a*' или '.*')
	if len(p) >= 2 && p[1] == '*' {
		// Рассмотрим два случая:
		// 1. пропустить '*' и предыдущий символ (например, 'a*' 
		//превращается в '')
		// 2. продолжить сопоставление, если есть совпадение
		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
	} else {
		// Если нет '*' просто проверить совпадение первого символа и 
		//рекурсивно перейти к остальным
		return firstMatch && isMatch(s[1:], p[1:])
	}
}

func main() {
	// Примеры использования
	fmt.Println(isMatch("aa", "a"))                   // false
	fmt.Println(isMatch("aa", "a*"))                  // true
	fmt.Println(isMatch("ab", ".*"))                  // true
	fmt.Println(isMatch("aab", "c*a*b"))              // true
	fmt.Println(isMatch("mississippi", "mis*is*p*.")) // false
}
