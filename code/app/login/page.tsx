'use client'

import { useState } from 'react'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Card } from '@/components/ui/card'
import { Dumbbell } from 'lucide-react'
import Link from 'next/link'

export default function LoginPage() {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault()
    // TODO: Connect to your users API
    console.log('[v0] Login attempt:', { email, password })
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-background p-4">
      <div className="absolute inset-0 bg-[url('/athletic-person-working-out-in-modern-gym.jpg')] bg-cover bg-center opacity-10" />
      
      <Card className="relative w-full max-w-md p-8 border-2">
        <div className="flex flex-col items-center mb-8">
          <div className="flex h-16 w-16 items-center justify-center rounded-xl bg-primary mb-4">
            <Dumbbell className="h-8 w-8 text-primary-foreground" />
          </div>
          <h1 className="font-mono text-3xl font-bold uppercase tracking-tight">MyGym</h1>
          <p className="text-muted-foreground text-sm mt-2">Sign in to your account</p>
        </div>

        <form onSubmit={handleLogin} className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="email">Email</Label>
            <Input
              id="email"
              type="email"
              placeholder="your@email.com"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
              className="h-12"
            />
          </div>

          <div className="space-y-2">
            <Label htmlFor="password">Password</Label>
            <Input
              id="password"
              type="password"
              placeholder="Enter your password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
              className="h-12"
            />
          </div>

          <Button type="submit" className="w-full h-12" size="lg">
            Sign In
          </Button>
        </form>

        <div className="mt-6 text-center space-y-2">
          <Link href="/registro" className="text-sm text-muted-foreground hover:text-foreground transition-colors block">
            Don't have an account? <span className="text-primary font-medium">Sign up</span>
          </Link>
          <Link href="/login-admin" className="text-sm text-muted-foreground hover:text-foreground transition-colors block">
            Admin Login
          </Link>
        </div>
      </Card>
    </div>
  )
}
