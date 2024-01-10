import { Syllable } from "./song_text/structure";
import {
  NOTENAMES,
  SongState,
  addStanceToSong,
  addSyllableToSong,
  addVerseToSong,
  addWordToSong,
  setSongFromSyllable,
} from "./song_text/songState";
import { useState } from "react";
import NoteSelector from "./NoteSelector";
import AddButton from "./AddButton";

export type SyllableProps = {
  songState: SongState;
  syllable: Syllable;
  stanceIndex: number;
  verseIndex: number;
  wordIndex: number;
  syllableIndex: number;
};

const Syllables = (s: SyllableProps) => {
  const { syllable } = s;
  console.log(syllable);
  const [text, setText] = useState<string>(syllable.text);
  return (
    <div className="syllable">
      <AddButton className="chord">
        <NoteSelector
          noteValue={
            syllable.chord ? NOTENAMES[syllable.chord.root.index] : NOTENAMES[0]
          }
          noteOnChange={(n: string) => {
            if (syllable.chord) {
              syllable.chord.root = {
                ...syllable.chord.root,
                index: NOTENAMES.indexOf(n),
              };
              setSongFromSyllable({ ...s, syllable });
            }
          }}
          accidentValue={syllable.chord ? syllable.chord.root.alteration : 0}
          accidentOnChange={(alteration: number) => {
            if (syllable.chord) {
              syllable.chord.root = {
                ...syllable.chord.root,
                alteration,
              };
              setSongFromSyllable({ ...s, syllable });
            }
          }}
        />
      </AddButton>
      <AddButton className="chordBase">
        {syllable.chord && syllable.chord.base && (
          <NoteSelector
            noteValue={NOTENAMES[syllable.chord.base.index]}
            noteOnChange={(n) => {
              if (syllable.chord && syllable.chord.base) {
                syllable.chord.base = {
                  ...syllable.chord.base,
                  index: NOTENAMES.indexOf(n),
                };
                setSongFromSyllable({ ...s, syllable });
              }
            }}
            accidentValue={syllable.chord.base.alteration}
            accidentOnChange={(alteration) => {
              if (syllable.chord && syllable.chord.base) {
                syllable.chord.base = {
                  ...syllable.chord.base,
                  alteration,
                };
                setSongFromSyllable({ ...s, syllable });
              }
            }}
          />
        )}
      </AddButton>
      <input
        autoFocus
        type="text"
        placeholder="Sílaba..."
        value={text}
        onChange={(e) => {
          switch (e.target.value[e.target.value.length - 1]) {
            case "+":
              addSyllableToSong(s);
              break;
            case " ":
              addWordToSong(s);
              break;
            case "|":
              addVerseToSong(s);
              break;
            case "\\":
              addStanceToSong(s);
              break;
            default:
              setText(e.target.value);
              syllable.text = text;
              setSongFromSyllable({ ...s, syllable });
          }
        }}
      />
    </div>
  );
};

export default Syllables;
