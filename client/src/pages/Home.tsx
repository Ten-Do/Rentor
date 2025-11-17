import { useLoaderData } from 'react-router-dom'
import { AdvertisementsGrid } from '../features/Advertisements'
import { SearchForm } from '../features/Search'
import type { HomeLoaderData } from '../dataLoaders/homeLoader'

export const Home = () => {
  const { ads } = useLoaderData() as HomeLoaderData

  return (
    <div>
      <h1 className="text-3xl font-bold mb-8 text-white">Advertisements</h1>
      <SearchForm />
      <AdvertisementsGrid ads={ads} />
    </div>
  )
}
