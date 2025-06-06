/*import React, { useEffect, useState } from 'react';
import './App.css';

function App() {
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
    <div className="App">
      <header className="App-header">
        <h1></h1>
        <ul>
          {actividades.map(act => (
            <li key={act.id}>{act.nombre} - {act.descripcion}</li>
          ))}
        </ul>

        <h1>hola mundo</h1>
        <body>

          <p>chuaaaa</p>
      </body>
        
      </header>
      
      
      <footer>
          
          
          <p>este parrafo es de prueba para ver</p>
        </footer>
    </div>
  );
}

export default App;*/

// src/components/Login.jsx
import { useNavigate } from "react-router-dom";

function Login() {
  const navigate = useNavigate();

  const irAHome = () => {
    navigate("/Home");
  };

  return (
    <div>
      <h2>Login</h2>
      <p>Bienvenido. Iniciá sesión para continuar.</p>
      <button onClick={irAHome}>Ir a Home →</button>
    </div>
  );
}

export default Login;

