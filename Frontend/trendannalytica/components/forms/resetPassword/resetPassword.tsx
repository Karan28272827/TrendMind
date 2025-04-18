"use client"; // Mark as a Client Component

import { useState } from "react";
import { useSearchParams } from "next/navigation";
import "./ResetPassword.css";

export default function ResetPasswordPage() {
  const searchParams = useSearchParams();
  const token = searchParams.get("token");
  const [newPassword, setNewPassword] = useState("");
  const [message, setMessage] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);

    try {
      const response = await fetch(
        "http://localhost:8080/auth/reset-password",
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ token, newPassword }),
        },
      );

      const data = await response.json();
      if (response.ok) {
        setMessage("Password reset successfully!");
      } else {
        setMessage(data.message || "Error resetting password");
      }
    } catch (error) {
      setMessage("Network error. Please try again.");
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="container">
      <h1 className="title">Reset Your Password</h1>
      {message && (
        <div
          className={`message ${message.includes("success") ? "success" : "error"}`}
        >
          {message}
        </div>
      )}
      <form onSubmit={handleSubmit} className="form">
        <div className="formGroup">
          <label htmlFor="password" className="password">
            New Password
          </label>
          <input
            id="password"
            type="password"
            value={newPassword}
            onChange={(e) => setNewPassword(e.target.value)}
            className="input"
            required
            minLength={8}
          />
        </div>
        <button type="submit" className="button" disabled={isLoading}>
          {isLoading ? "Processing..." : "Reset Password"}
        </button>
      </form>
    </div>
  );
}
