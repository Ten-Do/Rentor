import { useActionState, useState } from 'react'
import { Modal } from '../Modal'
import { Button } from '../Button'
import { Input } from '../Input'
import { SubmitButton } from '../SubmitButton'
import { useModal } from '../../contexts/ModalContext'
import { sendOtpAction, verifyOtpAction } from '../../api/auth'

type LoginStep = 'email' | 'otp'
type EmailActionState = { emailError?: string } | null
type OtpActionState = { otpError?: string } | null

export const LoginModal = () => {
    const { close } = useModal()
    const [step, setStep] = useState<LoginStep>('email')

    const [emailActionState, sendOtp] = useActionState<EmailActionState, FormData>(
        async (_prev, formData) => {
            const email = String(formData.get('email') || '')

            try {
                await sendOtpAction(email)
                setStep('otp')
                return null
            } catch (e) {
                const message = e instanceof Error ? e.message : 'Failed to send OTP code'
                return { emailError: message }
            }
        },
        null
    )

    const [otpActionState, verifyOtp] = useActionState<OtpActionState, FormData>(
        async (_prev, formData) => {
            const email = String(formData.get('email') || '')
            const otp_code = String(formData.get('otp_code') || '')

            try {
                await verifyOtpAction(email, otp_code)
                close('login')
                return null
            } catch (e) {
                const message = e instanceof Error ? e.message : 'Invalid OTP code'
                return { otpError: message }
            }
        },
        null
    )

    const handleBackToEmail = () => {
        setStep('email')
    }

    return (
        <Modal name="login" className="p-6">
            <div className="space-y-6">
                <div>
                    <h2 className="text-2xl font-bold text-white mb-2">Login</h2>
                    <p className="text-gray-400 text-sm">
                        {step === 'email'
                            ? 'Enter your email to receive an OTP code'
                            : 'Enter the OTP code sent to your email'}
                    </p>
                </div>

                <form action={step === 'email' ? sendOtp : verifyOtp} className="space-y-4">
                    <Input
                        label="Email"
                        id="email"
                        name="email"
                        type="email"
                        required
                        readOnly={step !== 'email'}
                        error={step === 'email' ? emailActionState?.emailError ?? null : null}
                        autoComplete="email"
                        autoFocus={step === 'email'}
                        disabled={step !== 'email'}
                    />
                    {step === 'otp' && (
                        <Input
                            label="OTP Code"
                            id="otp_code"
                            name="otp_code"
                            type="text"
                            required
                            inputMode="numeric"
                            error={otpActionState?.otpError ?? null}
                            placeholder="000000"
                            maxLength={6}
                            pattern="[0-9]{6}"
                            autoComplete="one-time-code"
                            autoFocus={step === 'otp'}
                        />
                    )}

                    <div className="flex gap-4">
                        {step === 'otp' ? (
                            <>
                                <Button
                                    type="button"
                                    variant="outline"
                                    onClick={handleBackToEmail}
                                    className="flex-1"
                                >
                                    Back
                                </Button>
                                <SubmitButton className="flex-1" idleText="Verify Code" pendingText="Verifying..." />
                            </>
                        ) : (
                            <SubmitButton className="w-full" idleText="Send OTP Code" pendingText="Sending..." />
                        )}
                    </div>
                </form>
            </div>
        </Modal>
    )
}
