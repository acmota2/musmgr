package songs

import (
	"fmt"
	"strings"
)

const (
	C = 0
	D = 2
	E = 4
	F = 5
	G = 7
	A = 9
	B = 11
)

var noteOrder = []int{C, D, E, F, G, A, B}

type Note struct {
	Index      int `json:"index"`      // index of noteOrder
	Alteration int `json:"alteration"` // semitones after or bellow note
}

func (n Note) String() string {
	alteration := ""
	if n.Alteration < 0 {
		alteration = strings.Repeat("♭", n.Alteration*-1)
	} else {
		alteration = strings.Repeat("#", n.Alteration)
	}
	notes := []string{"C", "D", "E", "F", "G", "A", "B"}
	return fmt.Sprintf("%s%s", notes[n.Index], alteration)
}

func (n *Note) AbsoluteCount() int {
	return n.Index + n.Alteration
}

type Tonality struct {
	Name string `json:"name"`
	Note Note   `json:"note"`
}

type Interval struct {
	distanceToNoteIndex int
	semitones           int
}

func (n *Note) transpose(interval Interval) {
	semitones := noteOrder[n.Index] + noteOrder[n.Index+interval.distanceToNoteIndex]
	n.Index = n.Index + interval.distanceToNoteIndex
	n.Alteration = semitones - interval.semitones
}

const transpositionNumber = 12

var sensicalTranspose = []Note{
	{0, 0},  // Dó
	{1, -1}, // Réb
	{1, 0},  // Ré
	{2, -1}, // Mib
	{2, 0},  // Mi
	{3, 0},  // Fá
	{3, 1},  // Fá#
	{4, 0},  // Sol
	{5, -1}, // Láb
	{5, 0},  // Lá
	{6, -1}, // Sib
	{6, 0},  // Si
}

func (t Tonality) TransposeNote(n Note) Interval {
	distanceToNote := t.Note.Index - n.Index
	semitones := n.AbsoluteCount() - t.Note.AbsoluteCount()

	return Interval{distanceToNote, semitones}
}

func (t Tonality) TransposeSemitones(interval int) Interval {
	target := sensicalTranspose[(t.Note.AbsoluteCount()+interval)%transpositionNumber]
	distanceToNote := target.Index - t.Note.Index

	return Interval{distanceToNote, interval}
}
