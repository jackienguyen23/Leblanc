import axios from 'axios'

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE || 'http://localhost:3000',
})

export const getDrinks = () => api.get('/drinks').then((res) => res.data)

export const recoFromFeatures = (payload) =>
  api.post('/reco/from-features', payload).then((res) => res.data)

export const createBooking = (booking) =>
  api.post('/bookings', booking).then((res) => res.data)

export default api
