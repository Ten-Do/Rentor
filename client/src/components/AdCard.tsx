import { Link } from './Link'
import { Image } from './Image'
import {
  PROPERTY_TYPE_MAP,
  ROOMS_LABEL_MAP,
  CURRENCY_SYMBOL,
} from '../utils/constants'
import type { AdvertisementLite } from '../types'

export type AdCardProps = AdvertisementLite & { to?: string }

export const AdCard = ({
  title,
  price,
  type,
  rooms,
  city,
  image_url,
  to,
}: AdCardProps) => {
  const propertyType = PROPERTY_TYPE_MAP[type]
  const roomsLabel =
    rooms === 'studio'
      ? ROOMS_LABEL_MAP.studio
      : `${rooms} ${ROOMS_LABEL_MAP[rooms]}`

  return (
    <Link
      to={to}
      className="block bg-gray-900 rounded-lg overflow-hidden border border-gray-800 hover:border-indigo-600 transition-colors duration-200"
    >
      <div className="aspect-video w-full overflow-hidden bg-gray-800">
        <Image
          src={image_url}
          alt={title}
          className="w-full h-full object-cover"
          loading="lazy"
        />
      </div>
      <div className="p-4 space-y-2">
        <h3 className="text-lg font-semibold text-white line-clamp-2">
          {title}
        </h3>
        <div className="flex items-center justify-between">
          <span className="text-2xl font-bold text-indigo-400">
            {price.toLocaleString('en-US')} {CURRENCY_SYMBOL}
          </span>
        </div>
        <div className="flex items-center gap-3 text-sm text-gray-400">
          <span>{city}</span>
          <span>•</span>
          <span>{roomsLabel}</span>
          <span>•</span>
          <span>{propertyType}</span>
        </div>
      </div>
    </Link>
  )
}
