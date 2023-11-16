package songs

import "fmt"

const Chorus = 0

type SongText struct {
	Tonality Tonality `json:"tonality"`
	Stances  []Stance `json:"stances"`
}

type StanceType int

type Stance struct {
	Text []Word     `json:"verses"`
	Type StanceType `json:"type"`
}

type Word []Syllable

type Syllable struct {
	Text  string `json:"syllable"`
	Chord *Chord `json:"chord"`
}

type Chord struct {
	Root    Note   `json:"root"`
	Base    *Note  `json:"base"`
	Details string `json:"details"`
}

// Chord String
func (c Chord) String() string {
	if c.Base == nil {
		return fmt.Sprintf("%v%s", c.Root, c.Details)
	}
	return fmt.Sprintf("%v%s/%v", c.Root, c.Details, c.Base)
}

// Transposition chain
func (s *SongText) TransposeAllStances(interval Interval) {
	for i := 0; i < len(s.Stances); i++ {
		s.Stances[i].transposeAllWords(interval)
	}
}

func (s *Stance) transposeAllWords(interval Interval) {
	for i := 0; i < len(s.Text); i++ {
		s.Text[i].transposeAllChords(interval)
	}
}

func (w *Word) transposeAllChords(interval Interval) {
	for i := 0; i < len(*w); i++ {
		(*w)[i].transposeChord(interval)
	}
}

func (s *Syllable) transposeChord(i Interval) {
	s.Chord.Transpose(i)
}

func (c *Chord) Transpose(i Interval) {
	if c.Base == nil {
		c.Base.Transpose(i)
	}
	c.Root.Transpose(i)
}
