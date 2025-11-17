import { useState, type FormEvent } from 'react'
import { useAuth } from '../hooks/data/useAuth'
import { Input, Button } from '../components'
import type { UserProfileUpdate } from '../types'
import { updateUserProfileAction } from '../api'
import { useModal } from '../contexts'

const formatDate = (dateString: string): string => {
  const date = new Date(dateString)
  const day = String(date.getDate()).padStart(2, '0')
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const year = date.getFullYear()
  return `${day}.${month}.${year}`
}

export const Profile = () => {
  const { user } = useAuth()
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [success, setSuccess] = useState(false)
  const { open } = useModal()

  if (!user) {
    return (
      <div className="bg-gray-900 rounded-lg p-6 border border-gray-800 flex items-center justify-between gap-4">
        <p className="text-gray-300">Please log in to view your profile.</p>
        <Button onClick={() => open('login')} size="md">
          Login
        </Button>
      </div>
    )
  }

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    setError(null)
    setSuccess(false)
    setIsLoading(true)

    const formData = new FormData(e.currentTarget)
    const updateData: UserProfileUpdate = {
      first_name: formData.get('first_name') as string,
      surname: formData.get('surname') as string,
      patronymic: (formData.get('patronymic') as string) || null,
      phone_number: (formData.get('phone_number') as string) || null,
    }

    try {
      await updateUserProfileAction(updateData)
      setSuccess(true)
    } catch (err) {
      setError((err as Error).message || 'Failed to update profile')
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <div className="space-y-8">
      <h1 className="text-3xl font-bold text-white">Profile</h1>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div className="bg-gray-900 rounded-lg p-6 border border-gray-800">
          <h2 className="text-xl font-semibold text-white mb-4">
            Account Information
          </h2>
          <table className="w-full text-sm">
            <tbody>
              <tr>
                <td className="text-gray-400 py-2 pr-6">ID</td>
                <td className="text-white font-mono">{user.id}</td>
              </tr>
              <tr>
                <td className="text-gray-400 py-2 pr-6">Email</td>
                <td className="text-white">{user.email}</td>
              </tr>
              <tr>
                <td className="text-gray-400 py-2 pr-6">Created At</td>
                <td className="text-white">{formatDate(user.created_at)}</td>
              </tr>
            </tbody>
          </table>
        </div>

        <div className="bg-gray-900 rounded-lg p-6 border border-gray-800">
          <h2 className="text-xl font-semibold text-white mb-4">
            Personal Information
          </h2>
          <form onSubmit={handleSubmit} className="space-y-4">
            <div>
              <Input
                label="First Name"
                id="first_name"
                name="first_name"
                type="text"
                defaultValue={user.first_name}
                required
                maxLength={100}
                disabled={isLoading}
              />
            </div>

            <div>
              <Input
                label="Surname"
                id="surname"
                name="surname"
                type="text"
                defaultValue={user.surname}
                required
                maxLength={100}
                disabled={isLoading}
              />
            </div>

            <div>
              <Input
                label="Patronymic"
                id="patronymic"
                name="patronymic"
                type="text"
                defaultValue={user.patronymic || ''}
                maxLength={100}
                disabled={isLoading}
              />
            </div>

            <div>
              <Input
                label="Phone Number"
                id="phone_number"
                name="phone_number"
                type="tel"
                defaultValue={user.phone_number || ''}
                disabled={isLoading}
              />
            </div>

            {error && (
              <div className="bg-red-900/20 border border-red-500 rounded-lg p-3">
                <p className="text-red-400 text-sm">{error}</p>
              </div>
            )}

            {success && (
              <div className="bg-green-900/20 border border-green-500 rounded-lg p-3">
                <p className="text-green-400 text-sm">
                  Profile updated successfully!
                </p>
              </div>
            )}

            <Button type="submit" disabled={isLoading} className="w-full">
              {isLoading ? 'Saving...' : 'Save Changes'}
            </Button>
          </form>
        </div>
      </div>
    </div>
  )
}
