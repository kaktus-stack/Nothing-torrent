package ui

import "strings"

// RuToEnMap maps Russian Cyrillic keyboard characters to their QWERTY physical equivalents
var RuToEnMap = map[rune]rune{
	// Top row
	'й': 'q', 'Й': 'q',
	'ц': 'w', 'Ц': 'w',
	'у': 'e', 'У': 'e',
	'к': 'r', 'К': 'r',
	'е': 't', 'Е': 't',
	'н': 'y', 'Н': 'y',
	'г': 'u', 'Г': 'u',
	'ш': 'i', 'Ш': 'i',
	'щ': 'o', 'Щ': 'o',
	'з': 'p', 'З': 'p',
	'х': '[', 'Х': '[',
	'ъ': ']', 'Ъ': ']',

	// Middle row
	'ф': 'a', 'Ф': 'a',
	'ы': 's', 'Ы': 's',
	'в': 'd', 'В': 'd',
	'а': 'f', 'А': 'f',
	'п': 'g', 'П': 'g',
	'р': 'h', 'Р': 'h',
	'о': 'j', 'О': 'j',
	'л': 'k', 'Л': 'k',
	'д': 'l', 'Д': 'l',
	'ж': ';', 'Ж': ';',
	'э': '\'', 'Э': '\'',

	// Bottom row
	'я': 'z', 'Я': 'z',
	'ч': 'x', 'Ч': 'x',
	'с': 'c', 'С': 'c',
	'м': 'v', 'М': 'v',
	'и': 'b', 'И': 'b',
	'т': 'n', 'Т': 'n',
	'ь': 'm', 'Ь': 'm',
	'б': ',', 'Б': ',',
	'ю': '.', 'Ю': '.',
}

// NormalizeKey converts layout-dependent key input (e.g. Cyrillic) into standard QWERTY action keys
func NormalizeKey(key string) string {
	if len(key) == 0 {
		return key
	}

	runes := []rune(key)
	if len(runes) == 1 {
		if en, ok := RuToEnMap[runes[0]]; ok {
			return string(en)
		}
	}

	return strings.ToLower(key)
}
