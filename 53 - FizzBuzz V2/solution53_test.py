import unittest
from io import StringIO
import sys
from solution53 import fizzbuzz

class TestFizzBuzz(unittest.TestCase):
    def test_fizzbuzz(self):
        output = StringIO()
        sys.stdout = output

        fizzbuzz()

        sys.stdout = sys.__stdout__

        result = output.getvalue().strip().split('\n')

        expected = [
            '1', '2', 'fizz', '4', 'buzz', 'fizz', '7', '8', 'fizz', 'buzz',
            '11', 'fizz', '13', '14', 'fizzbuzz', '16', '17', 'fizz', '19', 'buzz'
        ]
        self.assertEqual(result[:20], expected)

if __name__ == '__main__':
    unittest.main()

