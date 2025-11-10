export type AdvertisementType = 'apartment' | 'house' | 'room'

export type AdvertisementRooms = 'studio' | '1' | '2' | '3' | '4' | '5' | '6+'

export type AdvertisementStatus = 'active' | 'paused'

export interface AdvertisementLite {
  id: string
  title: string
  price: number
  type: AdvertisementType
  rooms: AdvertisementRooms
  city: string
  image_url: string | null
}

export interface Advertisement extends AdvertisementLite {
  description: string
  address: string
  latitude: number | null
  longitude: number | null
  square: number | null
  image_urls: string[]
  images?: { id: string; url: string }[]
  landlord_name: string
  landlord_email: string
  landlord_phone: string | null
  status: AdvertisementStatus
  created_at: string
}

export interface Pagination {
  page: number
  limit: number
  total: number
  total_pages: number
}

export interface AdvertisementListResponse {
  data: AdvertisementLite[]
  pagination: Pagination
}

export interface User {
  id: string
  email: string
  phone_number: string | null
  first_name: string
  surname: string
  patronymic: string | null
  created_at: string
}

export interface UserProfileUpdate {
  first_name: string
  surname: string
  patronymic: string | null
  phone_number: string | null
}
