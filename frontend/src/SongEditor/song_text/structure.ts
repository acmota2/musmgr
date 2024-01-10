import { Note, Tonality, Interval, NoteName, newNote } from "./transposer";

export type SongText = {
  tonality: Tonality;
  stances: SongStance[];
};

export function transposeSongText(st: SongText, i: Interval): SongText {
  const n = st.tonality.note.transpose(st.tonality.note, i);
  st.tonality.note = n;
  st.stances = st.stances.map((s) => transposeStance(s, i));
  return st;
}

export type Chorus = 0;
export type StanceType = Chorus | number;
export type SongStance = {
  type: StanceType;
  text: Verse[];
};

export function emptySong(): SongText {
  return {
    tonality: {
      details: "Maior",
      note: newNote(NoteName.C, 0),
    },
    stances: [
      {
        type: 0,
        text: [[[{ text: "", chord: makeChord(newNote(NoteName.C, 0)) }]]],
      },
    ],
  };
}

export type Verse = Word[];

function transposeStance(s: SongStance, i: Interval): SongStance {
  s.text = s.text.map((v) => v.map((w) => transposeWords(w, i)));
  return s;
}

export type Word = Syllable[];

function transposeWords(w: Word, i: Interval): Word {
  return w.map((s) => {
    s.chord = s.chord ? s.chord.transpose(s.chord, i) : undefined;
    return s;
  });
}

export type Syllable = {
  text: string;
  chord?: Chord;
};

export function emptySyllable() {
  return {
    text: "",
  };
}

export interface Chord {
  root: Note;
  base?: Note;
  details?: string;
  toString(c: Chord): string;
  transpose(c: Chord, i: Interval): Chord;
}

export default function makeChord(
  root: Note,
  base?: Note,
  details?: string,
): Chord {
  return {
    root,
    base,
    details,
    toString: chordToString,
    transpose: transposeChord,
  };
}

function chordToString(c: Chord) {
  const { root, base, details } = c;
  return (
    root.toString(root) +
    (base ? "/" + base.toString(base) : "") +
    (details ? details : "")
  );
}

function transposeChord(c: Chord, i: Interval): Chord {
  let { root, base } = c;
  root = root.transpose(root, i);
  if (base) {
    base = base.transpose(base, i);
  }
  return {
    root,
    base,
    details: c.details,
    toString: c.toString,
    transpose: c.transpose,
  };
}
