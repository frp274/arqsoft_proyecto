import { useNavigate } from "react-router-dom";

function Home() {
  const navigate = useNavigate();

  return (
    <div>
      <h2>Home</h2>
      <button onClick={() => navigate("/")}>← Volver a Login</button>
      <button onClick={() => navigate("/Detalle")}>Ir a Detalle →</button>
    </div>
  );
}

export default Home;
