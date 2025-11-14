import { useNavigate } from "react-router-dom";
import ListadoActividades from "../components/listadoActividades";
import { useState } from "react";
import { Button } from "../components/ui/button";
import { Input } from "../components/ui/input";
import { Search, Dumbbell, Calendar } from "lucide-react";

function Home() {
  const navigate = useNavigate();
  const [filtro, setFiltro] = useState('');
  const [busqueda, setBusqueda] = useState('');
  const [refrescar, setRefrescar] = useState(false);

  const handleSearch = () => {
    setFiltro(busqueda);
  };

  const handleKeyPress = (e) => {
    if (e.key === 'Enter') {
      handleSearch();
    }
  };

  const categorias = ['Yoga', 'Pilates', 'Crossfit', 'Funcional', 'Spinning', 'Box', 'Natación', 'Zumba'];

  return (
    <div className="min-h-screen bg-background">
      {/* Header */}
      <header className="sticky top-0 z-50 w-full border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
        <div className="container flex h-16 items-center justify-between px-4 md:px-6">
          <div className="flex items-center gap-2">
            <Dumbbell className="h-6 w-6" />
            <h1 className="font-mono text-xl md:text-2xl font-bold uppercase tracking-tighter">
              GOOD GYM
            </h1>
          </div>
          <Button 
            onClick={() => navigate("/MisInscripciones")}
            variant="outline"
            className="gap-2"
          >
            <Calendar className="h-4 w-4" />
            <span className="hidden sm:inline">Mis Inscripciones</span>
          </Button>
        </div>
      </header>

      {/* Hero Section */}
      <section className="relative h-[60vh] overflow-hidden bg-gradient-to-br from-primary/10 via-background to-secondary/10">
        <div className="absolute inset-0 bg-grid-white/5" />
        <div className="relative container h-full flex flex-col justify-center px-4 md:px-6">
          <div className="max-w-3xl">
            <h1 className="font-mono text-4xl md:text-6xl lg:text-7xl font-bold uppercase tracking-tighter text-balance mb-6">
              Transform Your<br />Fitness Journey
            </h1>
            <p className="text-lg md:text-xl text-muted-foreground mb-8 max-w-xl">
              Descubre las mejores actividades fitness adaptadas a tu nivel y objetivos
            </p>
            <div className="flex flex-wrap gap-3">
              <Button 
                size="lg" 
                onClick={() => document.getElementById('search-section')?.scrollIntoView({ behavior: 'smooth' })}
              >
                Explorar Actividades
              </Button>
              <Button 
                size="lg" 
                variant="outline"
                onClick={() => navigate("/MisInscripciones")}
              >
                Ver Mi Agenda
              </Button>
            </div>
          </div>
        </div>
      </section>

      {/* Search Section */}
      <section id="search-section" className="py-16 md:py-24 border-b border-border">
        <div className="container px-4 md:px-6">
          <div className="max-w-4xl mx-auto">
            <h2 className="font-mono text-3xl md:text-4xl font-bold uppercase tracking-tighter mb-8">
              Encontrá Tu Actividad
            </h2>
            <div className="flex gap-2 mb-12">
              <div className="relative flex-1">
                <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-5 w-5 text-muted-foreground" />
                <Input 
                  placeholder="Buscar por nombre, instructor o tipo de actividad..."
                  className="pl-10 h-12 bg-card"
                  value={busqueda}
                  onChange={(e) => setBusqueda(e.target.value)}
                  onKeyPress={handleKeyPress}
                />
              </div>
              <Button size="lg" className="h-12" onClick={handleSearch}>
                Buscar
              </Button>
            </div>
            <div className="grid grid-cols-2 md:grid-cols-4 gap-3">
              {categorias.map((categoria) => (
                <Button 
                  key={categoria}
                  variant="outline" 
                  className="h-14 text-sm font-medium hover:bg-primary hover:text-primary-foreground hover:border-primary transition-all"
                  onClick={() => {
                    setBusqueda(categoria);
                    setFiltro(categoria);
                  }}
                >
                  {categoria}
                </Button>
              ))}
            </div>
          </div>
        </div>
      </section>

      {/* Activities Section */}
      <section className="py-16 md:py-24">
        <div className="container px-4 md:px-6">
          <div className="mb-12">
            <h2 className="font-mono text-3xl md:text-4xl font-bold uppercase tracking-tighter mb-4">
              Actividades Disponibles
            </h2>
            <p className="text-lg text-muted-foreground">
              Descubrí todas las clases y comenzá tu transformación hoy
            </p>
          </div>
          <ListadoActividades filtro={filtro} refrescar={refrescar} />
        </div>
      </section>
    </div>
  );
}

export default Home;

