import axios from 'axios';
import API_ENDPOINTS from '../constants/apiEndpoints';
import { LoginFormData, SignupFormData } from "@/types/formTypes";

export const signupUser = async (data: SignupFormData) => {
    return axios.post(API_ENDPOINTS.SIGNUP, data);
  };

export const verifyOtp = async (email: string, otp: string, name: string, password: string) => {
  return axios.post(API_ENDPOINTS.VERIFY_OTP, {
    email, otp, name, password
  });
};

export const loginUser = async(data: LoginFormData) => {
  return axios.post(API_ENDPOINTS.LOGIN, data);
}