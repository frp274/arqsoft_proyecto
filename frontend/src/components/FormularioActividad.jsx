// Nuevo archivo: src/components/FormularioActividad.jsx

import { useState } from 'react';

function FormularioActividad({ onSubmit }) {
  const [nombre, setNombre] = useState('');
  const [descripcion, setDescripcion] = useState('');
  const [profesor, setProfesor] = useState('');
  const [horarios, setHorarios] = useState([{ dia: '', horarioInicio: '', horarioFinal: '', cupo: 0 }]);

  const handleHorarioChange = (index, field, value) => {
    const nuevosHorarios = [...horarios];
    nuevosHorarios[index][field] = value;
    setHorarios(nuevosHorarios);
  };

  const agregarHorario = () => {
    setHorarios([...horarios, { dia: '', horarioInicio: '', horarioFinal: '', cupo: 0 }]);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit({ nombre, descripcion, profesor, horarios });
    setNombre('');
    setDescripcion('');
    setProfesor('');
    setHorarios([{ dia: '', horarioInicio: '', horarioFinal: '', cupo: 0 }]);
  };

  return (
    <form onSubmit={handleSubmit} className="formulario-crear-actividad">
      <label>Nombre:</label>
      <input value={nombre} onChange={e => setNombre(e.target.value)} required />

      <label>Descripción:</label>
      <input value={descripcion} onChange={e => setDescripcion(e.target.value)} required />

      <label>Profesor:</label>
      <input value={profesor} onChange={e => setProfesor(e.target.value)} required />

      <label>Horarios:</label>
      {horarios.map((horario, index) => (
        <div key={index} className="bloque-horario">
          <input placeholder="Día" value={horario.dia} onChange={e => handleHorarioChange(index, 'dia', e.target.value)} required />
          <input placeholder="Hora inicio" value={horario.horarioInicio} onChange={e => handleHorarioChange(index, 'horarioInicio', e.target.value)} required />
          <input placeholder="Hora fin" value={horario.horarioFinal} onChange={e => handleHorarioChange(index, 'horarioFinal', e.target.value)} required />
          <input placeholder="Cupo" type="number" value={horario.cupo} onChange={e => handleHorarioChange(index, 'cupo', e.target.value)} required />
        </div>
      ))}
      <button type="button" onClick={agregarHorario}>Agregar horario</button>

      <button type="submit">Guardar Actividad</button>
    </form>
  );
}

export default FormularioActividad;
