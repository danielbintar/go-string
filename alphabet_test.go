package go_string

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

type AlphabetCaseTest struct {
  letter rune
  result bool
}

func TestIsAlphabet(t *testing.T) {
  caseTests := []*AlphabetCaseTest{
    &AlphabetCaseTest{letter: 'a', result: true},
    &AlphabetCaseTest{letter: 'z', result: true},
    &AlphabetCaseTest{letter: 'A', result: true},
    &AlphabetCaseTest{letter: 'Z', result: true},
    &AlphabetCaseTest{letter: '9', result: false},
  }

  for _, caseTest := range caseTests {
    assert.Equal(t, IsAlphabet(caseTest.letter), caseTest.result)
  }
}
