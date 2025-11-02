import { createBrowserRouter } from 'react-router-dom';
import { Home } from './pages/Home';
import { Advertisement } from './pages/Advertisement';
import { Profile } from './pages/Profile';

export const router = createBrowserRouter([
  {
    path: '/',
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
]);

