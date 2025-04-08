import axios from 'axios';
import API_ENDPOINTS from '../constants/apiEndpoints';
import { SignupFormData } from "@/types/formTypes";

export const signupUser = async (data: SignupFormData) => {
    return axios.post("http://localhost:8080/auth/register", data);
  };

export const verifyOtp = async (email: string, otp: string, name: string, password: string) => {
  return axios.post(API_ENDPOINTS.VERIFY_OTP, {
    email, otp, name, password
  });
};
