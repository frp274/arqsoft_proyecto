import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom";
import './Home.css';
import Foto from '../components/Foto';
import ListaDesplegable from '../components/ListaDesplegable';
import Buscador from "../components/buscador";
import ListadoActividades from "../components/listadoActividades";

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

        <ListadoActividades filtro={filtro} refrescar={refrescar} />
      </div>
    </div>
  );
}

export default HomeAdm;
