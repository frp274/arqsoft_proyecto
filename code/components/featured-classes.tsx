import { Card } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Calendar, Clock, Users, ArrowRight } from 'lucide-react'

const classes = [
  {
    id: 1,
    name: 'POWER HOUR',
    type: 'HIIT Training',
    date: 'MON-FRI',
    time: '6:00 AM',
    spots: 12,
    instructor: 'Sarah Johnson',
    image: '/intense-gym-class-hiit-training.jpg'
  },
  {
    id: 2,
    name: 'YOGA FLOW',
    type: 'Vinyasa Yoga',
    date: 'TUE-THU',
    time: '7:30 PM',
    spots: 8,
    instructor: 'Michael Chen',
    image: '/yoga-class-peaceful-studio.jpg'
  },
  {
    id: 3,
    name: 'IRON TEMPLE',
    type: 'Strength Training',
    date: 'MON-WED-FRI',
    time: '5:00 PM',
    spots: 15,
    instructor: 'David Martinez',
    image: '/strength-training-weightlifting-gym.jpg'
  }
]

export function FeaturedClasses() {
  return (
    <section className="py-16 md:py-24 border-b border-border">
      <div className="container px-4 md:px-6">
        <div className="flex items-end justify-between mb-12">
          <div>
            <span className="text-sm font-medium text-primary uppercase tracking-wider mb-2 block">
              Popular This Week
            </span>
            <h2 className="font-mono text-3xl md:text-5xl font-bold uppercase tracking-tighter">
              Featured Classes
            </h2>
          </div>
          <Button variant="ghost" className="hidden md:flex items-center gap-2">
            View All Classes
            <ArrowRight className="h-4 w-4" />
          </Button>
        </div>
        <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
          {classes.map((classItem) => (
            <Card key={classItem.id} className="overflow-hidden group cursor-pointer border-2 hover:border-primary transition-all">
              <div className="relative h-64 overflow-hidden">
                <img 
                  src={classItem.image || "/placeholder.svg"}
                  alt={classItem.name}
                  className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
                />
                <div className="absolute inset-0 bg-gradient-to-t from-card via-card/40 to-transparent" />
                <div className="absolute top-4 left-4">
                  <span className="inline-block px-3 py-1 bg-primary text-primary-foreground text-xs font-bold uppercase tracking-wide rounded-md">
                    {classItem.type}
                  </span>
                </div>
              </div>
              <div className="p-6">
                <h3 className="font-mono text-2xl font-bold uppercase tracking-tight mb-4">
                  {classItem.name}
                </h3>
                <div className="space-y-2 mb-6">
                  <div className="flex items-center gap-2 text-sm text-muted-foreground">
                    <Calendar className="h-4 w-4" />
                    <span>{classItem.date}</span>
                    <span className="text-border">•</span>
                    <Clock className="h-4 w-4" />
                    <span>{classItem.time}</span>
                  </div>
                  <div className="flex items-center gap-2 text-sm text-muted-foreground">
                    <Users className="h-4 w-4" />
                    <span>{classItem.spots} spots available</span>
                  </div>
                  <p className="text-sm font-medium">with {classItem.instructor}</p>
                </div>
                <div className="flex gap-2">
                  <Button className="flex-1" size="lg">
                    Book Class
                  </Button>
                  <Button variant="outline" size="lg">
                    Details
                  </Button>
                </div>
              </div>
            </Card>
          ))}
        </div>
      </div>
    </section>
  )
}
