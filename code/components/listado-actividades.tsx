'use client'

import { Card } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Calendar, Clock, Users, MapPin } from 'lucide-react'
import Link from 'next/link'

// TODO: Replace with data from your activities API
const activities = [
  {
    id: '1',
    name: 'POWER HOUR',
    type: 'HIIT Training',
    schedule: 'MON-FRI',
    time: '6:00 AM - 7:00 AM',
    duration: '60 min',
    spots: 12,
    totalSpots: 20,
    instructor: 'Sarah Johnson',
    location: 'Studio A',
    image: '/intense-gym-class-hiit-training.jpg'
  },
  {
    id: '2',
    name: 'YOGA FLOW',
    type: 'Vinyasa Yoga',
    schedule: 'TUE-THU',
    time: '7:30 PM - 8:30 PM',
    duration: '60 min',
    spots: 8,
    totalSpots: 15,
    instructor: 'Michael Chen',
    location: 'Zen Room',
    image: '/yoga-class-peaceful-studio.jpg'
  },
  {
    id: '3',
    name: 'IRON TEMPLE',
    type: 'Strength Training',
    schedule: 'MON-WED-FRI',
    time: '5:00 PM - 6:30 PM',
    duration: '90 min',
    spots: 15,
    totalSpots: 15,
    instructor: 'David Martinez',
    location: 'Weight Room',
    image: '/strength-training-weightlifting-gym.jpg'
  },
  {
    id: '4',
    name: 'CYCLE BURN',
    type: 'Cycling',
    schedule: 'MON-WED-FRI',
    time: '6:30 AM - 7:15 AM',
    duration: '45 min',
    spots: 5,
    totalSpots: 25,
    instructor: 'Emma Wilson',
    location: 'Cycle Studio',
    image: '/cycling-class-indoor-bikes.jpg'
  },
  {
    id: '5',
    name: 'BOXING FURY',
    type: 'Boxing',
    schedule: 'TUE-THU-SAT',
    time: '7:00 PM - 8:00 PM',
    duration: '60 min',
    spots: 10,
    totalSpots: 12,
    instructor: 'Marcus Thompson',
    location: 'Fight Zone',
    image: '/boxing-training-gym-heavy-bag.jpg'
  },
  {
    id: '6',
    name: 'PILATES CORE',
    type: 'Pilates',
    schedule: 'MON-WED-FRI',
    time: '9:00 AM - 10:00 AM',
    duration: '60 min',
    spots: 6,
    totalSpots: 10,
    instructor: 'Sofia Park',
    location: 'Mat Room',
    image: '/pilates-class-studio-reformer.jpg'
  }
]

export function ListadoActividades() {
  return (
    <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
      {activities.map((activity) => (
        <Card key={activity.id} className="overflow-hidden group border-2 hover:border-primary transition-all">
          <Link href={`/actividades/${activity.id}`}>
            <div className="relative h-48 overflow-hidden">
              <img 
                src={activity.image || "/placeholder.svg"}
                alt={activity.name}
                className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
              />
              <div className="absolute inset-0 bg-gradient-to-t from-card via-card/40 to-transparent" />
              <div className="absolute top-4 left-4">
                <span className="inline-block px-3 py-1 bg-primary text-primary-foreground text-xs font-bold uppercase tracking-wide rounded-md">
                  {activity.type}
                </span>
              </div>
              <div className="absolute top-4 right-4">
                <span className="inline-block px-3 py-1 bg-background/90 backdrop-blur text-foreground text-xs font-bold rounded-md">
                  {activity.duration}
                </span>
              </div>
            </div>
          </Link>
          
          <div className="p-6">
            <Link href={`/actividades/${activity.id}`}>
              <h3 className="font-mono text-xl font-bold uppercase tracking-tight mb-3 hover:text-primary transition-colors">
                {activity.name}
              </h3>
            </Link>
            
            <div className="space-y-2 mb-4">
              <div className="flex items-center gap-2 text-sm text-muted-foreground">
                <Calendar className="h-4 w-4" />
                <span>{activity.schedule}</span>
              </div>
              <div className="flex items-center gap-2 text-sm text-muted-foreground">
                <Clock className="h-4 w-4" />
                <span>{activity.time}</span>
              </div>
              <div className="flex items-center gap-2 text-sm">
                <Users className="h-4 w-4" />
                <span className={activity.spots > 5 ? 'text-green-500' : 'text-orange-500'}>
                  {activity.spots} / {activity.totalSpots} spots available
                </span>
              </div>
              <div className="flex items-center gap-2 text-sm text-muted-foreground">
                <MapPin className="h-4 w-4" />
                <span>{activity.location}</span>
              </div>
              <p className="text-sm font-medium pt-2">Instructor: {activity.instructor}</p>
            </div>
            
            <div className="flex gap-2">
              <Button className="flex-1" size="lg" disabled={activity.spots === 0}>
                {activity.spots === 0 ? 'Full' : 'Book Now'}
              </Button>
              <Button variant="outline" size="lg" asChild>
                <Link href={`/actividades/${activity.id}`}>
                  Details
                </Link>
              </Button>
            </div>
          </div>
        </Card>
      ))}
    </div>
  )
}
