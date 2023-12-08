type SongText = {
    tonality: Tonality;
    stances: Stance[];
};

function transposeSongText(st: SongText, i: Interval): SongText {
    const n = st.tonality.note.transpose(st.tonality.note, i);
    st.tonality.note = n;
    st.stances = st.stances.map((s) => transposeStance(s, i));
    return st;
}

type Chorus = 0;
type StanceHelper = number;
type StanceType = Chorus | StanceHelper;
type Stance = {
    type: StanceType;
    text: Word[];
};

function transposeStance(s: Stance, i: Interval): Stance {
    s.text = s.text.map((w) => transposeWords(w, i));
    return s;
}

type Word = Syllable[];

function transposeWords(w: Word, i: Interval): Word {
    return w.map((s) => {
        s.chord = s.chord.transpose(i);
        return s;
    });
}

type Syllable = {
    text: string;
    chord: Chord;
};

interface Chord {
    root: Note;
    base?: Note;
    details?: string;
    toString(): string;
    transpose(i: Interval): Chord;
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
