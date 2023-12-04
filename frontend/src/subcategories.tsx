import useGet from "./useGet";
import { SubcategoryList, SubCategory } from "./SubcategoryList";
import Loading from "./loading";
import ErrorPage from "./errorpage";

export const SubCategories = () => {
  const { data, pending, err } = useGet<SubCategory[]>(
    `http://localhost:9001/category/subcategories`
  );

  return (
    <div className="subcategories">
      {err && <ErrorPage err={err}></ErrorPage>}
      {pending && <Loading />}
      {data && (
        <SubcategoryList subcategories={data} title="Bla"></SubcategoryList>
      )}
    </div>
  );
};

export default SubCategories;
