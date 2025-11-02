import { useNavigate } from 'react-router-dom'

export const Home = () => {
  const navigate = useNavigate()

  return (
    <>
      <div>Главная</div>
      <button onClick={() => navigate('/profile')}>Profile</button>
      <button onClick={() => navigate('/advertisement/1')}>
        Advertisement
      </button>
    </>
  )
}
