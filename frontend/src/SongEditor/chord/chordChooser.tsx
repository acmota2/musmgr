import NoteSelector from "../NoteSelector";
import { NOTENAMES } from "../song_text/songState";
import { Chord } from "../song_text/structure";
import { Note } from "../song_text/transposer";

type ChordChooserElement = {
  note: Note;
};

export default function Note({ note }: ChordChooserElement) {
  return (<div>
    <NoteSelector
      noteValue={NOTENAMES[note.index]}
      noteOnChange={(n: string) => }
  </div>);
}
