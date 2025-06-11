
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Login from './paginas/Login.jsx';
import Home from './paginas/Home.jsx';
import Detalle from './paginas/Detalle.jsx';
import './App.css';


function App() {
  return (
    <div className="App-header">
      <Router>
        <Routes >
          <Route path="/Login" element={<Login />} />
          <Route path="/Home" element={<Home />} />
          <Route path="/Detalle" element={<Detalle />} />
          <Route path="/Detalle/:id" element={<Detalle />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
