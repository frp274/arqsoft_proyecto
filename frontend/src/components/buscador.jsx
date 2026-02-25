import React from 'react';
import './buscador.css';

function Buscador({ setFiltro, classNameOverride }) {
  const [valor, setValor] = React.useState('');

  const manejarCambioActividad = (e) => setValor(e.target.value);

  const manejarTecla = (e) => {
    if (e.key === 'Enter') {
      setFiltro(valor);
    }
  };

  return (
    <div className='campos'>
      <input
        className={classNameOverride || 'placeholder'}
        type="text"
        value={valor}
        onChange={manejarCambioActividad}
        onKeyDown={manejarTecla}
        placeholder="-- Buscar actividad 🔎 --"
      />
    </div>
  );
}

export default Buscador;
