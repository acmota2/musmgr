import { useSearchParams } from "react-router-dom";
import AnyList from "./multipurpose/anylist";
import TileList from "./multipurpose/tilelist";
import TitlePage from "./multipurpose/titlepage";

export type SubCategory = {
  id: number;
  name: string;
  categoryID: number;
};

const SubCategoryFormatter = ({ data: category }: { data: SubCategory }) => (
  <p>{category.name}</p>
);

export const SubCategory = () => {
  const [params] = useSearchParams();

  return (
    <TitlePage title={`Subcategorias de ${params.get("name")}`}>
      <AnyList
        path={`/subcategories/category/${params.get("id")}`}
        generator={(data: SubCategory[]) => (
          <TileList
            dataList={data}
            linkMaker={(sc: SubCategory) => sc.name}
            linkFormatter={SubCategoryFormatter}
          />
        )}
      />
    </TitlePage>
  );
};

export default SubCategory;
