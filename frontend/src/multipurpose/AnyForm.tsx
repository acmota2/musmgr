import "../styles/AnyForm.scss";
import { ReactNode } from "react";
import usePost from "../hooks/usePost";
import { To } from "react-router-dom";

export type DataCreator<S, T> = (state: S) => T;
export type Redirector<T> = (data: T | void) => (To | null);

export interface AnyFormProps<S, T> {
  path: string;
  redirectTo: Redirector<T>;
  buttonText: string;
  state: S;
  dataCreator: DataCreator<S, T>;
  children: ReactNode;
}

const AnyForm = <S, T>({
  path,
  redirectTo,
  buttonText,
  state,
  dataCreator,
  children,
}: AnyFormProps<S, T>) => {
  const { pending, err, poster } = usePost<T>(path, redirectTo());

  const submitHandler = (e: React.FormEvent) => {
    e.preventDefault();
    poster(dataCreator(state));
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

export default AnyForm;
