"""
 * Escribe un programa que reciba un texto y transforme lenguaje natural a
 * "lenguaje hacker" (conocido realmente como "leet" o "1337"). Este lenguaje
 *  se caracteriza por sustituir caracteres alfanuméricos.
 * - Utiliza esta tabla (https://www.gamehouse.com/blog/leet-speak-cheat-sheet)
 *   con el alfabeto y los números en "leet".
 *   (Usa la primera opción de cada transformación. Por ejemplo "4" para la "a")
 """

LEET_DICT = {
    'a': '4', 'b': 'I3', 'c': '[', 'd': ')', 'e': '3', 'f': '|=', 'g': '6', 'h': '#',
    'i': '1', 'j': ',_|', 'k': '>|', 'l': '1', 'm': '/\\/\\', 'n': '^/', 'o': '0', 'p': '|*',
    'q': '(_,)', 'r': 'r', 's': '5', 't': '7', 'u': '(_)', 'v': '\\/', 'w': '\\/\\/', 'x': '><',
    'y': 'j', 'z': '2',
    '0': 'o', '1': 'L', '2': 'R', '3': 'E', '4': 'A', '5': 'S', '6': 'b', '7': 'T', '8': 'B', '9': 'g'
}

def to_leet_speak(text):
    return ''.join(LEET_DICT.get(char.lower(), char) for char in text)

if __name__ == "__main__":
    sample_text = "Hello World! This is an example."
    print(to_leet_speak(sample_text))