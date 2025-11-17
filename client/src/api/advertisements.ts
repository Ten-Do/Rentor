import type { Advertisement, AdvertisementListResponse } from '../types'
import { $api } from './$api'

export interface AdvertisementCreateData {
  title: string
  description: string
  price: number
  type: 'apartment' | 'house' | 'room'
  rooms: 'studio' | '1' | '2' | '3' | '4' | '5' | '6+'
  city: string
  address: string
  latitude?: number | null
  longitude?: number | null
  square?: number | null
}

export interface AdvertisementUpdateData
  extends Partial<AdvertisementCreateData> {
  status?: 'active' | 'paused'
}

export const getMyAdvertisements = async () =>
  $api.get<AdvertisementListResponse>('/v1/advertisements/my')

export const createAdvertisement = async (data: AdvertisementCreateData) =>
  $api.post<Advertisement>('/v1/advertisements', data)

export const updateAdvertisement = async (
  id: string,
  data: AdvertisementUpdateData
) => $api.put<Advertisement>(`/v1/advertisements/${id}`, data)

export const IMAGES_FIELD_NAME = 'files' as const

export interface ImageUploadResponse {
  image_url: string
  image_id: string
}

export interface ImagesUploadResponse {
  uploaded_images: ImageUploadResponse[]
  total_uploaded: number
}

export const uploadAdvertisementImages = async (
  id: string,
  files: File[]
): Promise<ImagesUploadResponse> => {
  const formData = new FormData()
  files.forEach((file) => formData.append(IMAGES_FIELD_NAME, file))

  const response = await fetch(`/v1/advertisements/${id}/images`, {
    method: 'POST',
    credentials: 'include',
    body: formData,
  })

  if (!response.ok) {
    const errorData = await response.json().catch(() => ({}))
    const message =
      (errorData as { message?: string }).message || 'Upload failed'
    throw new Error(message)
  }

  return (await response.json()) as ImagesUploadResponse
}

export const deleteAdvertisementImage = async (
  adId: string,
  imageId: string
): Promise<void> => {
  const response = await fetch(`/v1/advertisements/${adId}/images/${imageId}`, {
    method: 'DELETE',
    credentials: 'include',
  })

  if (!response.ok) {
    const errorData = await response.json().catch(() => ({}))
    const message =
      (errorData as { message?: string }).message || 'Delete failed'
    throw new Error(message)
  }

  return
}
