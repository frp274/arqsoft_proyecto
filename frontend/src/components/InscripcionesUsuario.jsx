import React, { useEffect, useState } from 'react';

function InscripcionesUsuario({ usuarioId }) {
  const [inscripciones, setInscripciones] = useState([]);
  const [actividades, setActividades] = useState({});
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchInscripciones = async () => {
      try {
        const res = await fetch('http://localhost:8080/inscripciones/usuario/${usuarioId}');
        const data = await res.json();
        setInscripciones(data);

        // Obtener detalles de actividades
        const actividadesMap = {};
        for (const insc of data) {
          if (!actividadesMap[insc.actividadId]) {
            const actRes = await fetch('http://localhost:8080/actividad/${insc.actividadId}');
            const actData = await actRes.json();
            actividadesMap[insc.actividadId] = actData;
          }
        }
        setActividades(actividadesMap);
      } catch (err) {
        console.error("Error al obtener inscripciones:", err);
      } finally {
        setLoading(false);
      }
    };

    if (usuarioId) {
      fetchInscripciones();
    }
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