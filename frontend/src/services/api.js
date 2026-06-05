import axios from 'axios'

const API_BASE = import.meta.env.VITE_API_URL || 'http://localhost:3000'

const api = axios.create({
  baseURL: `${API_BASE}/api/v1`,
  headers: {
    'Content-Type': 'application/json'
  }
})

export default {
  /* ── Products ── */
  getProducts() {
    return api.get('/products')
  },

  getProduct(id) {
    return api.get(`/products/${id}`)
  },

  createProduct(product) {
    return api.post('/products', product)
  },

  updateProduct(id, product) {
    return api.put(`/products/${id}`, product)
  },

  deleteProduct(id) {
    return api.delete(`/products/${id}`)
  }
}
