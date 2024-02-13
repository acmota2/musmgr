import { ReactNode, useEffect, useRef } from "react";

type ModalElement = {
  isOpened: boolean;
  onClose: () => void;
  xPos?: number;
  yPos?: number;
  children: ReactNode;
};

export default function Modal({
  isOpened,
  onClose,
  children,
  xPos,
  yPos,
}: ModalElement) {
  const dialogRef = useRef<HTMLDialogElement>(null);
  const x = xPos ? xPos : 0;
  const y = yPos ? yPos : 0;

  useEffect(() => {
    if (isOpened) {
      dialogRef.current?.showModal();
      document.body.classList.add("modal-open"); // prevent bg scroll
    } else {
      dialogRef.current?.close();
      document.body.classList.remove("modal-open");
    }
  }, [isOpened]);

  return (
    <div className="modal">
      <dialog ref={dialogRef} style={{ position: "absolute", left: x, top: y }}>
        <button
          onClick={(e) => {
            e.preventDefault();
            onClose();
          }}
        >
          x
        </button>
        {children}
      </dialog>
    </div>
  );
}
