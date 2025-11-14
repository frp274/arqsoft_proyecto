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
          Username: formData.username,
          Email: formData.email,
          Nombre: formData.nombre,
          Apellido: formData.apellido,
          Password: formData.password,
          EsAdmin: false
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
