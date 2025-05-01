"use client";

import { useRouter } from "next/navigation";
import React, { useState } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faSignInAlt,
  faPaperPlane,
  faSearch,
} from "@fortawesome/free-solid-svg-icons";
import Chat from "@/types/chatTypes";

const ChatInterface = () => {
  const router = useRouter();

  const [chatList] = useState([
    { id: 1, name: "Karen chat" },
    { id: 2, name: "Akkhu chat" },
  ]);
  const [currentChat, setCurrentChat] = useState(1);
  const [currentInput, setCurrentInput] = useState("");
  const [chatContent, setChatContent] = useState<Chat[]>([
    {
      id: 1,
      messages: [
        {
          sender: "user",
          content:
            "I am trying for this chick from first year and have still not gotten her, how do I get her?",
        },
        {
          sender: "assistant",
          content:
            "Sorry I can't assist you with that. Feel free to ask me other questions",
        },
      ],
    },
    {
      id: 2,
      messages: [],
    },
  ]);

  const handleSend = () => {
    if (!currentInput.trim()) return;
    setChatContent((prev) =>
      prev.map((chat) =>
        chat.id === currentChat
          ? {
              ...chat,
              messages: [
                ...chat.messages,
                { sender: "user", content: currentInput },
                { sender: "assistant", content: "I can assist you with that!" },
              ],
            }
          : chat
      )
    );
    setCurrentInput("");
  };

  const currentChatObj = chatContent.find((c) => c.id === currentChat);
  const currentChatName =
    chatList.find((c) => c.id === currentChat)?.name || "Chat";

  return (
    <div className="flex h-screen w-screen font-sans bg-gray-950 text-gray-100">
      <div className="w-[300px] bg-gray-850 p-5 flex flex-col h-full items-center">
        <div className="flex items-center justify-between mb-6 w-full">
          <h2 className="text-xl font-bold text-blue-400 flex items-center">
            Chats
          </h2>
          <button
            onClick={() => router.push("/login")}
            className="flex items-center text-sm font-medium text-blue-400 hover:text-blue-300 transition px-3 py-1 rounded-full border border-blue-500 hover:bg-blue-500/10 focus:outline-none focus:ring-2 focus:ring-blue-500/40"
          >
            <FontAwesomeIcon icon={faSignInAlt} className="mr-1" />
            Login
          </button>
        </div>

        <div className="relative mb-6 w-full px-2">
          <FontAwesomeIcon
            icon={faSearch}
            className="absolute left-5 top-1/2 -translate-y-1/2 text-gray-500"
          />
          <input
            type="text"
            placeholder="Search chats..."
            className="w-full bg-gray-800 text-gray-200 pl-10 pr-4 py-2 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500/40 placeholder-gray-500"
          />
        </div>

        <ul className="flex-1 min-h-0 overflow-y-auto custom-scrollbar w-full items-center flex flex-col">
          {chatList.map((chat) => (
            <li key={chat.id} className="w-[90%] mt-2">
              <button
                className={`w-full px-4 py-2 rounded-md text-left transition duration-200 ${
                  currentChat === chat.id
                    ? "bg-blue-600 text-white"
                    : "text-gray-300 hover:bg-gray-700"
                }`}
                onClick={() => setCurrentChat(chat.id)}
              >
                <span
                  className={`inline-block w-2 h-2 rounded-full mr-3 ${
                    currentChat === chat.id ? "bg-white" : "bg-blue-400"
                  }`}
                />
                <span className="truncate">{chat.name}</span>
              </button>
            </li>
          ))}
        </ul>
      </div>

      <div className="flex-1 flex flex-col align-text-bottom">
        <div className="h-16 px-6 flex items-center  border-b border-gray-800 bg-gray-875">
          <div className="flex items-center space-x-3">
            <div className="w-8 h-8 rounded-full bg-blue-500 flex items-center justify-center text-white font-medium">
              {currentChatName[0].toUpperCase()}
            </div>
            <h3 className="mt-3 h-10 font-medium justify-center ">{currentChatName}</h3>
          </div>
        </div>

        <div className="flex-1 p-6 overflow-y-auto space-y-5 custom-scrollbar bg-gray-900">
          {currentChatObj && currentChatObj.messages.length > 0 ? (
            currentChatObj.messages.map((message, idx) => (
              <div
                key={idx}
                className={`flex ${
                  message.sender === "user" ? "justify-end" : "justify-start"
                }`}
              >
                {message.sender !== "user" && (
                  <div className="w-8 h-8 rounded-full bg-indigo-500 flex items-center justify-center text-white font-medium mr-3 mt-1">
                    A
                  </div>
                )}
                <div
                  className={`px-5 py-3.5 rounded-2xl max-w-[70%] ${
                    message.sender === "user"
                      ? "bg-blue-600 text-white shadow-md shadow-blue-900/20"
                      : "bg-gray-800 text-gray-100 shadow-md shadow-black/5"
                  }`}
                >
                  <div className="whitespace-pre-wrap break-words">
                    {message.content}
                  </div>
                </div>
                {message.sender === "user" && (
                  <div className="w-8 h-8 rounded-full bg-blue-500 flex items-center justify-center text-white font-medium ml-3 mt-1">
                    {currentChatName[0].toUpperCase()}
                  </div>
                )}
              </div>
            ))
          ) : (
            <div className="flex flex-col items-center justify-center h-full text-center">
              <h3 className="text-xl font-medium text-gray-300 mb-2">
                Start a Conversation
              </h3>
              <p className="text-gray-500 max-w-sm">
                Ask anything you want...
              </p>
            </div>
          )}
        </div>

        <div className="p-4 border-t border-gray-800 bg-gray-875">
          <div className="flex gap-3 items-center">
            <input
              type="text"
              placeholder="Type a message..."
              className="flex-1 px-4 py-3.5 bg-gray-800 text-white rounded-full focus:outline-none focus:ring-2 focus:ring-blue-500/50 transition"
              value={currentInput}
              onChange={(e) => setCurrentInput(e.target.value)}
              onKeyDown={(e) => {
                if (e.key === "Enter") handleSend();
              }}
            />
            <button
              className="bg-blue-600 hover:bg-blue-500 transition text-white p-3.5 rounded-full font-medium shadow-lg shadow-blue-900/20 focus:ring-2 focus:ring-blue-600/50 focus:outline-none"
              onClick={handleSend}
            >
              <FontAwesomeIcon icon={faPaperPlane} />
            </button>
          </div>
        </div>
      </div>

      <style jsx global>{`
        .custom-scrollbar::-webkit-scrollbar {
          width: 6px;
          height: 6px;
        }
        .custom-scrollbar::-webkit-scrollbar-track {
          background: transparent;
        }
        .custom-scrollbar::-webkit-scrollbar-thumb {
          background-color: rgba(107, 114, 128, 0.3);
          border-radius: 20px;
        }
        .custom-scrollbar::-webkit-scrollbar-thumb:hover {
          background-color: rgba(107, 114, 128, 0.5);
        }

        /* Add custom colors for gray-850 and gray-875 */
        .bg-gray-850 {
          background-color: #1d1f27;
        }
        .bg-gray-875 {
          background-color: #23252e;
        }
      `}</style>
    </div>
  );
};

export default ChatInterface;
