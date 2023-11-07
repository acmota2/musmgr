package songs

const (
	Chorus = 0
)

type Song struct {
	Tonality Tonality
	Stances  []Stance
}

type StanceType int

type Stance struct {
	Text []Word
	Type StanceType
}

type Word []Syllable

type Syllable struct {
	Text  string
	Chord Chord
}

type Chord struct {
	Note    Note
	Details string
}

// Transposition chain
func (s *Song) Transpose(i *Interval) {
	for _, stance := range s.Stances {
		stance.transpose(i)
	}
}

func (s *Stance) transpose(i *Interval) {
	for _, w := range s.Text {
		w.transpose(i)
	}
}

func (w *Word) transpose(i *Interval) {
	for _, s := range *w {
		s.transpose(i)
	}
}

func (s *Syllable) transpose(i *Interval) {
	s.Chord.transpose(i)
}

func (c *Chord) transpose(i *Interval) {
	c.Note.transpose(i)
}
