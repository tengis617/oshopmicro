package loremsvc

import (
	"github.com/drhodes/golorem"
)

type Service interface {
	Word(min, max int) string

	Sentence(min, max int) string

	Paragraph(min, max int) string
}

type LoremService struct{}

// Word generates a random word with at least min letters and at most max letters.
func (LoremService) Word(min, max int) string {
	return lorem.Word(min, max)
}

// Sentence generates a sentence with at least min words and at most max words.
func (LoremService) Sentence(min, max int) string {
	return lorem.Sentence(min, max)
}

// Paragraph generates a paragraph with at least min sentences and at most max sentences.
func (LoremService) Paragraph(min, max int) string {
	return lorem.Paragraph(min, max)
}
