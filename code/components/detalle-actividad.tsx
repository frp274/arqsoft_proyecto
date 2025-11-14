'use client'

import { Card } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Calendar, Clock, Users, MapPin, Star, ArrowLeft } from 'lucide-react'
import Link from 'next/link'

interface DetalleActividadProps {
  activityId: string
}

export function DetalleActividad({ activityId }: DetalleActividadProps) {
  // TODO: Fetch from your activities API using activityId
  const activity = {
    id: activityId,
    name: 'POWER HOUR',
    type: 'HIIT Training',
    description: 'High-intensity interval training that combines cardio and strength exercises. Push your limits and transform your fitness in this energetic 60-minute session designed to burn maximum calories and build lean muscle.',
    schedule: 'MON-FRI',
    time: '6:00 AM - 7:00 AM',
    duration: '60 min',
    spots: 12,
    totalSpots: 20,
    instructor: {
      name: 'Sarah Johnson',
      bio: 'Certified HIIT specialist with 8 years of experience. Sarah has helped hundreds of members achieve their fitness goals through her motivating and challenging training style.',
      image: '/female-athlete-confident-portrait.jpg'
    },
    location: 'Studio A',
    level: 'Intermediate to Advanced',
    equipment: 'Dumbbells, kettlebells, resistance bands',
    image: '/intense-gym-class-hiit-training.jpg',
    rating: 4.8,
    reviews: 127
  }

  return (
    <div className="container px-4 md:px-6 py-12">
      <Button variant="ghost" asChild className="mb-6">
        <Link href="/actividades">
          <ArrowLeft className="h-4 w-4 mr-2" />
          Back to Activities
        </Link>
      </Button>

      <div className="grid lg:grid-cols-5 gap-8">
        <div className="lg:col-span-3 space-y-6">
          <div className="relative h-[400px] rounded-xl overflow-hidden border-2">
            <img 
              src={activity.image || "/placeholder.svg"}
              alt={activity.name}
              className="w-full h-full object-cover"
            />
            <div className="absolute top-4 left-4">
              <span className="inline-block px-4 py-2 bg-primary text-primary-foreground text-sm font-bold uppercase tracking-wide rounded-lg">
                {activity.type}
              </span>
            </div>
          </div>

          <div>
            <h1 className="font-mono text-4xl md:text-5xl font-bold uppercase tracking-tighter mb-4">
              {activity.name}
            </h1>
            <div className="flex items-center gap-3 mb-6">
              <div className="flex items-center gap-1">
                <Star className="h-5 w-5 fill-primary text-primary" />
                <span className="font-bold">{activity.rating}</span>
              </div>
              <span className="text-muted-foreground">({activity.reviews} reviews)</span>
            </div>
            <p className="text-lg text-muted-foreground leading-relaxed">
              {activity.description}
            </p>
          </div>

          <Card className="p-6 border-2">
            <h3 className="font-mono text-xl font-bold uppercase tracking-tight mb-4">
              Class Details
            </h3>
            <div className="grid md:grid-cols-2 gap-4">
              <div className="space-y-3">
                <div className="flex items-start gap-3">
                  <Calendar className="h-5 w-5 text-primary mt-0.5" />
                  <div>
                    <p className="font-medium">Schedule</p>
                    <p className="text-sm text-muted-foreground">{activity.schedule}</p>
                  </div>
                </div>
                <div className="flex items-start gap-3">
                  <Clock className="h-5 w-5 text-primary mt-0.5" />
                  <div>
                    <p className="font-medium">Time</p>
                    <p className="text-sm text-muted-foreground">{activity.time}</p>
                  </div>
                </div>
                <div className="flex items-start gap-3">
                  <MapPin className="h-5 w-5 text-primary mt-0.5" />
                  <div>
                    <p className="font-medium">Location</p>
                    <p className="text-sm text-muted-foreground">{activity.location}</p>
                  </div>
                </div>
              </div>
              <div className="space-y-3">
                <div>
                  <p className="font-medium mb-1">Duration</p>
                  <p className="text-sm text-muted-foreground">{activity.duration}</p>
                </div>
                <div>
                  <p className="font-medium mb-1">Level</p>
                  <p className="text-sm text-muted-foreground">{activity.level}</p>
                </div>
                <div>
                  <p className="font-medium mb-1">Equipment Needed</p>
                  <p className="text-sm text-muted-foreground">{activity.equipment}</p>
                </div>
              </div>
            </div>
          </Card>

          <Card className="p-6 border-2">
            <h3 className="font-mono text-xl font-bold uppercase tracking-tight mb-4">
              Your Instructor
            </h3>
            <div className="flex gap-4">
              <img 
                src={activity.instructor.image || "/placeholder.svg"}
                alt={activity.instructor.name}
                className="w-20 h-20 rounded-xl object-cover border-2"
              />
              <div className="flex-1">
                <h4 className="font-bold text-lg mb-1">{activity.instructor.name}</h4>
                <p className="text-sm text-muted-foreground leading-relaxed">
                  {activity.instructor.bio}
                </p>
              </div>
            </div>
          </Card>
        </div>

        <div className="lg:col-span-2">
          <Card className="p-6 border-2 sticky top-24">
            <div className="space-y-6">
              <div>
                <div className="flex items-center justify-between mb-2">
                  <span className="text-sm font-medium text-muted-foreground">Availability</span>
                  <span className="text-sm font-bold">{activity.spots} / {activity.totalSpots} spots</span>
                </div>
                <div className="w-full bg-border rounded-full h-2 overflow-hidden">
                  <div 
                    className="bg-primary h-full rounded-full transition-all"
                    style={{ width: `${(activity.spots / activity.totalSpots) * 100}%` }}
                  />
                </div>
              </div>

              <Button size="lg" className="w-full h-14 text-lg" disabled={activity.spots === 0}>
                {activity.spots === 0 ? 'Class Full' : 'Book This Class'}
              </Button>

              <div className="pt-4 border-t space-y-3">
                <h4 className="font-bold">What to bring:</h4>
                <ul className="space-y-2 text-sm text-muted-foreground">
                  <li className="flex items-start gap-2">
                    <span className="text-primary mt-0.5">•</span>
                    <span>Water bottle</span>
                  </li>
                  <li className="flex items-start gap-2">
                    <span className="text-primary mt-0.5">•</span>
                    <span>Workout towel</span>
                  </li>
                  <li className="flex items-start gap-2">
                    <span className="text-primary mt-0.5">•</span>
                    <span>Athletic shoes</span>
                  </li>
                  <li className="flex items-start gap-2">
                    <span className="text-primary mt-0.5">•</span>
                    <span>Positive attitude</span>
                  </li>
                </ul>
              </div>

              <div className="pt-4 border-t">
                <h4 className="font-bold mb-2">Cancellation Policy</h4>
                <p className="text-sm text-muted-foreground">
                  Cancel up to 2 hours before class for a full refund. Late cancellations will be charged.
                </p>
              </div>
            </div>
          </Card>
        </div>
      </div>
    </div>
  )
}
