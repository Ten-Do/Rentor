import type { User } from '../types'
import { $api } from '../api'

export interface AuthLoaderData {
  user: User | null
}

export const authLoader = async (): Promise<AuthLoaderData> => {
  const user = await $api.get<User>('/v1/user/profile').catch(() => null)
  return { user }
}
