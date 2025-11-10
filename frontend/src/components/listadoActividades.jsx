/*import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom";
import './listadoActividades.css';
import ListaDesplegable from './ListaDesplegable';


function ListadoActividades({ filtro }) {
  const [actividades, setActividades] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    fetch('${process.env.REACT_APP_API_ACTIVIDADES_URL}/actividad')
      .then(response => response.json())
      .then(data => setActividades(data))
      .catch(error => console.error('Error:', error));
  }, []);

  const actividadesFiltradas = actividades.filter(act =>
    act.nombre.toLowerCase().includes(filtro.toLowerCase())
  );

  return (
    <div>
      <div className='listas'>
        {filtro && actividadesFiltradas.length === 0 ? (
          <p className='error'>⚠ No se encontró ninguna actividad con el nombre buscado.</p>
        ) : (
          <ul>
            {actividadesFiltradas.map(act => (
              <li className='li' key={act.id}>
                <button
                  className='botones'
                  onClick={() => navigate(`/Detalle/${act.id}`)}
                >
                  <p className='texto'>
                    <p >Nombre: {act.nombre} ----------------------------------- Profesor: {act.profesor}</p>
                  </p>
                  <p> <ListaDesplegable horarios={act.horarios || []} /> </p>
                </button>
              </li>
            ))}
          </ul>
        )}
      </div>
    </div>
  );
}

export default ListadoActividades;*/

/*import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom";
import './listadoActividades.css';
import ListaDesplegable from './ListaDesplegable';

function ListadoActividades() {
  const [actividades, setActividades] = useState([]);
  const navigate = useNavigate();
  


  useEffect(() => {
    fetch('${process.env.REACT_APP_API_ACTIVIDADES_URL}/actividad')
      .then(response => response.json())
      .then(data => setActividades(data))
      .catch(error => console.error('Error:', error));
  }, []);

    const actividadesFiltradas = actividades.filter(act =>
    act.nombre.toLowerCase().includes(filtro.toLowerCase())
     );

  return (
    <div>
      <div className='listas'>
        {filtro && actividadesFiltradas.length === 0 ? (
          <p className='error'>⚠ No se encontró ninguna actividad con el nombre buscado.</p>
        ) : (
        <ul>
          {actividadesFiltradas.map(act => (
           <li className="li" key={act.id}>
                <button
                className="boton-actividad"
                onClick={() => navigate(`/Detalle/${act.id}`)}
                >
                    <div
                    className="contenido-actividad"
                    onClick={(e) => e.stopPropagation()} // esto previene el click solo si es dentro del select
                    >
                        <span className="actividad-nombre">Nombre: {act.nombre}</span>
                        <span className="actividad-profesor">Profesor: {act.profesor}</span>
                        <ListaDesplegable horarios={act.horarios} />
                    </div>
                </button>
            </li>

          ))}
        </ul>)}
      </div>
    </div>
  );
}

export default ListadoActividades;*/


/*import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom";
import './listadoActividades.css';
import ListaDesplegable from './ListaDesplegable';

function ListadoActividades({ filtro }) {
  const [actividades, setActividades] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    fetch('${process.env.REACT_APP_API_ACTIVIDADES_URL}/actividad')
      .then(response => response.json())
      .then(data => setActividades(data))
      .catch(error => console.error('Error:', error));
  }, []);

  const actividadesFiltradas = actividades.filter(act =>
    act.nombre.toLowerCase().includes(filtro.toLowerCase())
  );

  return (
    <div>
      <div className='listas'>
        {filtro && actividadesFiltradas.length === 0 ? (
          <p className='error'>⚠ No se encontró ninguna actividad con el nombre buscado.</p>
        ) : (
          <ul>
            {actividadesFiltradas.map(act => (
              <li className="li" key={act.id}>
                <button
                  className="boton-actividad"
                  onClick={() => navigate(`/Detalle/${act.id}`)}
                >
                  <div
                    className="contenido-actividad"
                    onClick={(e) => e.stopPropagation()}
                  >
                    <span className="actividad-nombre">Nombre: {act.nombre}</span>
                    <span className="actividad-profesor">Profesor: {act.profesor}</span>
                    
                    <div className="actividad-horarios">
                        {act.horarios && act.horarios.length > 0 ? (
                            act.horarios.map((h, idx) => (
                            <span key={idx} className="horario-item">
                                {h.dia}: {h.horarioInicio} - {h.horarioFinal}
                            </span>
                            ))
                        ) : (
                            <span className="horario-item">Sin horarios</span>
                        )}
                    </div>
                  </div>
                </button>
              </li>
            ))}
          </ul>
        )}
      </div>
    </div>
  );
}

export default ListadoActividades;*/

