package main

import (
	"fmt"
)

func Names() (string, string) { // простая функция,
	return "Foo", "Bar" // возвращающая два строковых значения
}

func main() {
	n1, n2 := Names() // приваивание двум переменным по порядку двух значчениий
	fmt.Println(n1, n2)
	n1, n2 = n2, n1
	fmt.Println(n1, n2)

	n3, _ := Names() // использование _ в качестве никогда не используемой переменной
	fmt.Println(n3)
	// n4  := Names() а такая запись будет ошибкой. Names возвращает два значения,
	// ни как не утискивающиеся в одну переменную
}
