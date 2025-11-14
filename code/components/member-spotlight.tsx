import { Card } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { ChevronLeft, ChevronRight } from 'lucide-react'

const members = [
  {
    id: 1,
    name: 'Emma Rodriguez',
    achievement: 'Marathon Finisher',
    joined: '2023',
    image: '/female-athlete-confident-portrait.jpg'
  },
  {
    id: 2,
    name: 'James Mitchell',
    achievement: 'Powerlifting Champion',
    joined: '2022',
    image: '/male-athlete-strong-portrait.jpg'
  },
  {
    id: 3,
    name: 'Sofia Park',
    achievement: 'Yoga Instructor',
    joined: '2024',
    image: '/yoga-instructor-peaceful-portrait.jpg'
  },
  {
    id: 4,
    name: 'Marcus Thompson',
    achievement: 'Boxing Pro',
    joined: '2021',
    image: '/boxer-determined-portrait.jpg'
  }
]

export function MemberSpotlight() {
  return (
    <section className="py-16 md:py-24">
      <div className="container px-4 md:px-6">
        <div className="flex items-end justify-between mb-12">
          <div>
            <span className="text-sm font-medium text-primary uppercase tracking-wider mb-2 block">
              Community
            </span>
            <h2 className="font-mono text-3xl md:text-5xl font-bold uppercase tracking-tighter">
              Meet Our Members
            </h2>
          </div>
          <div className="hidden md:flex gap-2">
            <Button size="icon" variant="outline">
              <ChevronLeft className="h-5 w-5" />
            </Button>
            <Button size="icon" variant="outline">
              <ChevronRight className="h-5 w-5" />
            </Button>
          </div>
        </div>
        <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
          {members.map((member) => (
            <Card key={member.id} className="overflow-hidden group cursor-pointer border-2 hover:border-primary transition-all">
              <div className="relative aspect-[3/4] overflow-hidden">
                <img 
                  src={member.image || "/placeholder.svg"}
                  alt={member.name}
                  className="w-full h-full object-cover grayscale group-hover:grayscale-0 group-hover:scale-105 transition-all duration-500"
                />
                <div className="absolute inset-0 bg-gradient-to-t from-card via-card/20 to-transparent" />
                <div className="absolute bottom-0 left-0 right-0 p-4">
                  <span className="inline-block px-2 py-1 bg-primary text-primary-foreground text-xs font-bold uppercase tracking-wide rounded mb-2">
                    {member.achievement}
                  </span>
                  <h3 className="font-mono text-lg font-bold uppercase tracking-tight">
                    {member.name}
                  </h3>
                  <p className="text-xs text-muted-foreground">Member since {member.joined}</p>
                </div>
              </div>
            </Card>
          ))}
        </div>
        <div className="mt-12 text-center">
          <Button size="lg" variant="outline">
            View Full Community
          </Button>
        </div>
      </div>
    </section>
  )
}
