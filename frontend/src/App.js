
/*import './App.css';
import { componente1 } from './componets/componente1';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <h1>PROYECTO FINAL</h1>
    
      <footer>
        
        <h2>hola mundo</h2>
      <p>este parrafo es de prueba para ver</p>
      </footer>
      </header>
      
      
      
    </div>
  );
}

export default App;*/


// Ejemplo para obtener actividades y mostrarlas en tu App.js
import React, { useEffect, useState } from 'react';
import './App.css';

function App() {
  const [actividades, setActividades] = useState([]);

 useEffect(() => {
  fetch('http://localhost:8080/actividad')
    .then(response => response.json())
    .then(data => {
      console.log('Actividades:', data); // <-- Agrega este log
      setActividades(data);
    })
    .catch(error => console.error('Error:', error));
}, []);

  

  return (
    <div className="App">
      <header className="App-header">
        <h1>PROYECTO FINAL</h1>
        <ul>
          {actividades.map(act => (
            <li key={act.id}>{act.nombre} - {act.descripcion}</li>
          ))}
        </ul>
        <footer>
          <h2>hola mundo</h2>
          <p>este parrafo es de prueba para ver</p>
        </footer>
      </header>
    </div>
  );
}

export default App;