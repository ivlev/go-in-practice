/*
1. Возвращение заранее объявленых переменных
2. Попытка вернуть из функции одну переменную
3. Вывод об использовании переменных в golang ( спойлеров не будет, раскомментируйте и попробуйте )
Matt Butcher, Matt Farina, "Go in Practice". Chapter 1
*/

package main

import (
	"fmt"
)

func Names() (first string, second string) { // Переменные для возврата описываются в объявлении функции
	first = "Foo"  // собственно
	second = "Bar" // объявление пременных, обратите внимание на "=" а не на использование ":="
	return         // возвращаем всё! так бы я описал этот оператор
}

func main() {
	n1, n2 := Names()   // переприсваивание пременных в основной функции
	fmt.Println(n1, n2) // печать извлеченных переменных
}

/*
// для лучшего понимания примера стоит попробовать следуюй код

func Names() (first string) { // Мы говорим что будем извлекать только переменную "first" типа string
	first = "Foo" // так же задаем
	second = "Bar" // две переменных
	return // пытаемся вернуть одну пременную
}

func main() {
	n1 := Names() // тащим одну переменную
	fmt.Println(n1) // печать извлеченную переменную
}

*/
