import { useLoaderData, Link } from 'react-router-dom'
import type { AdvertisementLite } from '../types'
import { Button } from '../components/Button'
import { AdCard } from '../components/AdCard'

export const MyAds = () => {
    const { ads } = useLoaderData() as { ads: AdvertisementLite[] }

    return (
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
            <div className="flex items-center justify-between mb-6">
                <h1 className="text-2xl font-semibold text-white">My ads</h1>
                <Link to="/my/ads/new">
                    <Button size="md">Add advertisement</Button>
                </Link>
            </div>

            {ads.length === 0 ? (
                <div className="text-center py-16 border border-dashed border-gray-800 rounded-lg bg-gray-900">
                    <p className="text-gray-400 text-lg mb-6">You don’t have any ads yet</p>
                </div>
            ) : (
                <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
                    {ads.map((ad) => (
                        <AdCard key={ad.id} {...ad} to={`/my/ads/${ad.id}/edit`} />
                    ))}
                </div>
            )}
        </div>
    )
}


