import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom";
import Foto from '../components/Foto';
import Buscador from "../components/buscador";
import ListadoActividades from "../components/listadoActividades";
import { Dumbbell, PlusCircle, XCircle, Search, Calendar, Users, LogOut, CheckSquare } from "lucide-react";

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
    return {
      id: payload.jti || null,
      es_admin: payload.es_admin || false
    };
  } catch (e) {
    return null;
  }
}
  
function HomeAdm() {
  const navigate = useNavigate();
  if (getUserInfoFromToken()?.es_admin === false){
    navigate("/Home");
  }
  
  const [actividades, setActividades] = useState([]);
  const [filtro, setFiltro] = useState('');
  const [mostrarFormulario, setMostrarFormulario] = useState(false);
  const [nombre, setNombre] = useState('');
  const [descripcion, setDescripcion] = useState('');
  const [profesor, setProfesor] = useState('');
  const [horarios, setHorarios] = useState([{ dia: '', horarioInicio: '', horarioFinal: '', cupo: 0 }]);
  const [refrescar, setRefrescar] = useState(false);
  const [errorText, setErrorText] = useState('');

  const handleAgregarHorario = () => {
    setHorarios([...horarios, { dia: '', horarioInicio: '', horarioFinal: '', cupo: 0 }]);
  };

  const handleHorarioChange = (index, field, value) => {
    const nuevosHorarios = [...horarios];
    nuevosHorarios[index][field] = value;
    setHorarios(nuevosHorarios);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const nuevaActividad = {
      nombre,
      descripcion,
      profesor,
      horarios: horarios.map(h => ({
        dia: h.dia,
        horarioInicio: h.horarioInicio,
        horarioFinal: h.horarioFinal,
        cupo: parseInt(h.cupo, 10)
      }))
    };

    try {
      const response = await fetch(`${process.env.REACT_APP_API_ACTIVIDADES_URL}/actividad`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(nuevaActividad)
      });

      if (response.ok) {
        alert("Actividad creada correctamente");
        setRefrescar(prev => !prev); 
        setNombre('');
        setDescripcion('');
        setProfesor('');
        setHorarios([{ dia: '', horarioInicio: '', horarioFinal: '', cupo: 0 }]);
        setMostrarFormulario(false);
        setErrorText('');
      } else {
        setErrorText("Error al crear actividad");
      }
    } catch (error) {
      console.error("Error en la solicitud:", error);
      setErrorText("Error en la solicitud al servidor");
    }
  };

  useEffect(() => {
    setRefrescar(prev => !prev);
  }, []);

  const handleLogout = () => {
    document.cookie = "token=; path=/; expires=Thu, 01 Jan 1970 00:00:00 UTC;";
    navigate("/Login");
  };

  return (
    <div className="min-h-screen bg-background text-foreground pb-20">
      
      {/* HEADER NAVBAR */}
      <header className="sticky top-0 z-50 w-full border-b border-white/5 bg-background/80 backdrop-blur-xl supports-[backdrop-filter]:bg-background/60">
        <div className="container mx-auto max-w-7xl flex h-16 items-center justify-between px-4 sm:px-6">
          <div className="flex items-center gap-3">
            <div className="hidden sm:flex h-9 w-9 rounded-full bg-primary items-center justify-center shadow-[0_0_15px_-3px_hsl(var(--primary))]">
              <Dumbbell className="h-5 w-5 text-primary-foreground" />
            </div>
            <h1 className="font-mono text-xl sm:text-2xl font-bold uppercase tracking-tighter flex items-center gap-2 text-white">
              GOOD GYM <span className="text-primary opacity-80 text-sm hidden sm:inline">- ADMIN HUB</span>
            </h1>
          </div>
          <div className="flex items-center gap-2 sm:gap-4">
            <button 
              onClick={() => navigate("/MisInscripciones")}
              className="px-3 py-2 sm:px-4 sm:py-2 text-xs sm:text-sm font-semibold rounded-md bg-secondary/80 text-secondary-foreground hover:bg-secondary transition-colors flex items-center gap-2 border border-border"
            >
              <CheckSquare className="w-4 h-4 text-primary" />
              <span className="hidden sm:inline">Mis Inscripciones</span>
            </button>
            <button 
              onClick={handleLogout}
              className="p-2 sm:px-4 sm:py-2 text-xs sm:text-sm font-semibold rounded-md bg-destructive/10 text-destructive hover:bg-destructive/20 transition-colors flex items-center gap-2"
              title="Cerrar sesión"
            >
              <LogOut className="w-4 h-4" />
              <span className="hidden sm:inline">Salir</span>
            </button>
          </div>
        </div>
      </header>

      {/* MAIN LAYOUT */}
      <main className="container mx-auto max-w-7xl px-4 sm:px-6 mt-8 space-y-12">
        
        {/* HERO BANNER */}
        <div className="relative rounded-2xl overflow-hidden border border-white/10 shadow-2xl bg-card">
          <div className="absolute inset-0 bg-gradient-to-r from-background/90 via-background/60 to-transparent z-10" />
          <div className="absolute top-0 right-0 p-8 z-0 opacity-40 mix-blend-screen overflow-hidden max-h-full">
             <Foto />
          </div>
          <div className="relative z-20 p-8 sm:p-12 flex flex-col justify-center h-full max-w-2xl">
            <div className="inline-flex items-center rounded-full border border-primary/30 bg-primary/10 px-3 py-1 text-sm font-medium text-primary mb-6 w-max">
              <SparklesIcon className="mr-1 h-3 w-3" /> Panel de Control
            </div>
            <h2 className="text-3xl sm:text-5xl font-mono uppercase font-bold tracking-tight mb-4 text-white">
              Gestión de <span className="text-primary">Clases</span>
            </h2>
            <p className="text-muted-foreground text-lg mb-8 max-w-md">
              Administra las actividades, horarios y profesores de toda la plataforma desde un solo lugar.
            </p>
            <button 
              onClick={() => setMostrarFormulario(!mostrarFormulario)}
              className="inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 shadow-[0_0_20px_-5px_hsl(var(--primary)/0.4)] hover:shadow-[0_0_30px_-5px_hsl(var(--primary)/0.6)] bg-primary text-primary-foreground hover:bg-primary/90 h-10 px-6 py-2 w-max gap-2"
            >
              {mostrarFormulario ? (
                <><XCircle className="w-4 h-4" /> Cancelar creación</>
              ) : (
                <><PlusCircle className="w-4 h-4" /> Nueva Actividad</>
              )}
            </button>
          </div>
        </div>

        {/* FORMULARIO AGREGAR ACTIVIDAD */}
        {mostrarFormulario && (
          <div className="bg-card border border-border rounded-xl p-6 sm:p-8 shadow-xl animate-in slide-in-from-top-4 duration-500">
            <div className="mb-6 border-b border-border pb-4">
              <h3 className="text-xl font-bold uppercase tracking-tight text-white flex items-center gap-2">
                <PlusCircle className="text-primary w-5 h-5" />
                Crear Nueva Actividad
              </h3>
            </div>
            
            <form onSubmit={handleSubmit} className="space-y-6">
              <div className="grid sm:grid-cols-2 lg:grid-cols-3 gap-6">
                <div className="space-y-2">
                  <label className="text-sm font-medium text-muted-foreground">Nombre</label>
                  <input type="text" className="flex h-10 w-full rounded-md border border-input bg-background/50 px-3 py-2 text-sm focus-visible:outline-none focus-visible:border-primary/50 transition-colors" placeholder="Ej: Crossfit" value={nombre} onChange={e => setNombre(e.target.value)} required />
                </div>
                <div className="space-y-2 lg:col-span-2">
                  <label className="text-sm font-medium text-muted-foreground">Descripción</label>
                  <input type="text" className="flex h-10 w-full rounded-md border border-input bg-background/50 px-3 py-2 text-sm focus-visible:outline-none focus-visible:border-primary/50 transition-colors" placeholder="Breve descripción de la clase" value={descripcion} onChange={e => setDescripcion(e.target.value)} required />
                </div>
                <div className="space-y-2">
                  <label className="text-sm font-medium text-muted-foreground">Profesor</label>
                  <input type="text" className="flex h-10 w-full rounded-md border border-input bg-background/50 px-3 py-2 text-sm focus-visible:outline-none focus-visible:border-primary/50 transition-colors" placeholder="Nombre del coach" value={profesor} onChange={e => setProfesor(e.target.value)} required />
                </div>
              </div>

              <div className="space-y-4 pt-4 border-t border-border">
                <h4 className="text-sm font-medium text-white flex items-center gap-2">
                  <Calendar className="w-4 h-4 text-primary" />
                  Horarios Disponibles
                </h4>
                
                {horarios.map((h, index) => (
                  <div key={index} className="grid sm:grid-cols-4 gap-4 items-end bg-background/30 p-4 rounded-lg border border-border/50">
                    <div className="space-y-2">
                      <label className="text-xs text-muted-foreground">Día</label>
                      <input type="text" className="flex h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-sm focus-visible:outline-none focus-visible:border-primary/50 transition-colors" placeholder="Lunes" value={h.dia} onChange={e => handleHorarioChange(index, 'dia', e.target.value)} required />
                    </div>
                    <div className="space-y-2">
                      <label className="text-xs text-muted-foreground">Hora Inicio</label>
                      <input type="text" className="flex h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-sm focus-visible:outline-none focus-visible:border-primary/50 transition-colors" placeholder="08:00" value={h.horarioInicio} onChange={e => handleHorarioChange(index, 'horarioInicio', e.target.value)} required />
                    </div>
                    <div className="space-y-2">
                      <label className="text-xs text-muted-foreground">Hora Fin</label>
                      <input type="text" className="flex h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-sm focus-visible:outline-none focus-visible:border-primary/50 transition-colors" placeholder="09:00" value={h.horarioFinal} onChange={e => handleHorarioChange(index, 'horarioFinal', e.target.value)} required />
                    </div>
                    <div className="space-y-2">
                      <label className="text-xs text-muted-foreground">Cupo Máx</label>
                      <input type="number" className="flex h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-sm focus-visible:outline-none focus-visible:border-primary/50 transition-colors" placeholder="20" value={h.cupo} onChange={e => handleHorarioChange(index, 'cupo', e.target.value)} required />
                    </div>
                  </div>
                ))}
                
                <button type="button" onClick={handleAgregarHorario} className="text-sm text-primary hover:text-primary/80 font-medium flex items-center gap-1">
                  <PlusCircle className="w-3.5 h-3.5" /> Añadir otro horario
                </button>
              </div>

              <div className="pt-6 flex items-center gap-4">
                <button type="submit" className="inline-flex items-center justify-center rounded-md text-sm font-medium transition-colors bg-primary text-primary-foreground hover:bg-primary/90 h-10 px-8 shadow-sm">
                  Guardar Clase
                </button>
                {errorText && <span className="text-sm font-medium text-destructive">{errorText}</span>}
              </div>
            </form>
          </div>
        )}

        {/* LISTADO Y BUSCADOR */}
        <div className="space-y-6">
          <div className="flex flex-col sm:flex-row gap-4 items-start sm:items-center justify-between">
            <div>
              <h3 className="text-2xl font-mono uppercase font-bold tracking-tight text-white">Actividades Disponibles</h3>
              <p className="text-sm text-muted-foreground">Explora y gestiona el catálogo de clases.</p>
            </div>
            
            <div className="w-full sm:w-auto relative group">
              <Search className="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-muted-foreground group-focus-within:text-primary transition-colors" />
               <Buscador setFiltro={setFiltro} classNameOverride="pl-10 h-10 w-full sm:w-64 rounded-full border border-border bg-card text-sm focus:outline-none focus:border-primary/50 transition-colors" />
            </div>
          </div>

          <div className="min-h-[400px]">
            <ListadoActividades filtro={filtro} refrescar={refrescar} esAdmin={true} />
          </div>
        </div>
        
      </main>
    </div>
  );
}

