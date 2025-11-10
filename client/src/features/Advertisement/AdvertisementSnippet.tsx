import type { Advertisement } from '../../types'
import {
  PROPERTY_TYPE_MAP,
  ROOMS_LABEL_MAP,
  CURRENCY_SYMBOL,
} from '../../utils/constants'
import { Button } from '../../components/Button'
import { useModal } from '../../contexts'
import { useAuth } from '../../hooks/data/useAuth'

export interface AdvertisementSnippetProps {
  advertisement: Advertisement
}

export const AdvertisementSnippet = ({
  advertisement,
}: AdvertisementSnippetProps) => {
  const {
    title,
    price,
    type,
    rooms,
    city,
    square,
    landlord_name,
    landlord_email,
    landlord_phone,
  } = advertisement
  const { open } = useModal()
  const { isAuth } = useAuth()

  const handleOpenLandlordModal = () => {
    if (isAuth) {
      open('landlord', {
        landlord_name,
        landlord_email,
        landlord_phone,
      })
    } else {
      open('login')
    }
  }

  const propertyType = PROPERTY_TYPE_MAP[type]
  const roomsLabel =
    rooms === 'studio'
      ? ROOMS_LABEL_MAP.studio
      : `${rooms} ${ROOMS_LABEL_MAP[rooms]}`

  return (
    <div className="bg-gray-900 rounded-lg p-6 border border-gray-800 flex flex-col gap-4">
      <h1 className="text-3xl font-bold text-white">{title}</h1>
      <div className="flex flex-col gap-4">
        <div className="flex items-center gap-4">
          <span className="text-3xl font-bold text-indigo-400">
            {price.toLocaleString('en-US')} {CURRENCY_SYMBOL}
          </span>
        </div>
        <div className="flex flex-wrap items-center gap-3 text-sm text-gray-400">
          <span>{city}</span>
          <span>•</span>
          <span>{roomsLabel}</span>
          <span>•</span>
          <span>{propertyType}</span>
          {square && (
            <>
              <span>•</span>
              <span>{square} m²</span>
            </>
          )}
        </div>
      </div>
      <div className="flex-wrap flex gap-4 mt-auto pt-4 border-t border-gray-800">
        <Button
          onClick={handleOpenLandlordModal}
          size="lg"
          className="whitespace-nowrap flex-1"
        >
          Contact by Email
        </Button>
        {landlord_phone && (
          <Button
            onClick={handleOpenLandlordModal}
            size="lg"
            variant="outline"
            className="whitespace-nowrap flex-1"
          >
            Call
          </Button>
        )}
      </div>
    </div>
  )
}