/*import { useEffect, useState } from "react";

function ListadoActividades({ filtro }) {
  const [actividades, setActividades] = useState([]);

  useEffect(() => {
    const fetchActividades = async () => {
      try {
        const res = await fetch(`${process.env.REACT_APP_API_BUSQUEDAS_URL}/search/actividades?nombre=${filtro}`);
        const data = await res.json();
        setActividades(data);
      } catch (error) {
        console.error("Error al obtener actividades:", error);
      }
    };

    fetchActividades();
  }, [filtro]);

  return (
    <ul>
      {actividades.map((act) => (
        <li key={act.id}>{act.nombre}</li>
      ))}
    </ul>
  );
}

export default ListadoActividades;*/

/*import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom";
import './listadoActividades.css';

function ListadoActividades({ filtro }) {
  const [actividades, setActividades] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchActividades = async () => {
      try {
        const res = await fetch(`${process.env.REACT_APP_API_BUSQUEDAS_URL}/search/actividades?nombre=${filtro}`);
        const data = await res.json();
        setActividades(data);
      } catch (error) {
        console.error("Error al obtener actividades:", error);
      }
    };

    fetchActividades();
  }, [filtro]);

  return (
    <div>
      <div className='listas'>
        {filtro && actividades.length === 0 ? (
          <p className='error'>⚠ No se encontró ninguna actividad con el nombre buscado.</p>
        ) : (
          <ul>
            {actividades.map(act => (
              <li className="li" key={act.id}>
                <button
                  className="boton-actividad"
                  onClick={() => navigate(`/Detalle/${act.id}`)}
                >
                  <div
                    className="contenido-actividad"
                    onClick={(e) => e.stopPropagation()}
                  >
                    <span className="actividad-nombre">Nombre: {act.nombre}</span>
                    <span className="actividad-profesor">Profesor: {act.profesor}</span>
                    
                    <div className="actividad-horarios">
                      {act.horarios && act.horarios.length > 0 ? (
                        act.horarios.map((h, idx) => (
                          <span key={idx} className="horario-item">
                            {h.dia}: {h.horarioInicio} - {h.horarioFinal}
                          </span>
                        ))
                      ) : (
                        <span className="horario-item">Sin horarios</span>
                      )}
                    </div>
                  </div>
                </button>
              </li>
            ))}
          </ul>
        )}
      </div>
    </div>
  );
}

export default ListadoActividades;*/

/*import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom";
import './listadoActividades.css';

function ListadoActividades({ filtro, refrescar, esAdmin }) {
  const [actividades, setActividades] = useState([]);
  const navigate = useNavigate();

  const fetchActividades = async () => {
    try {
      const res = await fetch(`${process.env.REACT_APP_API_BUSQUEDAS_URL}/search/actividades?nombre=${filtro}`);
      if (!res.ok) throw new Error("Error en la respuesta del servidor");
      const data = await res.json();
      setActividades(data);
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
      const res = await fetch(`${process.env.REACT_APP_API_ACTIVIDADES_URL}/actividad/${id}`, {
        method: 'DELETE'
      });

      if (res.ok) {
        alert("Actividad eliminada correctamente");
        fetchActividades(); // Actualiza lista
      } else {
        alert("Error al eliminar la actividad");
      }
    } catch (error) {
      console.error("Error al eliminar:", error);
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
              <button
                className="boton-actividad"
                onClick={() => navigate(`/Detalle/${act.id}`)}
              >
                <div
                  className="contenido-actividad"
                  onClick={(e) => e.stopPropagation()}
                >
                  <span className="actividad-nombre">Nombre: {act.nombre}</span>
                  <span className="actividad-profesor">Profesor: {act.profesor}</span>
                  <div className="actividad-horarios">
                    {Array.isArray(act.horarios) && act.horarios.length > 0 ? (
                      act.horarios.map((h, idx) => (
                        <span key={idx} className="horario-item">
                          {h.dia}: {h.horarioInicio} - {h.horarioFinal} | Cupo: {h.cupo}
                        </span>
                      ))
                    ) : (
                      <span className="horario-item">Sin horarios</span>
                    )}
                  </div>
                  {esAdmin && (
                    <button
                      className="boton-eliminar"
                      onClick={(e) => {
                        e.stopPropagation();
                        handleEliminar(act.id);
                      }}
                    >
                      ❌ Eliminar
                    </button>
                  )}
                </div>
              </button>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default ListadoActividades;*/

