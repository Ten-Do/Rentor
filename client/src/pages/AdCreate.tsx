import { useNavigate } from 'react-router-dom'
import { AdFormFields } from '../components/AdFormFields'
import { SubmitButton } from '../components/SubmitButton'
import { createAdvertisement, uploadAdvertisementImages } from '../api'

export const AdCreate = () => {
    const navigate = useNavigate()

    const onSubmit: React.FormEventHandler<HTMLFormElement> = async (e) => {
        e.preventDefault()
        const form = e.currentTarget
        const data = new FormData(form)
        const price = Number(data.get('price') || 0)

        const created = await createAdvertisement({
            title: String(data.get('title') || ''),
            description: String(data.get('description') || ''),
            price,
            type: String(data.get('type') || 'apartment') as any,
            rooms: String(data.get('rooms') || 'studio') as any,
            city: String(data.get('city') || ''),
            address: String(data.get('address') || ''),
        })
        if (created?.id) {
            const files = data.getAll('files').filter((f): f is File => f instanceof File && f.size > 0)
            if (files.length > 0) {
                await uploadAdvertisementImages(created.id, files)
            }
        }
        navigate('/my/ads')
    }

    return (
        <div className="max-w-2xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
            <h1 className="text-2xl font-semibold text-white mb-6">Add advertisement</h1>
            <form onSubmit={onSubmit} className="space-y-4">
                <AdFormFields initData={null} />

                <div className="pt-2">
                    <SubmitButton idleText="Create" pendingText="Creating..." />
                </div>
            </form>
        </div>
    )
}


