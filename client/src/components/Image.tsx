import { PhotoIcon } from '@heroicons/react/24/outline'
import { useState } from 'react'

export interface ImageProps
  extends Omit<React.ImgHTMLAttributes<HTMLImageElement>, 'onError' | 'src'> {
  src: string | null
  alt: string
}

export const Image = ({ src, className = '', ...props }: ImageProps) => {
  const [imgSrc, setImgSrc] = useState<string | null>(src)

  const handleError = () => {
    if (imgSrc) {
      setImgSrc(null)
    }
  }

  if (!imgSrc) {
    return (
      <div
        className={`flex items-center justify-center bg-gray-800 text-gray-500 ${className}`}
      >
        <PhotoIcon className="w-1/3 aspect-square text-gray-500" />
      </div>
    )
  }

  return (
    <img src={imgSrc} className={className} onError={handleError} {...props} />
  )
}
