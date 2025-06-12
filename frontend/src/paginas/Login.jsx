// import React, { useState } from "react";
// import { useNavigate } from "react-router-dom"; // Si usás react-router

// const Login = () => {
//   const [usuario, setUsuario] = useState("");
//   const [contrasenia, setContrasenia] = useState("");
//   const [error, setError] = useState("");
//   const navigate = useNavigate(); // para redireccionar

//   const handleLogin = async (e) => {
//     e.preventDefault();
//     setError(""); // limpia error

//     // Validación en el FRONT
//     if (!usuario || !contrasenia) {
//       setError("Debe completar los campos.");
//       return; // No hace la request
//     }

//     // Si pasa la validación, hace la request al backend
//     try {
//       const response = await fetch("http://localhost:8080/login", {
//         method: "POST",
//         headers: { "Content-Type": "application/json" },
//         body: JSON.stringify({ UserName: usuario, Password:contrasenia }),
//       });

//       if (response.ok) {
//         const data = await response.json();
//         localStorage.setItem("token", data.token); // Guarda el token
//         navigate("/home"); // Redirige a Home si login OK
//       } else {
//         setError("Usuario o contraseña incorrectos."); // Error de backend
//       }

//     } catch (err) {
//       setError("Error de conexión al servidor.");
//     }
//   };

//   return (
//     <div className="login-container">
//       <form onSubmit={handleLogin}>
//         <input
//           type="text"
//           placeholder="Usuario"
//           value={usuario}
//           onChange={(e) => setUsuario(e.target.value)}
//         />
//         <input
//           type="password"
//           placeholder="Contraseña"
//           value={contrasenia}
//           onChange={(e) => setContrasenia(e.target.value)}
//         />
//         <button type="submit">Ingresar</button>
//         {error && <div className="error">{error}</div>}
//       </form>
 
//     </div>
//   );
// };

// export default Login;



// // import { useNavigate } from "react-router-dom";
// // import DosCampos from '../components/camposLogin';
// // import './Login.css';

// // function Login() {
// //   const navigate = useNavigate();

// //   const irAHome = () => {
// //     navigate("/Home");
// //   };

// //   return (
// //     <div className="login">
      
// //       <h2 className="titulo">GOOD GYM</h2>
// //       <p>Bienvenido. Ingrese su usuario para acceder : </p>
      

// //       <hr/>
// //       <p/>
// //       <DosCampos></DosCampos>
// //       <p/>
// //       <div className="boton">
// //       <button onClick={irAHome} className="ingresar" >  I N G R E S A R  </button>
// //       </div>
// //     </div>
// //   );
// // }

// // export default Login;

import React, { useState } from "react";
import "./Login.css";
import { useNavigate } from "react-router-dom";

function Login() {
  const [usuario, setUsuario] = useState("");
  const [contrasenia, setContrasenia] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleLogin = async (event) => {
    event.preventDefault();

    if (!usuario || !contrasenia) {
      setError("Debe completar ambos campos.");
      return;
    }

    try {
      const response = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          Username: usuario,
          Password: contrasenia,
        }),
      });
      // const data = await response.json();
      // console.log("Respuesta del backend:", data); // <-- AGREGÁ ESTO
      // localStorage.setItem("token", data.token);
      
      if (response.ok) {
        const data = await response.json();
        console.log("RESPUESTA DEL LOGIN:", data); // <-- AGREGÁ ESTA LÍNEA
        console.log(localStorage.getItem("token"))
        localStorage.setItem("token", data.token);
        navigate("/Home");
      } else {
        setError("Usuario o contraseña incorrectos.");
      }
    } catch (error) {
      setError("Error al conectar con el servidor.");
    }
  };

  return (
    <div className="login-container">
      <form className="login-form" onSubmit={handleLogin}>
        <h1 className="titulo"> 💪🏼 GOOD GYM 🦵🏼 </h1>
        
        <div className="boton">
          <input
            className="usuario"
            type="text"
            placeholder="Usuario"
            value={usuario}
            onChange={(e) => setUsuario(e.target.value)}
          />
          <p/>
          <input
            className="contra"
            type="password"
            placeholder="Contraseña"
            value={contrasenia}
            onChange={(e) => setContrasenia(e.target.value)}
          />
          <p/>
        
          <button className="ingresar" type="submit">Ingresar</button>
          {error && <p className="login-error">{error}</p>}
        </div>
      </form>
    </div>
  );
}

export default Login;
