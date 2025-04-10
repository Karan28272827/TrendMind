const BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;

const API_ENDPOINTS = {
  SIGNUP: `${BASE_URL}/register`,
  LOGIN: `${BASE_URL}/login`,
  VERIFY_OTP: `${BASE_URL}/verify-otp`,
};

export default API_ENDPOINTS;
