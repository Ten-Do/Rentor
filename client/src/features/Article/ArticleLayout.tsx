import { Outlet, useLocation } from 'react-router-dom'

const formatDate = () => {
  return new Date().toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  })
}

const articleTitles: Record<string, string> = {
  'privacy-policy': 'Privacy Policy',
  'terms-of-service': 'Terms of Service',
  'cookie-policy': 'Cookie Policy',
}

export const ArticleLayout = () => {
  const location = useLocation()
  const pathSegment = location.pathname.split('/').pop() || ''
  const title = articleTitles[pathSegment] || 'Article'

  return (
    <>
      <div className="flex justify-between items-center flex-wrap gap-4 mb-6">
        <h1 className="text-4xl font-bold text-indigo-400">{title}</h1>
        <p className="text-sm text-gray-400">Last updated: {formatDate()}</p>
      </div>
      <div className="prose prose-invert max-w-none space-y-4 text-gray-300">
        <Outlet />
      </div>
    </>
  )
}
