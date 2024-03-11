import { useSearchParams } from "react-router-dom";
import AnyList from "./multipurpose/anylist";
import TileList from "./multipurpose/tilelist";
import TitlePage from "./multipurpose/titlepage";
import { useState } from "react";
import CreateNode from "./multipurpose/CreateNode";

export type SubCategory = {
  id: number;
  name: string;
  category_name: string;
};

const subCategoryFormatter = ({ data: category }: { data: SubCategory }) => (
  <p>{category.name}</p>
);

export const SubCategory = () => {
  const [params] = useSearchParams();
  const subCategoryState = { stateName: useState<string>("") };
  const [name, setName] = subCategoryState.stateName;

  return (
    <TitlePage title={`Subcategorias de ${params.get("name")}`} renderRight={1}>
      <CreateNode
        path={"/subcategory"}
        redirectTo={() => `/category?name=${params.get("name")}`}
        buttonText="Criar subcategoria"
        state={subCategoryState}
        dataCreator={({ stateName: [name] }) => {
          return {
            id: 0,
            name,
            category_name: params.get("name"),
          };
        }}
      >
        <p>Nome da subcategoria:</p>
        <input
          autoFocus
          type="text"
          placeholder="SubCategoria..."
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
      </CreateNode>
      <AnyList
        path={`/subcategories/category/${params.get("name")}`}
        generator={(data: SubCategory[]) => (
          <TileList
            dataList={data}
            linkMaker={(sc: SubCategory) =>
              `/songs?name=${sc.name}&id=${sc.id}&subcategory=true`
            }
            linkFormatter={subCategoryFormatter}
          />
        )}
      />
    </TitlePage>
  );
};

export default SubCategory;
