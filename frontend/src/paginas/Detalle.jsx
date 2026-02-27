
import { useNavigate, useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import axios from "axios";
import { ArrowLeft, Clock, User, Users, CheckCircle, Info, Dumbbell } from "lucide-react";
import { Button } from "../components/ui/button";
import { Badge } from "../components/ui/badge";

function getUserIdFromToken() {
  const token = getCookie("token");
  if (!token) return null;

  const parts = token.split('.');
  if (parts.length !== 3) return null;

  try {
    const payload = JSON.parse(atob(parts[1]));
    return payload.jti || null;
  } catch (e) {
    return null;
  }
}

function getIsAdminFromToken() {
  const token = getCookie("token");
  if (!token) return false;
  const parts = token.split('.');
  if (parts.length !== 3) return false;

  try {
    const payload = JSON.parse(atob(parts[1]));
    return payload.es_admin || false;
  } catch (e) {
    return false;
  }
}

function getCookie(name) {
  const nameEQ = `${name}=`;
  const ca = document.cookie.split(';');
  for (let i = 0; i < ca.length; i++) {
    let c = ca[i].trim();
    if (c.indexOf(nameEQ) === 0) return c.substring(nameEQ.length, c.length);
  }
  return null;
}

function Detalle() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [actividad, setActividad] = useState(null);
  const [mensaje, setMensaje] = useState('');
  const [errorStatus, setErrorStatus] = useState(false);

  useEffect(() => {
    axios.get(`${process.env.REACT_APP_API_BUSQUEDAS_URL}/actividad/${id}`)
      .then((res) => setActividad(res.data))
      .catch((err) => console.error(err));
  }, [id]);

  const inscribirseHorario = (horarioId) => {
    const usuarioId = getUserIdFromToken();
    if (!usuarioId) {
      setErrorStatus(true);
      setMensaje("Debe iniciar sesión para inscribirse.");
      setTimeout(() => navigate("/Login"), 2000);
      return;
    }

    const horarioSeleccionado = horarios.find(h => `${h.dia || h.Dia}-${h.horarioInicio || h.horarioinicio || h.HorarioInicio}` === horarioId);
    console.log("Horario Seleccionado:", horarioSeleccionado); // <--- Agrega esto
    if (!horarioSeleccionado || (horarioSeleccionado.cupo || horarioSeleccionado.Cupo) <= 0) {
      setErrorStatus(true);
      setMensaje("No hay cupos disponibles para este horario.");
      return;
    }

    const actividadIdStr = id.toString();
    const horarioIdStr = horarioId.toString();
    const usuarioidint = parseInt(usuarioId, 10);

    axios.post(`${process.env.REACT_APP_API_USUARIOS_URL}/inscripcion`, {
      usuario_id: usuarioidint,
      actividad_id: actividadIdStr,
      horario_id: horarioIdStr
    }, {
      headers: {
        Authorization: `Bearer ${getCookie("token")}`,
        "Content-Type": "application/json"
      }
    })
      .then((res) => {
        setErrorStatus(false);
        setMensaje('¡Felicitaciones! Te inscribiste correctamente a la actividad.');
        setTimeout(() => {
          axios.get(`${process.env.REACT_APP_API_ACTIVIDADES_URL}/actividad/${id}`)
            .then((response) => setActividad(response.data))
            .catch((err) => console.error('Error al recargar:', err));
        }, 1500);
      })
      .catch((err) => {
        setErrorStatus(true);
        const backendError = err.response?.data?.error;

        if (backendError === "Ya estás inscrito en este horario") {
          setMensaje("Ya te encuentras inscripto a esta actividad");
        } else {
          setMensaje(backendError || err.response?.data?.message || 'Error al inscribirse. Intenta nuevamente.');
        }
      });
  };

  if (!actividad) return (
    <div className="min-h-screen bg-background flex items-center justify-center">
      <div className="flex flex-col items-center gap-4">
        <Dumbbell className="w-12 h-12 text-primary animate-spin" />
        <h1 className="text-xl font-mono uppercase font-bold text-white tracking-widest">Cargando actividad...</h1>
      </div>
    </div>
  );

  const horarios = actividad.horarios || actividad.Horarios || [];
  const nombreAct = actividad.nombre || actividad.Nombre;
  const descripcionAct = actividad.descripcion || actividad.Descripcion;
  const profesorAct = actividad.profesor || actividad.Profesor;

  const obtenerImagenActividad = (nombre) => {
    if (!nombre) return "/logo192.png";
    const nombreNormalizado = nombre.toLowerCase();
    switch (nombreNormalizado) {
      case "pilates": return "/pilates.png";
      case "mma": return "/mma.png";
      case "musculacion": return "/musculacion.png";
      case "zumba": return "/zumba.png";
      case "spinning": return "/spining.png";
      default: return "/logo192.png";
    }
  };

  return (
    <div className="min-h-screen bg-background text-foreground pb-20">
      <header className="sticky top-0 z-50 w-full border-b border-border/50 bg-background/80 backdrop-blur-xl supports-[backdrop-filter]:bg-background/60">
        <div className="container mx-auto max-w-5xl flex h-16 items-center px-4 sm:px-6">
          <Button
            variant="ghost"
            className="gap-2 text-muted-foreground hover:text-white hover:bg-white/5 transition-colors"
            onClick={() => {
              if (getIsAdminFromToken()) navigate("/Admin");
              else navigate("/Home");
            }}
          >
            <ArrowLeft className="w-4 h-4" /> Volver
          </Button>
        </div>
      </header>

      <main className="container mx-auto max-w-5xl px-4 sm:px-6 mt-8 sm:mt-12">
        <div className="grid lg:grid-cols-2 gap-8 lg:gap-12">

          <div className="space-y-6">
            <div className="relative rounded-2xl overflow-hidden border border-border bg-card/50 aspect-video lg:aspect-square flex items-center justify-center shadow-[0_0_50px_-12px_hsl(var(--primary)/0.15)] glow-container">
              <div className="absolute inset-0 bg-gradient-to-tr from-background/80 via-transparent to-transparent z-10" />
              <img
                src={obtenerImagenActividad(nombreAct)}
                alt={nombreAct}
                className="w-full h-full object-cover relative z-0"
              />
            </div>
          </div>

          <div className="flex flex-col">
            <div className="mb-6">
              <Badge variant="outline" className="mb-4 bg-primary/10 text-primary border-primary/30 uppercase tracking-widest">
                Detalle de Clase
              </Badge>
              <h1 className="text-4xl sm:text-5xl font-mono uppercase font-bold tracking-tighter text-white mb-4">
                {nombreAct}
              </h1>
              <div className="flex items-center gap-3 text-muted-foreground bg-secondary/50 w-max px-4 py-2 rounded-lg border border-border shadow-sm">
                <User className="h-5 w-5 text-primary" />
                <span className="font-medium text-white">{profesorAct}</span>
              </div>
            </div>

            <div className="mb-8 flex-1">
              <h3 className="text-lg font-semibold text-white mb-2 flex items-center gap-2">
                <Info className="w-5 h-5 text-muted-foreground" />
                Sobre esta actividad
              </h3>
              <p className="text-muted-foreground leading-relaxed">
                {descripcionAct}
              </p>
            </div>

            {mensaje && (
              <div className={`p-4 rounded-lg mb-8 flex border gap-3 font-medium shadow-sm items-start text-sm ${errorStatus ? 'bg-destructive/10 border-destructive/30 text-destructive' : 'bg-primary/10 border-primary/30 text-primary'}`}>
                {errorStatus ? <Info className="w-5 h-5 mt-0.5 shrink-0" /> : <CheckCircle className="w-5 h-5 mt-0.5 shrink-0" />}
                <p>{mensaje}</p>
              </div>
            )}

            <div className="bg-card/50 border border-border rounded-xl p-6 shadow-md backdrop-blur-sm">
              <h3 className="text-xl font-bold uppercase tracking-tight text-white mb-4 flex items-center gap-2 border-b border-border/50 pb-4">
                <Clock className="w-5 h-5 text-primary" />
                Horarios Disponibles
              </h3>

              <div className="space-y-3">
                {horarios.length > 0 ? horarios.map((h, idx) => {
                  const dia = h.dia || h.Dia;
                  const inicio = h.horarioInicio || h.horarioinicio || h.HorarioInicio;
                  const fin = h.horarioFinal || h.horariofinal || h.HorarioFinal;
                  const cupos = h.cupo || h.Cupo;
                  const hId = `${dia}-${inicio}`;
                  const sinCupo = cupos <= 0;

                  return (
                    <div key={idx} className="flex flex-col sm:flex-row sm:items-center justify-between gap-4 p-4 rounded-lg bg-background/50 border border-border/50 hover:bg-background transition-colors">
                      <div className="flex flex-col sm:flex-row sm:items-center gap-2 sm:gap-6">
                        <div className="font-medium text-white min-w-[100px]">{dia}</div>
                        <div className="text-muted-foreground font-mono text-sm">{inicio} - {fin}</div>
                        <Badge variant="secondary" className={`w-max ${sinCupo ? 'bg-destructive/20 text-destructive border-destructive/20' : 'bg-secondary text-secondary-foreground'}`}>
                          <Users className="w-3 h-3 mr-1" />
                          {cupos} cupos
                        </Badge>
                      </div>

                      <Button
                        onClick={() => inscribirseHorario(hId)}
                        disabled={sinCupo}
                        className={`font-semibold shadow-[0_0_15px_-3px_hsl(var(--primary)/0.3)] transition-all ${sinCupo ? 'opacity-50 cursor-not-allowed hidden' : 'hover:shadow-[0_0_20px_-3px_hsl(var(--primary)/0.5)]'}`}
                      >
                        Inscribirse
                      </Button>
                    </div>
                  );
                }) : (
                  <p className="text-muted-foreground text-center py-4">No hay horarios cargados para esta actividad.</p>
                )}
              </div>
            </div>

          </div>
        </div>
      </main>
    </div>
  );
}

export default Detalle;

