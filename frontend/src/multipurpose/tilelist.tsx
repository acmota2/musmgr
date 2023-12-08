import { Link } from "react-router-dom";
import "../styles/TileList.scss";

export type linker<T> = (form: T) => string;

const TileList = <T,>({
  dataList,
  linkMaker,
  linkFormatter: LinkFormatter,
}: {
  dataList: T[];
  linkMaker: linker<T>;
  linkFormatter: JSX.ElementType;
}) => {
  return (
    <ul className="tileList">
      {dataList.map((data, index) => {
        return (
          <li key={`${index}`}>
            <Link
              className="linkList"
              to={linkMaker(data)}
              key={`link_${index}`}
            >
              <LinkFormatter data={data} />
            </Link>
          </li>
        );
      })}
    </ul>
  );
};

export default TileList;
