
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import Login from './paginas/Login.jsx';
import Registro from './paginas/Registro.jsx';
import Home from './paginas/Home.jsx';
import Detalle from './paginas/Detalle.jsx';
import HomeAdm from './paginas/HomeAdm';
import MisInscripciones from './paginas/MisInscripciones.jsx';
import './App.css';


function App() {
  return (
    <div className="App-header">
      <Router>
        <Routes >
          <Route path="/" element={<Navigate to="/Login" />} />
          <Route path="/Login" element={<Login />} />
          <Route path="/Registro" element={<Registro />} />
          <Route path="/Home" element={<Home />} />
          <Route path="/Detalle" element={<Detalle />} />
          <Route path="/Detalle/:id" element={<Detalle />} />
          <Route path="/Admin" element={<HomeAdm />} />
          <Route path="/MisInscripciones" element={<MisInscripciones />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
