<script setup>
import { computed, inject, onMounted, ref } from 'vue'
import { getDrinks } from '@/api'

const fallbackDrinks = [
  {
    _id: 'espresso',
    name: 'Espresso',
    desc: 'Shot of espresso with your choice of bean: arabica, robusta, moka, culi.',
    price: 35000,
    image: 'https://images.unsplash.com/photo-1510626176961-4b37d0b4e904?auto=format&fit=crop&w=900&q=80',
    caffeine: 'High',
    temp: 'Hot',
    sweetness: 1,
    kind: 'Coffee',
    isAlcoholic: false,
  },
  {
    _id: 'nau-da',
    name: 'Nâu đá',
    desc: 'Vietnamese iced milk coffee. Beans: arabica, robusta, moka, culi.',
    price: 40000,
    image: 'https://images.unsplash.com/photo-1509042239860-f550ce710b93?auto=format&fit=crop&w=900&q=80',
    caffeine: 'High',
    temp: 'Cold',
    sweetness: 3,
    kind: 'Coffee',
    isAlcoholic: false,
  },
  {
    _id: 'bac-siu',
    name: 'Bạc sỉu',
    desc: 'Saigon-style coffee heavy on milk. Bean choices: arabica, robusta, moka, culi.',
    price: 38000,
    image: 'https://images.unsplash.com/photo-1512568400610-62da28bc8a13?auto=format&fit=crop&w=900&q=80',
    caffeine: 'Med',
    temp: 'Hot',
    sweetness: 4,
    kind: 'Coffee',
    isAlcoholic: false,
  },
  {
    _id: 'latte',
    name: 'Latte',
    desc: 'Silky espresso with steamed milk. Add cốm, vani, sữa đặc, sữa kem, hoặc chuối.',
    price: 45000,
    image: 'https://images.unsplash.com/photo-1470337458703-46ad1756a187?auto=format&fit=crop&w=900&q=80',
    caffeine: 'Med',
    temp: 'Either',
    sweetness: 2,
    kind: 'Coffee',
    isAlcoholic: false,
  },
  {
    _id: 'mocha',
    name: 'Mocha',
    desc: 'Espresso, cocoa, and velvety milk. Bean options: arabica, robusta, moka, culi.',
    price: 48000,
    image: 'https://images.unsplash.com/photo-1459257868276-5e65389e2722?auto=format&fit=crop&w=900&q=80',
    caffeine: 'Med',
    temp: 'Either',
    sweetness: 3,
    kind: 'Coffee',
    isAlcoholic: false,
  },
  {
    _id: 'macchiato',
    name: 'Macchiato',
    desc: 'Espresso marked with microfoam. Flavor with vani, sữa kem, chuối, hoặc cốm.',
    price: 47000,
    image: 'https://images.unsplash.com/photo-1495474472287-4d71bcdd2085?auto=format&fit=crop&w=900&q=80',
    caffeine: 'Med',
    temp: 'Hot',
    sweetness: 1,
    kind: 'Coffee',
    isAlcoholic: false,
  },
  {
    _id: 'cappuccino',
    name: 'Cappuccino',
    desc: 'Balanced espresso, steamed milk, and foam. Beans: arabica, robusta, moka, culi.',
    price: 45000,
    image: 'https://images.unsplash.com/photo-1485808191679-5f86510681a2?auto=format&fit=crop&w=900&q=80',
    caffeine: 'Med',
    temp: 'Hot',
    sweetness: 2,
    kind: 'Coffee',
    isAlcoholic: false,
  },
]

