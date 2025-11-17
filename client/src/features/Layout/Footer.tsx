import { Link } from '../../components/Link'

export const Footer = () => {
  return (
    <footer className="bg-gray-900 border-t border-gray-800">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="flex flex-col md:flex-row justify-between items-center gap-4">
          <div className="text-gray-400 text-sm">
            © 2024 Rentor. All rights reserved.
          </div>
          <div className="flex flex-wrap gap-4 justify-center text-sm">
            <Link to="/articles/privacy-policy" colorScheme="neutral" size="sm">
              Privacy Policy
            </Link>
            <Link
              to="/articles/terms-of-service"
              colorScheme="neutral"
              size="sm"
            >
              Terms of Service
            </Link>
            <Link to="/articles/cookie-policy" colorScheme="neutral" size="sm">
              Cookie Policy
            </Link>
            <a
              href="https://t.me/Yura_Rrrr"
              target="_blank"
              className="text-gray-400 hover:text-gray-300"
            >
              Contact Us
            </a>
          </div>
        </div>
      </div>
    </footer>
  )
}
