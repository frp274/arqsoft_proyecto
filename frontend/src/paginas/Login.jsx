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
import { useNavigate } from "react-router-dom";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "../components/ui/card";
import { Button } from "../components/ui/button";
import { Input } from "../components/ui/input";
import { Label } from "../components/ui/label";
import { Badge } from "../components/ui/badge";
import { Dumbbell, ArrowRight, Sparkles, Users, Calendar } from "lucide-react";

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
    <div className="min-h-screen bg-gradient-to-br from-background via-secondary/20 to-primary/10 flex items-center justify-center p-4">
      <div className="w-full max-w-6xl grid md:grid-cols-2 gap-8 items-center">
        
        {/* Panel izquierdo: branding */}
        <div className="hidden md:block space-y-8">
          <div className="flex items-center gap-3">
            <div className="h-12 w-12 rounded-full bg-primary flex items-center justify-center">
              <Dumbbell className="h-6 w-6 text-primary-foreground" />
            </div>
            <div>
              <h1 className="font-mono text-2xl font-bold uppercase tracking-tighter">
                GOOD GYM
              </h1>
              <Badge variant="secondary" className="mt-1">
                <Sparkles className="h-3 w-3 mr-1" />
                Fitness Platform
              </Badge>
            </div>
          </div>

          <div className="space-y-4">
            <h2 className="font-mono text-4xl font-bold uppercase tracking-tighter leading-tight">
              Transform Your<br />Fitness Journey
            </h2>
            <p className="text-lg text-muted-foreground max-w-md">
              Iniciá sesión para explorar actividades, gestionar tus inscripciones 
              y alcanzar tus objetivos fitness.
            </p>
          </div>

          <div className="flex flex-wrap gap-3">
            <Badge variant="outline" className="gap-2 py-2 px-4">
              <Users className="h-4 w-4" />
              Rutinas dinámicas
            </Badge>
            <Badge variant="outline" className="gap-2 py-2 px-4">
              <Calendar className="h-4 w-4" />
              Inscripciones online
            </Badge>
            <Badge variant="outline" className="gap-2 py-2 px-4">
              <Sparkles className="h-4 w-4" />
              Modo administrador
            </Badge>
          </div>

          <div className="pt-8">
            <p className="text-sm text-muted-foreground">
              ¿Todavía no tenés cuenta?{" "}
              <Button 
                variant="link" 
                className="p-0 h-auto font-semibold"
                onClick={irARegistro}
              >
                Crear cuenta
                <ArrowRight className="ml-1 h-4 w-4" />
              </Button>
            </p>
          </div>
        </div>

        {/* Panel derecho: formulario */}
        <Card className="w-full shadow-lg">
          <CardHeader>
            <CardTitle className="font-mono text-2xl uppercase tracking-tight">
              Iniciar sesión
            </CardTitle>
            <CardDescription>
              Ingresá tus credenciales para continuar
            </CardDescription>
          </CardHeader>
          <CardContent>
            <form onSubmit={handleLogin} className="space-y-6">
              <div className="space-y-2">
                <Label htmlFor="usuario">Usuario</Label>
                <Input
                  id="usuario"
                  type="text"
                  placeholder="Tu nombre de usuario"
                  value={usuario}
                  onChange={(e) => setUsuario(e.target.value)}
                  className="h-11"
                />
              </div>

              <div className="space-y-2">
                <Label htmlFor="contrasenia">Contraseña</Label>
                <Input
                  id="contrasenia"
                  type="password"
                  placeholder="••••••••"
                  value={contrasenia}
                  onChange={(e) => setContrasenia(e.target.value)}
                  className="h-11"
                />
              </div>

              {error && (
                <div className="text-sm text-destructive bg-destructive/10 border border-destructive/30 rounded-md p-3">
                  {error}
                </div>
              )}

              <Button type="submit" className="w-full h-11 font-semibold" size="lg">
                Ingresar
              </Button>

              <div className="md:hidden text-center pt-4">
                <p className="text-sm text-muted-foreground">
                  ¿No tienes cuenta?{" "}
                  <Button 
                    variant="link" 
                    className="p-0 h-auto"
                    onClick={irARegistro}
                    type="button"
                  >
                    Regístrate aquí
                  </Button>
                </p>
              </div>
            </form>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}

export default Login;
