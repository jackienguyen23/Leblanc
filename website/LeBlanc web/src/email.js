// Template helper for sending verification emails using EmailJS.
// Prereqs:
// 1) Install: npm install @emailjs/browser
// 2) Set env in .env:
//    VITE_EMAILJS_SERVICE_ID=...
//    VITE_EMAILJS_TEMPLATE_ID=...
//    VITE_EMAILJS_PUBLIC_KEY=...
// 3) Configure your EmailJS template to accept variables: to_email, verify_link, token, expires_at.

import emailjs from '@emailjs/browser'
import { requestVerify, verifyToken } from '@/api'

const SERVICE_ID = import.meta.env.VITE_EMAILJS_SERVICE_ID
const TEMPLATE_ID = import.meta.env.VITE_EMAILJS_TEMPLATE_ID
const PUBLIC_KEY = import.meta.env.VITE_EMAILJS_PUBLIC_KEY

const emailConfigured = SERVICE_ID && TEMPLATE_ID && PUBLIC_KEY

// If token/verifyLink are provided (from backend register response), reuse them.
export const sendVerificationEmail = async (email, name = '', opts = {}) => {
  if (!emailConfigured) throw new Error('EmailJS is not configured (missing env vars).')
  let token = opts.token
  let expiresAt = opts.expiresAt
  let verifyLink = opts.verifyLink

  if (!token) {
    const res = await requestVerify({ email })
    token = res.token
    expiresAt = res.expiresAt
  }
  if (!verifyLink) {
    verifyLink = `${window.location.origin}/verify?token=${encodeURIComponent(token)}`
  }
  await emailjs.send(
    SERVICE_ID,
    TEMPLATE_ID,
    {
      to_email: email,
      to: email,
      name,
      to_name: name,
      user_name: name,
      username: name,
      verify_link: verifyLink,
      verifyUrl: verifyLink,
      token,
      expires_at: expiresAt,
    },
    PUBLIC_KEY
  )
  return { token, expiresAt, verifyLink }
}


export const checkVerificationToken = async (token) => {
  return verifyToken({ token })
}

export const isEmailReady = () => Boolean(emailConfigured)
