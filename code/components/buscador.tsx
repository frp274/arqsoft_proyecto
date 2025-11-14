'use client'

import { useState } from 'react'
import { Search, SlidersHorizontal } from 'lucide-react'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Card } from '@/components/ui/card'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'

export function Buscador() {
  const [searchTerm, setSearchTerm] = useState('')
  const [category, setCategory] = useState('all')

  const handleSearch = () => {
    // TODO: Connect to your search API
    console.log('[v0] Search:', { searchTerm, category })
  }

  return (
    <Card className="p-6 mb-8 border-2">
      <div className="flex flex-col md:flex-row gap-4">
        <div className="relative flex-1">
          <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-5 w-5 text-muted-foreground" />
          <Input 
            placeholder="Search activities, instructors, or times..."
            className="pl-10 h-12 bg-background"
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
          />
        </div>
        
        <Select value={category} onValueChange={setCategory}>
          <SelectTrigger className="w-full md:w-48 h-12">
            <SelectValue placeholder="Category" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="all">All Categories</SelectItem>
            <SelectItem value="yoga">Yoga</SelectItem>
            <SelectItem value="hiit">HIIT</SelectItem>
            <SelectItem value="strength">Strength</SelectItem>
            <SelectItem value="cycling">Cycling</SelectItem>
            <SelectItem value="boxing">Boxing</SelectItem>
            <SelectItem value="pilates">Pilates</SelectItem>
          </SelectContent>
        </Select>

        <Button size="lg" className="h-12 md:w-32" onClick={handleSearch}>
          <Search className="h-5 w-5 md:mr-2" />
          <span className="hidden md:inline">Search</span>
        </Button>

        <Button size="lg" variant="outline" className="h-12 md:w-32">
          <SlidersHorizontal className="h-5 w-5 md:mr-2" />
          <span className="hidden md:inline">Filters</span>
        </Button>
      </div>

      <div className="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-8 gap-2 mt-6">
        {['Yoga', 'HIIT', 'Strength', 'Cycling', 'Boxing', 'Pilates', 'Swimming', 'Dance'].map((activity) => (
          <Button 
            key={activity}
            variant="outline" 
            size="sm"
            className="text-xs font-medium hover:bg-primary hover:text-primary-foreground hover:border-primary transition-all"
          >
            {activity}
          </Button>
        ))}
      </div>
    </Card>
  )
}
