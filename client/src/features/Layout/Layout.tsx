import { Outlet } from 'react-router-dom'
import { Header } from './Header'
import { Footer } from './Footer'
import { ModalProvider } from '../../contexts'

export const Layout = () => {
  return (
    <ModalProvider>
      <div className="bg-gray-950 text-white flex flex-col min-h-screen">
        <Header />
        <main className="flex-1 w-full max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
          <Outlet />
        </main>
        <Footer />
      </div>
    </ModalProvider>
  )
}
