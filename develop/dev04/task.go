package main

import (
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

func FindAnagrams(words *[]string) *map[string][]string {
	anagrams := make(map[string][]string)
	sortedWords := make(map[string]string)

	for _, word := range *words {
		// Приводим слово к нижнему регистру
		word = strings.ToLower(word)

		// Преобразуем слово в отсортированную строку
		sortedWord := sortString(word)

		// Если первый раз встречаем отсортированное слово, то добавляем его в мапу sortedWords
		//и создаем новую группу анаграмм
		if firstWord, found := sortedWords[sortedWord]; !found {
			sortedWords[sortedWord] = word
			anagrams[word] = append(anagrams[word], word)
		} else {
			// Это анаграмма, добавляем ее в соответствующую группу
			anagrams[firstWord] = append(anagrams[firstWord], word)
		}
	}

	// Удаляем группы с одним элементом или пустые ключи
	for key, group := range anagrams {
		if len(group) <= 1 || key == "" {
			delete(anagrams, key)
		} else {
			// Сортируем элементы в группе
			sort.Strings(group)
		}
	}

	return &anagrams
}

// Функция для сортировки символов в строке
func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagramGroups := FindAnagrams(&words)

	// Вывод результатов
	for key, group := range *anagramGroups {
		println(key, ":", strings.Join(group, ", "))
	}
}
