"use client"

import axios from "axios";
import React from "react";

const TestProtectedRoute = () => {

    const testFunc = async () => {
        const res = await axios.get("http://localhost:8080/auth/profile")
        console.log("hello lol  " + res)
    }

    return(
        <div>
            <h1>Test route</h1>
            <button onClick={testFunc}>Test API</button>
        </div>
    )
}

export default TestProtectedRoute   