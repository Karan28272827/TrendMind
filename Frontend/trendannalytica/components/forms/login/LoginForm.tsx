"use client";

import React, { useState } from "react";
import { useRouter } from "next/navigation";
import { LoginFormData } from "@/types/formTypes";
import { loginUser } from "@/services/authService";
import api from "@/constants/api";
import axios from "axios";
import "./LoginForm.css"

// import "./SignupForm.css"; 

const LoginForm = () => {
    const router = useRouter();
    const [error, setError] = useState("");
    const [formData, setFormData] = useState<LoginFormData>({
        email: "",
        password: "",
    });

    // Handle input change
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleGoogleSignIn = () => {
        window.location.href = "http://localhost:8080/auth/google";
    }


    // Handle form submission
    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setError(""); // Clear previous error

        try {
            console.log("Login Attempt:", {
                backendUrl: api.defaults.baseURL,
                endpoint: "/login",
                email: formData.email,
            });

            const response = await loginUser(formData);
            console.log("Response Data:", response.data);

            if (response.status === 200) {
                // If the response is just a string (the token), store it directly
                const token = response.data.JWT || response.data;

                localStorage.setItem("token", token);

                // Since the backend doesn’t send user details, you can manually store them from the form
                const user = {
                    email: formData.email,
                    password: formData.password || "", // assuming name is part of the formData
                };
                localStorage.setItem("user", JSON.stringify(user));

                alert("Login successful!");
                router.push("/trends");
            }
        } catch (error: unknown) {
            if (axios.isAxiosError(error)) {
                console.error("Axios Login Error Details:", {
                    fullError: error,
                    errorName: error.name,
                    errorMessage: error.message,
                    errorCode: error.code,
                    networkError: error.code === "ERR_NETWORK",
                    responseData: error.response?.data,
                    responseStatus: error.response?.status,
                    requestConfig: error.config,
                });

                if (error.code === "ERR_NETWORK") {
                    setError("Network Error: Cannot connect to server. Check your connection.");
                } else if (error.response) {
                    setError(error.response.data.message || "Login failed. Please try again.");
                } else if (error.request) {
                    setError("No response from server. Please check the backend.");
                } else {
                    setError("An unexpected error occurred during login.");
                }
            } else if (error instanceof Error) {
                console.error("General Error:", error.message);
                setError("An unknown error occurred. Please try again.");
            } else {
                console.error("Unknown error object:", error);
                setError("An unknown error occurred. Please try again.");
            }
        }
    };


    return (
        <div className="login-container">
            <div className="login-card">
                <div className="top-right-links">
                    <span>New User? <a href="/signup">Sign Up</a></span>
                </div>
                <h1 className="logo">TrendMind</h1>
                <h2>Welcome Back!</h2>
                <p>Login to continue</p>

                {error && <p className="error-message">{error}</p>}

                <form className="login-form" onSubmit={handleSubmit}>
                    <div className="input-group">
                        <input
                            type="email"
                            name="email"
                            placeholder="username11@gmail.com"
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
                    <button type="submit" className="login-button">LOGIN</button>
                    <div className="footer-links">
                        <a href="/forgot-password">Forgot Password?</a>
                    </div>
                </form>

                <div className="login-social">
                    <p>Login with</p>
                    <div className="social-icons" style={{ display: "flex", gap: "1rem", justifyContent: "center" }}>
                        <i
                            className="mdi mdi-google"
                            style={{ fontSize: "30px", color: "red", cursor: "pointer" }}
                            title="Sign in with Google"
                            onClick={handleGoogleSignIn}
                        />
                        <i
                            className="mdi mdi-github"
                            style={{ fontSize: "30px", color: "black", cursor: "pointer" }}
                            title="Sign in with Facebook"
                        />
                        <i
                            className="mdi mdi-twitter"
                            style={{ fontSize: "30px", color: "skyblue", cursor: "pointer" }}
                            title="Sign in with Twitter"
                        />
                    </div>
                </div>

            </div>
        </div>
    );
}
export default LoginForm;