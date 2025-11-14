/*import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import "./Registro.css";

function Registro() {
  const [formData, setFormData] = useState({
    username: "",
    email: "",
    nombre: "",
    apellido: "",
    password: "",
    confirmPassword: ""
  });
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");

    // Validaciones
    if (!formData.username || !formData.email || !formData.nombre || 
        !formData.apellido || !formData.password || !formData.confirmPassword) {
      setError("Todos los campos son obligatorios.");
      return;
    }

    // Validar email
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(formData.email)) {
      setError("El email no es válido.");
      return;
    }

    // Validar contraseñas
    if (formData.password !== formData.confirmPassword) {
      setError("Las contraseñas no coinciden.");
      return;
    }

    if (formData.password.length < 6) {
      setError("La contraseña debe tener al menos 6 caracteres.");
      return;
    }

    setLoading(true);

    try {
      const response = await fetch(`${process.env.REACT_APP_API_USUARIOS_URL}/usuario`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          username: formData.username,
          email: formData.email,
          nombre: formData.nombre,
          apellido: formData.apellido,
          password: formData.password,
          es_admin: false
        }),
      });


      if (response.ok) {
        alert("¡Registro exitoso! Ahora puedes iniciar sesión.");
        navigate("/Login");
      } else {
        const errorData = await response.json();
        setError(errorData.message || "Error al registrar el usuario. El nombre de usuario o email ya existe.");
      }
    } catch (error) {
      setError("Error al conectar con el servidor.");
      console.error("Error de registro:", error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="login-container">
      <form className="login-form" onSubmit={handleSubmit}>
        <h1 className="titulo">💪🏼 GOOD GYM 🦵🏼</h1>
        <h2 style={{ fontSize: "1.2rem", marginBottom: "20px" }}>Crear Cuenta Nueva</h2>

        <div className="boton">
          <input
            className="usuario"
            type="text"
            name="username"
            placeholder="Usuario *"
            value={formData.username}
            onChange={handleChange}
            disabled={loading}
          />
          <p />
          
          <input
            className="usuario"
            type="email"
            name="email"
            placeholder="Email *"
            value={formData.email}
            onChange={handleChange}
            disabled={loading}
          />
          <p />

          <input
            className="usuario"
            type="text"
            name="nombre"
            placeholder="Nombre *"
            value={formData.nombre}
            onChange={handleChange}
            disabled={loading}
          />
          <p />

          <input
            className="usuario"
            type="text"
            name="apellido"
            placeholder="Apellido *"
            value={formData.apellido}
            onChange={handleChange}
            disabled={loading}
          />
          <p />

          <input
            className="contra"
            type="password"
            name="password"
            placeholder="Contraseña *"
            value={formData.password}
            onChange={handleChange}
            disabled={loading}
          />
          <p />

          <input
            className="contra"
            type="password"
            name="confirmPassword"
            placeholder="Confirmar Contraseña *"
            value={formData.confirmPassword}
            onChange={handleChange}
            disabled={loading}
          />
          <p />

          <button className="ingresar" type="submit" disabled={loading}>
            {loading ? "Registrando..." : "Crear Cuenta"}
          </button>

          {error && <p className="login-error">{error}</p>}

          <p style={{ marginTop: "15px", fontSize: "0.9rem" }}>
            ¿Ya tienes cuenta?{" "}
            <span
              onClick={() => navigate("/Login")}
              style={{
                color: "#4CAF50",
                cursor: "pointer",
                textDecoration: "underline"
              }}
            >
              Inicia sesión aquí
            </span>
          </p>
        </div>
      </form>
    </div>
  );
}

export default Registro;
*/

import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import "./Registro.css";



