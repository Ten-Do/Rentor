import {
  type ReactNode,
  createContext,
  useContext,
  useState,
  useCallback,
} from 'react'
import { LandlordModal } from '../components/modals/LandlordModal'
import { LoginModal } from '../components/modals/LoginModal'

export type ModalName = 'login' | 'landlord'

interface ModalContextType {
  open: (name: ModalName, props?: any) => void
  close: (name: ModalName) => void
}

const initModals: Record<ModalName, any> = {
  login: null,
  landlord: null,
}

const ModalContext = createContext<ModalContextType>({
  open: () => { },
  close: () => { },
})

export const ModalProvider = ({ children }: { children: ReactNode }) => {
  const [modals, setModals] = useState<Record<ModalName, any>>(initModals)

  const open = useCallback((name: ModalName, props?: any) => {
    setModals((prev) => ({
      ...prev,
      [name]: props || true,
    }))
  }, [])

  const close = useCallback((name: ModalName) => {
    setModals((prev) => {
      const newModals = { ...prev }
      newModals[name] = null
      return newModals
    })
  }, [])

  return (
    <ModalContext.Provider value={{ open, close }}>
      {children}
      {modals.login && <LoginModal />}
      {modals.landlord && <LandlordModal {...modals.landlord} />}
    </ModalContext.Provider>
  )
}

export const useModal = () => {
  const context = useContext(ModalContext)
  if (!context) {
    throw new Error('useModal must be used within ModalProvider')
  }
  return context
}
