"use client"

import { useRouter } from "next/navigation";
import React, { useState } from "react";
import axios from "axios";


const ChatInterface = () => {
    const router = useRouter();

    const [chatList, setChatList] = useState([{ "id": 1, "name": "Karen chat" }, { "id": 2, "name": "Akkhu chat" }, { "id": 3, "name": "Mehendale chat" }]);
    const [currentChat, setCurrentChat] = useState(1);
    const [chatContent, setChatContent] = useState<{ [key: number]: { sender: string; content: string }[] }>({
        1: [
          { sender: "user", content: "How to get a chick I like from my First year of college" },
          { sender: "assistant", content: "I'm sorry, I can't assist with that." }
        ],
        2: [
          { sender: "user", content: "Can you help me with React?" },
          { sender: "assistant", content: "Sure! What do you need help with?" }
        ]
      });



    return (
        <div style={{ display: 'flex', height: '100vh', width: '100vw', fontFamily: 'Arial, sans-serif' }}>
            {/* Sidebar */}
            <div style={{ width: '250px', background: '#1f2937', color: 'white', padding: '1rem', overflowY: 'auto' }}>
                <h2 style={{ fontSize: '1.2rem', marginBottom: '1rem', color: '#60a5fa' }}>Chats</h2>
                <ul style={{ listStyle: 'none', padding: 0 }}>
                    {chatList.map((chat, index) => (
                        <li
                            key={index}
                            style={{
                                padding: '0.75rem',
                                marginBottom: '0.5rem',
                                background: '#374151',
                                borderRadius: '5px',
                                cursor: 'pointer',
                                color: '#e5e7eb',
                            }}
                        >
                            {chat.name}
                        </li>
                    ))}
                </ul>
            </div>

            {/* Main Chat Window */}
            <div style={{ flex: 1, background: '#111827', display: 'flex', flexDirection: 'column' }}>
                <div style={{ flex: 1, padding: '1rem', overflowY: 'auto' }}>
                    {chatContent[currentChat]?.map((message, index) => (
                        <div
                            key={index}
                            style={{
                                display: 'flex',
                                justifyContent: message.sender === 'user' ? 'flex-end' : 'flex-start',
                                marginBottom: '1rem',
                            }}
                        >
                            <div
                                style={{
                                    background: message.sender === 'user' ? '#3b82f6' : '#374151',
                                    color: 'white',
                                    padding: '0.75rem 1rem',
                                    borderRadius: '10px',
                                    maxWidth: '70%',
                                }}
                            >
                                {message.content}
                            </div>
                        </div>
                    ))}
                </div>

                {/* Chat Input */}
                <div style={{ padding: '1rem', borderTop: '1px solid #374151', background: '#1f2937' }}>
                    <div style={{ display: 'flex', alignItems: 'center' }}>
                        <input
                            type="text"
                            placeholder="Send a message..."
                            style={{
                                flex: 1,
                                padding: '0.75rem 1rem',
                                borderRadius: '8px',
                                border: 'none',
                                outline: 'none',
                                background: '#374151',
                                color: 'white',
                                fontSize: '1rem',
                            }}
                        />
                        <button
                            style={{
                                marginLeft: '0.5rem',
                                background: '#3b82f6',
                                border: 'none',
                                color: 'white',
                                padding: '0.75rem 1rem',
                                borderRadius: '8px',
                                cursor: 'pointer',
                            }}
                        >
                            Send
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );


}

export default ChatInterface;