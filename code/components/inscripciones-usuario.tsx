'use client'

import { Card } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Calendar, Clock, MapPin, X } from 'lucide-react'
import Link from 'next/link'

// TODO: Replace with data from your users API (user's bookings)
const bookings = [
  {
    id: '1',
    activityId: '1',
    activityName: 'POWER HOUR',
    type: 'HIIT Training',
    date: 'Mon, Nov 18, 2025',
    time: '6:00 AM - 7:00 AM',
    instructor: 'Sarah Johnson',
    location: 'Studio A',
    status: 'confirmed',
    image: '/intense-gym-class-hiit-training.jpg'
  },
  {
    id: '2',
    activityId: '2',
    activityName: 'YOGA FLOW',
    type: 'Vinyasa Yoga',
    date: 'Tue, Nov 19, 2025',
    time: '7:30 PM - 8:30 PM',
    instructor: 'Michael Chen',
    location: 'Zen Room',
    status: 'confirmed',
    image: '/yoga-class-peaceful-studio.jpg'
  },
  {
    id: '3',
    activityId: '3',
    activityName: 'IRON TEMPLE',
    type: 'Strength Training',
    date: 'Wed, Nov 20, 2025',
    time: '5:00 PM - 6:30 PM',
    instructor: 'David Martinez',
    location: 'Weight Room',
    status: 'confirmed',
    image: '/strength-training-weightlifting-gym.jpg'
  }
]

export function InscripcionesUsuario() {
  const handleCancelBooking = (bookingId: string) => {
    // TODO: Connect to your activities API to cancel booking
    console.log('[v0] Cancel booking:', bookingId)
  }

  if (bookings.length === 0) {
    return (
      <Card className="p-12 text-center border-2">
        <div className="max-w-md mx-auto">
          <h3 className="font-mono text-2xl font-bold uppercase tracking-tight mb-4">
            No Bookings Yet
          </h3>
          <p className="text-muted-foreground mb-6">
            Start your fitness journey by booking your first class
          </p>
          <Button size="lg" asChild>
            <Link href="/actividades">Browse Activities</Link>
          </Button>
        </div>
      </Card>
    )
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <p className="text-muted-foreground">
          You have {bookings.length} upcoming {bookings.length === 1 ? 'class' : 'classes'}
        </p>
        <Button variant="outline" asChild>
          <Link href="/actividades">Book More Classes</Link>
        </Button>
      </div>

      <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
        {bookings.map((booking) => (
          <Card key={booking.id} className="overflow-hidden border-2 hover:border-primary transition-all">
            <div className="relative h-40 overflow-hidden">
              <img 
                src={booking.image || "/placeholder.svg"}
                alt={booking.activityName}
                className="w-full h-full object-cover"
              />
              <div className="absolute inset-0 bg-gradient-to-t from-card via-card/60 to-transparent" />
              <div className="absolute top-3 left-3">
                <span className="inline-block px-3 py-1 bg-primary text-primary-foreground text-xs font-bold uppercase tracking-wide rounded-md">
                  {booking.type}
                </span>
              </div>
              <div className="absolute top-3 right-3">
                <span className="inline-block px-3 py-1 bg-green-500 text-white text-xs font-bold uppercase rounded-md">
                  Confirmed
                </span>
              </div>
            </div>
            
            <div className="p-5">
              <Link href={`/actividades/${booking.activityId}`}>
                <h3 className="font-mono text-lg font-bold uppercase tracking-tight mb-3 hover:text-primary transition-colors">
                  {booking.activityName}
                </h3>
              </Link>
              
              <div className="space-y-2 mb-4">
                <div className="flex items-center gap-2 text-sm text-muted-foreground">
                  <Calendar className="h-4 w-4" />
                  <span>{booking.date}</span>
                </div>
                <div className="flex items-center gap-2 text-sm text-muted-foreground">
                  <Clock className="h-4 w-4" />
                  <span>{booking.time}</span>
                </div>
                <div className="flex items-center gap-2 text-sm text-muted-foreground">
                  <MapPin className="h-4 w-4" />
                  <span>{booking.location}</span>
                </div>
                <p className="text-sm font-medium pt-1">Instructor: {booking.instructor}</p>
              </div>
              
              <Button 
                variant="destructive" 
                size="sm" 
                className="w-full"
                onClick={() => handleCancelBooking(booking.id)}
              >
                <X className="h-4 w-4 mr-2" />
                Cancel Booking
              </Button>
            </div>
          </Card>
        ))}
      </div>
    </div>
  )
}
