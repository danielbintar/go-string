package go_string

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

type CaseCaseTest struct {
  letter rune
  result bool
}

type TransformCaseCaseTest struct {
  letter rune
  result rune
}

func TestIsUpperCase(t *testing.T) {
  caseTests := []*CaseCaseTest{
    &CaseCaseTest{letter: 'a', result: false},
    &CaseCaseTest{letter: 'z', result: false},
    &CaseCaseTest{letter: 'A', result: true},
    &CaseCaseTest{letter: 'Z', result: true},
    &CaseCaseTest{letter: '9', result: false},
  }

  for _, caseTest := range caseTests {
    assert.Equal(t, IsUpperCase(caseTest.letter), caseTest.result)
  }
}

func TestIsLowerCase(t *testing.T) {
  caseTests := []*CaseCaseTest{
    &CaseCaseTest{letter: 'a', result: true},
    &CaseCaseTest{letter: 'z', result: true},
    &CaseCaseTest{letter: 'A', result: false},
    &CaseCaseTest{letter: 'Z', result: false},
    &CaseCaseTest{letter: '9', result: false},
  }

  for _, caseTest := range caseTests {
    assert.Equal(t, IsLowerCase(caseTest.letter), caseTest.result)
  }
}

func TestUpCase(t *testing.T) {
  caseTests := []*TransformCaseCaseTest{
    &TransformCaseCaseTest{letter: 'a', result: 'A'},
    &TransformCaseCaseTest{letter: 'z', result: 'Z'},
    &TransformCaseCaseTest{letter: 'A', result: 'A'},
    &TransformCaseCaseTest{letter: 'Z', result: 'Z'},
    &TransformCaseCaseTest{letter: '9', result: '9'},
  }

  for _, caseTest := range caseTests {
    assert.Equal(t, UpCase(caseTest.letter), caseTest.result)
  }
}

func TestDownCase(t *testing.T) {
  caseTests := []*TransformCaseCaseTest{
    &TransformCaseCaseTest{letter: 'a', result: 'a'},
    &TransformCaseCaseTest{letter: 'z', result: 'z'},
    &TransformCaseCaseTest{letter: 'A', result: 'a'},
    &TransformCaseCaseTest{letter: 'Z', result: 'z'},
    &TransformCaseCaseTest{letter: '9', result: '9'},
  }

  for _, caseTest := range caseTests {
    assert.Equal(t, DownCase(caseTest.letter), caseTest.result)
  }
}
