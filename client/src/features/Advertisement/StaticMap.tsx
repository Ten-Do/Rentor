export interface StaticMapProps {
  latitude: number | null
  longitude: number | null
}

export const StaticMap = ({ latitude, longitude }: StaticMapProps) => {
  if (!latitude || !longitude) {
    return (
      <div className="w-full aspect-video bg-gray-800 rounded-lg flex items-center justify-center">
        <span className="text-gray-500">Location not available</span>
      </div>
    )
  }

  const mapUrl = `https://static-maps.yandex.ru/1.x/?ll=${longitude},${latitude}&size=600,400&z=15&l=map&pt=${longitude},${latitude},pm2rdm`

  return (
    <div className="w-full aspect-video rounded-lg overflow-hidden bg-gray-800">
      <img
        src={mapUrl}
        alt="Property location"
        className="w-full h-full object-cover"
      />
    </div>
  )
}
