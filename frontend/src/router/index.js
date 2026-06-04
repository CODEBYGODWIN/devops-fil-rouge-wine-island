import { createRouter, createWebHistory } from 'vue-router'
import ProductList from '../views/ProductList.vue'
import ProductDetail from '../views/ProductDetail.vue'
import ProductForm from '../views/ProductForm.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: ProductList
  },
  {
    path: '/products/new',
    name: 'product-create',
    component: ProductForm
  },
  {
    path: '/products/:id',
    name: 'product-detail',
    component: ProductDetail,
    props: true
  },
  {
    path: '/products/:id/edit',
    name: 'product-edit',
    component: ProductForm,
    props: true
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
