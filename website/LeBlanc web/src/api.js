import axios from 'axios'
import {
  getUsersGraphQL,
  getDrinksGraphQL,
  getDrinkGraphQL,
  getBookingsGraphQL,
  createBookingGraphQL,
  registerUserGraphQL,
  loginUserGraphQL,
  recoFromFeaturesGraphQL,
} from './graphql'

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE  ,
})

// Configuration: Set to true to use GraphQL, false to use REST API
const USE_GRAPHQL = import.meta.env.VITE_USE_GRAPHQL === 'true' || false

// REST API functions
const getUsersREST = () => api.get('/users').then((res) => res.data)
const getDrinksREST = () => api.get('/drinks').then((res) => res.data)

const recoFromFeaturesREST = (payload) =>
  api.post('/reco/from-features', payload).then((res) => res.data)

const createBookingREST = (booking) =>
  api.post('/bookings', booking).then((res) => res.data)

const registerUserREST = (payload) =>
  api.post('/auth/register', payload).then((res) => res.data)

const loginUserREST = (payload) =>
  api.post('/auth/login', payload).then((res) => res.data)

const requestVerifyREST = (payload) =>
  api.post('/auth/request-verify', payload).then((res) => res.data)

const verifyTokenREST = (payload) =>
  api.post('/auth/verify', payload).then((res) => res.data)

// Unified API - switches between REST and GraphQL based on configuration
export const getUsers = () => {
  return USE_GRAPHQL ? getUsersGraphQL() : getUsersREST()
}

export const getDrinks = () => {
  return USE_GRAPHQL ? getDrinksGraphQL() : getDrinksREST()
}

export const getDrink = (id) => {
  return USE_GRAPHQL ? getDrinkGraphQL(id) : api.get(`/drinks/${id}`).then((res) => res.data)
}

export const getBookings = () => {
  return USE_GRAPHQL ? getBookingsGraphQL() : api.get('/bookings').then((res) => res.data)
}

export const recoFromFeatures = (payload) => {
  if (USE_GRAPHQL) {
    // Convert REST payload to GraphQL format
    // Assuming payload has emotion fit and optional preferences
    const emotionFit = payload.emotionFit || {
      calm: payload.calm || 0,
      happy: payload.happy || 0,
      stressed: payload.stressed || 0,
      sad: payload.sad || 0,
      adventurous: payload.adventurous || 0,
    }
    return recoFromFeaturesGraphQL(
      emotionFit,
      payload.caffeine,
      payload.temp,
      payload.sweetness
    )
  }
  return recoFromFeaturesREST(payload)
}

export const createBooking = (booking) => {
  if (USE_GRAPHQL) {
    // Convert booking data to GraphQL format if needed
    const input = {
      name: booking.name,
      phone: booking.phone,
      time: booking.time,
      items: booking.items.map(item => ({
        drinkId: item.drinkId,
        qty: item.qty,
        options: item.options ? JSON.stringify(item.options) : ''
      })),
      channel: booking.channel || 'web',
    }
    return createBookingGraphQL(input)
  }
  return createBookingREST(booking)
}

export const registerUser = (payload) => {
  if (USE_GRAPHQL) {
    const input = {
      name: payload.name,
      email: payload.email,
      password: payload.password,
    }
    return registerUserGraphQL(input)
  }
  return registerUserREST(payload)
}

export const loginUser = (payload) => {
  if (USE_GRAPHQL) {
    const input = {
      name: payload.name || payload.nameOrEmail,
      password: payload.password,
    }
    return loginUserGraphQL(input)
  }
  return loginUserREST(payload)
}

export const requestVerify = (payload) => {
  return requestVerifyREST(payload)
}

export const verifyToken = (payload) => {
  return verifyTokenREST(payload)
}

// Export GraphQL functions directly for specific use cases
export {
  getDrinksGraphQL,
  getDrinkGraphQL,
  getBookingsGraphQL,
  createBookingGraphQL,
  registerUserGraphQL,
  loginUserGraphQL,
  recoFromFeaturesGraphQL,
  getUsersGraphQL,
}

// Export REST functions directly for specific use cases
export {
  getDrinksREST,
  recoFromFeaturesREST,
  createBookingREST,
  registerUserREST,
  loginUserREST,
  getUsersREST,
  requestVerifyREST,
  verifyTokenREST,
}

export default api
