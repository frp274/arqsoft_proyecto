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



























/*import React, { useState } from "react";
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
      const response = await fetch(`${process.env.REACT_APP_API_USUARIOS_URL}/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          username: usuario,
          password: contrasenia,
        }),
      });
      // const data = await response.json();
      // console.log("Respuesta del backend:", data); // <-- AGREGÁ ESTO
      // localStorage.setItem("token", data.token);
      
// Después de recibir la respuesta del login
      if (response.ok) {
        const data = await response.json();
        console.log("RESPUESTA DEL LOGIN:", data); // Verifica la respuesta

        // Guardar el token en la cookie
        // document.cookie = `token=${data.token}; path=/; SameSite=Strict`;
        document.cookie = `token=; path=/; expires=Thu, 01 Jan 1970 00:00:00 UTC;`;
        document.cookie = `token=${data.token}; path=/; SameSite=Strict; Secure`;

        // Guardar el usuarioId en la cookie (opcional)
        //document.cookie = `userId=${data.id}; path=/; Secure; HttpOnly`;

        // Redirigir al Home
        if (data.es_admin === true){
          navigate("/Admin");
        }
        else{
          navigate("/home");
        }
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

          <p style={{ marginTop: "15px", fontSize: "0.9rem" }}>
            ¿No tienes cuenta?{" "}
            <span
              onClick={() => navigate("/Registro")}
              style={{
                color: "#4CAF50",
                cursor: "pointer",
                textDecoration: "underline"
              }}
            >
              Regístrate aquí
            </span>
          </p>
        </div>
      </form>
    </div>
  );
}

export default Login;
*/

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
      const response = await fetch(
        `${process.env.REACT_APP_API_USUARIOS_URL}/login`,
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            username: usuario,
            password: contrasenia,
          }),
        }
      );

      if (response.ok) {
        const data = await response.json();
        console.log("RESPUESTA DEL LOGIN:", data);

        // limpiar token anterior y setear el nuevo
        document.cookie =
          "token=; path=/; expires=Thu, 01 Jan 1970 00:00:00 UTC;";
        document.cookie = `token=${data.token}; path=/; SameSite=Strict; Secure`;

        // redirigir según rol
        if (data.es_admin === true) {
          navigate("/Admin");
        } else {
          navigate("/Home");
        }
      } else {
        setError("Usuario o contraseña incorrectos.");
      }
    } catch (error) {
      console.error("Error al conectar con el servidor:", error);
      setError("Error al conectar con el servidor.");
    }
  };

  const irARegistro = () => navigate("/Registro");

  return (
    <div className="login-page">
      <div className="login-shell">
        {/* Panel izquierdo: info / branding */}
        <div className="login-info-panel">
          <div className="login-logo-pill">
            <span className="login-logo-dot" />
            <span>GOOD GYM · RETRO FUTURE</span>
          </div>

          <h1 className="login-title">Bienvenido a GOOD GYM</h1>

          <p className="login-subtitle">
            Iniciá sesión para explorar actividades, gestionar tus
            inscripciones y entrenar con estilo retro-futurista.
          </p>

          <div className="login-chips">
            <span className="login-chip">Rutinas dinámicas</span>
            <span className="login-chip">Modo administrador</span>
            <span className="login-chip">Inscripciones online</span>
          </div>

          <p className="login-small-text">
            ¿Todavía no tenés cuenta?
            <button
              type="button"
              className="login-link-button"
              onClick={irARegistro}
            >
              Crear cuenta
            </button>
          </p>
        </div>

        {/* Panel derecho: formulario de login */}
        <div className="login-form-wrapper">
          <form className="login-form" onSubmit={handleLogin}>
            <h2 className="login-form-title">Iniciar sesión</h2>
            <p className="login-form-subtitle">
              Ingresá tus credenciales para continuar.
            </p>

            <div className="login-field-group">
              <label className="login-label" htmlFor="usuario">
                Usuario
              </label>
              <input
                id="usuario"
                className="login-input"
                type="text"
                placeholder="Tu nombre de usuario"
                value={usuario}
                onChange={(e) => setUsuario(e.target.value)}
              />
            </div>

            <div className="login-field-group">
              <label className="login-label" htmlFor="contrasenia">
                Contraseña
              </label>
              <input
                id="contrasenia"
                className="login-input"
                type="password"
                placeholder="••••••••"
                value={contrasenia}
                onChange={(e) => setContrasenia(e.target.value)}
              />
            </div>

            {error && <p className="login-error">{error}</p>}

            <button className="login-submit" type="submit">
              Ingresar
            </button>

            <p className="login-bottom-text">
              ¿No tienes cuenta?
              <button
                type="button"
                className="login-link-inline"
                onClick={irARegistro}
              >
                Regístrate aquí
              </button>
            </p>
          </form>
        </div>
      </div>
    </div>
  );
}

export default Login;
