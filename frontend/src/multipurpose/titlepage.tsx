import { Children, ReactNode } from "react";
import "../styles/TitlePage.scss";

export interface TitlePageProps {
  title: string;
  renderRight?: number;
  children: ReactNode;
}

function TitlePage({ title, renderRight, children }: TitlePageProps) {
  if (!renderRight) {
    renderRight = 0;
  }
  console.log(renderRight);
  const childArray = Children.toArray(children);
  const first = childArray.slice(0, renderRight);
  console.log(first);
  const second = childArray.slice(renderRight);
  return (
    <div className="titlePage">
      <div className="titlePart">
        <h1>{title}</h1>
        {first}
      </div>
      {second}
    </div>
  );
}

export default TitlePage;
