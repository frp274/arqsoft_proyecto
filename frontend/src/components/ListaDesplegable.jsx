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

import { useState } from 'react';
import './lista.css';

function ListaDesplegable() {
  const [seleccion, setSeleccion] = useState('');

  const manejarCambio = (e) => {
    setSeleccion(e.target.value);
  };

  return (
    <div>
      <label htmlFor="opciones">Seleccioná una opción:  </label>
      <select id="opciones" value={seleccion} onChange={manejarCambio} className='lista'>
        <option value="">-- Elegir --</option>
        <option value="a">Opción A</option>
        <option value="b">Opción B</option>
        <option value="c">Opción C</option>
      </select>

      <p>Seleccionaste: {seleccion}</p>
    </div>
  );
}

export default ListaDesplegable;