import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom";
import './Detalle.css';

function getCookie(name) {
  const nameEQ = `${name}=`;
  const ca = document.cookie.split(';');
  
  for (let i = 0; i < ca.length; i++) {
    let c = ca[i].trim();
    if (c.indexOf(nameEQ) === 0) {
      return c.substring(nameEQ.length, c.length);
    }
  }
  return null;
}

function getUserInfoFromToken() {
  const token = getCookie("token");
  if (!token) {
    console.log("No se encontró el token en las cookies");
    return null;
  }

  const parts = token.split('.');
  if (parts.length !== 3) {
    console.error("Token JWT inválido");
    return null;
  }

  try {
    const payload = JSON.parse(atob(parts[1]));
    return {
      id: payload.jti || null,
      es_admin: payload.es_admin || false
    };
  } catch (e) {
    console.error("Error al decodificar el token:", e);
    return null;
  }
}

function MisInscripciones() {
  const navigate = useNavigate();
  const [inscripciones, setInscripciones] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const usuarioInfo = getUserInfoFromToken();
  const usuarioId = usuarioInfo?.id;
  const esAdmin = usuarioInfo?.es_admin;

  useEffect(() => {
    if (!usuarioId) {
      setError("Debes iniciar sesión para ver tus inscripciones.");
      setLoading(false);
      setTimeout(() => navigate("/Login"), 2000);
      return;
    }

    const fetchInscripciones = async () => {
      try {
        const response = await fetch(
          `${process.env.REACT_APP_API_BUSQUEDAS_URL}/inscripciones/usuario/${usuarioId}`,
          {
            headers: {
              Authorization: `Bearer ${getCookie("token")}`
            }
          }
        );

        if (response.ok) {
          const data = await response.json();
          console.log("Inscripciones recibidas:", data);
          setInscripciones(Array.isArray(data) ? data : []);
        } else if (response.status === 404) {
          setInscripciones([]);
        } else {
          setError("Error al cargar las inscripciones.");
        }
      } catch (err) {
        console.error("Error al obtener inscripciones:", err);
        setError("Error de conexión al servidor.");
      } finally {
        setLoading(false);
      }
    };

    fetchInscripciones();
  }, [usuarioId, navigate]);

  const obtenerImagenActividad = (nombre) => {
    if (!nombre) return "/logo192.png";
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

  const handleVerDetalle = (actividadId) => {
    navigate(`/Detalle/${actividadId}`);
  };

  const volverAlHome = () => {
    if (esAdmin) {
      navigate("/Admin");
    } else {
      navigate("/Home");
    }
  };

  if (loading) {
    return (
      <div className="detalles">
        <h1>Cargando inscripciones...</h1>
      </div>
    );
  }

  return (
    <div>
      <button className="boton-home" onClick={volverAlHome}>
        ← Volver al Inicio
      </button>

      <div className="detalles">
        <h1>MIS INSCRIPCIONES</h1>

        {error && <p className="mensaje-exito" style={{ color: 'red' }}>{error}</p>}

        {!error && inscripciones.length === 0 ? (
          <div className="detalle-card">
            <p className="desc">No tienes inscripciones activas en este momento.</p>
            <p className="desc" style={{ marginTop: '20px' }}>
              ¡Explora nuestras actividades y encuentra la perfecta para ti!
            </p>
            <button 
              className="inscribirse-btn" 
              onClick={volverAlHome}
              style={{ marginTop: '20px' }}
            >
              Ver Actividades Disponibles
            </button>
          </div>
        ) : (
          <div className="inscripciones-grid">
            {inscripciones.map((actividad) => (
              <div key={actividad.Id || actividad.id} className="detalle-card">
                <img 
                  src={obtenerImagenActividad(actividad.Nombre || actividad.nombre)} 
                  alt={actividad.Nombre || actividad.nombre}
                  style={{ 
                    width: "100%", 
                    maxWidth: "300px", 
                    height: "200px", 
                    objectFit: "cover",
                    borderRadius: "8px",
                    marginBottom: "15px"
                  }}
                />
                
                <h3 className="nact">{actividad.Nombre || actividad.nombre}</h3>
                <p className="desc">
                  <strong>Descripción:</strong> {actividad.Descripcion || actividad.descripcion}
                </p>
                <p className="desc">
                  <strong>Profesor:</strong> {actividad.Profesor || actividad.profesor}
                </p>

                {(actividad.Horarios || actividad.horarios) && (
                  <div>
                    <p className="desc"><strong>Horarios:</strong></p>
                    <ul className="ul">
                      {(actividad.Horarios || actividad.horarios).map((h, idx) => (
                        <li key={idx}>
                          {h.Dia || h.dia}: {h.HorarioInicio || h.horarioInicio} - {h.HorarioFinal || h.horarioFinal}
                        </li>
                      ))}
                    </ul>
                  </div>
                )}

                <button 
                  className="inscribirse-btn" 
                  onClick={() => handleVerDetalle(actividad.Id || actividad.id)}
                  style={{ marginTop: '15px' }}
                >
                  Ver Detalles
                </button>
              </div>
            ))}
          </div>
        )}
      </div>

      <style jsx>{`
        .inscripciones-grid {
          display: grid;
          grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
          gap: 30px;
          padding: 20px;
          max-width: 1200px;
          margin: 0 auto;
        }

        @media (max-width: 768px) {
          .inscripciones-grid {
            grid-template-columns: 1fr;
          }
        }
      `}</style>
    </div>
  );
}

export default MisInscripciones;
