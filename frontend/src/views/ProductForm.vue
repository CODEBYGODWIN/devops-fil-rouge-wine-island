<template>
  <section class="form-page">
    <div class="container">
      <router-link to="/" class="back-link animate-fade-in" id="form-back-link">
        ← Return to Collection
      </router-link>

      <div class="form-layout animate-fade-in-up">
        <!-- Left: header & preview -->
        <div class="form-left">
          <p class="form-eyebrow">{{ isEditing ? 'Edit Wine' : 'New Addition' }}</p>
          <h1 class="form-title">
            {{ isEditing ? 'Update' : 'Add to' }}<br/>
            <em>the Collection</em>
          </h1>
          <div class="divider" style="margin: 24px 0 32px"></div>

          <!-- Live preview -->
          <div class="live-preview" v-if="form.name || form.image_url">
            <div class="preview-image-frame">
              <img
                v-if="form.image_url && !previewError"
                :src="form.image_url"
                alt="Preview"
                class="preview-image"
                @error="previewError = true"
              />
              <div v-else class="preview-placeholder">🍷</div>
            </div>
            <div class="preview-meta">
              <h3 class="preview-name">{{ form.name || 'Wine Name' }}</h3>
              <span class="preview-price" v-if="form.price">
                {{ formatPrice(form.price) }}
              </span>
            </div>
          </div>
        </div>

        <!-- Right: form -->
        <div class="form-right">
          <form @submit.prevent="handleSubmit">
            <div class="input-group">
              <label for="product-name">Name</label>
              <input
                id="product-name"
                v-model="form.name"
                type="text"
                class="input-field"
                placeholder="e.g. Château Margaux 2015"
                required
              />
            </div>

            <div class="input-group">
              <label for="product-description">Description</label>
              <textarea
                id="product-description"
                v-model="form.description"
                class="input-field"
                placeholder="Tasting notes, origin, grape variety…"
                rows="5"
                required
              ></textarea>
            </div>

            <div class="form-row">
              <div class="input-group">
                <label for="product-price">Price (USD)</label>
                <input
                  id="product-price"
                  v-model.number="form.price"
                  type="number"
                  step="0.01"
                  min="0"
                  class="input-field"
                  placeholder="89.00"
                  required
                />
              </div>

              <div class="input-group">
                <label for="product-image">Image URL</label>
                <input
                  id="product-image"
                  v-model="form.image_url"
                  type="url"
                  class="input-field"
                  placeholder="https://…"
                  @input="previewError = false"
                />
              </div>
            </div>

            <div class="divider-wide" style="margin: 32px 0"></div>

            <div class="form-actions">
              <router-link to="/" class="btn btn-ghost" id="form-cancel-btn">
                Cancel
              </router-link>
              <button
                type="submit"
                class="btn btn-primary"
                :disabled="submitting"
                id="form-submit-btn"
              >
                <span v-if="submitting" class="spinner"></span>
                {{ submitting ? 'Saving…' : (isEditing ? 'Update Wine' : 'Add Wine') }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import api from '../services/api'
import { showToast } from '../components/ToastNotification.vue'

export default {
  name: 'ProductForm',
  props: {
    id: {
      type: [String, Number],
      default: null
    }
  },
  data() {
    return {
      form: {
        name: '',
        description: '',
        price: null,
        image_url: ''
      },
      submitting: false,
      previewError: false
    }
  },
  computed: {
    isEditing() {
      return !!this.id
    }
  },
  mounted() {
    if (this.isEditing) {
      this.loadProduct()
    }
  },
  methods: {
    async loadProduct() {
      try {
        const { data } = await api.getProduct(this.id)
        this.form = {
          name: data.name,
          description: data.description,
          price: data.price,
          image_url: data.image_url || ''
        }
      } catch {
        showToast('Failed to load wine data', 'error')
        this.$router.push('/')
      }
    },
    async handleSubmit() {
      this.submitting = true
      try {
        if (this.isEditing) {
          await api.updateProduct(this.id, this.form)
          showToast('Wine updated successfully', 'success')
          this.$router.push(`/products/${this.id}`)
        } else {
          const { data } = await api.createProduct(this.form)
          showToast('Wine added to collection', 'success')
          this.$router.push(`/products/${data.id}`)
        }
      } catch (err) {
        const msg = err?.response?.data?.error || 'Something went wrong'
        showToast(msg, 'error')
      } finally {
        this.submitting = false
      }
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
/* Layout */
.form-layout {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 80px;
  align-items: start;
}

/* Left side */
.form-eyebrow {
  font-size: 0.65rem;
  letter-spacing: 0.3em;
  text-transform: uppercase;
  color: var(--gold);
  font-weight: 500;
  margin-bottom: 16px;
}

.form-title {
  font-size: 2.8rem;
  font-weight: 400;
  line-height: 1.1;
}

.form-title em {
  font-style: italic;
  color: var(--white-muted);
}

/* Preview */
.live-preview {
  display: flex;
  align-items: center;
  gap: 18px;
  padding: 20px;
  background: var(--black-card);
  border: 1px solid var(--border);
  border-radius: var(--radius-sm);
}

.preview-image-frame {
  width: 72px;
  height: 96px;
  flex-shrink: 0;
  overflow: hidden;
  border-radius: var(--radius-xs);
  background: var(--black-elevated);
}

.preview-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.preview-placeholder {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2rem;
  opacity: 0.15;
}

.preview-name {
  font-family: var(--font-display);
  font-size: 0.95rem;
  font-weight: 500;
  color: var(--white-soft);
  margin-bottom: 4px;
}

.preview-price {
  font-size: 0.82rem;
  color: var(--gold);
  font-weight: 500;
}

/* Right side: form */
.form-right {
  padding-top: 8px;
}

.form-right form {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

/* Spinner */
.spinner {
  width: 14px;
  height: 14px;
  border: 2px solid rgba(10, 10, 10, 0.2);
  border-top-color: var(--black);
  border-radius: 50%;
  animation: spin 0.6s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

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

@media (max-width: 768px) {
  .form-layout {
    grid-template-columns: 1fr;
    gap: 48px;
  }

  .form-title {
    font-size: 2rem;
  }

  .form-row {
    grid-template-columns: 1fr;
  }
}
</style>
