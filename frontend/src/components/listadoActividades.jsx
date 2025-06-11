import React, { useEffect, useState } from 'react';
import './listadoActividades.css';
import ListaDesplegable from './ListaDesplegable';


function ListadoActividades() {
  const [actividades, setActividades] = useState([]);

 useEffect(() => {
  fetch('http://localhost:8080/actividad')
    .then(response => response.json())
    .then(data => {
      console.log('Actividades:', data); 
      setActividades(data);
    })
    .catch(error => console.error('Error:', error));
}, []);

  

  return (
    <div>
     
        
        <div className='listas'>
            <ul>
            {actividades.map(act => (
                <li className='li' key={act.id}>
                    <button className='botones'>
                    <p className='texto'> Nombre: {act.nombre} ----------------------------------------------------------------------------- Profesor: {act.profesor}</p> 
                    <p> <ListaDesplegable/> </p>
                    </button>
                    </li>
            ))}
            </ul>
        </div>

    </div>
  );
}

export default ListadoActividades;