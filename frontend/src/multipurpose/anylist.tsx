import useGet from "../hooks/useGet";
import ErrorPage from "./errorpage";
import Loading from "./loading";
import "../styles/AnyList.scss";
import { ReactNode } from "react";
import { DependencyFunction } from "../hooks/hooktypes";

export interface ListType<T> {
  path: string;
  generator: (t: T[]) => ReactNode;
  dependencies?: DependencyFunction<T[]>;
}

function AnyList<T>({ path, generator, dependencies }: ListType<T>) {
  const { data, pending, err } = useGet<T[]>(path, dependencies);
  return (
    <ul className="anyList">
      {err && <ErrorPage err={err} />}
      {pending && <Loading />}
      {data && generator(data)}
    </ul>
  );
}

export default AnyList;
