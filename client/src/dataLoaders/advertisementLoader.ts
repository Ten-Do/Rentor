import type { Advertisement } from '../types'
import type { Params } from 'react-router-dom'
import { $api } from '../api'

export interface AdvertisementLoaderData {
  advertisement: Advertisement
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

  const endpoint = `/v1/advertisements/${id}`
  const advertisement = await $api.get<Advertisement>(endpoint)

  return { advertisement }
}