const nightSpecials = [
  {
    _id: 'sangria',
    name: 'Sangria - Tình đơn phương',
    desc: 'Red wine, citrus, stone fruits; ngọt nhẹ như một mối tình đơn phương.',
    price: 120000,
    image: 'https://images.unsplash.com/photo-1504674900247-0877df9cc836?auto=format&fit=crop&w=900&q=80',
    caffeine: 'None',
    temp: 'Cold',
    sweetness: 3,
    kind: 'Cocktail',
    isAlcoholic: true,
    baseSpirit: 'Wine',
  },
  {
    _id: 'margarita',
    name: 'Margarita - Rối lòng',
    desc: 'Tequila, lime, salted rim—cú xoáy rối lòng nhưng đầy tỉnh táo.',
    price: 130000,
    image: 'https://images.unsplash.com/photo-1504674900247-0877df9cc836?auto=format&fit=crop&w=900&q=80&sat=-10',
    caffeine: 'None',
    temp: 'Cold',
    sweetness: 2,
    kind: 'Cocktail',
    isAlcoholic: true,
    baseSpirit: 'Tequila',
  },
  {
    _id: 'mojito',
    name: 'Mojito - Thanh lọc',
    desc: 'Rum, bạc hà, lime, soda—thanh lọc và sảng khoái.',
    price: 110000,
    image: 'https://images.unsplash.com/photo-1544145945-19cc90f9c6bf?auto=format&fit=crop&w=900&q=80',
    caffeine: 'None',
    temp: 'Cold',
    sweetness: 3,
    kind: 'Cocktail',
    isAlcoholic: true,
    baseSpirit: 'Rum',
  },
  {
    _id: 'pimms',
    name: 'Pimm’s - Ngậm ngùi',
    desc: 'Pimm’s, dưa leo, táo, gừng ale—ngậm ngùi nhưng dễ uống.',
    price: 120000,
    image: 'https://images.unsplash.com/photo-1504674900247-0877df9cc836?auto=format&fit=crop&w=900&q=80&sat=5',
    caffeine: 'None',
    temp: 'Cold',
    sweetness: 2,
    kind: 'Cocktail',
    isAlcoholic: true,
    baseSpirit: 'Pimm’s',
  },
  {
    _id: 'pina-colada',
    name: 'Piña Colada - Tức giận',
    desc: 'Rum, dứa, cream of coconut—tropical nhưng dư vị hơi “tức giận”.',
    price: 115000,
    image: 'https://images.unsplash.com/photo-1504674900247-0877df9cc836?auto=format&fit=crop&w=900&q=80&sat=-5',
    caffeine: 'None',
    temp: 'Cold',
    sweetness: 3,
    kind: 'Cocktail',
    isAlcoholic: true,
    baseSpirit: 'Rum',
  },
  {
    _id: 'caipirinha',
    name: 'Caipirinha - Ích kỉ tự trách',
    desc: 'Cachaça, lime, đường mía—sắc chua gắt như lúc tự trách.',
    price: 120000,
    image: 'https://images.unsplash.com/photo-1497534446932-c925b458314e?auto=format&fit=crop&w=900&q=80',
    caffeine: 'None',
    temp: 'Cold',
    sweetness: 2,
    kind: 'Cocktail',
    isAlcoholic: true,
    baseSpirit: 'Cachaça',
  },
  {
    _id: 'negroni',
    name: 'Negroni - Đắng cay',
    desc: 'Gin, Campari, vermouth đỏ—đắng cay nhưng sâu sắc.',
    price: 140000,
    image: 'https://images.unsplash.com/photo-1470337458703-46ad1756a187?auto=format&fit=crop&w=900&q=80&sat=-10',
    caffeine: 'None',
    temp: 'Cold',
    sweetness: 1,
    kind: 'Cocktail',
    isAlcoholic: true,
    baseSpirit: 'Gin',
  },
]

// Local fallbacks when a drink has no image URL from the API.
const placeholders = [
  'https://images.unsplash.com/photo-1447933601403-0c6688de566e?auto=format&fit=crop&w=600&q=80',
  'https://images.unsplash.com/photo-1497534446932-c925b458314e?auto=format&fit=crop&w=600&q=80',
  'https://images.unsplash.com/photo-1498804103079-a6351b050096?auto=format&fit=crop&w=600&q=80',
]

const theme = inject('theme', ref('day'))

const loading = ref(true)
const error = ref('')
const dayDrinks = ref([])
const nightDrinks = ref(nightSpecials)

