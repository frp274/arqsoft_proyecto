/*import React, { useEffect, useState, useContext } from 'react';
import { useNavigate } from "react-router-dom";
import './Home.css';
import Foto from '../components/Foto';
import ListaDesplegable from '../components/ListaDesplegable';
import Buscador from "../components/buscador";
import ListadoActividades from "../components/listadoActividades";
import InscripcionesUsuario from '../components/InscripcionesUsuario';


// function getUserIdFromToken() {
  
//   const token = getCookie("token");
//   if (token) {
//     console.log("Token obtenido desde la cookie:", token);
//   } else {
//     console.log("No se encontró el token en las cookies");
//   }

//   // Verifica si el token tiene tres partes (header, payload, signature)
//   const parts = token.split('.');
//   if (parts.length !== 3) {
//     console.error("Token JWT inválido");
//     return null;
//   }

//   try {
//     const payload = JSON.parse(atob(parts[1]));  // Decodificar el payload del token
//     console.log(payload); // Verifica si contiene el 'id' que buscas
//     return payload.jti || null;  // Retorna el ID del usuario
//   } catch (e) {
//     console.error("Error al decodificar el token:", e);
//     return null;
//   }
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
// const json_info  = getUserInfoFromToken();
// const usuario_id = json_info.id;
// const usuario_es_admin = json_info.es_admin;


  
function HomeAdm() {
  const navigate = useNavigate();
  if (getUserInfoFromToken().es_admin === false){
    navigate("/Home");
  }
  const [filtro, setFiltro] = useState('');
  const [mostrarFormulario, setMostrarFormulario] = useState(false);
  const [nombre, setNombre] = useState('');
  const [descripcion, setDescripcion] = useState('');
  const [profesor, setProfesor] = useState('');
  const [horarios, setHorarios] = useState([{ dia: '', horarioInicio: '', horarioFinal: '', cupo: 0 }]);
  const [refrescar, setRefrescar] = useState(false);
<<<<<<< HEAD
  const [errorText, setErrorText] = useState('');
  const usuarioData = localStorage.getItem("usuario");
  const usuario = usuarioData ? JSON.parse(usuarioData) : null;

  



=======
  const [errorText, setErrorText] = useState('');  
>>>>>>> e43fbceea6b284179ea7a0af3f9f64d24a2875ac
  const handleAgregarHorario = () => {
    setHorarios([...horarios, { dia: '', horarioInicio: '', horarioFinal: '', cupo: 0 }]);
  };

  const handleHorarioChange = (index, field, value) => {
    const nuevosHorarios = [...horarios];
    nuevosHorarios[index][field] = value;
    setHorarios(nuevosHorarios);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const nuevaActividad = {
      nombre,
      descripcion,
      profesor,
      horarios: horarios.map(h => ({
        dia: h.dia,
        horarioInicio: h.horarioInicio,
        horarioFinal: h.horarioFinal,
        cupo: parseInt(h.cupo, 10)
      }))
    };

    try {
      const response = await fetch('http://localhost:8080/actividad', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(nuevaActividad)
      });

      if (response.ok) {
        alert("Actividad creada correctamente");
        setRefrescar(prev => !prev); // Fuerza recarga
        setNombre('');
        setDescripcion('');
        setProfesor('');
        setHorarios([{ dia: '', horarioInicio: '', horarioFinal: '', cupo: 0 }]);
        setMostrarFormulario(false);
        setErrorText('');
      } else {
        setErrorText("Error al crear actividad");
      }
    } catch (error) {
      console.error("Error en la solicitud:", error);
      setErrorText("Error en la solicitud al servidor");
    }
  };

  useEffect(() => {
    setRefrescar(prev => !prev); // Inicial para cargar desde el inicio
  }, []);

 

  return (
    <div className="home">
      <hr />
      <h1 className="Titulo">G O O D   G Y M - ADMIN</h1>

      <div className="foto">
        <Foto />

        <div style={{ textAlign: 'right', margin: '10px' }}>
          <button onClick={() => setMostrarFormulario(!mostrarFormulario)}>
            {mostrarFormulario ? 'Cerrar Formulario' : 'Agregar Actividad'}
          </button>
        </div>

        {mostrarFormulario && (
          <form onSubmit={handleSubmit} className="formulario-actividad">
            <h3>Formulario de Nueva Actividad</h3>
            <input type="text" placeholder="Nombre" value={nombre} onChange={e => setNombre(e.target.value)} required />
            <input type="text" placeholder="Descripción" value={descripcion} onChange={e => setDescripcion(e.target.value)} required />
            <input type="text" placeholder="Profesor" value={profesor} onChange={e => setProfesor(e.target.value)} required />

            {horarios.map((h, index) => (
              <div key={index} style={{ marginBottom: '10px' }}>
                <input type="text" placeholder="Día" value={h.dia} onChange={e => handleHorarioChange(index, 'dia', e.target.value)} required />
                <input type="text" placeholder="Hora Inicio" value={h.horarioInicio} onChange={e => handleHorarioChange(index, 'horarioInicio', e.target.value)} required />
                <input type="text" placeholder="Hora Fin" value={h.horarioFinal} onChange={e => handleHorarioChange(index, 'horarioFinal', e.target.value)} required />
                <input type="number" placeholder="Cupo" value={h.cupo} onChange={e => handleHorarioChange(index, 'cupo', e.target.value)} required />
              </div>
            ))}
            <button type="button" onClick={handleAgregarHorario}>Agregar otro horario</button>
            <br />
            <button type="submit">Guardar Actividad</button>
            {errorText && <p className="error" style={{ color: 'red' }}>{errorText}</p>}
          </form>
        )}

        <p className="espacio" />
        <h1 className='subtitulo'>ACTIVIDADES DISPONIBLES</h1>
        <p className="espacio" />

        <Buscador setFiltro={setFiltro} />
        <p className="espacio" />

        <ListadoActividades filtro={filtro} refrescar={refrescar} esAdmin={true} />

        <p className="espacio" />

        <div style={{ marginTop: '5rem', marginLeft: '5rem' }}>
          {usuario ? (
            <InscripcionesUsuario usuarioId={usuario.id} />
          ) : (
            <p>No se pudo obtener el usuario</p>
          )}

        </div>

      </div>
    </div>
  );
}

export default HomeAdm;*/

