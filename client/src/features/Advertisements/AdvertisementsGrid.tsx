import { AdCard } from '../../components'
import type { AdvertisementLite } from '../../types'

interface AdvertisementsGridProps {
  ads: AdvertisementLite[]
}

export const AdvertisementsGrid = ({ ads }: AdvertisementsGridProps) => {
  if (ads.length === 0) {
    return (
      <div className="text-center py-12">
        <p className="text-gray-400 text-lg">No advertisements found</p>
      </div>
    )
  }

  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
      {ads.map((ad) => (
        <AdCard key={ad.id} {...ad} to={`/advertisement/${ad.id}`} />
      ))}
    </div>
  )
}
