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
import DosCampos from '../components/camposLogin';
import './Login.css';

function Login() {
  const navigate = useNavigate();

  const irAHome = () => {
    navigate("/Home");
  };

  return (
    <div className="login">
      
      <h2 className="titulo">💪🏼 GOOD GYM 🦵🏼</h2>
      <div className="sb">
      <p className="subtitulo">Bienvenido. Ingrese su usuario para acceder : </p>
      </div>
      
      <p className="espacio"/>
      <DosCampos></DosCampos>
      <p/>
      <div className="boton">
      <button onClick={irAHome} className="ingresar" >  I N G R E S A R  </button>
      </div>
    </div>
  );
}

export default Login;

