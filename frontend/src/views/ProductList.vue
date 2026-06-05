<template>
  <section class="catalog-page">
    <!-- Hero -->
    <header class="hero container">
      <div class="hero-content animate-fade-in-up">
        <p class="hero-eyebrow">E s t .&nbsp;&nbsp;2 0 2 6</p>
        <h1 class="hero-title">
          Our<br/>
          <em>Collection</em>
        </h1>
        <div class="divider" style="margin-top: 24px"></div>
        <p class="hero-tagline">
          Discover rare and exceptional wines from the world's most<br class="hide-mobile" />
          celebrated vineyards, handpicked for the discerning palate.
        </p>
      </div>
    </header>

    <div class="container">
      <!-- Controls bar -->
      <div class="controls-bar animate-fade-in-up" style="animation-delay: 200ms">
        <div class="search-wrapper">
          <input
            type="text"
            class="search-input"
            placeholder="Search the collection…"
            v-model="searchQuery"
            id="search-products"
          />
        </div>
        <span class="wine-count" v-if="!loading">
          {{ filteredProducts.length }} wine{{ filteredProducts.length !== 1 ? 's' : '' }}
        </span>
      </div>

      <!-- Loading -->
      <div v-if="loading" class="wine-grid stagger">
        <div v-for="i in 8" :key="i" class="skeleton-card animate-fade-in-up">
          <div class="skeleton skeleton-img"></div>
          <div class="skeleton-body">
            <div class="skeleton skeleton-line" style="width: 50%; height: 1px; margin-bottom: 14px"></div>
            <div class="skeleton skeleton-line" style="width: 70%; height: 16px"></div>
            <div class="skeleton skeleton-line" style="width: 45%; height: 12px; margin-top: 8px"></div>
            <div class="skeleton skeleton-line" style="width: 30%; height: 14px; margin-top: 14px"></div>
          </div>
        </div>
      </div>

      <!-- Error -->
      <div v-else-if="error" class="empty-state">
        <div class="empty-state-icon">🍷</div>
        <h2>Cellar Unavailable</h2>
        <p>{{ error }}</p>
        <button class="btn btn-primary" @click="fetchProducts" id="retry-btn">
          Try Again
        </button>
      </div>

      <!-- Empty -->
      <div v-else-if="filteredProducts.length === 0 && !searchQuery" class="empty-state">
        <div class="empty-state-icon">🏝️</div>
        <h2>The Cellar Awaits</h2>
        <p>Add your first wine to begin curating the collection.</p>
        <router-link to="/products/new" class="btn btn-primary" id="empty-add-btn">
          Add First Wine
        </router-link>
      </div>

      <!-- No results -->
      <div v-else-if="filteredProducts.length === 0 && searchQuery" class="empty-state">
        <div class="empty-state-icon">🔍</div>
        <h2>No Wines Found</h2>
        <p>No wines match "{{ searchQuery }}"</p>
        <button class="btn btn-secondary" @click="searchQuery = ''" id="clear-search-btn">
          Clear Search
        </button>
      </div>

      <!-- Wine grid -->
      <div v-else class="wine-grid stagger">
        <ProductCard
          v-for="product in filteredProducts"
          :key="product.id"
          :product="product"
        />
      </div>
    </div>
  </section>
</template>

<script>
import api from '../services/api'
import ProductCard from '../components/ProductCard.vue'

export default {
  name: 'ProductList',
  components: {
    ProductCard
  },
  data() {
    return {
      products: [],
      loading: true,
      error: null,
      searchQuery: ''
    }
  },
  computed: {
    filteredProducts() {
      if (!this.searchQuery) return this.products
      const q = this.searchQuery.toLowerCase()
      return this.products.filter(p =>
        p.name.toLowerCase().includes(q) ||
        (p.description && p.description.toLowerCase().includes(q))
      )
    }
  },
  mounted() {
    this.fetchProducts()
  },
  methods: {
    async fetchProducts() {
      this.loading = true
      this.error = null
      try {
        const { data } = await api.getProducts()
        this.products = data || []
      } catch (err) {
        this.error = 'Could not connect to the cellar. Please ensure the backend is running.'
      } finally {
        this.loading = false
      }
    }
  }
}
</script>

<style scoped>
/* Hero */
.hero {
  padding-bottom: 72px;
}

.hero-eyebrow {
  font-size: 0.65rem;
  letter-spacing: 0.35em;
  text-transform: uppercase;
  color: var(--gold);
  margin-bottom: 20px;
  font-weight: 500;
}

.hero-title {
  font-size: 4rem;
  font-weight: 400;
  line-height: 1.05;
  color: var(--white);
}

.hero-title em {
  font-style: italic;
  color: var(--white-soft);
}

.hero-tagline {
  margin-top: 24px;
  font-size: 0.95rem;
  line-height: 1.8;
  max-width: 480px;
}

/* Controls */
.controls-bar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 48px;
  border-bottom: 1px solid var(--border);
  margin-bottom: 48px;
  gap: 24px;
}

.search-wrapper {
  flex: 1;
  max-width: 360px;
}

.search-input {
  width: 100%;
  padding: 12px 0;
  background: none;
  border: none;
  border-bottom: 1px solid var(--border-light);
  color: var(--white);
  font-family: var(--font-body);
  font-size: 0.9rem;
  font-weight: 300;
  outline: none;
  transition: border-color var(--t-fast);
}

.search-input::placeholder {
  color: var(--white-faint);
}

.search-input:focus {
  border-bottom-color: var(--gold);
}

.wine-count {
  font-size: 0.72rem;
  color: var(--white-faint);
  letter-spacing: 0.1em;
  text-transform: uppercase;
  white-space: nowrap;
}

/* Grid */
.wine-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 32px 28px;
}

/* Skeleton */
.skeleton-card {
  display: flex;
  flex-direction: column;
}

.skeleton-img {
  width: 100%;
  aspect-ratio: 3 / 4;
  border-radius: var(--radius-xs);
}

.skeleton-body {
  padding: 20px 4px 0;
}

.skeleton-line {
  border-radius: 2px;
}

/* Responsive */
@media (max-width: 1024px) {
  .wine-grid {
    grid-template-columns: repeat(3, 1fr);
  }
}

@media (max-width: 768px) {
  .hero-title {
    font-size: 2.8rem;
  }

  .wine-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 24px 20px;
  }

  .hide-mobile {
    display: none;
  }
}

@media (max-width: 480px) {
  .wine-grid {
    grid-template-columns: 1fr;
    max-width: 320px;
    margin: 0 auto;
  }

  .controls-bar {
    flex-direction: column;
    align-items: flex-start;
  }

  .search-wrapper {
    max-width: 100%;
  }
}
</style>
