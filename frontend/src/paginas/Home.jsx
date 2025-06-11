import { useNavigate } from "react-router-dom";
import './Home.css';
import Foto from '../components/Foto';
import ListaDesplegable from '../components/ListaDesplegable';


function Home() {
  const navigate = useNavigate();

  return (
    <div className="home">
      <h2>Home</h2>
      <button onClick={() => navigate("/")} className="botonRedondoVolver">← Volver a Login</button>
      <button onClick={() => navigate("/Detalle")} className="botonRedondoAdelante">Ir a Detalle →</button>
      <hr/>
      <h1 className="Titulo">GIMNACIO HELL NAHHHHH 😩</h1>
     <Foto/>
     <ListaDesplegable/>
     
     

      
    </div>
  );
}

export default Home;

/*import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

function Home() {
  const navigate = useNavigate();
  const [actividades, setActividades] = useState([]);
  const [seleccionada, setSeleccionada] = useState("");

  useEffect(() => {
    // Llamada al backend para traer las actividades
    fetch("http://localhost:8080/actividades")
      .then((res) => res.json())
      .then((data) => setActividades(data))
      .catch((error) => console.error("Error al obtener actividades:", error));
  }, []);

  return (
    <div>
      <h2>Home</h2>

      <button onClick={() => navigate("/")}>← Volver a Login</button>
      <button onClick={() => navigate("/Detalle")}>Ir a Detalle →</button>

      <h3>Lista de actividades</h3>
      <select
        value={seleccionada}
        onChange={(e) => setSeleccionada(e.target.value)}
      >
        <option value="">-- Seleccioná una actividad --</option>
        {actividades.map((actividad, index) => (
          <option key={index} value={actividad.id}>
            {actividad.nombre}
          </option>
        ))}
      </select>

      {seleccionada && (
        <p>
          Actividad seleccionada:{" "}
          {
            actividades.find((a) => a.id.toString() === seleccionada)?.nombre
          }
        </p>
      )}


    </div>
  );
}

export default Home;*/

