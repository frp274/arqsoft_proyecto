import React from 'react';
import './buscador.css';

function Buscador({ setFiltro }) {
  const manejarCambioActividad = (e) => setFiltro(e.target.value);

  return (
    <div className='campos'>
      <input 
        className='placeholder'
        type="text" 
        onChange={manejarCambioActividad} 
        placeholder="-- Buscar actividad 🔎 --"
      />
    </div>
  );
}

export default Buscador;
