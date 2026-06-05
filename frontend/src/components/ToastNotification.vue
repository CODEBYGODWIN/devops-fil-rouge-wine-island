<template>
  <div class="toast-container">
    <transition-group name="toast-anim">
      <div
        v-for="toast in toasts"
        :key="toast.id"
        :class="['toast', `toast-${toast.type}`]"
      >
        {{ toast.message }}
      </div>
    </transition-group>
  </div>
</template>

<script>
import { ref } from 'vue'

const toasts = ref([])
let toastId = 0

export function showToast(message, type = 'success', duration = 3500) {
  const id = ++toastId
  toasts.value.push({ id, message, type })
  setTimeout(() => {
    toasts.value = toasts.value.filter(t => t.id !== id)
  }, duration)
}

export default {
  name: 'ToastNotification',
  setup() {
    return { toasts }
  }
}
</script>

<style scoped>
.toast-anim-enter-active {
  animation: fadeInUp 0.35s var(--ease-out) both;
}

.toast-anim-leave-active {
  animation: fadeIn 0.2s var(--ease-out) reverse both;
}
</style>
