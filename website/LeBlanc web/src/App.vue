<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div class="site">
    <header class="site-header">
      <RouterLink to="/" class="brand">Le'Blanc</RouterLink>
      <nav class="nav">
        <RouterLink to="/" class="nav-link" exact-active-class="active">About</RouterLink>
        <RouterLink to="/menu" class="nav-link" exact-active-class="active">Menu</RouterLink>
        <RouterLink to="/booking" class="nav-link" exact-active-class="active">Booking</RouterLink>
        <RouterLink to="/mood-booker" class="nav-link" exact-active-class="active">Mood-Booker</RouterLink>
        <a href="#contact" class="nav-link">Contact</a>
        <RouterLink to="/login" class="nav-link" exact-active-class="active">Login</RouterLink>
        <RouterLink to="/register" class="cta" exact-active-class="active">Reserve</RouterLink>
      </nav>
    </header>

    <main class="page">
      <RouterView />
    </main>
  </div>
</template>

<style scoped>
:global(html) {
  font-size: 18px;
  zoom: 200%;
}

:global(body) {
  margin: 0;
  font-family: 'Playfair Display', 'Inter', 'Segoe UI', system-ui, sans-serif;
  background: #f6efe4;
  color: #1c0f05;
  font-size: 1rem;
  line-height: 1.6;
}

.site {
  min-height: 100vh;
}

.site-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 18px clamp(16px, 5vw, 80px);
  border-bottom: 1px solid #1f1208;
  background: #fff;
  position: sticky;
  top: 0;
  z-index: 5;
}

.brand {
  font-size: 1.6rem;
  text-decoration: none;
  color: #1f1208;
  font-weight: 700;
}

.nav {
  display: flex;
  gap: 18px;
  align-items: center;
  flex-wrap: wrap;
}

.nav-link {
  text-decoration: none;
  color: #1f1208;
  font-weight: 500;
  font-size: 1rem;
}

.nav-link.active {
  text-decoration: underline;
}

.cta {
  border: 1px solid #1f1208;
  padding: 12px 26px;
  border-radius: 999px;
  text-decoration: none;
  color: #1f1208;
  font-weight: 600;
  font-size: 1rem;
}

.page {
  padding: 80px clamp(32px, 8vw, 160px);
  max-width: min(1440px, 94vw);
  margin: 0 auto;
  width: 100%;
}

@media (max-width: 720px) {
  .site-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 12px;
  }
  .nav {
    gap: 12px;
  }
}
</style>
