import { useFormStatus } from 'react-dom'
import { Button } from './Button'

export function SubmitButton({
    className,
    idleText,
    pendingText,
}: {
    className?: string
    idleText: string
    pendingText: string
}) {
    const { pending } = useFormStatus()
    return (
        <Button type="submit" disabled={pending} className={className}>
            {pending ? pendingText : idleText}
        </Button>
    )
}
