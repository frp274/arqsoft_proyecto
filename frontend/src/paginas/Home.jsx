import { useNavigate } from "react-router-dom";
import './Home.css';
import Foto from '../components/Foto';
import ListaDesplegable from '../components/ListaDesplegable';
import Buscador from "../components/buscador";
import ListadoActividades from "../components/listadoActividades";


function Home() {
  const navigate = useNavigate();

  return (
    <div className="home">
      <top className="top">

        <button onClick={() => navigate("/")} className="botonRedondoVolver">← Volver a Login</button>
        <button onClick={() => navigate("/Detalle")} className="botonRedondoAdelante">Ir a Detalle →</button>
        <p className="espacio"/>
        <h2>Home</h2>

      </top>
      <hr/>
      <h1 className="Titulo">    G O O D  G Y M    </h1>
      <div className="foto">

      <Foto/>

      <p className="espacio"/>
      <h1 className='subtitulo'>ACTIVIDADES DISPONIBLES</h1>
      <p className="espacio"/>

      <Buscador/>

      <p className="espacio"/>

      <ListadoActividades/>


     </div>
     
     
    </div>
  );
}

export default Home;


