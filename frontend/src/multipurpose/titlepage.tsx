import { ReactNode } from "react";
import "../styles/TitlePage.scss";

export interface TitlePageProps {
  title: string;
  children: ReactNode;
}

function TitlePage({ title, children }: TitlePageProps) {
  return (
    <div className="titlePage">
      <h1>{title}</h1>
      {children}
    </div>
  );
}

export default TitlePage;