function SparklesIcon(props) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M9.937 15.5A2 2 0 0 0 8.5 14.063l-6.135-1.582a.5.5 0 0 1 0-.962L8.5 9.936A2 2 0 0 0 9.937 8.5l1.582-6.135a.5.5 0 0 1 .963 0L14.063 8.5A2 2 0 0 0 15.5 9.937l6.135 1.581a.5.5 0 0 1 0 .964L15.5 14.063a2 2 0 0 0-1.437 1.437l-1.582 6.135a.5.5 0 0 1-.963 0z" />
    </svg>
  )
}

export default HomeAdm;

// import React, { useEffect, useState } from 'react';
// import { useNavigate } from "react-router-dom";
// import './Home.css';
// import Foto from '../components/Foto';
// import ListaDesplegable from '../components/ListaDesplegable';
// import Buscador from "../components/buscador";
// import ListadoActividades from "../components/listadoActividades";
// import InscripcionesUsuario from '../components/InscripcionesUsuario';



// /*function getCookiee(name) {
//   const nameEQ = `${name}=`;
//   const ca = document.cookie.split(';');
  
//   for (let i = 0; i < ca.length; i++) {
//     let c = ca[i].trim();
//     if (c.indexOf(nameEQ) === 0) {
//       return c.substring(nameEQ.length, c.length);
//     }
//   }
//   return null; // Si no se encuentra la cookie, devuelve null
// }

