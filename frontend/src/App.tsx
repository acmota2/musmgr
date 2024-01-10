import "./styles/App.scss";
import Navbar from "./Navbar";
import Sidebar from "./Sidebar";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Home from "./Home";
import Categories from "./Categories";
import EventTypes from "./eventtypes";
import Songs from "./songs";
import Category from "./Category";
import CreateSong from "./CreateSong";
import MainEditor from "./SongEditor/SongEditor";

export default function App() {
  return (
    <BrowserRouter>
      <div className="App">
        <Navbar />
        <div className="center">
          <Sidebar />
          <div className="body">
            <Routes>
              <Route path="/song-editor" element={<MainEditor />} />
              <Route path="/song-wizard" element={<CreateSong />} />
              <Route path="/category" element={<Category />} />
              <Route path="/categories" element={<Categories />} />
              <Route path="/event-types" element={<EventTypes />} />
              <Route path="/songs" element={<Songs />} />
              <Route path="/" element={<Home />} />
            </Routes>
          </div>
        </div>
      </div>
    </BrowserRouter>
  );
}