/*import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom";
import './listadoActividades.css';

function ListadoActividades({ filtro, refrescar, esAdmin }) {
  const [actividades, setActividades] = useState([]);
  const [editando, setEditando] = useState(null);
  const [formData, setFormData] = useState({ nombre: '', descripcion: '', profesor: '', horarios: [] });
  const navigate = useNavigate();

  const fetchActividades = async () => {
    try {
      const res = await fetch(`${process.env.REACT_APP_API_BUSQUEDAS_URL}/search/actividades?nombre=${filtro}`);
      if (!res.ok) throw new Error("Error en la respuesta del servidor");
      const data = await res.json();
      setActividades(data);
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
      const res = await fetch(`${process.env.REACT_APP_API_ACTIVIDADES_URL}/actividad/${id}`, {
        method: 'DELETE'
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
      horarios: actividad.horarios || []
    });
  };

  const handleHorarioChange = (index, field, value) => {
    const nuevosHorarios = [...formData.horarios];
    nuevosHorarios[index][field] = value;
    setFormData({ ...formData, horarios: nuevosHorarios });
  };

  const handleGuardarEdicion = async (id) => {
    try {
      const res = await fetch(`${process.env.REACT_APP_API_ACTIVIDADES_URL}/actividad/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(formData)
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
              <button
                className="boton-actividad"
                onClick={() => navigate(`/Detalle/${act.id}`)}
              >
                <div
                  className="contenido-actividad"
                  onClick={(e) => e.stopPropagation()}
                >
                  {editando === act.id ? (
                    <div className="form-edicion">
                      <input type="text" value={formData.nombre} onChange={(e) => setFormData({ ...formData, nombre: e.target.value })} />
                      <input type="text" value={formData.descripcion} onChange={(e) => setFormData({ ...formData, descripcion: e.target.value })} />
                      <input type="text" value={formData.profesor} onChange={(e) => setFormData({ ...formData, profesor: e.target.value })} />
                      {formData.horarios.map((h, i) => (
                        <div key={i} className="edit-horario">
                          <input type="text" value={h.dia} onChange={(e) => handleHorarioChange(i, 'dia', e.target.value)} placeholder="Día" />
                          <input type="text" value={h.horarioInicio} onChange={(e) => handleHorarioChange(i, 'horarioInicio', e.target.value)} placeholder="Inicio" />
                          <input type="text" value={h.horarioFinal} onChange={(e) => handleHorarioChange(i, 'horarioFinal', e.target.value)} placeholder="Fin" />
                          <input type="number" value={h.cupo} onChange={(e) => handleHorarioChange(i, 'cupo', e.target.value)} placeholder="Cupo" />
                        </div>
                      ))}
                      <button type="button" onClick={(e) => { e.stopPropagation(); handleGuardarEdicion(act.id); }}>💾 Guardar</button>
                      <button type="button" onClick={(e) => { e.stopPropagation(); setEditando(null); }}>✖ Cancelar</button>
                    </div>
                  ) : (
                    <>
                      <span className="actividad-nombre">Nombre: {act.nombre}</span>
                      <span className="actividad-profesor">Profesor: {act.profesor}</span>
                      <div className="actividad-horarios">
                        {Array.isArray(act.horarios) && act.horarios.length > 0 ? (
                          act.horarios.map((h, idx) => (
                            <span key={idx} className="horario-item">
                              {h.dia}: {h.horarioInicio} - {h.horarioFinal} | Cupo: {h.cupo}
                            </span>
                          ))
                        ) : (
                          <span className="horario-item">Sin horarios</span>
                        )}
                      </div>
                      {esAdmin && (
                        <>
                          <button
                            className="boton-eliminar"
                            onClick={(e) => {
                              e.stopPropagation();
                              handleEliminar(act.id);
                            }}
                          >❌ Eliminar</button>
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
              </button>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default ListadoActividades;*/

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
      const res = await fetch(`${process.env.REACT_APP_API_BUSQUEDAS_URL}/search/actividades?nombre=${filtro}`);
      if (!res.ok) throw new Error("Error en la respuesta del servidor");
      const data = await res.json();
      setActividades(data);
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
      const res = await fetch(`${process.env.REACT_APP_API_ACTIVIDADES_URL}/actividad/${id}`, {
        method: 'DELETE'
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
        id: h.id,
        dia: h.dia || '',
        horaInicio: h.horarioInicio || '',  // del DTO original
        horaFin: h.horarioFinal || '',      // del DTO original
        cupo: h.cupo !== undefined ? Number(h.cupo) : 0
      }))
    });
  };

  const handleHorarioChange = (index, field, value) => {
    setFormData((prev) => {
      const nuevosHorarios = [...prev.horarios];
      nuevosHorarios[index] = {
        ...nuevosHorarios[index],
        [field]: field === 'cupo' ? Number(value) : value
      };
      return { ...prev, horarios: nuevosHorarios };
    });
  };

  const handleGuardarEdicion = async (id) => {
    try {
      const res = await fetch(`${process.env.REACT_APP_API_ACTIVIDADES_URL}/actividad/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          id,
          nombre: formData.nombre,
          descripcion: formData.descripcion,
          profesor: formData.profesor,
          horarios: formData.horarios.map(h => ({
          id: h.id,
          dia: h.dia,
          horarioInicio: h.horaInicio,  // 👈 correcto
          horarioFinal: h.horaFin,      // 👈 correcto
          cupo: Number(h.cupo)

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
                  onClick={(e) => e.stopPropagation()}
                  onKeyDown={(e) => {
                    if (e.key === ' ') e.stopPropagation();
                  }}
                >
                  {editando === act.id ? (
                    <div className="form-edicion">
                      <input type="text" value={formData.nombre} onChange={(e) => setFormData({ ...formData, nombre: e.target.value })} />
                      <input type="text" value={formData.descripcion} onChange={(e) => setFormData({ ...formData, descripcion: e.target.value })} />
                      <input type="text" value={formData.profesor} onChange={(e) => setFormData({ ...formData, profesor: e.target.value })} />
                      {formData.horarios.map((h, i) => (
                        <div key={i} className="edit-horario">
                          <input type="text" value={h.dia} onChange={(e) => handleHorarioChange(i, 'dia', e.target.value)} placeholder="Día" />
                          <input type="text" value={h.horaInicio} onChange={(e) => handleHorarioChange(i, 'horaInicio', e.target.value)} placeholder="Inicio" />
                          <input type="text" value={h.horaFin} onChange={(e) => handleHorarioChange(i, 'horaFin', e.target.value)} placeholder="Fin" />
                          <input type="number" value={h.cupo} onChange={(e) => handleHorarioChange(i, 'cupo', e.target.value)} placeholder="Cupo" />
                        </div>
                      ))}
                      <button type="button" onClick={(e) => { e.stopPropagation(); handleGuardarEdicion(act.id); }}>💾 Guardar</button>
                      <button type="button" onClick={(e) => { e.stopPropagation(); setEditando(null); }}>✖ Cancelar</button>
                    </div>
                  ) : (
                    <>
                      <span className="actividad-nombre">Nombre: {act.nombre}</span>
                      <span className="actividad-profesor">Profesor: {act.profesor}</span>
                      <div className="actividad-horarios">
                        {Array.isArray(act.horarios) && act.horarios.length > 0 ? (
                          act.horarios.map((h, idx) => (
                            <span key={idx} className="horario-item">
                              {h.dia}: {h.horarioInicio} - {h.horarioFinal} | Cupo: {h.cupo}
                            </span>
                          ))
                        ) : (
                          <span className="horario-item">Sin horarios</span>
                        )}
                      </div>
                      {esAdmin && (
                        <>
                          <button
                            className="boton-eliminar"
                            onClick={(e) => {
                              e.stopPropagation();
                              handleEliminar(act.id);
                            }}
                          >❌ Eliminar</button>
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


