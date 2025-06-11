// import { useNavigate, useParams } from "react-router-dom";
// import { useEffect, useState } from "react";
// import axios from "axios";
// import './Detalle.css';

// function Detalle() {
//   const { id } = useParams(); // Tomamos el ID de la URL
//   const navigate = useNavigate();
//   const [actividad, setActividad] = useState(null);
//   const [mensaje, setMensaje] = useState('');

//   // Cargar detalles de la actividad
//   useEffect(() => {
//     axios.get(`http://localhost:8080/actividad/${id}`)
//       .then((res) => setActividad(res.data))
//       .catch((err) => console.error(err));
//   }, [id]);

//   // Función para inscribirse
//   const inscribirse = () => {
//     axios.post('http://localhost:8080/inscripcion', { actividadId: id })
//       .then(() => {
//         setMensaje('¡Inscripción exitosa!');
//         setTimeout(() => navigate('/Home'), 2000); // Vuelve a Home después de 2 segundos
//       })
//       .catch((err) => console.error(err));
//   };

//   if (!actividad) return <p>Cargando actividad...</p>;

//   return (
//     <div>
//       <h2>Detalles</h2>
//       <p>Listado o gestión de tareas.</p>
//       <button onClick={() => navigate("/Home")}>← Volver a Home</button>
//       <h1>DETALLES DE LA ACTIVIDAD</h1>

//       <div className="detalle-card">
//         <h3>{actividad.nombre}</h3>
//         <p><strong>Descripción:</strong> {actividad.descripcion}</p>
//         <p><strong>Horarios:</strong></p>
//         <ul>
//           {actividad.horarios && actividad.horarios.map((h, idx) => (
//             <li key={idx}>
//               {h.dia}: {h.horarioInicio} - {h.horarioFinal}
//             </li>
//           ))}
//         </ul>

//         <p><strong>Profesor:</strong> {actividad.profesor}</p>

//         <button onClick={inscribirse} className="inscribirse-btn">
//           Inscribirme
//         </button>

//         {mensaje && <p className="mensaje-exito">{mensaje}</p>}
//       </div>
//     </div>
//   );
// }

// export default Detalle;


import { useNavigate, useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import axios from "axios";
import './Detalle.css';

// Función para extraer el userId del token
function getUserIdFromToken() {
  const token = localStorage.getItem("token");
  if (!token) return null;
  try {
    const payload = JSON.parse(atob(token.split('.')[1]));
    // Usá el nombre correcto del campo según cómo guardás el id en el token
    return payload.id || payload.userId || payload.UserId || null;
  } catch (e) {
    return null;
  }
}

function Detalle() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [actividad, setActividad] = useState(null);
  const [mensaje, setMensaje] = useState('');

  useEffect(() => {
    axios.get(`http://localhost:8080/actividad/${id}`)
      .then((res) => setActividad(res.data))
      .catch((err) => console.error(err));
  }, [id]);

  // Nueva función: inscribirse a un horario puntual
  const inscribirseHorario = (horarioId) => {
    const usuarioId = getUserIdFromToken();
    if (!usuarioId) {
      setMensaje("Debe iniciar sesión para inscribirse.");
      setTimeout(() => navigate("/Login"), 2000);
      return;
    }
    axios.post('http://localhost:8080/inscripcion', {
      actividadId: id,
      horarioId: horarioId,
      usuarioId: usuarioId
    }, {
      headers: { Authorization: `Bearer ${localStorage.getItem("token")}` }
    })
      .then(() => {
        setMensaje('¡Inscripción exitosa!');
        setTimeout(() => navigate('/Home'), 2000);
      })
      .catch((err) => {
        setMensaje("Error al inscribirse.");
        console.error(err);
      });
  };

  if (!actividad) return <p>Cargando actividad...</p>;

  const horarios = actividad.horarios || actividad.Horarios || [];

  return (
    <div>
      <h2>Detalles</h2>
      <p>Listado o gestión de tareas.</p>
      <button onClick={() => navigate("/Home")}>← Volver a Home</button>
      <h1>DETALLES DE LA ACTIVIDAD</h1>

      <div className="detalle-card">
        <h3>{actividad.nombre || actividad.Nombre}</h3>
        <p><strong>Descripción:</strong> {actividad.descripcion || actividad.Descripcion}</p>
        <p><strong>Horarios:</strong></p>
        <ul>
          {horarios.length > 0 ? horarios.map((h, idx) => (
            <li key={idx} style={{display:"flex", alignItems:"center", justifyContent:"space-between"}}>
              <span>
                {h.dia || h.Dia}: {h.horarioInicio || h.horarioinicio || h.HorarioInicio} - {h.horarioFinal || h.horariofinal || h.HorarioFinal}
              </span>
              <button className="inscribirse-btn" onClick={() => inscribirseHorario(h.id || h.Id)}>
                Inscribirme a este horario
              </button>
            </li>
          )) : <li>No hay horarios cargados.</li>}
        </ul>

        <p><strong>Profesor:</strong> {actividad.profesor || actividad.Profesor}</p>

        {mensaje && <p className="mensaje-exito">{mensaje}</p>}
      </div>
    </div>
  );
}

export default Detalle;

