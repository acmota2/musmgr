import "../styles/AnyForm.scss";
import { ReactNode } from "react";
import usePost from "../hooks/usePost";
import { To } from "react-router-dom";

export type DataCreator<S, T> = (state: FormState<S>) => T;
export type Redirector<T> = (data: T | void) => To | null;

export type FormState<T> = {
  [P in keyof T]: [T[P], React.Dispatch<React.SetStateAction<T[P]>>];
};

export interface AnyPostFormProps<S, T> {
  path: string;
  redirectTo: Redirector<T>;
  buttonText: string;
  state?: FormState<S>;
  dataCreator?: DataCreator<S, T>;
  children: ReactNode;
}

const AnyPostForm = <S, T>({
  path,
  redirectTo,
  buttonText,
  state,
  dataCreator,
  children,
}: AnyPostFormProps<S, T>) => {
  const { pending, err, poster } = usePost<T>(path, redirectTo());

  const submitHandler = (e: React.FormEvent) => {
    e.preventDefault();
    if (dataCreator && state) {
      poster(dataCreator(state));
    }
  };

  return (
    <form onSubmit={submitHandler}>
      {children}
      {!pending && <button>{buttonText}</button>}
      {pending && <button disabled>A carregar...</button>}
      {err && <p>{err.message}</p>}
    </form>
  );
};

export default AnyPostForm;
