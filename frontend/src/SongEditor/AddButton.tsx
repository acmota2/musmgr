import { useState } from "react";

type AddButton = {
  title?: string;
  className: string;
};

const AddButton = ({ title, className }: AddButton) => {
  let [wasClicked, setWasClicked] = useState(false);
  return (
    <div className={className}>
      {!wasClicked && (
        <button
          onClick={(e) => {
            e.preventDefault();
            setWasClicked(true);
          }}
        >
          {title || "+"}
        </button>
      )}
      {wasClicked && <input />}
    </div>
  );
};

export default AddButton;
