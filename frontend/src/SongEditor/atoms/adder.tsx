import "../../styles/atoms/adder.scss";

type AdderElement = {
  elemName: string;
  onAdd: () => void;
  onRemove: () => void;
};

export default function Adder({ elemName, onAdd, onRemove }: AdderElement) {
  return (
    <div className="adder">
      <div className="buttons">
        <button
          className="add"
          onClick={(e) => {
            e.preventDefault();
            onAdd();
          }}
        >
          +
        </button>
        <p className="hide">{elemName}</p>
        <button
          className="remove"
          onClick={(e) => {
            e.preventDefault();
            onRemove();
          }}
        >
          -
        </button>
      </div>
    </div>
  );
}
