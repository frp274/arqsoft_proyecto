import { Header } from '@/components/header'
import { InscripcionesUsuario } from '@/components/inscripciones-usuario'

export default function MisInscripcionesPage() {
  return (
    <main className="min-h-screen bg-background">
      <Header />
      <div className="container px-4 md:px-6 py-12">
        <h1 className="font-mono text-4xl md:text-5xl font-bold uppercase tracking-tighter mb-8">
          My Bookings
        </h1>
        <InscripcionesUsuario />
      </div>
    </main>
  )
}
