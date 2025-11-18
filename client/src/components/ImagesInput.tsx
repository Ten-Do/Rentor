import { useMemo, useRef, useState } from 'react'
import { Button } from './Button'
import { Image } from './Image'

export interface ImagesInputProps {
    name?: string
    label?: string
}

export const ImagesInput = ({
    name = 'images',
    label = 'Images',
}: ImagesInputProps) => {
    const inputRef = useRef<HTMLInputElement | null>(null)
    const [files, setFiles] = useState<File[]>([])

    const previews = useMemo(() => files.map((f) => ({ file: f, url: URL.createObjectURL(f) })), [files])

    const handleChange: React.ChangeEventHandler<HTMLInputElement> = (e) => {
        const selected = Array.from(e.currentTarget.files || [])
        const next = [...files, ...selected]
        setFiles(next)
    }

    const removeAt = (idx: number) => {
        const next = files.filter((_, i) => i !== idx)
        setFiles(next)
    }

    const clear = () => {
        setFiles([])
        if (inputRef.current) inputRef.current.value = ''
    }

    return (
        <div className="space-y-2">
            <label className="block text-sm font-medium text-gray-300">{label}</label>
            <input
                ref={inputRef}
                type="file"
                name={name}
                accept="image/*"
                multiple
                className="block w-full text-sm text-gray-300 file:mr-4 file:py-2 file:px-4 file:rounded-md file:border-0 file:text-sm file:font-medium file:bg-indigo-600 file:text-white hover:file:bg-indigo-500"
                onChange={handleChange}
            />

            {previews.length > 0 && (
                <div className="space-y-2">
                    <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-3">
                        {previews.map((p, idx) => (
                            <div key={idx} className="relative group border border-gray-800 rounded-lg overflow-hidden bg-gray-900">
                                <Image src={p.url} alt={p.file.name} className="w-full aspect-square object-cover" />
                                <div className="absolute top-2 right-2">
                                    <Button type="button" size="sm" variant="ghost" colorScheme="neutral" onClick={() => removeAt(idx)}>
                                        Remove
                                    </Button>
                                </div>
                            </div>
                        ))}
                    </div>
                    <div className="flex gap-2">
                        <Button type="button" variant="outline" colorScheme="neutral" size="sm" onClick={clear}>
                            Clear all
                        </Button>
                    </div>
                </div>
            )}
        </div>
    )
}


