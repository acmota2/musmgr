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
  const [note, setNote] = useState<string>(noteValue);
  const [accident, setAccident] = useState<number>(accidentValue);

  return (
    <div className="note">
      <select
        value={note}
        onChange={(e) => {
          setNote(e.target.value);
          noteOnChange(note);
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
      <input
        type="number"
        value={accident >= 0 ? "#".repeat(accident) : "♭".repeat(-accident)}
        onChange={(e) => {
          const v = e.target.value;
          setAccident(v[0] === "#" ? v.length : -v.length);
          accidentOnChange(accident);
        }}
      />
    </div>
  );
};

export default NoteSelector;
