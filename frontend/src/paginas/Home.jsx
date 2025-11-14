import { useNavigate } from "react-router-dom";
import './Home.css';
import Buscador from "../components/buscador";
import ListadoActividades from "../components/listadoActividades";
import { useState } from "react";

function Home() {
  const navigate = useNavigate();
  const [filtro, setFiltro] = useState('');
  const [refrescar, setRefrescar] = useState(false);

  return (
    <div className="home">
      {/* Header moderno con gradiente */}
      <header className="home-header">
        <div className="header-content">
          <h1 className="logo-title">
            <span className="logo-icon">💪</span>
            GOOD GYM
          </h1>
          <button 
            onClick={() => navigate("/MisInscripciones")}
            className="btn-inscripciones"
          >
            <span className="btn-icon">📋</span>
            Mis Inscripciones
          </button>
        </div>
      </header>

      {/* Hero Section */}
      <section className="hero-section">
        <div className="hero-content">
          <h2 className="hero-title">Transforma tu cuerpo,<br />Transforma tu vida</h2>
          <p className="hero-subtitle">Descubre las mejores actividades fitness adaptadas a tu nivel</p>
        </div>
      </section>

      {/* Buscador destacado */}
      <div className="search-section">
        <Buscador setFiltro={setFiltro} />
      </div>

      {/* Título de sección */}
      <div className="activities-header">
        <h3 className="activities-title">✨ Actividades Disponibles</h3>
        <p className="activities-subtitle">Elige tu próximo desafío fitness</p>
      </div>

      {/* Grid de actividades */}
      <ListadoActividades filtro={filtro} refrescar={refrescar} />
    </div>
  );
}

export default Home;

