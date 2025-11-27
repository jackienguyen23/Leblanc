<script setup>
import { RouterLink } from 'vue-router'
</script>

<template>
  <div class="page">
    <section class="hero">
      <div class="hero-media hero-coffee" aria-hidden="true"></div>
      <div class="hero-blur" aria-hidden="true"></div>
      <div class="hero-overlay"></div>
      <div class="hero-body">
        <p class="hero-kicker">Le'Blanc Mix Artistry</p>
        <h1 class="hero-title">Mood-first drinks, made for clear nights and bright mornings.</h1>
        <p class="hero-lede">
          Crafted mocktails and botanical coffees that keep you <strong>centered, creative, and clear</strong> without
          losing the ritual.
        </p>
        <div class="hero-actions">
          <RouterLink to="/menu" class="btn-oval">Explore the menu</RouterLink>
          <RouterLink to="/booking" class="btn-ghost">Book a table</RouterLink>
        </div>
      </div>
    </section>

    <section class="feature">
      <div class="feature-image" aria-hidden="true"></div>
      <div class="feature-inner">
        <span class="feature-kicker">NEW</span>
        <h2 class="feature-title">PRIMAVERA</h2>
        <p class="feature-sub">A BOLD, SPICED APERITIF FEATURING CARDAMOM, GINGER, AND CINNAMON.</p>
      </div>
    </section>

    <footer class="footer">
      <div class="footer-logo">LE'BLANC</div>
      <div class="footer-bottom">
        <span>Find stockists</span>
        <span>hello@leblanc.ai</span>
        <span>(646) 555-4567</span>
      </div>
    </footer>
  </div>
</template>

<style scoped>
.page {
  display: grid;
  gap: 0;
  padding-bottom: 64px;
}

.hero {
  width: 100vw;
  margin-left: calc(50% - 50vw);
  margin-right: calc(50% - 50vw);
  position: relative;
  border-radius: 0;
  overflow: hidden;
  min-height: 480px;
  box-shadow: none;
}

.hero-media {
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  background-size: cover;
  background-repeat: no-repeat;
  background-position: center;
  z-index: 0;
}

.hero-coffee {
  background-image: url('https://images.unsplash.com/photo-1509042239860-f550ce710b93?auto=format&fit=crop&w=1600&q=80');
}

.hero-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(120deg, rgba(12, 12, 12, 0.82) 0%, rgba(12, 12, 12, 0.42) 48%, rgba(12, 12, 12, 0.18) 70%);
  mix-blend-mode: multiply;
  z-index: 2;
}

