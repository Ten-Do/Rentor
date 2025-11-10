import type { AdvertisementType, AdvertisementRooms } from '../types'

export const PROPERTY_TYPE_MAP: Record<AdvertisementType, string> = {
  apartment: 'Apartment',
  house: 'House',
  room: 'Room',
} as const

export const ROOMS_LABEL_MAP: Record<AdvertisementRooms, string> = {
  studio: 'Studio',
  '1': 'Room',
  '2': 'Rooms',
  '3': 'Rooms',
  '4': 'Rooms',
  '5': 'Rooms',
  '6+': 'Rooms',
} as const

export const CURRENCY_SYMBOL = '₸'
