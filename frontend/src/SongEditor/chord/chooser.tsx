type ChooserElement = {
  chordText: string;
  onClick: () => void;
};

export default function Chooser({ chordText, onClick }: ChooserElement) {
  return (
    <button
      onClick={(e) => {
        e.preventDefault();
        onClick();
      }}
    >
      {chordText !== "" ? chordText : "+"}
    </button>
  );
}
