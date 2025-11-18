import { useState } from 'react'
import { Swiper, SwiperSlide } from 'swiper/react'
import { Navigation, Thumbs } from 'swiper/modules'
import type { Swiper as SwiperType } from 'swiper'
import { ChevronLeftIcon, ChevronRightIcon } from '@heroicons/react/24/outline'
import { Button } from '../../components/Button'
import { Image } from '../../components/Image'
import { API_BASE_URL } from '../../utils/constants'
import 'swiper/swiper.css'

export interface ImageSliderProps {
  images: string[]
}

const btnNextClass = 'gallery-next'
const btnPrevClass = 'gallery-prev'

const btnNextSelector = `.${btnNextClass}`
const btnPrevSelector = `.${btnPrevClass}`

export const ImageSlider = ({ images }: ImageSliderProps) => {
  const [thumbsSwiper, setThumbsSwiper] = useState<SwiperType | null>(null)

  if (!images || images.length === 0) {
    return (
      <div className="w-full aspect-video bg-gray-800 rounded-lg flex items-center justify-center">
        <span className="text-gray-500">No images available</span>
      </div>
    )
  }

  return (
    <div className="w-full space-y-4">
      <div className="relative w-full aspect-video rounded-lg overflow-hidden">
        <Swiper
          spaceBetween={8}
          modules={[Thumbs]}
          thumbs={{ swiper: thumbsSwiper }}
          className="h-full w-full"
        >
          {images.map((image, index) => (
            <SwiperSlide key={index}>
              <Image
                src={image}
                alt={`Property image ${index + 1}`}
                className="w-full h-full object-cover"
              />
            </SwiperSlide>
          ))}
        </Swiper>
      </div>

      {images.length > 1 && (
        <div className="w-full relative">
          <Swiper
            modules={[Thumbs, Navigation]}
            onSwiper={setThumbsSwiper}
            spaceBetween={8}
            slidesPerView={3}
            navigation={{
              nextEl: btnNextSelector,
              prevEl: btnPrevSelector,
            }}
            freeMode
            watchSlidesProgress
            className="h-24"
            breakpoints={{
              480: {
                slidesPerView: 4,
              },
              640: {
                slidesPerView: 5,
              },
              1024: {
                slidesPerView: 6,
              },
            }}
          >
            {images.map((image, index) => {
              const imageUrl =
                image?.startsWith('http://') || image?.startsWith('https://')
                  ? image
                  : `${API_BASE_URL}${image?.startsWith('/') ? image : `/${image}`}`
              return (
                <SwiperSlide key={index} className="cursor-pointer">
                  <div className="h-full rounded overflow-hidden opacity-50 transition-opacity hover:opacity-75">
                    <div
                      role="img"
                      aria-label={`Thumbnail ${index + 1}`}
                      className="w-full h-full bg-center bg-cover bg-gray-600"
                      style={{ backgroundImage: `url('${imageUrl}')` }}
                    />
                  </div>
                </SwiperSlide>
              )
            })}
            <Button
              variant="ghost"
              colorScheme="neutral"
              className={`${btnPrevClass} absolute -left-4 top-1/2 -translate-y-1/2 z-10 rounded-full p-3`}
              aria-label="Previous image"
            >
              <ChevronLeftIcon className="w-6 h-6" />
            </Button>
            <Button
              variant="ghost"
              colorScheme="neutral"
              className={`${btnNextClass} absolute -right-4 top-1/2 -translate-y-1/2 z-10 rounded-full p-3`}
              aria-label="Next image"
            >
              <ChevronRightIcon className="w-6 h-6" />
            </Button>
          </Swiper>
        </div>
      )}
    </div>
  )
}
