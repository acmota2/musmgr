package songs

const (
	Chorus = 0
)

type Chord interface {
	Note
	Transpose(Interval)
}

type Song[T Chord] struct {
	Tonality Tonality    `json:"tonality"`
	Stances  []Stance[T] `json:"stances"`
}

type StanceType int

type Stance[T Chord] struct {
	Text []Word[T]  `json:"verses"`
	Type StanceType `json:"type"`
}

type Word[T Chord] []Syllable[T]

type Syllable[T Chord] struct {
	Text  string `json:"syllable"`
	Chord T      `json:"chord"`
}

type SimpleChord struct {
	Root    Note   `json:"root"`
	Details string `json:"details"`
}

type SlashChord struct {
	Root    Note   `json:"root"`
	Base    Note   `json:"base"`
	Details string `json:"details"`
}

// Transposition chain
func (s *Song[T]) TransposeAllStances(interval Interval) {
	for i := 0; i < len(s.Stances); i++ {
		s.Stances[i].transposeAllWords(interval)
	}
}

func (s *Stance[T]) transposeAllWords(interval Interval) {
	for i := 0; i < len(s.Text); i++ {
		s.Text[i].transposeAllChords(interval)
	}
}

func (w *Word[T]) transposeAllChords(interval Interval) {
	for i := 0; i < len(*w); i++ {
		(*w)[i].transposeChord(interval)
	}
}

func (s *Syllable[T]) transposeChord(i Interval) {
	s.Chord.Transpose(i)
}

func (c *SimpleChord) Transpose(i Interval) {
	c.Root.transpose(i)
}

func (c *SlashChord) Transpose(i Interval) {
	c.Base.transpose(i)
	c.Root.transpose(i)
}