// function getUserIdFromTokenn() {
//   const token = getCookiee("token");
//   if (token) {
//     console.log("Token obtenido desde la cookie:", token);
//   } else {
//     console.log("No se encontró el token en las cookies");
//     return null;
//   }

//   const parts = token.split('.');
//   if (parts.length !== 3) {
//     console.error("Token JWT inválido");
//     return null;
//   }

//   try {
//     const payload = JSON.parse(atob(parts[1]));
//     console.log("Payload del token:", payload);
//     return payload.jti || null;
//   } catch (e) {
//     console.error("Error al decodificar el token:", e);
//     return null;
//   }
// }*/


// function HomeAdm() {
//   const navigate = useNavigate();
//   const [filtro, setFiltro] = useState('');
//   const [mostrarFormulario, setMostrarFormulario] = useState(false);
//   const [nombre, setNombre] = useState('');
//   const [descripcion, setDescripcion] = useState('');
//   const [profesor, setProfesor] = useState('');
//   const [horarios, setHorarios] = useState([{ dia: '', horarioInicio: '', horarioFinal: '', cupo: 0 }]);
//   const [refrescar, setRefrescar] = useState(false);
//   const [errorText, setErrorText] = useState('');

  

//   const handleAgregarHorario = () => {
//     setHorarios([...horarios, { dia: '', horarioInicio: '', horarioFinal: '', cupo: 0 }]);
//   };

