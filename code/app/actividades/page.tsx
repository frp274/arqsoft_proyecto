import { Buscador } from '@/components/buscador'
import { ListadoActividades } from '@/components/listado-actividades'
import { Header } from '@/components/header'

export default function ActividadesPage() {
  return (
    <main className="min-h-screen bg-background">
      <Header />
      <div className="container px-4 md:px-6 py-12">
        <div className="mb-8">
          <h1 className="font-mono text-4xl md:text-5xl font-bold uppercase tracking-tighter mb-4">
            All Activities
          </h1>
          <p className="text-muted-foreground text-lg">
            Browse and book all available classes and training sessions
          </p>
        </div>
        <Buscador />
        <ListadoActividades />
      </div>
    </main>
  )
}
