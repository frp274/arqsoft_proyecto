import React, { useState } from 'react';
import './camposLogin.css';



function DosCampos() {
  const [usuario, setUsuario] = useState('');
  const [contrasenia, setContrasenia] = useState('');

  const manejarCambioUsuario = (e) => setUsuario(e.target.value);
  const manejarCambioContrasenia = (e) => setContrasenia(e.target.value);

  return (
    <div className='campos'>
      <input 
        className='placeholder'
        type="text" 
        value={usuario} 
        onChange={manejarCambioUsuario} 
        placeholder="-- Escribe tu Usuario --"
      />

      <p/>

      <input 
        className='placeholder'
        type="password" 
        value={contrasenia} 
        onChange={manejarCambioContrasenia} 
        placeholder="-- Escribe tu Contrasenia --"
      />

    </div>
  );
}

export default DosCampos;