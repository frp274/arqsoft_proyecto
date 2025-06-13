
import { useNavigate, useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import axios from "axios";
import './Detalle.css';



function getUserIdFromToken() {
  
  const token = getCookie("token");
  if (token) {
    console.log("Token obtenido desde la cookie:", token);
  } else {
    console.log("No se encontró el token en las cookies");
  }

  // Verifica si el token tiene tres partes (header, payload, signature)
  const parts = token.split('.');
  if (parts.length !== 3) {
    console.error("Token JWT inválido");
    return null;
  }

  try {
    const payload = JSON.parse(atob(parts[1]));  // Decodificar el payload del token
    console.log(payload); // Verifica si contiene el 'id' que buscas
    return payload.jti || null;  // Retorna el ID del usuario
  } catch (e) {
    console.error("Error al decodificar el token:", e);
    return null;
  }
}

// function getCookie(name) {
//   const value = `; ${document.cookie}`;
//   const parts = value.split(`; ${name}=`);
//   if (parts.length === 2) return parts.pop().split(';').shift();
// }

function getCookie(name) {
  const nameEQ = `${name}=`;
  const ca = document.cookie.split(';');
  
  for (let i = 0; i < ca.length; i++) {
    let c = ca[i].trim();
    if (c.indexOf(nameEQ) === 0) {
      return c.substring(nameEQ.length, c.length);
    }
  }
  return null; // Si no se encuentra la cookie, devuelve null
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

    const cupo = actividad.cupo || actividad.Cupo;
    if (cupo === 0) {
      setMensaje("No hay cupos disponibles para esta actividad.");
      return;
    }
    const Id = parseInt(id, 10);
    const horarioidint = parseInt(horarioId, 10);
    const usuarioidint = parseInt(usuarioId, 10);
    axios.post('http://localhost:8080/inscripcion', {
      actividad_id: Id,
      horario_id: horarioidint,
      usuario_id: usuarioidint
    }, {
      headers: { Authorization: `Bearer ${getCookie("token")}`, "Content-Type": "application/json"
                  
    }
    })
      .then(() => {
        setMensaje('¡Inscripción exitosa!');
        setTimeout(() => window.location.reload(), 2000);  // Recarga la página después de 2 segundos
      })
      .catch((err) => {
        setMensaje("Error al inscribirse.");
        console.error(err);
      });
  };

  if (!actividad) return <p>Cargando actividad...</p>;

  const horarios = actividad.horarios || actividad.Horarios || [];

    const obtenerImagenActividad = (nombre) => {
    if (!nombre) return "/default.jpg";
    const nombreNormalizado = nombre.toLowerCase();
    switch (nombreNormalizado) {
      case "pilates":
        return "/pilates.png";
      case "mma":
        return "/mma.png";
      case "musculacion":
        return "/musculacion.png";
      case "zumba":
        return "/zumba.png";
      case "spinning":
        return "/spining.png";
      default:
        return "/logo192.png";
    }
  };

  return (
    <div>
      
      <button className="boton-home" onClick={() => navigate("/Home")}>← Volver a Home</button>
      <div className="detalles">
        <h1>DETALLES DE LA ACTIVIDAD</h1>

        <div className="detalle-card">
          <h3 className="nact">{actividad.nombre || actividad.Nombre}</h3>
          <p className="desc"><strong>Descripción:</strong> {actividad.descripcion || actividad.Descripcion}</p>
          <p className="desc"><strong>Profesor:</strong> {actividad.profesor || actividad.Profesor}</p>
          <p className="desc"><strong>Cupo disponible:</strong> {actividad.cupo || actividad.Cupo}</p>

          <p><strong>Horarios:</strong></p>
          <ul className="ul">
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

          

          {mensaje && <p className="mensaje-exito">{mensaje}</p>}

          <img 
            src={obtenerImagenActividad(actividad.nombre || actividad.Nombre)} 
            alt="Imagen de la actividad" 
            style={{ width: "400px", marginTop: "20px", borderRadius: "8px" }}
          />

        </div>
      </div>
    </div>
  );
}

export default Detalle;

