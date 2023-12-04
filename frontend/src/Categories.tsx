import ErrorPage from "./errorpage";
import Loading from "./loading";
import TileList from "./tilelist";
import useGet from "./useGet";

export type Category = {
  id: number;
  name: string;
  description: string;
};

export const Categories = () => {
  const {
    data: categories,
    pending,
    err,
  } = useGet<Category[]>(`http://localhost:9001/category/subcategories`);

  return (
    <div className="categories">
      {err && <ErrorPage err={err}></ErrorPage>}
      {pending && <Loading />}
      {categories && (
        <TileList
          dataList={categories}
          title="Categorias"
          linkFormatter={(c) => c.name}
        />
      )}
    </div>
  );
};

export default Categories;
