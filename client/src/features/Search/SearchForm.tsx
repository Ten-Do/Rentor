import { useSearchParams, useNavigate } from 'react-router-dom'
import { Input, Select, Button } from '../../components'

export const AD_TYPE_OPTIONS = [
    { value: '', label: 'All types' },
    { value: 'apartment', label: 'Apartment' },
    { value: 'house', label: 'House' },
    { value: 'room', label: 'Room' },
] as const

export const ROOMS_OPTIONS = [
    { value: '', label: 'Any rooms' },
    { value: 'studio', label: 'Studio' },
    { value: '1', label: '1 room' },
    { value: '2', label: '2 rooms' },
    { value: '3', label: '3 rooms' },
    { value: '4', label: '4 rooms' },
    { value: '5', label: '5 rooms' },
    { value: '6+', label: '6+ rooms' },
] as const

export const SearchForm = () => {
    const [searchParams, setSearchParams] = useSearchParams()
    const navigate = useNavigate()

    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        const formData = new FormData(e.currentTarget)
        const params = new URLSearchParams()

        const keywords = formData.get('keywords')?.toString().trim()
        const city = formData.get('city')?.toString().trim()
        const type = formData.get('type')?.toString()
        const rooms = formData.get('rooms')?.toString()
        const minPrice = formData.get('minPrice')?.toString().trim()
        const maxPrice = formData.get('maxPrice')?.toString().trim()

        if (keywords) params.set('keywords', keywords)
        if (city) params.set('city', city)
        if (type) params.set('type', type)
        if (rooms) params.set('rooms', rooms)
        if (minPrice) params.set('minPrice', minPrice)
        if (maxPrice) params.set('maxPrice', maxPrice)

        setSearchParams(params)
    }

    return (
        <form onSubmit={handleSubmit} className="mb-8 bg-gray-800 p-6 rounded-lg border border-gray-700">
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                <Input
                    name="keywords"
                    placeholder="Search by title..."
                    defaultValue={searchParams.get('keywords') || ''}
                    label="Keywords"
                />

                <Input
                    name="city"
                    placeholder="City..."
                    defaultValue={searchParams.get('city') || ''}
                    label="City"
                />

                <div>
                    <label htmlFor="type" className="block text-sm font-medium text-gray-300 mb-2">
                        Type
                    </label>
                    <Select
                        id="type"
                        name="type"
                        defaultValue={searchParams.get('type') || ''}
                    >
                        {AD_TYPE_OPTIONS.map(({ value, label }) => (
                            <option key={value} value={value}>
                                {label}
                            </option>
                        ))}
                    </Select>
                </div>

                <div>
                    <label htmlFor="rooms" className="block text-sm font-medium text-gray-300 mb-2">
                        Rooms
                    </label>
                    <Select
                        id="rooms"
                        name="rooms"
                        defaultValue={searchParams.get('rooms') || ''}
                    >
                        {ROOMS_OPTIONS.map(({ value, label }) => (
                            <option key={value} value={value}>
                                {label}
                            </option>
                        ))}
                    </Select>
                </div>

                <Input
                    type="number"
                    name="minPrice"
                    placeholder="Min price..."
                    defaultValue={searchParams.get('minPrice') || ''}
                    label="Min price"
                    min="0"
                />

                <Input
                    type="number"
                    name="maxPrice"
                    placeholder="Max price..."
                    defaultValue={searchParams.get('maxPrice') || ''}
                    label="Max price"
                    min="0"
                />
            </div>

            <div className="flex gap-3 mt-4">
                <Button type="submit" variant="solid" colorScheme="primary">
                    Search
                </Button>
                <Button
                    type="reset"
                    variant="outline"
                    colorScheme="neutral"
                    onClick={() => {
                        navigate('/')
                    }}
                >
                    Clear
                </Button>
            </div>
        </form>
    )
}

