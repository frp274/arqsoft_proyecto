'use client'

import { Search } from 'lucide-react'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Card } from '@/components/ui/card'

export function SearchActivities() {
  return (
    <section className="py-16 md:py-24 border-b border-border">
      <div className="container px-4 md:px-6">
        <div className="max-w-4xl mx-auto">
          <h2 className="font-mono text-3xl md:text-4xl font-bold uppercase tracking-tighter mb-8">
            Find Your Activity
          </h2>
          <div className="flex gap-2 mb-12">
            <div className="relative flex-1">
              <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-5 w-5 text-muted-foreground" />
              <Input 
                placeholder="Search classes, trainers, or activities..."
                className="pl-10 h-12 bg-card"
              />
            </div>
            <Button size="lg" className="h-12">
              Search
            </Button>
          </div>
          <div className="grid grid-cols-2 md:grid-cols-4 gap-3">
            {['Yoga', 'HIIT', 'Strength', 'Cycling', 'Boxing', 'Pilates', 'Swimming', 'Dance'].map((activity) => (
              <Button 
                key={activity}
                variant="outline" 
                className="h-14 text-sm font-medium hover:bg-primary hover:text-primary-foreground hover:border-primary transition-all"
              >
                {activity}
              </Button>
            ))}
          </div>
        </div>
      </div>
    </section>
  )
}
