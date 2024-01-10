import { ReactNode } from "react";

type AddButton = {
  title?: string;
  className: string;
  children: ReactNode;
};

const AddButton = ({ title, className, children }: AddButton) => {
  let wasClicked = false;
  return (
    <div className={className}>
      {!wasClicked && (
        <button
          onClick={(e) => {
            e.preventDefault();
            wasClicked = true;
          }}
        >
          {title || "+"}
        </button>
      )}
      {wasClicked && children}
    </div>
  );
};

export default AddButton;
