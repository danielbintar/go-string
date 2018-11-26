package go_string

func IsAlphabet(letter rune) bool {
  return (letter >= 65 && letter <= 90) || (letter >= 97 && letter <= 122)
}
