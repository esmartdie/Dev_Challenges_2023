import unittest
from solution54 import to_leet_speak

class TestLeetSpeak(unittest.TestCase):
    def test_leet_speak(self):
        self.assertEqual(to_leet_speak("Hello World!"), "#3110 \\/\\/0r1)!")
        self.assertEqual(to_leet_speak("Python is fun!"), "|*j7#0^/ 15 |=(_)^/!")
        self.assertEqual(to_leet_speak("1234567890"), "LREASbTBgo")
        self.assertEqual(to_leet_speak("leet"), "1337")

if __name__ == '__main__':
    unittest.main()
