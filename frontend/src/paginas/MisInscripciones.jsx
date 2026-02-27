import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom";
import { ArrowLeft, CalendarDays, Clock, User, Dumbbell } from 'lucide-react';
import { Button } from "../components/ui/button";

function getCookie(name) {
  const nameEQ = `${name}=`;
  const ca = document.cookie.split(';');

  for (let i = 0; i < ca.length; i++) {
    let c = ca[i].trim();
    if (c.indexOf(nameEQ) === 0) {
      return c.substring(nameEQ.length, c.length);
    }
  }
  return null;
}

function getUserInfoFromToken() {
  const token = getCookie("token");
  if (!token) {
    return null;
  }

  const parts = token.split('.');
  if (parts.length !== 3) {
    return null;
  }

  try {
    const payload = JSON.parse(atob(parts[1]));
    let id = payload.jti || null;
    if (id && typeof id === 'string' && id.includes(':')) {
      id = id.split(':')[0];
    }
    return {
      id: id,
      es_admin: payload.es_admin || false
    };
  } catch (e) {
    return null;
  }
}

function MisInscripciones() {
  const navigate = useNavigate();
  const [inscripciones, setInscripciones] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const usuarioInfo = getUserInfoFromToken();
  const usuarioId = usuarioInfo?.id;
  const esAdmin = usuarioInfo?.es_admin;

  useEffect(() => {
    if (!usuarioId) {
      setError("Debes iniciar sesión para ver tus inscripciones.");
      setLoading(false);
      setTimeout(() => navigate("/Login"), 2000);
      return;
    }

    const fetchInscripciones = async () => {
      try {
        const response = await fetch(
          `${process.env.REACT_APP_API_BUSQUEDAS_URL}/inscripciones/usuario/${usuarioId}`,
          {
            headers: {
              Authorization: `Bearer ${getCookie("token")}`
            }
          }
        );

        if (response.ok) {
          const data = await response.json();
          setInscripciones(Array.isArray(data) ? data : []);
        } else if (response.status === 404) {
          setInscripciones([]);
        } else {
          setError("Error al cargar las inscripciones.");
        }
      } catch (err) {
        setError("Error de conexión al servidor.");
      } finally {
        setLoading(false);
      }
    };

    fetchInscripciones();
  }, [usuarioId, navigate]);

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

  const handleVerDetalle = (actividadId) => navigate(`/Detalle/${actividadId}`);

  const volverAlHome = () => {
    if (esAdmin) navigate("/Admin");
    else navigate("/Home");
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-background flex items-center justify-center">
        <div className="flex flex-col items-center gap-4">
          <Dumbbell className="w-12 h-12 text-primary animate-spin" />
          <h1 className="text-xl font-mono uppercase font-bold text-white tracking-widest">Cargando inscripciones...</h1>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-background text-foreground pb-20">

      {/* Header */}
      <header className="sticky top-0 z-50 w-full border-b border-border/50 bg-background/80 backdrop-blur-xl supports-[backdrop-filter]:bg-background/60">
        <div className="container mx-auto max-w-7xl flex h-16 items-center px-4 sm:px-6">
          <Button
            variant="ghost"
            className="gap-2 text-muted-foreground hover:text-white hover:bg-white/5 transition-colors"
            onClick={volverAlHome}
          >
            <ArrowLeft className="w-4 h-4" />
            Volver al Inicio
          </Button>
        </div>
      </header>

      <main className="container mx-auto max-w-7xl px-4 sm:px-6 mt-12">
        <div className="mb-12 border-b border-border pb-6 flex items-center gap-4">
          <div className="h-12 w-12 rounded-full bg-primary/20 flex items-center justify-center border border-primary/50 shadow-[0_0_15px_-3px_hsl(var(--primary)/0.5)]">
            <CalendarDays className="h-6 w-6 text-primary" />
          </div>
          <div>
            <h1 className="text-3xl sm:text-4xl font-mono uppercase font-bold tracking-tight text-white mb-2">Mis Inscripciones</h1>
            <p className="text-muted-foreground text-sm sm:text-base">Consulta tus clases activas y administra tu agenda de entrenamiento.</p>
          </div>
        </div>

        {error && (
          <div className="bg-destructive/10 border border-destructive/50 text-destructive px-4 py-3 rounded-lg mb-8 flex items-center gap-2">
            <span className="w-2 h-2 rounded-full bg-destructive animate-pulse"></span>
            {error}
          </div>
        )}

        {!error && inscripciones.length === 0 ? (
          <div className="flex flex-col items-center justify-center py-20 text-center bg-card/30 border border-border/50 rounded-2xl">
            <CalendarDays className="w-16 h-16 text-muted-foreground/30 mb-6" />
            <p className="text-xl font-medium text-white mb-2">No tienes inscripciones activas</p>
            <p className="text-muted-foreground mb-8 max-w-md mx-auto">
              ¡Explora nuestras actividades y encuentra tu clase perfecta para empezar a entrenar hoy mismo!
            </p>
            <Button size="lg" onClick={volverAlHome} className="h-12 px-8 font-bold tracking-wide shadow-[0_0_20px_-5px_hsl(var(--primary)/0.4)]">
              Ver Actividades Disponibles
            </Button>
          </div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {inscripciones.map((actividad) => {
              const actID = actividad.Id || actividad.id;
              const nombre = actividad.Nombre || actividad.nombre;
              const descripcion = actividad.Descripcion || actividad.descripcion;
              const profesor = actividad.Profesor || actividad.profesor;
              const horarios = actividad.Horarios || actividad.horarios;

              return (
                <div key={actID} className="group relative bg-card/60 backdrop-blur-sm border border-border rounded-xl overflow-hidden hover:border-primary/50 hover:shadow-[0_0_30px_-5px_hsl(var(--primary)/0.15)] transition-all duration-300 flex flex-col">

                  <div className="relative h-48 overflow-hidden">
                    <div className="absolute inset-0 bg-gradient-to-t from-background to-transparent z-10" />
                    <img
                      src={obtenerImagenActividad(nombre)}
                      alt={nombre}
                      className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
                    />
                    <div className="absolute bottom-4 left-4 z-20">
                      <h3 className="text-2xl font-mono uppercase font-bold text-white tracking-tight">{nombre}</h3>
                    </div>
                  </div>

                  <div className="p-6 flex-1 flex flex-col">
                    <p className="text-sm text-muted-foreground line-clamp-2 mb-4 flex-1">
                      {descripcion}
                    </p>

                    <div className="space-y-3 mb-6">
                      <div className="flex items-center gap-2 text-sm text-muted-foreground bg-secondary/30 px-3 py-2 rounded-md">
                        <User className="w-4 h-4 text-primary" />
                        <span className="font-medium text-white truncate">{profesor}</span>
                      </div>

                      {horarios && horarios.length > 0 && (
                        <div className="space-y-1">
                          <p className="text-xs font-semibold text-muted-foreground mb-2 flex items-center gap-2">
                            <Clock className="w-3 h-3" /> Horarios de tu clase
                          </p>
                          {horarios.map((h, idx) => (
                            <div key={idx} className="flex justify-between items-center text-sm border-l-2 border-primary pl-3 py-1">
                              <span className="font-medium text-white">{h.Dia || h.dia}</span>
                              <span className="text-muted-foreground">{h.HorarioInicio || h.horarioInicio} - {h.HorarioFinal || h.horarioFinal}</span>
                            </div>
                          ))}
                        </div>
                      )}
                    </div>

                    <Button
                      variant="outline"
                      className="w-full border-primary/20 hover:bg-primary hover:text-primary-foreground transition-colors group-hover:border-primary/50"
                      onClick={() => handleVerDetalle(actID)}
                    >
                      Ver todos los detalles
                    </Button>
                  </div>
                </div>
              );
            })}
          </div>
        )}
      </main>
    </div>
  );
}

export default MisInscripciones;
