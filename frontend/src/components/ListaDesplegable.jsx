/*import react from 'react';

function ListaDesplegable() {
  return (
    <div>
      <label htmlFor="opciones">Elige una opción:  </label>
      <select id="opciones">
        <option value="opcion1">Opción 1</option>
        <option value="opcion2">Opción 2</option>
        <option value="opcion3">Opción 3</option>
      </select>
    </div>
  );
}*/

import React, { useState } from 'react';
import './lista.css';

function ListaDesplegable({ horarios }) {
  const [seleccion, setSeleccion] = useState('');

  const horariosSeguros = Array.isArray(horarios) ? horarios : [];

  const manejarCambio = (e) => {
    e.stopPropagation();
    setSeleccion(e.target.value);
  };

  const manejarClick = (e) => {
    e.stopPropagation();
  };

  return (
    <div onClick={manejarClick}>
      <select
        id="opciones"
        value={seleccion}
        onChange={manejarCambio}
        className="lista"
        onClick={(e) => e.stopPropagation()} 
      >
        <option value="">Horarios</option>
        {horariosSeguros.map((h, index) => (
          <option key={index} value={`${h.dia}-${h.hora}`}>
            {h.dia} - {h.hora}
          </option>
        ))}
      </select>
    </div>
  );
}

export default ListaDesplegable; // 🔥 Este export es obligatorio
