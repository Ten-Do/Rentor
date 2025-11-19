import { PhotoIcon } from '@heroicons/react/24/outline'
import { useState, useEffect } from 'react'
import { API_BASE_URL } from '../utils/constants'

export interface ImageProps
  extends Omit<React.ImgHTMLAttributes<HTMLImageElement>, 'onError' | 'src'> {
  src: string | null
  alt: string
}

const getImageUrl = (src: string | null): string | null => {
  if (!src) return null
  if (src.startsWith('http://') || src.startsWith('https://') || src.startsWith('blob:')) return src
  const cleanSrc = src.startsWith('/') ? src : `/${src}`
  return `${API_BASE_URL}${cleanSrc}`
}

export const Image = ({ src, className = '', ...props }: ImageProps) => {
  const [imgSrc, setImgSrc] = useState<string | null>(getImageUrl(src))
  const [hasError, setHasError] = useState(false)

  useEffect(() => {
    setImgSrc(getImageUrl(src))
    setHasError(false)
  }, [src])

  const handleError = () => {
    setHasError(true)
    setImgSrc(null)
  }

  if (!imgSrc || hasError) {
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
