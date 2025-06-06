import { useNavigate } from "react-router-dom";

function Detalle() {
  const navigate = useNavigate();

  return (
    <div>
      <h2>Detalles</h2>
      <p>Listado o gestión de tareas.</p>
      <button onClick={() => navigate("/Home")}>← Volver a Home</button>
    </div>
  );
}

export default Detalle;
