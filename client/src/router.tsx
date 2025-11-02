import { createBrowserRouter } from 'react-router-dom'
import { Layout } from './features/Layout'
import { ArticleLayout } from './features/Article'
import { Home } from './pages/Home'
import { Advertisement } from './pages/Advertisement'
import { Profile } from './pages/Profile'
import { PrivacyPolicy } from './pages/PrivacyPolicy'
import { TermsOfService } from './pages/TermsOfService'
import { CookiePolicy } from './pages/CookiePolicy'

export const router = createBrowserRouter([
  {
    path: '/',
    element: <Layout />,
    children: [
      {
        index: true,
        element: <Home />,
      },
      {
        path: '/advertisement/:id',
        element: <Advertisement />,
      },
      {
        path: '/profile',
        element: <Profile />,
      },
      {
        path: '/articles',
        element: <ArticleLayout />,
        children: [
          {
            path: 'privacy-policy',
            element: <PrivacyPolicy />,
          },
          {
            path: 'terms-of-service',
            element: <TermsOfService />,
          },
          {
            path: 'cookie-policy',
            element: <CookiePolicy />,
          },
        ],
      },
    ],
  },
])