function Registro() {
  const [formData, setFormData] = useState({
    username: "",
    email: "",
    nombre: "",
    apellido: "",
    password: "",
    confirmPassword: ""
  });
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");

    // Validaciones
    if (
      !formData.username ||
      !formData.email ||
      !formData.nombre ||
      !formData.apellido ||
      !formData.password ||
      !formData.confirmPassword
    ) {
      setError("Todos los campos son obligatorios.");
      return;
    }

    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(formData.email)) {
      setError("El email no es válido.");
      return;
    }

    if (formData.password !== formData.confirmPassword) {
      setError("Las contraseñas no coinciden.");
      return;
    }

    if (formData.password.length < 6) {
      setError("La contraseña debe tener al menos 6 caracteres.");
      return;
    }

    setLoading(true);

    try {
      const response = await fetch(
        `${process.env.REACT_APP_API_USUARIOS_URL}/usuario`,
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            username: formData.username,
            email: formData.email,
            nombre: formData.nombre,
            apellido: formData.apellido,
            password: formData.password,
            es_admin: false
          })
        }
      );

      if (response.ok) {
        alert("¡Registro exitoso! Ahora puedes iniciar sesión.");
        navigate("/Login");
      } else {
        const errorData = await response.json().catch(() => ({}));
        setError(
          errorData.error ||
            errorData.message ||
            "Error al registrar el usuario. El nombre de usuario o email ya existe."
        );
      }
    } catch (error) {
      console.error("Error de registro:", error);
      setError("Error al conectar con el servidor.");
    } finally {
      setLoading(false);
    }
  };

  const irALogin = () => navigate("/Login");

  return (
    <div className="login-page">
      <div className="login-shell">
        {/* Panel izquierdo: info / branding */}
        <div className="login-info-panel">
          <div className="login-logo-pill">
            <span className="login-logo-dot" />
            <span>GOOD GYM · RETRO FUTURE</span>
          </div>

          <h1 className="login-title">Crear cuenta en GOOD GYM</h1>

          <p className="login-subtitle">
            Registrate para poder buscar actividades, administrar tus clases e
            inscribirte online en el gimnasio con estética retro-futurista.
          </p>

          <div className="login-chips">
            <span className="login-chip">Registro rápido</span>
            <span className="login-chip">Perfil de usuario</span>
            <span className="login-chip">Acceso al sistema</span>
          </div>

          <p className="login-small-text">
            ¿Ya tenés una cuenta?
            <button
              type="button"
              className="login-link-button"
              onClick={irALogin}
            >
              Iniciar sesión
            </button>
          </p>
        </div>

        {/* Panel derecho: formulario de registro */}
        <div className="login-form-wrapper">
          <form className="login-form" onSubmit={handleSubmit}>
            <h2 className="login-form-title">Crear cuenta nueva</h2>
            <p className="login-form-subtitle">
              Completá tus datos para registrarte.
            </p>

            <div className="login-field-group">
              <label className="login-label" htmlFor="username">
                Usuario
              </label>
              <input
                id="username"
                className="login-input"
                type="text"
                name="username"
                placeholder="Usuario *"
                value={formData.username}
                onChange={handleChange}
                disabled={loading}
              />
            </div>

            <div className="login-field-group">
              <label className="login-label" htmlFor="email">
                Email
              </label>
              <input
                id="email"
                className="login-input"
                type="email"
                name="email"
                placeholder="Email *"
                value={formData.email}
                onChange={handleChange}
                disabled={loading}
              />
            </div>

            <div className="login-field-group">
              <label className="login-label" htmlFor="nombre">
                Nombre
              </label>
              <input
                id="nombre"
                className="login-input"
                type="text"
                name="nombre"
                placeholder="Nombre *"
                value={formData.nombre}
                onChange={handleChange}
                disabled={loading}
              />
            </div>

            <div className="login-field-group">
              <label className="login-label" htmlFor="apellido">
                Apellido
              </label>
              <input
                id="apellido"
                className="login-input"
                type="text"
                name="apellido"
                placeholder="Apellido *"
                value={formData.apellido}
                onChange={handleChange}
                disabled={loading}
              />
            </div>

            <div className="login-field-group">
              <label className="login-label" htmlFor="password">
                Contraseña
              </label>
              <input
                id="password"
                className="login-input"
                type="password"
                name="password"
                placeholder="Contraseña *"
                value={formData.password}
                onChange={handleChange}
                disabled={loading}
              />
            </div>

            <div className="login-field-group">
              <label className="login-label" htmlFor="confirmPassword">
                Confirmar contraseña
              </label>
              <input
                id="confirmPassword"
                className="login-input"
                type="password"
                name="confirmPassword"
                placeholder="Confirmar contraseña *"
                value={formData.confirmPassword}
                onChange={handleChange}
                disabled={loading}
              />
            </div>

            {error && <p className="login-error">{error}</p>}

            <button className="login-submit" type="submit" disabled={loading}>
              {loading ? "Registrando..." : "Crear cuenta"}
            </button>

            <p className="login-bottom-text">
              ¿Ya tienes cuenta?
              <button
                type="button"
                className="login-link-inline"
                onClick={irALogin}
              >
                Inicia sesión aquí
              </button>
            </p>
          </form>
        </div>
      </div>
    </div>
  );
}

export default Registro;
