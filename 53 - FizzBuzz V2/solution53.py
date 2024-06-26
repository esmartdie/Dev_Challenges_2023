"""
 * Escribe un programa que muestre por consola (con un print) los
 * números de 1 a 100 (ambos incluidos y con un salto de línea entre
 * cada impresión), sustituyendo los siguientes:
 * - Múltiplos de 3 por la palabra "fizz".
 * - Múltiplos de 5 por la palabra "buzz".
 * - Múltiplos de 3 y de 5 a la vez por la palabra "fizzbuzz".
 
 """

def fizzbuzz():
    def get_value(number):
        return "fizzbuzz" if number % 3 == 0 and number % 5 == 0 else \
               "fizz" if number % 3 == 0 else \
               "buzz" if number % 5 == 0 else \
               str(number)
    
    for i in range(1, 101):
        print(get_value(i))

if __name__ == "__main__":
    fizzbuzz()