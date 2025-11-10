import type { AdvertisementLite, AdvertisementListResponse } from '../types'
import { $api } from '../api'

export interface MyAdsLoaderData {
  ads: AdvertisementLite[]
}

export const myAdsLoader = async (): Promise<MyAdsLoaderData> => {
  const response = await $api.get<AdvertisementListResponse>(
    '/v1/advertisements/my'
  )
  return { ads: response?.data || [] }
}
