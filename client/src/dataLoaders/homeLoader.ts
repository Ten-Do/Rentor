import type { AdvertisementListResponse, AdvertisementLite } from '../types'
import { $api } from '../api'

export interface HomeLoaderData {
  ads: AdvertisementLite[]
}

export const homeLoader = async ({
  request,
}: {
  request: Request
}): Promise<HomeLoaderData> => {
  const responseData =
    await $api.get<AdvertisementListResponse>('/v1/advertisements')
  let ads = responseData?.data || []

  const url = new URL(request.url)
  const keywords = url.searchParams.get('keywords')
  const city = url.searchParams.get('city')
  const type = url.searchParams.get('type')
  const rooms = url.searchParams.get('rooms')
  const minPrice = url.searchParams.get('minPrice')
  const maxPrice = url.searchParams.get('maxPrice')

  if (keywords) {
    const keywordsLower = keywords.toLowerCase()
    ads = ads.filter((ad) => ad.title.toLowerCase().includes(keywordsLower))
  }

  if (city) {
    ads = ads.filter((ad) => ad.city === city)
  }

  if (type) {
    ads = ads.filter((ad) => ad.type === type)
  }

  if (rooms) {
    ads = ads.filter((ad) => ad.rooms === rooms)
  }

  if (minPrice) {
    const min = Number.parseFloat(minPrice)
    if (!Number.isNaN(min)) {
      ads = ads.filter((ad) => ad.price >= min)
    }
  }

  if (maxPrice) {
    const max = Number.parseFloat(maxPrice)
    if (!Number.isNaN(max)) {
      ads = ads.filter((ad) => ad.price <= max)
    }
  }

  return { ads }
}
