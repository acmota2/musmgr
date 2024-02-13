import "../../styles/editor/chord.scss";
import { useState } from "react";
import NoteSelector from "../NoteSelector";
import { NOTENAMES } from "../song_text/songState";
import makeChord, { Chord } from "../song_text/structure";
import { Note, NoteName, newNote } from "../song_text/transposer";

type ChordProps = {
  onChange: (c: Chord) => void;
  chord: Chord | undefined;
};

export default function ChordElement({ onChange, chord }: ChordProps) {
  const [chordState, setChordState] = useState<Chord>(
    chord ? chord : makeChord(newNote(NoteName.C, 0)),
  );
  const [base, setBase] = useState<Note | undefined>(
    chordState.base ? chordState.base : undefined,
  );
  const [details, setDetails] = useState<string>(
    chordState.details ? chordState.details : "",
  );
  const [root, setRoot] = useState<Note>(chordState.root);
  const buttonText = "+";

  return (
    <div className="chord">
      <div className="root">
        <NoteSelector
          noteValue={NOTENAMES[chordState.root.index]}
          noteOnChange={(n: string) => {
            console.log("from ChordElement", n);
            const note = {
              ...chordState.root,
              index: NOTENAMES.indexOf(n),
            };
            setRoot(note);
            const newChord = { ...chordState, root: note };
            setChordState(newChord);
            onChange(newChord);
          }}
          accidentValue={chordState ? chordState.root.alteration : 0}
          accidentOnChange={(alteration: number) => {
            const n = {
              ...root,
              alteration,
            };
            setRoot(n);
            const newChord = { ...chordState, root: n };
            setChordState(newChord);
            onChange(newChord);
          }}
        />
        <input
          type="text"
          value={details}
          onChange={(e) => {
            const d = e.target.value;
            setDetails(d);
            const newChord = { ...chordState, details: d };
            setChordState(newChord);
            onChange(newChord);
          }}
        />
      </div>
      {!base && (
        <button
          onClick={(e) => {
            e.preventDefault();
            const n = newNote(NoteName.C, 0);
            setBase(n);
            const newChord = { ...chordState, base: n };
            setChordState(newChord);
            onChange(newChord);
          }}
        >
          {buttonText}
        </button>
      )}
      {base && chordState && (
        <>
          <div className="base">
            <p>/</p>
            <NoteSelector
              noteValue={NOTENAMES[base.index]}
              noteOnChange={(n: string) => {
                const note = {
                  ...chordState!.base!,
                  index: NOTENAMES.indexOf(n),
                };
                setBase(note);
                const newChord = { ...chordState, base: note };
                setChordState(newChord);
                onChange(newChord);
              }}
              accidentValue={base.alteration}
              accidentOnChange={(alteration: number) => {
                const n = {
                  ...base,
                  alteration,
                };
                setBase(n);
                const newChord = { ...chordState, base: n };
                setChordState(newChord);
                onChange(newChord);
              }}
            />
          </div>
          <button
            onClick={(e) => {
              e.preventDefault();
              setBase(undefined);
              const newChord = { ...chordState, base: undefined };
              setChordState(newChord);
              onChange(newChord);
            }}
          >
            -
          </button>
        </>
      )}
    </div>
  );
}