.hero-blur {
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  background-image: url('https://images.unsplash.com/photo-1509042239860-f550ce710b93?auto=format&fit=crop&w=1600&q=80');
  background-size: cover;
  background-position: center;
  filter: blur(10px);
  opacity: 0.78;
  z-index: 1;
  pointer-events: none;
  -webkit-mask-image: linear-gradient(90deg, #000 0%, #000 56%, rgba(0, 0, 0, 0) 74%);
  mask-image: linear-gradient(90deg, #000 0%, #000 56%, rgba(0, 0, 0, 0) 74%);
}

.hero-body {
  position: relative;
  z-index: 3;
  display: grid;
  gap: 16px;
  padding: clamp(28px, 6vw, 64px);
  max-width: 820px;
  color: #f8f5ed;
}

.hero-kicker {
  text-transform: uppercase;
  letter-spacing: 0.22em;
  font-size: 0.9rem;
  margin: 0;
  font-family: 'Georgia', 'Times New Roman', serif;
  font-style: italic;
  color: #e9d7b6;
}

.hero-title {
  margin: 0;
  font-size: clamp(2.5rem, 3vw + 1rem, 3.6rem);
  line-height: 1.1;
  font-weight: 800;
  letter-spacing: 0.01em;
}

.hero-lede {
  margin: 0;
  font-size: 1.1rem;
  line-height: 1.6;
  color: #f1e9de;
}

.hero-actions {
  display: flex;
  gap: 26px;
  flex-wrap: wrap;
  margin-top: 10px;
  margin-bottom: 6px;
}

.btn-oval {
  padding: 14px 22px;
  background: var(--tan);
  color: #0b0b0b;
  border-radius: 999px;
  text-decoration: none;
  font-weight: 900;
  letter-spacing: 0.02em;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.16);
  transform: scale(1.1);
  transition: transform 0.18s ease, box-shadow 0.18s ease;
}

.btn-ghost {
  padding: 14px 22px;
  border-radius: 999px;
  border: 1px solid rgba(255, 255, 255, 0.68);
  color: #f8f5ed;
  text-decoration: none;
  font-weight: 800;
  letter-spacing: 0.02em;
  backdrop-filter: blur(2px);
  transform: scale(1.1);
  transition: transform 0.18s ease, box-shadow 0.18s ease, background-color 0.18s ease;
}

.btn-oval:hover,
.btn-ghost:hover {
  transform: scale(1.18);
  box-shadow: 0 16px 34px rgba(0, 0, 0, 0.22);
}

.btn-ghost {
  padding: 12px 20px;
  border-radius: 999px;
  border: 1px solid rgba(255, 255, 255, 0.68);
  color: #f8f5ed;
  text-decoration: none;
  font-weight: 800;
  letter-spacing: 0.02em;
  backdrop-filter: blur(2px);
}

.feature {
  background: #000;
  color: #fff;
  padding: 28px 18px 32px;
  max-width: 100vw;
  margin: 0 calc(50% - 50vw);
  text-align: center;
  display: grid;
  gap: 18px;
  border-radius: 0;
}

.grid {
  max-width: min(1600px, 98vw);
  margin: 28px auto 0;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
  gap: 24px;
}

.feature-image {
  height: 420px;
  background: radial-gradient(circle at 50% 20%, rgba(255, 255, 255, 0.06), transparent 40%),
    url('@/assets/primavera.jpg') center/contain no-repeat;
}

.feature-inner {
  display: grid;
  gap: 12px;
  margin-top: -12px;
}

.feature-kicker {
  letter-spacing: 0.2em;
  text-transform: uppercase;
  font-size: 0.8rem;
}

.feature-title {
  margin: 0;
  font-size: 2.6rem;
  letter-spacing: 0.02em;
}

.feature-sub {
  margin: 0;
  font-size: 1.1rem;
  font-weight: 700;
}

.grid-card {
  display: grid;
  gap: 10px;
  background: #fff;
  border-radius: 6px;
  padding: 14px;
  box-shadow: var(--shadow);
}

.grid-label {
  margin: 0;
  font-weight: 900;
  letter-spacing: 0.12em;
  text-transform: uppercase;
  font-size: 0.95rem;
}

.grid-art {
  height: 200px;
  border-radius: 6px;
  position: relative;
  overflow: hidden;
}

.grid-art::after {
  content: '';
  position: absolute;
  inset: 12px;
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 8px;
}

.lavande {
  background: url('@/assets/lavande.jpg') center/cover no-repeat;
}

.primavera {
  background: url('@/assets/primavera.jpg') center/cover no-repeat;
}

.spezia {
  background: url('@/assets/spezia.jpg') center/cover no-repeat;
}

.grid-desc {
  margin: 0;
  font-size: 0.95rem;
  color: var(--ink);
}

.grid-link {
  text-decoration: underline;
  font-weight: 800;
  color: var(--dark);
}

.find {
  max-width: min(1600px, 98vw);
  margin: 0 auto;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
  gap: 22px;
  padding: 32px;
  background: linear-gradient(120deg, #c39452, #b17c36);
  border-radius: 10px;
  box-shadow: var(--shadow);
  color: var(--ink);
}

.find-left {
  display: grid;
  gap: 8px;
}

.find-words {
  margin: 0;
  font-size: clamp(2.2rem, 4vw, 2.8rem);
  letter-spacing: 0.05em;
  text-transform: uppercase;
  font-weight: 900;
  line-height: 1;
}

.find-title {
  margin: 0;
  font-size: clamp(2rem, 3.4vw, 2.6rem);
  letter-spacing: 0.02em;
  text-transform: uppercase;
}

.find-body {
  margin: 0;
  font-size: 1rem;
  font-weight: 700;
}

.find-actions {
  display: flex;
  gap: 12px;
  flex-wrap: wrap;
  margin-top: 8px;
}

.btn-dark {
  padding: 12px 16px;
  border-radius: 8px;
  background: var(--dark);
  color: #fff;
  text-decoration: none;
  font-weight: 800;
}

.btn-link {
  color: var(--dark);
  text-decoration: underline;
  font-weight: 800;
}

.find-right {
  min-height: 360px;
  border-radius: 8px;
  background: url('@/assets/find.jpg') center/cover no-repeat;
  position: relative;
  overflow: hidden;
}

.footer {
  max-width: min(1600px, 98vw);
  margin: 0 auto;
  padding: 18px;
  display: grid;
  gap: 10px;
  color: var(--ink);
}

.footer-logo {
  font-size: 1.6rem;
  font-weight: 900;
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

.checker-sm {
  height: 10px;
  width: 100%;
  background-image: repeating-linear-gradient(
    90deg,
    var(--dark) 0,
    var(--dark) 14px,
    transparent 14px,
    transparent 28px
  );
}

.footer-bottom {
  display: flex;
  gap: 12px;
  font-size: 0.95rem;
  letter-spacing: 0.05em;
  text-transform: uppercase;
  flex-wrap: wrap;
}
</style>
