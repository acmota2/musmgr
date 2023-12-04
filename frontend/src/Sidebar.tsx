import { Link } from "react-router-dom";
import "./styles/Sidebar.scss";

type OrganizationTile = {
  name: string;
  route: string;
};

const Sidebar = () => {
  const organizationTiles: OrganizationTile[] = [
    { name: "Categorias", route: "/categories" },
    { name: "Eventos", route: "/event-types" },
    { name: "Músicas", route: "/songs" },
  ];

  return (
    <div className="Sidebar">
      {organizationTiles.map((tile) => (
        <Link to={tile.route}>
          <div className="organization-tile">
            <p>{tile.name}</p>
          </div>
        </Link>
      ))}
    </div>
  );
};

export default Sidebar;
