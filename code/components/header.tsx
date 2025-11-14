import { Button } from '@/components/ui/button'
import { Search, Menu, User, Calendar } from 'lucide-react'
import Link from 'next/link'

export function Header() {
  return (
    <header className="sticky top-0 z-50 border-b border-border bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
      <div className="container flex h-16 items-center justify-between px-4 md:px-6">
        <div className="flex items-center gap-8">
          <Link href="/" className="flex items-center gap-2">
            <div className="flex h-10 w-10 items-center justify-center rounded-xl bg-primary">
              <span className="font-mono text-xl font-bold text-primary-foreground">MG</span>
            </div>
            <span className="text-xl font-bold uppercase tracking-tight">MyGym</span>
          </Link>
          <nav className="hidden md:flex items-center gap-6">
            <Link href="/actividades" className="text-sm font-medium text-muted-foreground hover:text-foreground transition-colors">
              Classes
            </Link>
            <Link href="/mis-inscripciones" className="text-sm font-medium text-muted-foreground hover:text-foreground transition-colors">
              My Bookings
            </Link>
            <a href="#trainers" className="text-sm font-medium text-muted-foreground hover:text-foreground transition-colors">
              Trainers
            </a>
            <a href="#membership" className="text-sm font-medium text-muted-foreground hover:text-foreground transition-colors">
              Membership
            </a>
          </nav>
        </div>
        <div className="flex items-center gap-3">
          <Button variant="ghost" size="icon" className="hidden md:flex" asChild>
            <Link href="/actividades">
              <Search className="h-5 w-5" />
            </Link>
          </Button>
          <Button variant="ghost" size="icon" className="hidden md:flex" asChild>
            <Link href="/login">
              <User className="h-5 w-5" />
            </Link>
          </Button>
          <Button variant="default" className="hidden md:inline-flex" asChild>
            <Link href="/actividades">Book Now</Link>
          </Button>
          <Button variant="ghost" size="icon" className="md:hidden">
            <Menu className="h-6 w-6" />
          </Button>
        </div>
      </div>
    </header>
  )
}
