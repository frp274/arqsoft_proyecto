import { useNavigate, useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import axios from "axios";
import './Detalle.css';

function Detalle() {
  const { id } = useParams(); // Tomamos el ID de la URL
  const navigate = useNavigate();
  const [actividad, setActividad] = useState(null);
  const [mensaje, setMensaje] = useState('');

  // Cargar detalles de la actividad
  useEffect(() => {
    axios.get(`http://localhost:8080/actividad/${id}`)
      .then((res) => setActividad(res.data))
      .catch((err) => console.error(err));
  }, [id]);

  // Función para inscribirse
  const inscribirse = () => {
    axios.post('http://localhost:8080/inscripcion', { actividadId: id })
      .then(() => {
        setMensaje('¡Inscripción exitosa!');
        setTimeout(() => navigate('/Home'), 2000); // Vuelve a Home después de 2 segundos
      })
      .catch((err) => console.error(err));
  };

  if (!actividad) return <p>Cargando actividad...</p>;

  return (
    <div>
      <h2>Detalles</h2>
      <p>Listado o gestión de tareas.</p>
      <button onClick={() => navigate("/Home")}>← Volver a Home</button>
      <h1>DETALLES DE LA ACTIVIDAD</h1>

      <div className="detalle-card">
        <h3>{actividad.nombre}</h3>
        <p><strong>Descripción:</strong> {actividad.descripcion}</p>
        <p><strong>Horario:</strong> {actividad.horario}</p>
        <p><strong>Profesor:</strong> {actividad.profesor}</p>

        <button onClick={inscribirse} className="inscribirse-btn">
          Inscribirme
        </button>

        {mensaje && <p className="mensaje-exito">{mensaje}</p>}
      </div>
    </div>
  );
}

export default Detalle;



