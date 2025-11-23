<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div class="site">
    <div class="glow glow-amber" aria-hidden="true"></div>
    <div class="glow glow-forest" aria-hidden="true"></div>

    <header class="site-header">
      <RouterLink to="/" class="brand">
        <span class="brand-mark">Le'</span>
        <div class="brand-title">
          <span class="brand-name">Blanc</span>
          <span class="brand-note">Maison culinaire</span>
        </div>
      </RouterLink>

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

    <section class="page-shell">
      <div class="page-meta">
        <p class="page-kicker">Élevé · Intime · Néobistro</p>
        <p class="page-intro">
          A luminous shell for each page—content from every route is framed inside the glass card below.
        </p>
      </div>
      <main class="page">
        <RouterView />
      </main>
    </section>
  </div>
</template>

<style scoped>
:global(:root) {
  --sand: #f4e8d8;
  --bone: #efe5d7;
  --ink: #0e0a07;
  --forest: #1a4a3e;
  --sage: #4f7c6d;
  --amber: #e7a44d;
  --card: rgba(255, 255, 255, 0.86);
  --border: #e1d5c3;
  --shadow: 0 24px 70px rgba(12, 10, 7, 0.18);
}

:global(html) {
  font-size: 16px;
}

:global(body) {
  margin: 0;
  font-family: 'Inter', 'Playfair Display', 'Segoe UI', system-ui, sans-serif;
  background: radial-gradient(1200px 1200px at 14% 18%, #fbeee0, #f5e6d6 45%, #ede0cf 70%, #e3d5c0 100%)
      fixed,
    linear-gradient(145deg, #fdf6ec 0%, #f2e6d7 45%, #e7d8c6 100%) fixed;
  color: var(--ink);
  font-size: 1rem;
  line-height: 1.6;
}

.site {
  min-height: 100vh;
  position: relative;
  overflow: hidden;
  padding: 22px 0 64px;
}

.glow {
  position: absolute;
  filter: blur(80px);
  opacity: 0.7;
  pointer-events: none;
}

.glow-amber {
  width: 420px;
  height: 420px;
  background: radial-gradient(circle, #f6c46d, transparent 60%);
  top: -120px;
  right: -80px;
}

.glow-forest {
  width: 360px;
  height: 360px;
  background: radial-gradient(circle, #3c806c, transparent 60%);
  bottom: -120px;
  left: -100px;
}

.site-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 18px;
  padding: 16px 20px;
  margin: 0 auto 24px;
  max-width: min(1200px, 94vw);
  border: 1px solid var(--border);
  background: linear-gradient(120deg, rgba(255, 255, 255, 0.94), rgba(253, 248, 240, 0.92))
    border-box;
  position: sticky;
  top: 18px;
  z-index: 5;
  box-shadow: var(--shadow);
  border-radius: 20px;
  backdrop-filter: blur(18px);
}

.brand {
  display: inline-flex;
  align-items: center;
  gap: 12px;
  text-decoration: none;
  color: var(--ink);
  font-weight: 800;
  letter-spacing: 0.01em;
}

.brand-mark {
  width: 42px;
  height: 42px;
  border-radius: 12px;
  background: linear-gradient(135deg, #1a4a3e, #2b7a66);
  color: #fff;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  font-size: 1.05rem;
  box-shadow: 0 12px 25px rgba(26, 74, 62, 0.38);
}

.brand-title {
  display: flex;
  flex-direction: column;
  line-height: 1.1;
}

.brand-name {
  font-size: 1.4rem;
}

.brand-note {
  font-size: 0.85rem;
  color: var(--sage);
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

.nav {
  display: flex;
  gap: 10px;
  align-items: center;
  flex-wrap: wrap;
  padding: 6px 8px;
  background: rgba(255, 255, 255, 0.7);
  border-radius: 999px;
  border: 1px solid var(--border);
  box-shadow: 0 12px 22px rgba(12, 10, 7, 0.06);
  backdrop-filter: blur(12px);
}

.nav-link {
  padding: 10px 14px;
  border-radius: 999px;
  text-decoration: none;
  color: var(--ink);
  font-weight: 600;
  font-size: 1rem;
  position: relative;
  transition: all 0.2s ease;
}

.nav-link:hover {
  background: rgba(26, 74, 62, 0.07);
  color: var(--forest);
}

.nav-link.active {
  background: rgba(26, 74, 62, 0.12);
  color: var(--forest);
  box-shadow: inset 0 0 0 1px rgba(26, 74, 62, 0.2);
}

.cta {
  padding: 11px 20px;
  border-radius: 14px;
  text-decoration: none;
  color: #fff;
  font-weight: 700;
  font-size: 1rem;
  background: linear-gradient(135deg, #1a4a3e, #2b7a66);
  box-shadow: 0 12px 28px rgba(26, 74, 62, 0.32);
  border: 1px solid rgba(26, 74, 62, 0.2);
}

.page-shell {
  max-width: min(1200px, 94vw);
  margin: 0 auto;
  display: grid;
  gap: 14px;
}

.page-meta {
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
  gap: 14px;
  flex-wrap: wrap;
  padding: 0 6px;
}

.page-kicker {
  margin: 0;
  font-size: 0.95rem;
  letter-spacing: 0.14em;
  text-transform: uppercase;
  color: var(--forest);
  font-weight: 700;
}

.page-intro {
  margin: 0;
  color: #574b3f;
  max-width: 560px;
  line-height: 1.5;
}

.page {
  padding: clamp(32px, 6vw, 64px);
  width: 100%;
  background: var(--card);
  border-radius: 28px;
  box-shadow: var(--shadow);
  border: 1px solid var(--border);
  backdrop-filter: blur(18px);
}

@media (max-width: 720px) {
  .site-header {
    flex-direction: column;
    align-items: flex-start;
    position: sticky;
    top: 12px;
  }
  .nav {
    gap: 8px;
    width: 100%;
    overflow-x: auto;
    padding: 8px;
  }
  .page-meta {
    align-items: flex-start;
    flex-direction: column;
  }
}
</style>
