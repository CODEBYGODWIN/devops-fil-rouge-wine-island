<template>
  <div class="app-shell">
    <AppNavbar />
    <main class="app-main">
      <router-view v-slot="{ Component }">
        <transition name="page" mode="out-in">
          <component :is="Component" />
        </transition>
      </router-view>
    </main>
    <footer class="app-footer">
      <div class="container footer-inner">
        <div class="divider-wide"></div>
        <div class="footer-content">
          <span class="footer-brand">Wine Island</span>
          <span class="footer-copy">© {{ year }} — Curated Fine Wines</span>
        </div>
      </div>
    </footer>
    <ToastNotification />
  </div>
</template>

<script>
import AppNavbar from './components/AppNavbar.vue'
import ToastNotification from './components/ToastNotification.vue'

export default {
  name: 'App',
  components: {
    AppNavbar,
    ToastNotification
  },
  computed: {
    year() {
      return new Date().getFullYear()
    }
  }
}
</script>

<style>
/* Page transition */
.page-enter-active {
  animation: fadeInUp 0.5s var(--ease-out) both;
}

.page-leave-active {
  animation: fadeIn 0.15s var(--ease-out) reverse both;
}

.app-main {
  padding-top: 120px;
  padding-bottom: 40px;
  min-height: 100vh;
}

/* Footer */
.app-footer {
  padding: 0 0 48px;
}

.footer-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 32px;
}

.footer-brand {
  font-family: var(--font-display);
  font-size: 1rem;
  font-weight: 500;
  color: var(--white-dim);
  letter-spacing: 0.04em;
}

.footer-copy {
  font-size: 0.75rem;
  color: var(--white-faint);
  letter-spacing: 0.06em;
  text-transform: uppercase;
}

@media (max-width: 480px) {
  .footer-content {
    flex-direction: column;
    gap: 8px;
    text-align: center;
  }
}
</style>
