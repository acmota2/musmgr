import "./styles/App.scss";
import Navbar from "./Navbar";
import Sidebar from "./Sidebar";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Home from "./Home";
import Categories from "./Categories";

export default function App() {
  return (
    <BrowserRouter>
      <div className="App">
        <Navbar />
        <div className="center">
          <Sidebar />
          <div className="body">
            <Routes>
              <Route path="/categories" element={<Categories />} />
              <Route path="/" element={<Home />} />
            </Routes>
          </div>
        </div>
      </div>
    </BrowserRouter>
  );
}
