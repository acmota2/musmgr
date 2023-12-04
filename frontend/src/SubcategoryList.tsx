export type SubCategory = {
  id: number;
  name: string;
  subCategoryID: number;
};

export const SubcategoryList = ({
  subcategories,
  title,
}: {
  subcategories: SubCategory[];
  title: string;
}) => {
  return (
    <div className="subCategoryList">
      <p>{title}</p>
      {subcategories.map((subcategory) => {
        const subcategoryRef = `/subcategory/${subcategory.name}`;

        return (
          <div>
            <a href={subcategoryRef}>{subcategory.name}</a>
          </div>
        );
      })}
    </div>
  );
};

export default SubcategoryList;
