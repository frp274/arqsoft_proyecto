import { Header } from '@/components/header'
import { DetalleActividad } from '@/components/detalle-actividad'

export default function DetallePage({ params }: { params: { id: string } }) {
  return (
    <main className="min-h-screen bg-background">
      <Header />
      <DetalleActividad activityId={params.id} />
    </main>
  )
}
