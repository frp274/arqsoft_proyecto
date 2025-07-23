
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import Login from './paginas/Login.jsx';
import Home from './paginas/Home.jsx';
import Detalle from './paginas/Detalle.jsx';
import HomeAdm from './paginas/HomeAdm';
import './App.css';


function App() {
  return (
    <div className="App-header">
      <Router>
        <Routes >
          <Route path="/" element={<Navigate to="/Login" />} />
          <Route path="/Login" element={<Login />} />
          <Route path="/Home" element={<Home />} />
          <Route path="/Detalle" element={<Detalle />} />
          <Route path="/Detalle/:id" element={<Detalle />} />
          <Route path="/Admin" element={<HomeAdm />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
