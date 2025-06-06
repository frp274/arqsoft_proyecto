
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Login from './components/Login.jsx';
import Home from './components/Home.jsx';
import Detalle from './components/Detalle.jsx';


function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/Home" element={<Home />} />
        <Route path="/Detalle" element={<Detalle />} />
      </Routes>
    </Router>
  );
}

export default App;
