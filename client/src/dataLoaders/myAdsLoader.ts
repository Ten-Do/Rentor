import type { AdvertisementLite } from '../types'
import { $api } from '../api'
import { redirect } from 'react-router-dom'

interface BackendAdPreview {
  id: number
  title: string
  city: string
  price: number
  type: string
  rooms: string
  square: number
  imageUrl: { imageId: number; imageUrl: string } | null
}

interface BackendResponse {
  items: BackendAdPreview[]
  total: number
  page: number
  limit: number
}

export interface MyAdsLoaderData {
  ads: AdvertisementLite[]
}

export const myAdsLoader = async (): Promise<MyAdsLoaderData> => {
  try {
    const response = await $api.get<BackendResponse>('/advertisements/my')

    const ads: AdvertisementLite[] = (response?.items || []).map((item) => ({
      id: String(item.id),
      title: item.title,
      price: item.price,
      type: item.type as AdvertisementLite['type'],
      rooms: item.rooms as AdvertisementLite['rooms'],
      city: item.city,
      image_url: item.imageUrl?.imageUrl || null,
    }))

    return { ads }
  } catch {
    throw redirect('/')
  }
}
