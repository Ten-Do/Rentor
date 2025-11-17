import { useRouteLoaderData } from 'react-router-dom'
import type { AuthLoaderData } from '../../dataLoaders/authLoader'

export const useAuth = () => {
  const { user } = useRouteLoaderData('root') as AuthLoaderData
  return { user, isAuth: Boolean(user) }
}
