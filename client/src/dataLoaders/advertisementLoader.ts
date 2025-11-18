import type { Advertisement } from '../types'
import type { Params } from 'react-router-dom'
import { $api } from '../api'

export interface AdvertisementLoaderData {
  advertisement: Advertisement
}

interface BackendImageUrl {
  imageId: number
  imageUrl: string
}

interface BackendAdvertisementResponse {
  id: number
  title: string
  description: string | null
  price: number
  type: string
  rooms: string
  city: string
  address: string
  latitude: number | null
  longitude: number | null
  square: number
  status: string
  landlordName: string | null
  landlordEmail: string
  landlordPhone: string | null
  imageUrls: BackendImageUrl[]
}

export const advertisementLoader = async ({
  params,
}: {
  params: Params
}): Promise<AdvertisementLoaderData> => {
  const { id } = params

  if (!id) {
    throw new Response('Advertisement ID is required', { status: 400 })
  }

  const endpoint = `/advertisements/${id}`
  const backendAd = await $api.get<BackendAdvertisementResponse>(endpoint)

  const advertisement: Advertisement = {
    id: String(backendAd.id),
    title: backendAd.title,
    description: backendAd.description || '',
    price: backendAd.price,
    type: backendAd.type as Advertisement['type'],
    rooms: backendAd.rooms as Advertisement['rooms'],
    city: backendAd.city,
    image_url: backendAd.imageUrls[0]?.imageUrl || null,
    address: backendAd.address,
    latitude: backendAd.latitude,
    longitude: backendAd.longitude,
    square: backendAd.square || null,
    image_urls: backendAd.imageUrls.map((img) => img.imageUrl),
    images: backendAd.imageUrls.map((img) => ({
      id: String(img.imageId),
      url: img.imageUrl,
    })),
    landlord_name: backendAd.landlordName || '',
    landlord_email: backendAd.landlordEmail,
    landlord_phone: backendAd.landlordPhone,
    status: backendAd.status as Advertisement['status'],
    created_at: '',
  }

  return { advertisement }
}
