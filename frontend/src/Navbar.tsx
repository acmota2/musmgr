import { Link } from "react-router-dom";
import "./styles/Navbar.scss";

const Navbar = () => {
  return (
    <nav className="Navbar">
      <div className="links">
        <div className="homeButton">
          <Link to="/">SongMGR</Link>
        </div>
        <div className="searchBar">
          <input type="text" placeholder="Procurar música..." />
        </div>
        <div className="songWizard">
          <Link to="/song-wizard">Criar música</Link>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
