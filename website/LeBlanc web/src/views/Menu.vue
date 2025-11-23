<script setup>
import { ref, onMounted } from 'vue'
import { getDrinks } from '@/api'

const drinks = ref([])
const loading = ref(false)
const error = ref('')

// Local fallbacks when a drink has no image URL from the API.
const placeholders = [
  'https://images.unsplash.com/photo-1447933601403-0c6688de566e?auto=format&fit=crop&w=600&q=80',
  'https://images.unsplash.com/photo-1497534446932-c925b458314e?auto=format&fit=crop&w=600&q=80',
  'https://images.unsplash.com/photo-1498804103079-a6351b050096?auto=format&fit=crop&w=600&q=80',
]

const fetchDrinks = async () => {
  loading.value = true
  error.value = ''
  try {
    drinks.value = await getDrinks()
  } catch (err) {
    error.value = err?.message || 'Khong the tai menu luc nay.'
  } finally {
    loading.value = false
  }
}

const imgForIndex = (drink, idx) => drink.image || placeholders[idx % placeholders.length]
// Prices arrive as integers; format to Vietnamese locale currency.
const formatCurrency = (value) => `${(Number(value) || 0).toLocaleString('vi-VN')} VND`

onMounted(fetchDrinks)
</script>

<template>
  <section class="menu-page">
    <div class="panel heading">
      <div>
        <p class="kicker">The Creative Coffee Menu</p>
        <h1>Thoughtfully brewed for every mood.</h1>
      </div>
      <p>
        Every cup from our Leblanc coffee lab balances smooth texture, restorative botanicals, and mindful service.
      </p>
    </div>

    <div v-if="loading" class="state">Dang tai menu...</div>
    <div v-else-if="error" class="state error">{{ error }}</div>
    <div v-else class="grid">
      <article v-for="(drink, idx) in drinks" :key="drink._id || drink.name" class="card">
        <img :src="imgForIndex(drink, idx)" :alt="drink.name" loading="lazy" />
        <div>
          <div class="card-top">
            <h3>{{ drink.name }}</h3>
            <span>{{ formatCurrency(drink.price) }}</span>
          </div>
          <p>{{ drink.desc }}</p>
          <ul>
            <li>Caffeine: {{ drink.caffeine }}</li>
            <li>Temp: {{ drink.temp }}</li>
            <li>Sweetness: {{ drink.sweetness }}/5</li>
          </ul>
        </div>
      </article>
    </div>
  </section>
</template>

<style scoped>
.menu-page {
  display: flex;
  flex-direction: column;
  gap: 40px;
  color: var(--ink);
  background: linear-gradient(180deg, var(--cream) 0%, #ffffff 100%);
  border-radius: 24px;
  padding: 32px;
  border: 1px solid var(--cream-strong);
}

.panel {
  padding: clamp(36px, 5vw, 70px) 0;
  border: none;
}

.heading {
  display: flex;
  justify-content: space-between;
  gap: 24px;
  flex-wrap: wrap;
}

.kicker {
  text-transform: uppercase;
  letter-spacing: 0.18em;
  font-size: 0.8rem;
  color: #5b4635;
}

.grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
  gap: 26px;
}

.card {
  border: 1px solid #e7ddcf;
  background: linear-gradient(180deg, #ffffff 0%, #fff8f1 100%);
  display: flex;
  flex-direction: column;
  box-shadow: 0 18px 40px rgba(0, 0, 0, 0.08);
  border-radius: 16px;
  overflow: hidden;
}

.card img {
  width: 100%;
  height: 240px;
  object-fit: cover;
}

.card div {
  padding: 24px;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.card-top {
  display: flex;
  justify-content: space-between;
  font-weight: 600;
  color: var(--blue);
}

ul {
  padding-left: 18px;
  margin: 0;
}

.state {
  padding: 36px;
  text-align: center;
  background: linear-gradient(120deg, var(--blue-soft), #f7f4ef);
  border-radius: 12px;
}

.state.error {
  color: var(--orange-strong);
  background: #ffe8dd;
}
</style>
