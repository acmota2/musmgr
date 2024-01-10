import AnyList from "./multipurpose/anylist";
import TileList from "./multipurpose/tilelist";
import TitlePage from "./multipurpose/titlepage";

export type Category = {
  name: string;
  description: string;
};

const CategoryFormatter = ({ data: category }: { data: Category }) => (
  <p>{category.name}</p>
);
const linkMaker = (c: Category) => `/category?name=${c.name}`;

export const Categories = () => {
  return (
    <TitlePage title="Categorias">
      <AnyList
        path="/categories"
        generator={(data: Category[]) => (
          <TileList
            dataList={data}
            linkMaker={linkMaker}
            linkFormatter={CategoryFormatter}
          />
        )}
      />
    </TitlePage>
  );
};

export default Categories;
