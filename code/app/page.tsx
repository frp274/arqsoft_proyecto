import { Header } from '@/components/header'
import { Hero } from '@/components/hero'
import { FeaturedClasses } from '@/components/featured-classes'
import { SearchActivities } from '@/components/search-activities'
import { MemberSpotlight } from '@/components/member-spotlight'

export default function Home() {
  return (
    <main className="min-h-screen bg-background">
      <Header />
      <Hero />
      <SearchActivities />
      <FeaturedClasses />
      <MemberSpotlight />
      {/* Additional content can be added here if needed */}
    </main>
  )
}
