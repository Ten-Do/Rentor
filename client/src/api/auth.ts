import { $api } from './$api'

export type SendOtpResponse = unknown

export interface VerifyOtpResponse {
  access_token: string
  user: {
    id: number
    email: string
    phone: string | null
    created_at: string
  }
}

export async function sendOtpAction(email: string): Promise<SendOtpResponse> {
  return $api.post<SendOtpResponse>('/auth/send-otp', { email })
}

export async function verifyOtpAction(
  email: string,
  otp_code: string
): Promise<VerifyOtpResponse> {
  return $api.post<VerifyOtpResponse>('/auth/verify-otp', {
    email,
    otp_code,
  })
}

export async function logoutAction(): Promise<{ message: string }> {
  return $api.post<{ message: string }>('/auth/logout')
}
