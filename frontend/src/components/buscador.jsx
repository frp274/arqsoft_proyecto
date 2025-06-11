import React, { useState } from 'react';
import './buscador.css';



function Buscador() {
  const [actividad, setActividad] = useState('');
  

  const manejarCambioActividad = (e) => setActividad(e.target.value);
  

  return (
    <div className='campos'>
      <input 
        className='placeholder'
        type="text" 
        value={actividad} 
        onChange={manejarCambioActividad} 
        placeholder="-- Buscar actividad 🔎 --"
      />
    </div>
  );
}

export default Buscador;