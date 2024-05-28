/*
 * Escribe un programa que reciba un texto y transforme lenguaje natural a
 * "lenguaje hacker" (conocido realmente como "leet" o "1337"). Este lenguaje
 *  se caracteriza por sustituir caracteres alfanuméricos.
 * - Utiliza esta tabla (https://www.gamehouse.com/blog/leet-speak-cheat-sheet)
 *   con el alfabeto y los números en "leet".
 *   (Usa la primera opción de cada transformación. Por ejemplo "4" para la "a")
 */

package main

import (
	"fmt"
	"strings"
)

var hackerMap map[string]string

func init() {
	hackerMap = make(map[string]string)
	hackerMap["a"] = "4"
	hackerMap["b"] = "I3"
	hackerMap["c"] = "["
	hackerMap["d"] = ")"
	hackerMap["e"] = "3"
	hackerMap["f"] = "|="
	hackerMap["g"] = "6"
	hackerMap["h"] = "#"
	hackerMap["i"] = "1"
	hackerMap["j"] = ",_|"
	hackerMap["k"] = ">|"
	hackerMap["l"] = "1"
	hackerMap["m"] = "/\\/\\"
	hackerMap["n"] = "^/"
	hackerMap["o"] = "0"
	hackerMap["p"] = "|*"
	hackerMap["q"] = "(_,)"
	hackerMap["r"] = "r"
	hackerMap["s"] = "5"
	hackerMap["t"] = "7"
	hackerMap["u"] = "(_)"
	hackerMap["v"] = "\\/"
	hackerMap["w"] = "\\/\\/"
	hackerMap["x"] = "><"
	hackerMap["y"] = "j"
	hackerMap["z"] = "2"
	hackerMap["0"] = "o"
	hackerMap["1"] = "L"
	hackerMap["2"] = "R"
	hackerMap["3"] = "E"
	hackerMap["4"] = "A"
	hackerMap["5"] = "S"
	hackerMap["6"] = "b"
	hackerMap["7"] = "T"
	hackerMap["8"] = "B"
	hackerMap["9"] = "g"
}

func hackerProcessing(s string) string {
	s = strings.ToLower(s)
	var builder strings.Builder

	for _, char := range s {
		value, exists := hackerMap[string(char)]
		if exists {
			builder.WriteString(value)
		} else {
			builder.WriteRune(char)
		}
	}

	return builder.String()
}

func main() {
	original := "Hello World! This is an example text to convert to leet speak."
	process := hackerProcessing(original)
	fmt.Println("Hacker code: ", process)
}
