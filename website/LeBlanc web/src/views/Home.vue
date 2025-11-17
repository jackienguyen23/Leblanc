<script setup>
import { RouterLink } from 'vue-router'

const highlights = [
  {
    title: 'Signature Drinks',
    desc: 'Hand-crafted specialty drinks tailored to your mood.',
    image:
      'https://images.unsplash.com/photo-1504674900247-0877df9cc836?auto=format&fit=crop&w=700&q=80',
  },
  {
    title: 'Mood-Based Coffee',
    desc: 'Emotion-driven experiences in a cozy modern space.',
    image:
      'https://images.unsplash.com/photo-1470337458703-46ad1756a187?auto=format&fit=crop&w=700&q=80',
  },
  {
    title: 'Innovative Beverages',
    desc: 'Novelty and quality far beyond ordinary cafés.',
    image:
      'https://images.unsplash.com/photo-1497534446932-c925b458314e?auto=format&fit=crop&w=700&q=80',
  },
]
</script>

<template>
  <div class="home">
    <section class="panel hero">
      <div class="hero-copy">
        <p class="kicker">Bridge Mood & Taste through Healing Coffee</p>
        <h1>
          Bridge <em>Mood</em> & Taste through <span>Healing Coffee</span>
        </h1>
        <RouterLink to="/mood-booker" class="ghost-btn">Read More</RouterLink>
      </div>
      <div class="hero-media">
        <img
          src="https://images.unsplash.com/photo-1470337458703-46ad1756a187?auto=format&fit=crop&w=900&q=80"
          alt="LeBlanc interior"
        />
      </div>
    </section>

    <section class="panel about">
      <div>
        <p class="kicker">About Le'Blanc</p>
        <h2>Welcome to Le'Blanc</h2>
        <p>
          Our vision is to become Vietnam’s most creative coffee atelier, where every detail is sensory. We craft
          uplifting customer journeys by blending art and tea with slow brewing, high-quality, restorative beverages in a
          warm, inspiring atmosphere that connects with customers’ moods.
        </p>
        <RouterLink to="/menu" class="ghost-btn">Read More</RouterLink>
      </div>
      <div class="about-media">
        <img
          src="https://images.unsplash.com/photo-1511920170033-f8396924c348?auto=format&fit=crop&w=900&q=80"
          alt="Barista crafting drinks"
        />
      </div>
    </section>

    <section class="panel menu-preview">
      <div class="menu-header">
        <p class="kicker">The Creative Coffee Menu</p>
        <h2>The Creative Coffee Menu</h2>
        <RouterLink to="/menu" class="ghost-btn">View Menu</RouterLink>
      </div>
      <div class="highlight-grid">
        <article v-for="item in highlights" :key="item.title" class="highlight-card">
          <img :src="item.image" :alt="item.title" />
          <div class="text">
            <h3>{{ item.title }}</h3>
            <p>{{ item.desc }}</p>
          </div>
        </article>
      </div>
    </section>

    <section class="panel contact" id="contact">
      <div>
        <p class="kicker">Contact & Make Reservation</p>
        <h2>Contact & Make Reservation</h2>
        <p>
          Visit us at Le'Blanc and experience coffee that heals a story. We are here to brighten your day with every
          cup.
        </p>
      </div>
      <div class="details">
        <div>
          <strong>Location</strong>
          <p>Ho Chi Minh City, Vietnam</p>
        </div>
        <div>
          <strong>Hours</strong>
          <p>Mon - Sun: 7:00 AM - 02:00 PM</p>
        </div>
      </div>
      <RouterLink to="/booking" class="ghost-btn">Reserve Now</RouterLink>
    </section>
  </div>
</template>

<style scoped>
.home {
  display: flex;
  flex-direction: column;
  gap: 60px;
  color: #1c0f05;
}

.panel {
  padding: clamp(48px, 6vw, 82px);
  background: #f7f0e4;
  border-radius: 36px;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
  gap: 32px;
  border: 1px solid rgba(28, 15, 5, 0.08);
}

.hero,
.about {
  align-items: center;
}

.hero-copy h1 {
  font-size: clamp(2.6rem, 5.8vw, 4.4rem);
  font-weight: 700;
  margin-bottom: 28px;
}

.hero-copy em {
  font-style: italic;
}

.hero-copy span {
  text-decoration: underline;
}

.kicker {
  text-transform: uppercase;
  font-size: 1rem;
  letter-spacing: 0.32em;
  color: #7a6046;
  margin-bottom: 24px;
}

.hero-media img,
.about-media img,
.highlight-card img {
  width: 100%;
  display: block;
  border-radius: 4px;
  object-fit: cover;
  min-height: 220px;
}

.about-media img {
  min-height: 260px;
}

.ghost-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 16px 36px;
  border: 2px solid #1c0f05;
  border-radius: 999px;
  color: #1c0f05;
  text-decoration: none;
  font-weight: 600;
  width: fit-content;
  font-size: 1.05rem;
}

.menu-preview {
  grid-template-columns: 1fr;
}

.menu-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 16px;
}

.highlight-grid {
  margin-top: 24px;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
  gap: 26px;
}

.highlight-card {
  border: none;
  background: #fff;
  display: flex;
  flex-direction: column;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.08);
  border-radius: 16px;
  overflow: hidden;
}

.highlight-card .text {
  padding: 16px;
}

.contact {
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
}

.details {
  display: grid;
  gap: 12px;
}
</style>
