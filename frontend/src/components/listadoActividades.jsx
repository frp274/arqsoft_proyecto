import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom";
import './listadoActividades.css';

function ListadoActividades({ filtro, refrescar, esAdmin }) {
  const [actividades, setActividades] = useState([]);
  const [editando, setEditando] = useState(null);
  const [formData, setFormData] = useState({ nombre: '', descripcion: '', profesor: '', horarios: [] });
  const navigate = useNavigate();

  const fetchActividades = async () => {
    try {
      const res = await fetch(`${process.env.REACT_APP_API_BUSQUEDAS_URL}/search/actividades?nombre=${filtro || ''}`);
      if (!res.ok) throw new Error("Error en la respuesta del servidor");
      const data = await res.json();
      setActividades(data.actividades || []);
    } catch (error) {
      console.error("Error al obtener actividades:", error);
      setActividades([]);
    }
  };

  useEffect(() => {
    fetchActividades();
  }, [filtro, refrescar]);

  const handleEliminar = async (id) => {
    if (!window.confirm("¿Estás seguro de que deseas eliminar esta actividad?")) return;

    try {
      const tokenCookie = document.cookie.split('; ').find(row => row.startsWith('token='));
      const token = tokenCookie ? tokenCookie.split('=')[1] : null;

      const res = await fetch(`${process.env.REACT_APP_API_ACTIVIDADES_URL}/actividad/${id}`, {
        method: 'DELETE',
        headers: {
          'Authorization': `Bearer ${token}`
        }
      });

      if (res.ok) {
        alert("Actividad eliminada correctamente");
        fetchActividades();
      } else {
        alert("Error al eliminar la actividad");
      }
    } catch (error) {
      console.error("Error al eliminar:", error);
    }
  };

  const handleEditar = (actividad) => {
    setEditando(actividad.id);
    setFormData({
      nombre: actividad.nombre,
      descripcion: actividad.descripcion,
      profesor: actividad.profesor,
      horarios: actividad.horarios.map(h => ({
        dia: h.dia,
        horarioInicio: h.horarioInicio,
        horarioFinal: h.horarioFinal,
        cupo: h.cupo
      }))
    });
  };

  const handleCancelar = () => {
    setEditando(null);
    setFormData({ nombre: '', descripcion: '', profesor: '', horarios: [] });
  };

  const handleActualizar = async (id) => {
    try {
      const tokenCookie = document.cookie.split('; ').find(row => row.startsWith('token='));
      const token = tokenCookie ? tokenCookie.split('=')[1] : null;

      const res = await fetch(`${process.env.REACT_APP_API_ACTIVIDADES_URL}/actividad/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify({
          nombre: formData.nombre,
          descripcion: formData.descripcion,
          profesor: formData.profesor,
          horarios: formData.horarios.map(h => ({
            dia: h.dia,
            horarioInicio: h.horarioInicio,
            horarioFinal: h.horarioFinal,
            cupo: parseInt(h.cupo)
          }))
        })
      });

      if (res.ok) {
        alert("Actividad actualizada correctamente");
        setEditando(null);
        fetchActividades();
      } else {
        alert("Error al actualizar la actividad");
      }
    } catch (error) {
      console.error("Error al actualizar:", error);
    }
  };

  return (
    <div className="listas">
      {filtro && actividades && actividades.length === 0 ? (
        <p className="error">⚠ No se encontró ninguna actividad con el nombre buscado.</p>
      ) : (
        <ul>
          {actividades && actividades.map((act) => (
            <li className="li" key={act.id}>
              <div
                className="boton-actividad"
                onClick={() => navigate(`/Detalle/${act.id}`)}
              >
                <div
                  className="contenido-actividad"
                  onClick={(e) => {
                    if (editando === act.id) {
                      e.stopPropagation();
                    }
                  }}
                >
                  {editando === act.id ? (
                    <div className="formulario-edicion">
                      <input
                        type="text"
                        value={formData.nombre}
                        onChange={(e) => setFormData({ ...formData, nombre: e.target.value })}
                        placeholder="Nombre"
                      />
                      <textarea
                        value={formData.descripcion}
                        onChange={(e) => setFormData({ ...formData, descripcion: e.target.value })}
                        placeholder="Descripción"
                      />
                      <input
                        type="text"
                        value={formData.profesor}
                        onChange={(e) => setFormData({ ...formData, profesor: e.target.value })}
                        placeholder="Profesor"
                      />
                      
                      <div className="horarios-edicion">
                        <h4>Horarios:</h4>
                        {formData.horarios.map((h, i) => (
                          <div key={i} className="horario-item">
                            <input
                              type="text"
                              value={h.dia}
                              onChange={(e) => {
                                const nuevosHorarios = [...formData.horarios];
                                nuevosHorarios[i].dia = e.target.value;
                                setFormData({ ...formData, horarios: nuevosHorarios });
                              }}
                              placeholder="Día"
                            />
                            <input
                              type="time"
                              value={h.horarioInicio}
                              onChange={(e) => {
                                const nuevosHorarios = [...formData.horarios];
                                nuevosHorarios[i].horarioInicio = e.target.value;
                                setFormData({ ...formData, horarios: nuevosHorarios });
                              }}
                            />
                            <input
                              type="time"
                              value={h.horarioFinal}
                              onChange={(e) => {
                                const nuevosHorarios = [...formData.horarios];
                                nuevosHorarios[i].horarioFinal = e.target.value;
                                setFormData({ ...formData, horarios: nuevosHorarios });
                              }}
                            />
                            <input
                              type="number"
                              value={h.cupo}
                              onChange={(e) => {
                                const nuevosHorarios = [...formData.horarios];
                                nuevosHorarios[i].cupo = e.target.value;
                                setFormData({ ...formData, horarios: nuevosHorarios });
                              }}
                              placeholder="Cupo"
                            />
                          </div>
                        ))}
                      </div>

                      <div className="botones-edicion">
                        <button onClick={(e) => { e.stopPropagation(); handleActualizar(act.id); }}>
                          💾 Guardar
                        </button>
                        <button onClick={(e) => { e.stopPropagation(); handleCancelar(); }}>
                          ❌ Cancelar
                        </button>
                      </div>
                    </div>
                  ) : (
                    <>
                      <p className="texto">
                        <strong>Nombre:</strong> {act.nombre}
                        <br />
                        <strong>Profesor:</strong> {act.profesor}
                      </p>
                      <p className="descripcion">{act.descripcion}</p>
                      <div className="horarios-lista">
                        <strong>Horarios:</strong>
                        {act.horarios && act.horarios.map((h, idx) => (
                          <div key={idx} className="horario-item">
                            {h.dia}: {h.horarioInicio} - {h.horarioFinal} (Cupo: {h.cupo})
                          </div>
                        ))}
                      </div>
                      
                      {esAdmin && (
                        <>
                          <button
                            className="boton-eliminar"
                            onClick={(e) => {
                              e.stopPropagation();
                              handleEliminar(act.id);
                            }}
                          >🗑️ Eliminar</button>
                          <button
                            className="boton-editar"
                            onClick={(e) => {
                              e.stopPropagation();
                              handleEditar(act);
                            }}
                          >✏️ Editar</button>
                        </>
                      )}
                    </>
                  )}
                </div>
              </div>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default ListadoActividades;
