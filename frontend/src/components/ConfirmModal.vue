<template>
  <div class="modal-overlay" @click.self="$emit('close')">
    <div class="modal-content animate-scale-in">
      <h2 class="modal-title">{{ title }}</h2>
      <div class="divider" style="margin: 16px 0 20px"></div>
      <p class="modal-message">{{ message }}</p>
      <div class="modal-actions">
        <button class="btn btn-secondary" @click="$emit('close')" id="modal-cancel-btn">
          Cancel
        </button>
        <button class="btn btn-danger" @click="$emit('confirm')" id="modal-confirm-btn">
          {{ confirmLabel }}
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ConfirmModal',
  props: {
    title: {
      type: String,
      default: 'Are you sure?'
    },
    message: {
      type: String,
      default: 'This action cannot be undone.'
    },
    confirmLabel: {
      type: String,
      default: 'Remove'
    }
  },
  emits: ['close', 'confirm']
}
</script>

<style scoped>
.modal-title {
  font-weight: 400;
  font-size: 1.3rem;
}

.modal-message {
  margin-bottom: 32px;
  color: var(--white-muted);
  font-size: 0.9rem;
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}
</style>
