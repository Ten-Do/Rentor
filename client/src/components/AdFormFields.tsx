import type { Advertisement } from '../types'
import { Input } from './Input'
import { Textarea } from './Textarea'
import { Select } from './Select'
import { ImagesInput } from './ImagesInput'

export interface AdFormFieldsProps {
    initData: Partial<Advertisement> | null
}

export const AdFormFields = ({
    initData,
}: AdFormFieldsProps) => {
    return (
        <>
            <Input id="title" name="title" label="Title" required defaultValue={initData?.title} />
            <Textarea id="description" name="description" label="Description" required rows={6} defaultValue={initData?.description} />
            <Input id="price" name="price" label="Price" type="number" min={0} required defaultValue={initData?.price} />

            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
                <Select name="type" defaultValue={initData?.type || 'apartment'}>
                    <option value="apartment">Apartment</option>
                    <option value="house">House</option>
                    <option value="room">Room</option>
                </Select>
                <Select name="rooms" defaultValue={initData?.rooms || 'studio'}>
                    <option value="studio">Studio</option>
                    <option value="1">1 room</option>
                    <option value="2">2 rooms</option>
                    <option value="3">3 rooms</option>
                    <option value="4">4 rooms</option>
                    <option value="5">5 rooms</option>
                    <option value="6+">6+ rooms</option>
                </Select>
            </div>

            <Input id="city" name="city" label="City" required defaultValue={initData?.city} />
            <Input id="address" name="address" label="Address" required defaultValue={initData?.address} />

            {initData?.status && (
                <Select name="status" defaultValue={initData.status}>
                    <option value="active">Active</option>
                    <option value="paused">Paused</option>
                </Select>
            )}

            {/* New images uploader (used in both create and edit) */}
            <ImagesInput name="images" />
        </>
    )
}


