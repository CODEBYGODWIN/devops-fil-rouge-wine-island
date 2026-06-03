<template>
  <nav class="navbar" :class="{ scrolled: isScrolled }">
    <div class="navbar-inner container">
      <router-link to="/" class="navbar-brand" id="nav-brand">
        <span class="brand-text">Wine</span>
        <span class="brand-sep">·</span>
        <span class="brand-text brand-italic">Island</span>
      </router-link>

      <div class="navbar-links">
        <router-link to="/" class="nav-link" id="nav-catalog">Collection</router-link>
        <router-link to="/products/new" class="nav-link nav-link-cta" id="nav-add-product">
          Add Wine
        </router-link>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: 'AppNavbar',
  data() {
    return {
      isScrolled: false
    }
  },
  mounted() {
    window.addEventListener('scroll', this.onScroll)
  },
  beforeUnmount() {
    window.removeEventListener('scroll', this.onScroll)
  },
  methods: {
    onScroll() {
      this.isScrolled = window.scrollY > 20
    }
  }
}
</script>

<style scoped>
.navbar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 100;
  transition: all var(--t-base);
}

.navbar.scrolled {
  background: rgba(10, 10, 10, 0.92);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border-bottom: 1px solid var(--border);
}

.navbar-inner {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 88px;
}

/* Brand */
.navbar-brand {
  display: flex;
  align-items: baseline;
  gap: 6px;
  text-decoration: none;
  color: var(--white);
}

.brand-text {
  font-family: var(--font-display);
  font-size: 1.4rem;
  font-weight: 500;
  letter-spacing: 0.02em;
}

.brand-italic {
  font-style: italic;
  color: var(--gold);
}

.brand-sep {
  color: var(--gold);
  font-size: 1.2rem;
  opacity: 0.6;
}

/* Links */
.navbar-links {
  display: flex;
  align-items: center;
  gap: 36px;
}

.nav-link {
  font-size: 0.75rem;
  font-weight: 500;
  letter-spacing: 0.14em;
  text-transform: uppercase;
  color: var(--white-dim);
  text-decoration: none;
  transition: color var(--t-fast);
  position: relative;
}

.nav-link:hover {
  color: var(--white);
}

.nav-link::after {
  content: '';
  position: absolute;
  bottom: -4px;
  left: 0;
  width: 0;
  height: 1px;
  background: var(--gold);
  transition: width var(--t-fast);
}

.nav-link:hover::after {
  width: 100%;
}

.nav-link-cta {
  color: var(--gold);
}

.nav-link-cta:hover {
  color: var(--gold-light);
}

.nav-link-cta::after {
  display: none;
}

@media (max-width: 480px) {
  .navbar-inner {
    height: 72px;
  }

  .navbar-links {
    gap: 20px;
  }

  .nav-link {
    font-size: 0.7rem;
  }
}
</style>
