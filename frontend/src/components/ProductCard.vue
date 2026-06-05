<template>
  <router-link
    :to="`/products/${product.id}`"
    class="wine-card animate-fade-in-up"
    :id="`product-card-${product.id}`"
  >
    <div class="card-image-wrapper">
      <img
        v-if="product.image_url"
        :src="product.image_url"
        :alt="product.name"
        class="card-image"
        @error="onImageError"
      />
      <div v-else class="card-image-placeholder">
        <span class="placeholder-icon">🍷</span>
      </div>
    </div>

    <div class="card-body">
      <div class="card-divider"></div>
      <h3 class="card-name">{{ product.name }}</h3>
      <p class="card-desc">{{ truncate(product.description, 60) }}</p>
      <span class="card-price">{{ formatPrice(product.price) }}</span>
    </div>
  </router-link>
</template>

<script>
export default {
  name: 'ProductCard',
  props: {
    product: {
      type: Object,
      required: true
    }
  },
  methods: {
    formatPrice(price) {
      return new Intl.NumberFormat('en-US', {
        style: 'currency',
        currency: 'USD'
      }).format(price)
    },
    truncate(text, max) {
      if (!text) return ''
      return text.length > max ? text.substring(0, max) + '…' : text
    },
    onImageError(e) {
      e.target.style.display = 'none'
      e.target.parentElement.classList.add('image-failed')
    }
  }
}
</script>

<style scoped>
.wine-card {
  display: flex;
  flex-direction: column;
  text-decoration: none;
  color: inherit;
  cursor: pointer;
  transition: transform var(--t-base);
}

.wine-card:hover {
  transform: translateY(-6px);
}

/* Image */
.card-image-wrapper {
  position: relative;
  width: 100%;
  aspect-ratio: 3 / 4;
  overflow: hidden;
  background: var(--black-card);
  border-radius: var(--radius-xs);
}

.card-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.8s var(--ease-out), opacity var(--t-base);
  filter: grayscale(15%) contrast(1.05);
}

.wine-card:hover .card-image {
  transform: scale(1.04);
  filter: grayscale(0%) contrast(1.08);
}

.card-image-placeholder {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--black-card);
}

.placeholder-icon {
  font-size: 3.5rem;
  opacity: 0.15;
}

/* Body */
.card-body {
  padding: 20px 4px 8px;
}

.card-divider {
  width: 24px;
  height: 1px;
  background: var(--gold);
  margin-bottom: 14px;
  opacity: 0.5;
  transition: width var(--t-base), opacity var(--t-base);
}

.wine-card:hover .card-divider {
  width: 40px;
  opacity: 0.9;
}

.card-name {
  font-family: var(--font-display);
  font-size: 1.05rem;
  font-weight: 500;
  color: var(--white);
  margin-bottom: 6px;
  transition: color var(--t-fast);
}

.wine-card:hover .card-name {
  color: var(--gold-light);
}

.card-desc {
  font-size: 0.82rem;
  color: var(--white-faint);
  font-weight: 300;
  line-height: 1.5;
  margin-bottom: 12px;
}

.card-price {
  font-size: 0.9rem;
  font-weight: 500;
  color: var(--gold);
  letter-spacing: 0.04em;
}
</style>
