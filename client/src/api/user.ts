import { $api } from './$api'
import type { UserProfileUpdate } from '../types'

export const updateUserProfileAction = async (
  data: UserProfileUpdate
): Promise<void> => {
  await $api.put<void>('/v1/user/profile', data)
}
