import { Button } from '@/components/ui/button'
import { Play, Pause } from 'lucide-react'

export function Hero() {
  return (
    <section className="relative h-[70vh] overflow-hidden">
      <div 
        className="absolute inset-0 bg-cover bg-center"
        style={{
          backgroundImage: `url('/athletic-person-working-out-in-modern-gym.jpg')`,
        }}
      >
        <div className="absolute inset-0 bg-gradient-to-t from-background via-background/60 to-transparent" />
      </div>
      <div className="relative container h-full flex flex-col justify-end px-4 md:px-6 pb-16 md:pb-24">
        <div className="max-w-3xl">
          <h1 className="font-mono text-5xl md:text-7xl font-bold uppercase tracking-tighter text-balance mb-6">
            Transform Your<br />Fitness Journey
          </h1>
          <div className="flex flex-wrap gap-3 mb-8">
            <Button size="lg" variant="secondary">
              View Schedule
            </Button>
            <Button size="lg" variant="outline" className="gap-2">
              <Play className="h-4 w-4" />
              Watch Tour
            </Button>
          </div>
        </div>
        <div className="flex gap-3">
          <Button size="icon" variant="ghost" className="rounded-full">
            <Pause className="h-5 w-5" />
          </Button>
        </div>
      </div>
    </section>
  )
}