const menuMode = computed(() => (theme?.value === 'night' ? 'night' : 'day'))
const currentMenu = computed(() => (menuMode.value === 'night' ? nightDrinks.value : dayDrinks.value))

const normalizeDrink = (drink, fallbackKind = 'Coffee') => ({
  ...drink,
  kind: drink.kind || fallbackKind,
  isAlcoholic: typeof drink.isAlcoholic === 'boolean' ? drink.isAlcoholic : fallbackKind !== 'Coffee',
})

const filterByTag = (items, tag) => items.filter((d) => (d.tags || d.Tags || []).some((t) => t?.toLowerCase() === tag))

const fetchDrinks = async () => {
  loading.value = true
  error.value = ''
  try {
    const res = await getDrinks()
    const list = Array.isArray(res) && res.length ? res : []
    const dayList = list.length ? filterByTag(list, 'day') : []
    const nightList = list.length ? filterByTag(list, 'night') : []

    dayDrinks.value = (dayList.length ? dayList : fallbackDrinks).map((d) => normalizeDrink(d, 'Coffee'))
    nightDrinks.value = (nightList.length ? nightList : nightSpecials).map((d) => normalizeDrink(d, 'Cocktail'))
  } catch (err) {
    error.value = err?.message || 'Khong the tai menu luc nay. Hien thi menu mau.'
    dayDrinks.value = fallbackDrinks.map((d) => normalizeDrink(d, 'Coffee'))
    nightDrinks.value = nightSpecials.map((d) => normalizeDrink(d, 'Cocktail'))
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
      <div class="hero-copy">
        <p class="kicker" :class="{ night: menuMode === 'night' }">
          {{ menuMode === 'night' ? 'Nightfall Mixology' : 'The Creative Coffee Menu' }}
        </p>
        <h1 class="hero-title">
          {{ menuMode === 'night' ? 'Mocktails, cocktails, and spirited pours.' : 'Thoughtfully brewed for every mood.' }}
        </h1>
        <p class="lede">
          {{
            menuMode === 'night'
              ? 'After dark we shift into bar service with low-ABV mocktails, signature cocktails, and spirited coffee riffs.'
              : 'Every cup from our Leblanc coffee lab balances smooth texture, restorative botanicals, and mindful service.'
          }}
        </p>
      </div>
      <div class="mode-banner hero-card" :class="menuMode">
        <div class="mode-label">
          <span class="dot"></span>
          {{ menuMode === 'night' ? 'Night mode: mocktail, cocktail & boozy picks' : 'Day mode: coffee & non-alcohol' }}
        </div>
        <small>Use the Day/Night toggle in the top right to see the matching menu.</small>
      </div>
    </div>

    <div v-if="loading && menuMode === 'day'" class="state">Dang tai menu...</div>
    <div v-else-if="menuMode === 'day' && !currentMenu.length" class="state error">
      {{ error || 'Khong co du lieu menu.' }}
    </div>
    <div v-else>
      <div v-if="menuMode === 'day' && error" class="state error">{{ error }}</div>
      <div class="grid">
        <article v-for="(drink, idx) in currentMenu" :key="drink._id || drink.name" class="card">
          <img :src="imgForIndex(drink, idx)" :alt="drink.name" loading="lazy" />
          <div>
            <div class="card-top">
              <div class="title-wrap">
                <h3>{{ drink.name }}</h3>
                <span class="pill" :class="drink.kind ? drink.kind.toLowerCase().replace(/\s+/g, '-') : ''">
                  {{ drink.kind || (menuMode === 'night' ? 'Night' : 'Day') }}
                </span>
              </div>
              <span>{{ formatCurrency(drink.price) }}</span>
            </div>
            <p class="desc">{{ drink.desc }}</p>
            <ul>
              <li v-if="drink.isAlcoholic !== undefined">Alcohol: {{ drink.isAlcoholic ? 'Co' : 'Khong' }}</li>
              <li v-if="drink.baseSpirit">Base: {{ drink.baseSpirit }}</li>
              <li>Caffeine: {{ drink.caffeine }}</li>
              <li>Temp: {{ drink.temp }}</li>
              <li>Sweetness: {{ drink.sweetness }}/5</li>
            </ul>
          </div>
        </article>
      </div>
    </div>
  </section>
</template>

<style scoped>
.menu-page {
  display: flex;
  flex-direction: column;
  gap: 32px;
  color: var(--ink);
  background: var(--paper);
  padding: 32px;
  border: none;
}

.panel {
  padding: clamp(20px, 3vw, 48px) 0;
  border: none;
}

.heading {
  display: grid;
  grid-template-columns: minmax(0, 1.4fr) minmax(280px, 0.9fr);
  align-items: center;
  gap: 24px 32px;
}

.kicker {
  text-transform: uppercase;
  letter-spacing: 0.18em;
  font-size: 1.05rem;
  color: #5b4635;
  font-family: 'Georgia', 'Times New Roman', serif;
  font-style: italic;
  font-weight: 700;
}

:global(.theme-night) .kicker {
  color: #fff !important;
}

.kicker.night {
  color: #fff;
}

.hero-copy {
  display: flex;
  flex-direction: column;
  gap: 14px;
  max-width: 980px;
}

.hero-title {
  font-size: clamp(2.6rem, 3vw + 1.4rem, 4.2rem);
  line-height: 1.08;
  margin: 0;
}

.lede {
  margin: 0;
  font-size: 1.08rem;
  max-width: 820px;
}

.mode-banner {
  display: flex;
  flex-direction: column;
  gap: 10px;
  align-items: flex-start;
  padding: 18px 20px;
  border: 1px solid var(--cream-strong);
  border-radius: 14px;
  background: var(--paper);
  box-shadow: 0 16px 32px rgba(0, 0, 0, 0.06);
}

.mode-banner.night {
  border-color: #243145;
  background: #101826;
  color: #f5f1e8;
}

.hero-card {
  min-width: 280px;
  max-width: 420px;
  justify-self: end;
}

.mode-label {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  font-weight: 800;
  letter-spacing: 0.04em;
  font-size: 1rem;
}

.mode-banner small {
  opacity: 0.78;
  font-size: 0.95rem;
}

.dot {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  background: var(--tan);
  box-shadow: 0 0 0 6px rgba(216, 163, 90, 0.18);
}

.grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
  gap: 26px;
}

.card {
  border: 1px solid #e7ddcf;
  background: var(--paper);
  display: flex;
  flex-direction: column;
  box-shadow: 0 18px 40px rgba(0, 0, 0, 0.08);
  border-radius: 16px;
  overflow: hidden;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.card:hover {
  transform: translateY(-6px);
  box-shadow: 0 24px 60px rgba(0, 0, 0, 0.18);
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
  color: var(--ink);
  gap: 16px;
}

.title-wrap {
  display: inline-flex;
  gap: 10px;
  align-items: center;
  flex-wrap: wrap;
}

.pill {
  padding: 6px 10px;
  border-radius: 999px;
  background: #f3ede4;
  border: 1px solid #e1d4c3;
  font-size: 0.7rem;
  letter-spacing: 0.04em;
  text-transform: uppercase;
  font-weight: 800;
}

.pill.mocktail {
  background: #e9f7ef;
  border-color: #c1e9d4;
  color: #0f6b3f;
}

.pill.cocktail {
  background: #fef3e6;
  border-color: #f2d3a2;
  color: #9b5b14;
}

.pill.coffee {
  background: #f3ede4;
  border-color: #e1d4c3;
  color: #5b4635;
}

.desc {
  margin: 0;
}

ul {
  padding-left: 18px;
  margin: 0;
}

.state {
  padding: 36px;
  text-align: center;
  background: var(--paper);
  border-radius: 12px;
  border: 1px solid var(--cream-strong);
}

.state.error {
  color: var(--orange-strong);
  background: #ffe8dd;
}

@media (max-width: 900px) {
  .heading {
    grid-template-columns: 1fr;
  }

  .mode-banner {
    width: 100%;
  }
}
</style>