//   const handleHorarioChange = (index, field, value) => {
//     const nuevosHorarios = [...horarios];
//     nuevosHorarios[index][field] = value;
//     setHorarios(nuevosHorarios);
//   };

//   const handleSubmit = async (e) => {
//     e.preventDefault();

//     const nuevaActividad = {
//       nombre,
//       descripcion,
//       profesor,
//       horarios: horarios.map(h => ({
//         dia: h.dia,
//         horarioInicio: h.horarioInicio,
//         horarioFinal: h.horarioFinal,
//         cupo: parseInt(h.cupo, 10)
//       }))
//     };

//     try {
//       const response = await fetch('${process.env.REACT_APP_API_ACTIVIDADES_URL}/actividad', {
//         method: 'POST',
//         headers: { 'Content-Type': 'application/json' },
//         body: JSON.stringify(nuevaActividad)
//       });

//       if (response.ok) {
//         alert("Actividad creada correctamente");
//         setRefrescar(prev => !prev); // Fuerza recarga
//         setNombre('');
//         setDescripcion('');
//         setProfesor('');
//         setHorarios([{ dia: '', horarioInicio: '', horarioFinal: '', cupo: 0 }]);
//         setMostrarFormulario(false);
//         setErrorText('');
//       } else {
//         setErrorText("Error al crear actividad");
//       }
//     } catch (error) {
//       console.error("Error en la solicitud:", error);
//       setErrorText("Error en la solicitud al servidor");
//     }
//   };

//   useEffect(() => {
//     setRefrescar(prev => !prev); // Inicial para cargar desde el inicio
//   }, []);

//   /*useEffect(() => {
//     console.log("ID de usuario desde el token:", usuarioId);
//   }, [usuarioId]);*/


//   return (
//     <div className="home">
//       <hr />
//       <h1 className="Titulo">G O O D   G Y M - ADMIN</h1>

//       <div className="foto">
//         <Foto />

//         <div style={{ textAlign: 'right', margin: '10px' }}>
//           <button onClick={() => setMostrarFormulario(!mostrarFormulario)}>
//             {mostrarFormulario ? 'Cerrar Formulario' : 'Agregar Actividad'}
//           </button>
//         </div>

//         {mostrarFormulario && (
//           <form onSubmit={handleSubmit} className="formulario-actividad">
//             <h3>Formulario de Nueva Actividad</h3>
//             <input type="text" placeholder="Nombre" value={nombre} onChange={e => setNombre(e.target.value)} required />
//             <input type="text" placeholder="Descripción" value={descripcion} onChange={e => setDescripcion(e.target.value)} required />
//             <input type="text" placeholder="Profesor" value={profesor} onChange={e => setProfesor(e.target.value)} required />

//             {horarios.map((h, index) => (
//               <div key={index} style={{ marginBottom: '10px' }}>
//                 <input type="text" placeholder="Día" value={h.dia} onChange={e => handleHorarioChange(index, 'dia', e.target.value)} required />
//                 <input type="text" placeholder="Hora Inicio" value={h.horarioInicio} onChange={e => handleHorarioChange(index, 'horarioInicio', e.target.value)} required />
//                 <input type="text" placeholder="Hora Fin" value={h.horarioFinal} onChange={e => handleHorarioChange(index, 'horarioFinal', e.target.value)} required />
//                 <input type="number" placeholder="Cupo" value={h.cupo} onChange={e => handleHorarioChange(index, 'cupo', e.target.value)} required />
//               </div>
//             ))}
//             <button type="button" onClick={handleAgregarHorario}>Agregar otro horario</button>
//             <br />
//             <button type="submit">Guardar Actividad</button>
//             {errorText && <p className="error" style={{ color: 'red' }}>{errorText}</p>}
//           </form>
//         )}

//         <p className="espacio" />
//         <h1 className='subtitulo'>ACTIVIDADES DISPONIBLES</h1>
//         <p className="espacio" />

//         <Buscador setFiltro={setFiltro} />
//         <p className="espacio" />

//         <ListadoActividades filtro={filtro} refrescar={refrescar} esAdmin={true} />

//         <p className="espacio" />

        

//       </div>
//     </div>
//   );
//   }


// export default HomeAdm;


