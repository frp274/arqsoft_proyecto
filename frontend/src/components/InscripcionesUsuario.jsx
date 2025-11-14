import React, { useEffect, useState } from 'react';

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
    console.log("Payload del token:", payload);

    return {
      id: payload.jti || null,
      es_admin: payload.es_admin || false  // o 'Es_admin' si tu backend lo envía así
    };
  } catch (e) {
    console.error("Error al decodificar el token:", e);
    return null;
  }
}


function InscripcionesUsuario() {
  const [inscripciones, setInscripciones] = useState([]);
  const [actividades, setActividades] = useState({});
  const [loading, setLoading] = useState(true);
  const usuarioInfo = getUserInfoFromToken();
  const usuarioId = usuarioInfo?.id;

  useEffect(() => {
    const fetchInscripciones = async () => {
      if (!usuarioId) {
        setLoading(false);
        return;
      }

      try {
        const res = await fetch(`${process.env.REACT_APP_API_BUSQUEDAS_URL}/inscripciones/usuario/${usuarioId}`, {
          headers: {
            Authorization: `Bearer ${getCookie("token")}`
          }
        });
        
        if (res.ok) {
          const data = await res.json();
          setInscripciones(Array.isArray(data) ? data : []);
        } else {
          setInscripciones([]);
        }
      } catch (err) {
        console.error("Error al obtener inscripciones:", err);
        setInscripciones([]);
      } finally {
        setLoading(false);
      }
    };

    fetchInscripciones();
  }, [usuarioId]);

  if (loading) return <p>Cargando inscripciones...</p>;

  return (
    <div className="inscripciones-usuario">
      <h3>Mis Inscripciones</h3>
      {inscripciones.length === 0 ? (
        <p>No hay inscripciones recientes.</p>
      ) : (
        <ul>
          {inscripciones.map((insc) => {
            const actividad = actividades[insc.actividadId];
            return (
              <li key={insc.id}>
                <strong>{actividad?.nombre || 'Actividad desconocida'}:</strong> Horario ID {insc.horarioId}
              </li>
            );
          })}
        </ul>
      )}
    </div>
  );
}

export default InscripcionesUsuario;