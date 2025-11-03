import { useLoaderData } from 'react-router-dom'
import {
  ImageSlider,
  AdvertisementSnippet,
  StaticMap,
} from '../features/Advertisement'
import type { AdvertisementLoaderData } from '../dataLoaders/advertisementLoader'
import {
  PROPERTY_TYPE_MAP,
  ROOMS_LABEL_MAP,
  CURRENCY_SYMBOL,
} from '../utils/constants'

export const Advertisement = () => {
  const { advertisement } = useLoaderData() as AdvertisementLoaderData

  const propertyType = PROPERTY_TYPE_MAP[advertisement.type]
  const roomsLabel =
    advertisement.rooms === 'studio'
      ? ROOMS_LABEL_MAP.studio
      : `${advertisement.rooms} ${ROOMS_LABEL_MAP[advertisement.rooms]}`

  return (
    <div className="space-y-8">
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <ImageSlider images={advertisement.image_urls} />
        <AdvertisementSnippet advertisement={advertisement} />
      </div>

      <div className="space-y-8">
        <section>
          <h2 className="text-2xl font-bold text-white mb-4">Description</h2>
          <p className="text-gray-300 leading-relaxed">
            {advertisement.description}
          </p>
        </section>

        <section className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div className="bg-gray-900 rounded-lg p-6 border border-gray-800">
            <h3 className="text-xl font-semibold text-white mb-4">
              Characteristics
            </h3>
            <dl className="space-y-3">
              <div className="flex justify-between">
                <dt className="text-gray-400">Type</dt>
                <dd className="text-white">{propertyType}</dd>
              </div>
              <div className="flex justify-between">
                <dt className="text-gray-400">Rooms</dt>
                <dd className="text-white">{roomsLabel}</dd>
              </div>
              {advertisement.square && (
                <div className="flex justify-between">
                  <dt className="text-gray-400">Area</dt>
                  <dd className="text-white">{advertisement.square} m²</dd>
                </div>
              )}
              <div className="flex justify-between">
                <dt className="text-gray-400">Price</dt>
                <dd className="text-white font-semibold">
                  {advertisement.price.toLocaleString('en-US')}{' '}
                  {CURRENCY_SYMBOL}
                </dd>
              </div>
            </dl>
          </div>

          <div className="bg-gray-900 rounded-lg p-6 border border-gray-800">
            <h3 className="text-xl font-semibold text-white mb-4">Address</h3>
            <div className="space-y-2">
              <p className="text-gray-300">{advertisement.address}</p>
              <p className="text-gray-400 text-sm">{advertisement.city}</p>
            </div>
          </div>
        </section>

        <section>
          <h3 className="text-xl font-semibold text-white mb-4">Location</h3>
          <StaticMap
            latitude={advertisement.latitude}
            longitude={advertisement.longitude}
          />
        </section>
      </div>
    </div>
  )
}
