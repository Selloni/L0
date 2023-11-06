package main

import (
	"fmt"
	"sort"
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
	// забыл про нижний регист, грубое решение, зарее переводим в нижний регистр
	buff := make([]string, len(tmp))
	for i := range tmp {
		buff[i] = strings.ToLower(tmp[i])
	}
	// храню каждое слово и сумму букв по ascii
	countASCII := make(map[string]int32)
	for world := range buff {
		_, ok := countASCII[buff[world]]
		if !ok { // избаляемся от дублей
			for _, i := range buff[world] {
				countASCII[buff[world]] += i // складываем буквы по ascii
			}
		}
	}
	// результирующая мапа
	allMap := fillMyMap(countASCII, buff)
	for k := range allMap {
		if len(allMap[k]) < 2 { // удаляем мапу с единственным значением
			delete(allMap, k)
		}
		sort.Strings(allMap[k]) // сортируем множество
	}
	return allMap
}

func fillMyMap(countASCII map[string]int32, buff []string) map[string][]string {
	myMap := make(map[string][]string)
	for i := range buff {
		// костыль что бы избежать повторения мап
		var asciiNum int32
		_, ok := myMap[buff[i]]
		// запсиываем первое совпадение и удаляем для избежания повторений
		if !ok {
			myMap[(buff[i])] = []string{}
			asciiNum = countASCII[buff[i]]
			delete(countASCII, buff[i])
		}
		// сравниваме по сумме ascii символов, если совпадает значит принадлеит одному подмножеству
		for k, v := range countASCII {
			if asciiNum == v {
				myMap[buff[i]] = append(myMap[buff[i]], k)
				delete(countASCII, k)
			}

		}
	}
	return myMap
}
