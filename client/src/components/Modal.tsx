import { XMarkIcon } from '@heroicons/react/24/outline'
import { type ReactNode, useEffect, useRef } from 'react'
import { useModal, type ModalName } from '../contexts/ModalContext'

interface ModalProps {
  name: ModalName
  children: ReactNode
  className?: string
}

export const Modal = ({ name, children, className = '' }: ModalProps) => {
  const modalRef = useRef<HTMLDivElement>(null)
  const overlayRef = useRef<HTMLDivElement>(null)

  const { close } = useModal()

  const closeModal = () => {
    close(name)
  }

  useEffect(() => {
    const handleEscape = (e: KeyboardEvent) => {
      if (e.key === 'Escape') closeModal()
    }

    document.addEventListener('keydown', handleEscape)
    document.body.style.overflow = 'hidden'

    return () => {
      document.removeEventListener('keydown', handleEscape)
      document.body.style.overflow = ''
    }
  }, [])

  const handleOverlayClick = (e: React.MouseEvent<HTMLDivElement>) => {
    if (e.target === overlayRef.current) closeModal()
  }

  return (
    <div
      ref={overlayRef}
      className="fixed inset-0 z-50 flex items-center justify-center bg-black/60"
      onClick={handleOverlayClick}
    >
      <div
        ref={modalRef}
        className={`relative bg-gray-800 rounded-lg shadow-xl max-w-lg w-full mx-4 max-h-[90vh] overflow-y-auto ${className}`}
        onClick={(e) => e.stopPropagation()}
      >
        <button
          onClick={closeModal}
          className="absolute top-4 right-4 text-gray-400 hover:text-gray-300 transition-colors duration-200"
          aria-label="Закрыть"
        >
          <XMarkIcon className="w-6 h-6" />
        </button>
        {children}
      </div>
    </div>
  )
}
