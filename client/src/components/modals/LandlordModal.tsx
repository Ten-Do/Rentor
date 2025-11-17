import { Modal } from '../Modal'
import { Button } from '../Button'
import { ClipboardDocumentIcon } from '@heroicons/react/24/outline'

export interface LandlordModalProps {
  landlord_name: string
  landlord_email: string
  landlord_phone: string | null
}

interface CopyableFieldProps {
  label: string
  value: string
}

const CopyableField = ({ label, value }: CopyableFieldProps) => {
  const copyToClipboard = async (text: string) => {
    try {
      await navigator.clipboard.writeText(text)
    } catch (err) {
      console.error('Failed to copy:', err)
    }
  }

  return (
    <div>
      <span className="text-gray-400 text-sm block mb-2">{label}</span>
      <div
        onClick={() => copyToClipboard(value)}
        className="flex items-center justify-between gap-3 p-3 rounded-lg bg-gray-800 border border-gray-700 cursor-pointer hover:bg-gray-700 hover:border-indigo-500 transition-colors duration-200 group"
      >
        <p className="text-white text-lg">{value}</p>
        <ClipboardDocumentIcon className="w-5 h-5 text-gray-400 group-hover:text-indigo-400 transition-colors duration-200 flex-shrink-0" />
      </div>
    </div>
  )
}

export const LandlordModal = ({
  landlord_name,
  landlord_email,
  landlord_phone,
}: LandlordModalProps) => {
  return (
    <Modal name="landlord" className="p-6">
      <div className="space-y-6">
        <div>
          <h2 className="text-2xl font-bold text-white mb-6">
            Contact Landlord
          </h2>
        </div>

        <div className="space-y-4">
          <CopyableField label="Name" value={landlord_name} />
          <CopyableField label="Email" value={landlord_email} />
          {landlord_phone && (
            <CopyableField label="Phone" value={landlord_phone} />
          )}
        </div>

        <div className="flex flex-wrap gap-4 pt-4 border-t border-gray-700">
          <Button
            as="a"
            href={`mailto:${landlord_email}`}
            size="lg"
            className="whitespace-nowrap flex-1"
          >
            Contact by Email
          </Button>
          {landlord_phone && (
            <Button
              as="a"
              href={`tel:${landlord_phone}`}
              size="lg"
              variant="outline"
              className="whitespace-nowrap flex-1"
            >
              Call
            </Button>
          )}
        </div>
      </div>
    </Modal>
  )
}
