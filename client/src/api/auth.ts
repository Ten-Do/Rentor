import { $api } from './$api'

export type SendOtpResponse = unknown
export type VerifyOtpResponse = unknown

export async function sendOtpAction(email: string): Promise<SendOtpResponse> {
  return $api.post<SendOtpResponse>('/v1/auth/send-otp', { email })
}

export async function verifyOtpAction(
  email: string,
  otp_code: string
): Promise<VerifyOtpResponse> {
  return $api.post<VerifyOtpResponse>('/v1/auth/verify-otp', {
    email,
    otp_code,
  })
}
