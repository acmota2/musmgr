import { ReactNode, useRef, useState } from "react";
import AnyPostForm, { AnyPostFormProps } from "./AnyPostForm";
import Modal from "./Modal";

type CreateNodeElement<T, S> = AnyPostFormProps<S, T> & {
  buttonText: string;
  children: ReactNode;
};

export default function CreateNode<S, T>(node: CreateNodeElement<T, S>) {
  const [isOpened, setIsOpened] = useState<boolean>(false);
  const buttonRef = useRef<HTMLButtonElement>(null);
  const { buttonText, children } = node;
  return (
    <>
      <button onClick={() => setIsOpened(true)} ref={buttonRef}>
        {buttonText}
      </button>
      {buttonRef.current && (
        <Modal
          isOpened={isOpened}
          onClose={() => setIsOpened(false)}
          yPos={buttonRef.current.offsetTop + buttonRef.current.offsetHeight}
          xPos={buttonRef.current.offsetLeft}
        >
          <AnyPostForm {...node}>{children}</AnyPostForm>
        </Modal>
      )}
    </>
  );
}
