package go_string

func IsUpperCase(letter rune) bool {
  return IsAlphabet(letter) && (letter >= 65 && letter <= 90)
}

func IsLowerCase(letter rune) bool {
  return IsAlphabet(letter) && (letter >= 97 && letter <= 122)
}

func UpCase(letter rune) rune {
  if !IsAlphabet(letter) || IsUpperCase(letter) {
    return letter
  }
  return letter - 32
}

func DownCase(letter rune) rune {
  if !IsAlphabet(letter) || IsLowerCase(letter) {
    return letter
  }
  return letter + 32
}
