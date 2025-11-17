import { createBrowserRouter } from 'react-router-dom'
import { Layout } from './features/Layout'
import { ArticleLayout } from './features/Article'
import { Home } from './pages/Home'
import { Advertisement } from './pages/Advertisement'
import { Profile } from './pages/Profile'
import { PrivacyPolicy } from './pages/PrivacyPolicy'
import { TermsOfService } from './pages/TermsOfService'
import { CookiePolicy } from './pages/CookiePolicy'
import { homeLoader, advertisementLoader, authLoader, myAdsLoader } from './dataLoaders'
import { MyAds } from './pages/MyAds'
import { AdCreate } from './pages/AdCreate'
import { AdEdit } from './pages/AdEdit'

export const router = createBrowserRouter([
  {
    id: 'root',
    path: '/',
    element: <Layout />,
    loader: authLoader,
    children: [
      {
        index: true,
        element: <Home />,
        loader: homeLoader,
      },
      {
        path: '/advertisement/:id',
        element: <Advertisement />,
        loader: advertisementLoader,
      },
      {
        path: '/profile',
        element: <Profile />,
      },
      {
        path: '/my/ads',
        element: <MyAds />,
        loader: myAdsLoader,
      },
      {
        path: '/my/ads/new',
        element: <AdCreate />,
      },
      {
        path: '/my/ads/:id/edit',
        element: <AdEdit />,
        loader: advertisementLoader,
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
