<template>
  <section class="detail-page">
    <div class="container">
      <!-- Back -->
      <router-link to="/" class="back-link animate-fade-in" id="back-to-list">
        ← Return to Collection
      </router-link>

      <!-- Loading -->
      <div v-if="loading" class="detail-layout animate-fade-in">
        <div class="skeleton detail-skel-image"></div>
        <div class="detail-info">
          <div class="skeleton" style="width: 40%; height: 14px"></div>
          <div class="skeleton" style="width: 70%; height: 36px; margin-top: 16px"></div>
          <div class="skeleton" style="width: 25%; height: 28px; margin-top: 24px"></div>
          <div class="skeleton" style="width: 100%; height: 100px; margin-top: 32px"></div>
        </div>
      </div>

      <!-- Error -->
      <div v-else-if="error" class="empty-state">
        <div class="empty-state-icon">🍷</div>
        <h2>Wine Not Found</h2>
        <p>{{ error }}</p>
        <router-link to="/" class="btn btn-primary" id="go-home-btn">
          Browse Collection
        </router-link>
      </div>

      <!-- Detail -->
      <div v-else class="detail-layout">
        <!-- Image -->
        <div class="detail-image-col animate-fade-in-up">
          <div class="detail-image-frame">
            <img
              v-if="product.image_url"
              :src="product.image_url"
              :alt="product.name"
              class="detail-image"
              @error="imgFailed = true"
            />
            <div v-if="!product.image_url || imgFailed" class="detail-image-empty">
              <span>🍷</span>
            </div>
          </div>
        </div>

        <!-- Info -->
        <div class="detail-info-col">
          <div class="animate-fade-in-up" style="animation-delay: 100ms">
            <p class="detail-eyebrow">Wine Nº {{ product.id }}</p>
            <h1 class="detail-name">{{ product.name }}</h1>
            <div class="divider" style="margin: 24px 0"></div>
          </div>

          <div class="detail-price animate-fade-in-up" style="animation-delay: 180ms">
            {{ formatPrice(product.price) }}
          </div>

          <div class="detail-desc animate-fade-in-up" style="animation-delay: 250ms">
            <p class="desc-label">Tasting Notes</p>
            <p class="desc-text">{{ product.description || 'No description provided.' }}</p>
          </div>

          <div class="detail-actions animate-fade-in-up" style="animation-delay: 320ms">
            <router-link
              :to="`/products/${product.id}/edit`"
              class="btn btn-secondary"
              id="edit-product-btn"
            >
              Edit
            </router-link>
            <button
              class="btn btn-danger"
              @click="showDeleteModal = true"
              id="delete-product-btn"
            >
              Remove
            </button>
          </div>
        </div>
      </div>

      <!-- Delete modal -->
      <ConfirmModal
        v-if="showDeleteModal"
        title="Remove Wine?"
        :message="`This will permanently remove &quot;${product.name}&quot; from the collection.`"
        confirm-label="Remove"
        @close="showDeleteModal = false"
        @confirm="handleDelete"
      />
    </div>
  </section>
</template>

<script>
import api from '../services/api'
import ConfirmModal from '../components/ConfirmModal.vue'
import { showToast } from '../components/ToastNotification.vue'

export default {
  name: 'ProductDetail',
  components: { ConfirmModal },
  props: {
    id: {
      type: [String, Number],
      required: true
    }
  },
  data() {
    return {
      product: {},
      loading: true,
      error: null,
      showDeleteModal: false,
      imgFailed: false
    }
  },
  mounted() {
    this.fetchProduct()
  },
  methods: {
    async fetchProduct() {
      this.loading = true
      this.error = null
      try {
        const { data } = await api.getProduct(this.id)
        this.product = data
      } catch {
        this.error = 'This wine could not be found or the cellar is unreachable.'
      } finally {
        this.loading = false
      }
    },
    async handleDelete() {
      try {
        await api.deleteProduct(this.product.id)
        showToast('Wine removed from collection', 'success')
        this.$router.push('/')
      } catch {
        showToast('Failed to remove wine', 'error')
      }
      this.showDeleteModal = false
    },
    formatPrice(price) {
      return new Intl.NumberFormat('en-US', {
        style: 'currency',
        currency: 'USD'
      }).format(price)
    }
  }
}
</script>

<style scoped>
/* Back */
.back-link {
  display: inline-block;
  margin-bottom: 48px;
  font-size: 0.78rem;
  font-weight: 400;
  letter-spacing: 0.06em;
  color: var(--white-faint);
  text-decoration: none;
  transition: color var(--t-fast);
}

.back-link:hover {
  color: var(--gold);
}

/* Layout */
.detail-layout {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 80px;
  align-items: start;
}

/* Image */
.detail-image-frame {
  aspect-ratio: 3 / 4;
  overflow: hidden;
  background: var(--black-card);
  border-radius: var(--radius-xs);
}

.detail-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  filter: contrast(1.04);
}

.detail-image-empty {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 5rem;
  opacity: 0.12;
}

.detail-skel-image {
  aspect-ratio: 3 / 4;
  border-radius: var(--radius-xs);
}

/* Info */
.detail-info-col {
  padding-top: 16px;
}

.detail-eyebrow {
  font-size: 0.68rem;
  letter-spacing: 0.2em;
  text-transform: uppercase;
  color: var(--gold);
  font-weight: 500;
  margin-bottom: 12px;
}

.detail-name {
  font-size: 2.8rem;
  font-weight: 400;
  line-height: 1.1;
  color: var(--white);
}

.detail-price {
  font-family: var(--font-display);
  font-size: 2rem;
  font-weight: 500;
  color: var(--white-soft);
  margin-bottom: 32px;
}

/* Description */
.detail-desc {
  padding-top: 32px;
  border-top: 1px solid var(--border);
}

.desc-label {
  font-size: 0.68rem;
  letter-spacing: 0.2em;
  text-transform: uppercase;
  color: var(--white-dim);
  margin-bottom: 14px;
  font-weight: 500;
}

.desc-text {
  font-size: 0.95rem;
  line-height: 1.9;
  color: var(--white-muted);
  font-weight: 300;
}

/* Actions */
.detail-actions {
  margin-top: 40px;
  display: flex;
  gap: 14px;
}

@media (max-width: 768px) {
  .detail-layout {
    grid-template-columns: 1fr;
    gap: 40px;
  }

  .detail-name {
    font-size: 2rem;
  }

  .detail-price {
    font-size: 1.5rem;
  }
}
</style>
