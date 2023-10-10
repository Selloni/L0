package main

import (
	"fmt"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	ss := []string{"пятак", "пятка", "тяПка", "тяпка", "тЕрка", "терк", "Листок", "сЛИток", "столик", "лист", "литс"}
	fmt.Println(findAnagramme(ss))
}

func findAnagramme(tmp []string) map[string][]string {
	buff := make([]string, len(tmp))
	for i := range tmp {
		buff[i] = tmp[i]
	}
	countAscii := make(map[string]int32)
	for world, _ := range buff {
		_, ok := countAscii[buff[world]]
		if !ok {
			for _, i := range buff[world] {
				countAscii[buff[world]] += i
			}
		}
	}
	allMap := fillMyMap(countAscii, buff)
	for k, _ := range allMap {
		if len(allMap[k]) < 2 {
			delete(allMap, k)
		}
	}
	return allMap
}

func fillMyMap(countAscii map[string]int32, buff []string) map[string][]string {
	myMap := make(map[string][]string)
	for i, _ := range buff {
		var asciiNum int32
		_, ok := myMap[buff[i]]
		if !ok {
			myMap[strings.ToUpper(buff[i])] = []string{}
			asciiNum = countAscii[buff[i]]
			delete(countAscii, buff[i])
		}
		for k, v := range countAscii {
			if asciiNum == v {
				myMap[buff[i]] = append(myMap[buff[i]], strings.ToUpper(k))
				delete(countAscii, k)
			}
		}
	}
	return myMap
}
