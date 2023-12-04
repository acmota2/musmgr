import { Link } from "react-router-dom";

type formatter<T> = (form: T) => string;

const TileList = <T,>({
  dataList,
  title,
  linkFormatter,
}: {
  dataList: T[];
  title: string;
  linkFormatter: formatter<T>;
}) => {
  return (
    <div className="tileList">
      <h2>{title}</h2>
      {dataList.map((data) => {
        // arranjar forma de meter o texto direito em qualquer caso ali no meio
        return (
          <div>
            <Link to={linkFormatter(data)}></Link>
          </div>
        );
      })}
    </div>
  );
};

export default TileList;
