"use client";

import React, { useState } from "react";
import { useRouter } from "next/navigation";
import { SignupFormData } from "@/types/formTypes";
import { isPasswordValid } from "@/utils/validators";
import { signupUser, verifyOtp } from "@/services/authService";
// import { setLocalStorage } from "@/utils/localStorage";
import "./SignupForm.css"; 

const SignupForm = () => {
  const [formData, setFormData] = useState<SignupFormData>({
    name: "",
    email: "",
    password: "",
    retypepassword: "",
  });

  const [otp, setOtp] = useState("");
  const [showOtpField, setShowOtpField] = useState(false);
  const [error, setError] = useState("");
  const router = useRouter();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!isPasswordValid(formData.password)) {
      alert(
        "Password must be at least 8 characters long, include an uppercase letter, a lowercase letter, a number, and a special character."
      );
      return;
    }

    if (formData.password !== formData.retypepassword) {
      alert("Passwords do not match!");
      return;
    }

    try {
      const response = await signupUser(formData);

      if (response.status === 200) {
        alert("OTP sent to your email. Please verify.");
        setShowOtpField(true);
      }
    } catch (err: unknown) {
        if (err && typeof err === "object" && "response" in err) {
          const errorObj = err as { response?: { data?: { msg?: string } } };
          setError(errorObj.response?.data?.msg || "Signup failed.");
        } else {
          setError("Signup failed.");
        }
      }      
  };

  const handleVerifyOtp = async () => {
    try {
      const response = await verifyOtp(
        formData.email,
        otp,
        formData.name,
        formData.password
      );

      if (response.status === 201) {
        alert("Email verified successfully! Account created.");
        const firstLetter = formData.name.charAt(0).toUpperCase();
        const profileImageUrl =
          response.data.user.profileImage ||
          `https://ui-avatars.com/api/?name=${firstLetter}&background=random&color=fff&size=128`;

        const userData = {
          name: response.data.user.name,
          profileImage: profileImageUrl,
        };

        // setLocalStorage("user", userData);
        // setLocalStorage("token", response.data.token);

        window.dispatchEvent(new Event("storage"));
        router.push("/bookings");
      }
    } catch (err: unknown) {
        if (err && typeof err === "object" && "response" in err) {
          const errorObj = err as { response?: { data?: { msg?: string } } };
          setError(errorObj.response?.data?.msg || "Invalid OTP. Try again.");
        } else {
          setError("Invalid OTP. Try again.");
        }
      }
      
  };
  
  return (
    <div className="signup-container">
      <div className="signup-card">
        <h1 className="logo">TrendMind</h1>
        <div className="top-right-links">
          <span>
            Already have an account? <a href="/login">Login</a>
          </span>
        </div>
        <p>Register to continue</p>

        <form className="signup-form" onSubmit={handleSubmit}>
          <div className="input-group">
            <input
              type="text"
              name="name"
              placeholder="Full Name"
              value={formData.name}
              onChange={handleChange}
              required
            />
          </div>
          <div className="input-group">
            <input
              type="email"
              name="email"
              placeholder="Email Id"
              value={formData.email}
              onChange={handleChange}
              required
            />
          </div>
          <div className="input-group">
            <input
              type="password"
              name="password"
              placeholder="Enter Password"
              value={formData.password}
              onChange={handleChange}
              required
            />
          </div>
          <div className="input-group">
            <input
              type="password"
              name="retypepassword"
              placeholder="Confirm Password"
              value={formData.retypepassword}
              onChange={handleChange}
              required
            />
          </div>

          {!showOtpField ? (
            <button type="submit" className="signup-button">
              SIGN UP
            </button>
          ) : (
            <>
              <div className="input-group">
                <input
                  type="text"
                  placeholder="Enter OTP"
                  value={otp}
                  onChange={(e) => setOtp(e.target.value)}
                  required
                />
              </div>
              <button
                type="button"
                className="signup-button"
                onClick={handleVerifyOtp}
              >
                Verify OTP
              </button>
            </>
          )}

          {error && <p className="error-text">{error}</p>}
        </form>
      </div>
    </div>
  );
};

export default SignupForm;
