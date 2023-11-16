package tests

import (
	"backend/songs"
	"encoding/json"
	"testing"
)

var tonality songs.Tonality = songs.Tonality{
	Name: "minor",
	Note: songs.Note{Index: 2, Alteration: 0},
}

var note songs.Note = songs.Note{
	Index:      1,
	Alteration: 0,
}

var chord songs.Chord = songs.Chord{
	Root: note,
	Base: nil,
}

var words []songs.Word = []songs.Word{
	{
		songs.Syllable{
			Text:  "Is",
			Chord: &chord,
		},
		songs.Syllable{
			Text:  "to",
			Chord: nil,
		},
		songs.Syllable{
			Text:  "",
			Chord: nil,
		},
		songs.Syllable{
			Text:  "é",
			Chord: &chord,
		},
		songs.Syllable{
			Text:  "\n",
			Chord: nil,
		},
	},
	{
		songs.Syllable{
			Text:  "um",
			Chord: &chord,
		},
		songs.Syllable{
			Text:  " ",
			Chord: nil,
		},
		songs.Syllable{
			Text:  "tes",
			Chord: nil,
		},
		songs.Syllable{
			Text:  "te",
			Chord: &chord,
		},
		songs.Syllable{
			Text:  "\n",
			Chord: nil,
		},
	},
}

var song songs.SongText = songs.SongText{
	Tonality: tonality,
	Stances: []songs.Stance{
		{
			Text: words,
			Type: songs.Chorus,
		},
	},
}

func TestMarshal(t *testing.T) {
	out, err := json.MarshalIndent(song, "", "    ")
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}
	var toCompare songs.SongText
	err = json.Unmarshal(out, &toCompare)
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}
}

var major3rdDown songs.Interval = songs.Interval{
	DistanceToNoteIndex: -2,
	SemitoneDistance:    -4,
}

func TestTransposeInterval(t *testing.T) {
	do := songs.Note{
		Index:      0,
		Alteration: 0,
	}

	lab := songs.Note{
		Index:      5,
		Alteration: -1,
	}

	actual := do
	actual.Transpose(major3rdDown)

	if actual != lab {
		t.Errorf("Excepted %v, got %v\n", lab, actual)
	}
}
