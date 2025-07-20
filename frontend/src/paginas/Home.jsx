import { useNavigate } from "react-router-dom";
import './Home.css';
import Foto from '../components/Foto';
import ListaDesplegable from '../components/ListaDesplegable';
import Buscador from "../components/buscador";
import ListadoActividades from "../components/listadoActividades";
import { useState } from "react";

function Home() {
  const navigate = useNavigate();
  const [filtro, setFiltro] = useState('');

  return (
    <div className="home">
      

      <hr/>
      <h1 className="Titulo">    G O O D   G Y M    </h1>
      <div className="foto">
        <Foto/>

        <div style={{ textAlign: 'right', margin: '10px' }}> 
          <button onClick={() => navigate('/admin')}>
            Panel Administrador
          </button>
        </div>

       

        <p className="espacio"/>
        <h1 className='subtitulo'>ACTIVIDADES DISPONIBLES</h1>
        <p className="espacio"/>

        <Buscador setFiltro={setFiltro} />

        <p className="espacio"/>

        <ListadoActividades filtro={filtro} />
      </div>
    </div>
  );
}

export default Home;

