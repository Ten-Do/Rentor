import { Link } from '../../components/Link'

export const Header = () => {
  return (
    <header className="sticky top-0 z-50 bg-gray-900 border-b border-gray-800">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between h-16">
          <Link
            to="/"
            className="text-2xl font-bold text-indigo-400 hover:text-indigo-300 transition-colors"
          >
            Rentor
          </Link>

          <Link to="/profile" size="md" colorScheme="primary">
            Profile
          </Link>
        </div>
      </div>
    </header>
  )
}