import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom";
import './Home.css';
import Foto from '../components/Foto';
import ListaDesplegable from '../components/ListaDesplegable';
import Buscador from "../components/buscador";
import ListadoActividades from "../components/listadoActividades";
import InscripcionesUsuario from '../components/InscripcionesUsuario';



/*function getCookiee(name) {
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

function getUserIdFromTokenn() {
  const token = getCookiee("token");
  if (token) {
    console.log("Token obtenido desde la cookie:", token);
  } else {
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
    return payload.jti || null;
  } catch (e) {
    console.error("Error al decodificar el token:", e);
    return null;
  }
}*/


function HomeAdm() {
  const navigate = useNavigate();
  const [filtro, setFiltro] = useState('');
  const [mostrarFormulario, setMostrarFormulario] = useState(false);
  const [nombre, setNombre] = useState('');
  const [descripcion, setDescripcion] = useState('');
  const [profesor, setProfesor] = useState('');
  const [horarios, setHorarios] = useState([{ dia: '', horarioInicio: '', horarioFinal: '', cupo: 0 }]);
  const [refrescar, setRefrescar] = useState(false);
  const [errorText, setErrorText] = useState('');

  

  const handleAgregarHorario = () => {
    setHorarios([...horarios, { dia: '', horarioInicio: '', horarioFinal: '', cupo: 0 }]);
  };

  const handleHorarioChange = (index, field, value) => {
    const nuevosHorarios = [...horarios];
    nuevosHorarios[index][field] = value;
    setHorarios(nuevosHorarios);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const nuevaActividad = {
      nombre,
      descripcion,
      profesor,
      horarios: horarios.map(h => ({
        dia: h.dia,
        horarioInicio: h.horarioInicio,
        horarioFinal: h.horarioFinal,
        cupo: parseInt(h.cupo, 10)
      }))
    };

    try {
      const response = await fetch('http://localhost:8080/actividad', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(nuevaActividad)
      });

      if (response.ok) {
        alert("Actividad creada correctamente");
        setRefrescar(prev => !prev); // Fuerza recarga
        setNombre('');
        setDescripcion('');
        setProfesor('');
        setHorarios([{ dia: '', horarioInicio: '', horarioFinal: '', cupo: 0 }]);
        setMostrarFormulario(false);
        setErrorText('');
      } else {
        setErrorText("Error al crear actividad");
      }
    } catch (error) {
      console.error("Error en la solicitud:", error);
      setErrorText("Error en la solicitud al servidor");
    }
  };

  useEffect(() => {
    setRefrescar(prev => !prev); // Inicial para cargar desde el inicio
  }, []);

  /*useEffect(() => {
    console.log("ID de usuario desde el token:", usuarioId);
  }, [usuarioId]);*/


  return (
    <div className="home">
      <hr />
      <h1 className="Titulo">G O O D   G Y M - ADMIN</h1>

      <div className="foto">
        <Foto />

        <div style={{ textAlign: 'right', margin: '10px' }}>
          <button onClick={() => setMostrarFormulario(!mostrarFormulario)}>
            {mostrarFormulario ? 'Cerrar Formulario' : 'Agregar Actividad'}
          </button>
        </div>

        {mostrarFormulario && (
          <form onSubmit={handleSubmit} className="formulario-actividad">
            <h3>Formulario de Nueva Actividad</h3>
            <input type="text" placeholder="Nombre" value={nombre} onChange={e => setNombre(e.target.value)} required />
            <input type="text" placeholder="Descripción" value={descripcion} onChange={e => setDescripcion(e.target.value)} required />
            <input type="text" placeholder="Profesor" value={profesor} onChange={e => setProfesor(e.target.value)} required />

            {horarios.map((h, index) => (
              <div key={index} style={{ marginBottom: '10px' }}>
                <input type="text" placeholder="Día" value={h.dia} onChange={e => handleHorarioChange(index, 'dia', e.target.value)} required />
                <input type="text" placeholder="Hora Inicio" value={h.horarioInicio} onChange={e => handleHorarioChange(index, 'horarioInicio', e.target.value)} required />
                <input type="text" placeholder="Hora Fin" value={h.horarioFinal} onChange={e => handleHorarioChange(index, 'horarioFinal', e.target.value)} required />
                <input type="number" placeholder="Cupo" value={h.cupo} onChange={e => handleHorarioChange(index, 'cupo', e.target.value)} required />
              </div>
            ))}
            <button type="button" onClick={handleAgregarHorario}>Agregar otro horario</button>
            <br />
            <button type="submit">Guardar Actividad</button>
            {errorText && <p className="error" style={{ color: 'red' }}>{errorText}</p>}
          </form>
        )}

        <p className="espacio" />
        <h1 className='subtitulo'>ACTIVIDADES DISPONIBLES</h1>
        <p className="espacio" />

        <Buscador setFiltro={setFiltro} />
        <p className="espacio" />

        <ListadoActividades filtro={filtro} refrescar={refrescar} esAdmin={true} />

        <p className="espacio" />

        

      </div>
    </div>
  );
  }


export default HomeAdm;


