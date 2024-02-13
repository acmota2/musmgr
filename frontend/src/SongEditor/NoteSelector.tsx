import "../styles/editor/noteselector.scss";
import { useState } from "react";

type NoteSelector = {
  noteValue: string;
  noteOnChange: (n: string) => void;
  accidentValue: number;
  accidentOnChange: (alteration: number) => void;
};

const NoteSelector = ({
  noteValue,
  noteOnChange,
  accidentValue,
  accidentOnChange,
}: NoteSelector) => {
  const fromAccident = (a: number) => {
    if (a > 0) {
      return "#".repeat(a);
    } else if (a < 0) {
      return "♭".repeat(-a);
    } else {
      return "";
    }
  };

  const [note, setNote] = useState<string>(noteValue);
  const [accident, setAccident] = useState<number>(accidentValue);
  const [accidentText, setAccidentText] = useState<string>(
    fromAccident(accident),
  );

  return (
    <div className="note">
      <input
        max={2}
        min={-2}
        type="number"
        value={accident}
        onChange={(e) => {
          const n = parseInt(e.target.value, 10);
          setAccident(n);
          setAccidentText(fromAccident(n));
          accidentOnChange(n);
        }}
      />
      <select
        value={note}
        onChange={(e) => {
          const newNote = e.target.value;
          setNote(newNote);
          console.log("from NoteSelector", newNote);
          noteOnChange(newNote);
        }}
      >
        <option value="Dó">Dó</option>
        <option value="Ré">Ré</option>
        <option value="Mi">Mi</option>
        <option value="Fá">Fá</option>
        <option value="Sol">Sol</option>
        <option value="Lá">Lá</option>
        <option value="Si">Si</option>
      </select>
      <p>{accidentText}</p>
    </div>
  );
};

export default NoteSelector;
