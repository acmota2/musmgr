import "../styles/editor/syllables.scss";
import makeChord, { Syllable } from "./song_text/structure";
import {
  CurrentStanceState,
  SongState,
  /* addStanceToSong, */
  addSyllableToSong,
  addVerseToSong,
  addWordToSong,
  setSongFromSyllable,
} from "./song_text/songState";
import { useEffect, useRef, useState } from "react";
import ChordElement from "./chord/chord";
import { NoteName, newNote } from "./song_text/transposer";
import Modal from "../multipurpose/Modal";

export type SyllableProps = {
  songState: SongState;
  stanceState: CurrentStanceState;
  syllable: Syllable;
  stanceIndex: number;
  verseIndex: number;
  wordIndex: number;
  syllableIndex: number;
};

const Syllables = (s: SyllableProps) => {
  const { syllable } = s;
  const chord = syllable.chord;
  console.log(syllable);
  const [text, setText] = useState<string>(syllable.text);
  const [chordText, setChordText] = useState<string>(
    chord ? chord.toString(chord) : "",
  );
  const [modalOpen, setModalOpen] = useState<boolean>(false);

  const buttonRef = useRef<HTMLButtonElement>(null);
  const spanRef = useRef<HTMLSpanElement>(null);
  const [spanWidth, setSpanWidth] = useState<number>(
    chordText !== "" ? chordText.length : 0,
  );
  useEffect(
    () =>
      setSpanWidth(
        spanRef.current ? spanRef.current.offsetWidth : "...".length,
      ),
    [text],
  );

  return (
    <>
      {chordText === "" ? (
        <button
          className="chooser"
          onClick={(e) => {
            e.preventDefault();
            setModalOpen(true);
            const curChord = makeChord(newNote(NoteName.C, 0));
            setSongFromSyllable({
              ...s,
              syllable: { ...syllable, chord: curChord },
            });
            setChordText(curChord.toString(curChord));
          }}
          ref={buttonRef}
        >
          +
        </button>
      ) : (
        <div className="chooserDefined">
          <button
            className="chooser"
            onClick={(e) => {
              e.preventDefault();
              setModalOpen(true);
            }}
            ref={buttonRef}
          >
            {chordText}
          </button>
          <button
            className="chooser"
            onClick={(e) => {
              e.preventDefault();
              setSongFromSyllable({
                ...s,
                syllable: { ...syllable, chord: undefined },
              });
              setChordText("");
            }}
          >
            -
          </button>
        </div>
      )}
      {buttonRef.current && (
        <Modal
          isOpened={modalOpen}
          onClose={() => setModalOpen(false)}
          xPos={buttonRef.current.offsetLeft}
          yPos={buttonRef.current.offsetTop + buttonRef.current.offsetHeight}
        >
          <ChordElement
            chord={chord ? chord : makeChord(newNote(NoteName.C, 0))}
            onChange={(c) => {
              console.log("from Syllable", c);
              setSongFromSyllable({
                ...s,
                syllable: { ...syllable, chord: c },
              });
              setChordText(c.toString(c));
            }}
          />
        </Modal>
      )}
      <span className="input" ref={spanRef}></span>
      <input
        autoFocus
        className="syllableText"
        type="text"
        placeholder="..."
        value={text}
        style={{ width: spanWidth }}
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
            /* case "\\":
              addStanceToSong(s);
              break; */
            default:
              const newText = e.target.value;
              setText(newText);
              syllable.text = newText;
              setSongFromSyllable({ ...s, syllable });
          }
        }}
      />
    </>
  );
};

export default Syllables;
