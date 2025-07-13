/*import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom";
import './listadoActividades.css';
import ListaDesplegable from './ListaDesplegable';


function ListadoActividades({ filtro }) {
  const [actividades, setActividades] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    fetch('http://localhost:8080/actividad')
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
    fetch('http://localhost:8080/actividad')
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


import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom";
import './listadoActividades.css';
import ListaDesplegable from './ListaDesplegable';

function ListadoActividades({ filtro }) {
  const [actividades, setActividades] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    fetch('http://localhost:8080/actividad')
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

export default ListadoActividades;

