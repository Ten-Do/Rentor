import { useLoaderData, useNavigate } from 'react-router-dom'
import type { Advertisement } from '../types'
import { AdFormFields } from '../components/AdFormFields'
import { SubmitButton } from '../components/SubmitButton'
import { Button } from '../components/Button'
import { Image } from '../components/Image'
import { deleteAdvertisementImage, updateAdvertisement, uploadAdvertisementImages } from '../api'
import { useState } from 'react'

export const AdEdit = () => {
    const { advertisement } = useLoaderData() as { advertisement: Advertisement }
    const navigate = useNavigate()
    const [existingImages, setExistingImages] = useState<string[]>(
        advertisement.image_urls
    )


    const onSubmit: React.FormEventHandler<HTMLFormElement> = async (e) => {
        e.preventDefault()
        const form = e.currentTarget
        const data = new FormData(form)

        await updateAdvertisement(advertisement.id, {
            title: String(data.get('title') || ''),
            description: String(data.get('description') || ''),
            price: Number(data.get('price') || 0),
            type: String(data.get('type') || 'apartment') as any,
            rooms: String(data.get('rooms') || 'studio') as any,
            city: String(data.get('city') || ''),
            address: String(data.get('address') || ''),
            status: String(data.get('status') || 'active') as any,
        })
        const files = data.getAll('images') as File[]
        if (files.length > 0) {
            await uploadAdvertisementImages(advertisement.id, files)
        }
        navigate('/my/ads')
    }

    return (
        <div className="max-w-2xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
            <h1 className="text-2xl font-semibold text-white mb-6">Edit advertisement</h1>
            <form onSubmit={onSubmit} className="space-y-4">
                <AdFormFields initData={advertisement} />

                {existingImages.length > 0 && (
                    <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-3">
                        {existingImages.map((img) => (
                            <div key={img} className="relative border border-gray-800 rounded-lg overflow-hidden bg-gray-900">
                                <Image src={img} alt="ad image" className="w-full aspect-square object-cover" />
                                <Button
                                    className="absolute top-2 right-2"
                                    type="button"
                                    size="sm"
                                    variant="ghost"
                                    colorScheme="neutral"
                                    onClick={() => {
                                        const copy = [...existingImages]
                                        deleteAdvertisementImage(advertisement.id, img).catch(() => {
                                            setExistingImages(copy)
                                        })
                                        setExistingImages((prev) => prev.filter((x) => x !== img))
                                    }}
                                >
                                    Delete
                                </Button>
                            </div>
                        ))}
                    </div>
                )}

                <div className="pt-2">
                    <SubmitButton idleText="Save changes" pendingText="Saving..." />
                </div>
            </form>
        </div>
    )
}


