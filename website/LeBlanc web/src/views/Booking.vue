<script setup>
import { computed, inject, onBeforeUnmount, onMounted, ref, watch } from 'vue'
import { createBooking, getDrinks, recoFromFeatures } from '@/api'
import { isBookingEmailReady, sendBookingEmail } from '@/email'

const form = ref({
  name: '',
  phone: '',
  email: '',
  time: '',
  guests: 2,
})
const formDate = ref('')
const formClock = ref('')

const mood = ref('happy')
const caffeinePref = ref('')
const tempPref = ref('')
const sweetness = ref(5)
const nightType = ref('')
const nightBase = ref('')

const theme = inject('theme', ref('day'))
const isNight = computed(() => theme?.value === 'night')

const loadUser = () => {
  try {
    const raw = localStorage.getItem('leblancUser')
    return raw ? JSON.parse(raw) : null
  } catch (err) {
    console.warn('Could not parse stored user', err)
    return null
  }
}

const userPrefilled = ref(false)
const applyUser = (u) => {
  if (!u) return
  if (!userPrefilled.value || !form.value.name) form.value.name = u.name || form.value.name
  if (!userPrefilled.value || !form.value.email) form.value.email = u.email || form.value.email
  userPrefilled.value = true
}
const handleUserUpdated = (event) => applyUser(event?.detail)

const drinks = ref([])
const reco = ref([])
const selection = ref({})

const bookingLoading = ref(false)
const bookingOk = ref(false)
const bookingError = ref('')
const bookingEmailSent = ref(false)
const bookingEmailError = ref('')

const recoLoading = ref(false)
const recoError = ref('')

const canSubmit = computed(
  () => form.value.name && form.value.phone && form.value.email && formDate.value && formClock.value && form.value.time,
)
const bookingEmailReady = computed(() => isBookingEmailReady())

watch([formDate, formClock], ([date, clock]) => {
  if (date && clock) {
    const local = new Date(`${date}T${clock}`)
    if (!Number.isNaN(local.getTime())) {
      form.value.time = local.toISOString()
      return
    }
  }
  form.value.time = ''
})

const moodToEmotionFit = (val) => {
  switch (val) {
    case 'calm':
      return { calm: 0.9, happy: 0.4, stressed: 0.2, sad: 0.3, adventurous: 0.3 }
    case 'stressed':
      return { calm: 0.2, happy: 0.3, stressed: 0.9, sad: 0.2, adventurous: 0.3 }
    case 'sad':
      return { calm: 0.3, happy: 0.2, stressed: 0.2, sad: 0.9, adventurous: 0.3 }
    case 'adventurous':
      return { calm: 0.3, happy: 0.6, stressed: 0.3, sad: 0.2, adventurous: 0.9 }
    default:
      return { calm: 0.3, happy: 0.9, stressed: 0.2, sad: 0.2, adventurous: 0.4 }
  }
}

const fetchDrinks = async () => {
  try {
    drinks.value = await getDrinks()
  } catch (err) {
    console.warn('Could not load drinks', err)
  }
}

const resolveDrink = (id) => drinks.value.find((d) => d._id === id)

const fetchReco = async () => {
  recoLoading.value = true
  recoError.value = ''
  try {
    const result = await recoFromFeatures({
      emotionFit: moodToEmotionFit(mood.value),
      caffeine: caffeinePref.value || undefined,
      temp: tempPref.value || undefined,
      sweetness: sweetness.value,
    })
    // result may be array of {drinkId, score}; enrich with drink info if available
    let mapped = (result || []).map((item) => {
      const drink = resolveDrink(item.drinkId) || {}
      return {
        ...drink,
        drinkId: item.drinkId || drink._id,
        score: item.score,
      }
    })
    if (isNight.value) {
      const typeTag = nightType.value?.toLowerCase()
      const baseTag = nightBase.value?.toLowerCase()
      mapped = mapped.filter((item) => {
        const tags = (item.tags || item.Tags || []).map((t) => (t || '').toLowerCase())
        const okType = !typeTag || tags.includes(typeTag)
        const okBase = !baseTag || tags.includes(baseTag)
        return okType && okBase
      })
    }
    reco.value = mapped
  } catch (err) {
    recoError.value = err?.message || 'Không thể gợi ý lúc này.'
  } finally {
    recoLoading.value = false
  }
}

const addDrink = (drink) => {
  if (!drink?.drinkId && !drink?._id) return
  const id = drink.drinkId || drink._id
  const current = selection.value[id]?.qty || 0
  selection.value = {
    ...selection.value,
    [id]: { drink, qty: current + 1 },
  }
}

const updateQty = (id, delta) => {
  const current = selection.value[id]?.qty || 0
  const next = Math.max(0, current + delta)
  if (next === 0) {
    const { [id]: _, ...rest } = selection.value
    selection.value = rest
    return
  }
  selection.value = {
    ...selection.value,
    [id]: { drink: selection.value[id]?.drink, qty: next },
  }
}

const selectedItems = computed(() =>
  Object.entries(selection.value).map(([drinkId, entry]) => ({
    drinkId,
    drink: entry.drink || resolveDrink(drinkId) || {},
    qty: entry.qty,
  })),
)

const totalItems = computed(() =>
  selectedItems.value.reduce((sum, item) => sum + (item.qty || 0), 0),
)

const book = async () => {
  if (!canSubmit.value || bookingLoading.value) return
  bookingLoading.value = true
  bookingError.value = ''
  bookingOk.value = false
  bookingEmailSent.value = false
  bookingEmailError.value = ''
  try {
    const items = selectedItems.value.map((item) => ({
      drinkId: item.drinkId,
      qty: item.qty,
      options: {},
    }))
    const payload = {
      ...form.value,
      items,
      channel: 'web',
    }
    const res = await createBooking(payload)
    bookingOk.value = Boolean(res?.ok || res?._id)
    if (bookingOk.value) {
      if (form.value.email && bookingEmailReady.value) {
        const emailItems = selectedItems.value.map((item) => ({
          drinkId: item.drinkId,
          name: item.drink?.name || 'Drink',
          qty: item.qty,
        }))
        const bookingForEmail = {
          ...payload,
          items: emailItems,
          bookingId: res?.id || res?._id || '',
        }
        sendBookingEmail(bookingForEmail)
          .then(() => {
            bookingEmailSent.value = true
          })
          .catch((err) => {
            bookingEmailError.value = err?.message || 'Gửi email xác nhận thất bại.'
            console.warn('Booking email failed', err)
          })
      }
      selection.value = {}
    }
  } catch (err) {
    bookingError.value = err?.message || 'Không thể đặt lúc này.'
  } finally {
    bookingLoading.value = false
  }
}

onMounted(() => {
  fetchDrinks()
  fetchReco()
  applyUser(loadUser())
  window.addEventListener('leblanc-user-updated', handleUserUpdated)
})

onBeforeUnmount(() => {
  window.removeEventListener('leblanc-user-updated', handleUserUpdated)
})
</script>

<template>
  <section class="booking">
    <div class="panel form">
      <p class="eyebrow">Reservation</p>
      <h1>Book your table</h1>
      <p class="lede">Giữ chỗ và thêm đồ uống trước nếu bạn muốn.</p>

      <form class="form-fields" @submit.prevent="book">
        <label>
          Name
          <input v-model="form.name" placeholder="Nguyễn Văn A" />
        </label>
        <label>
          Phone
          <input v-model="form.phone" placeholder="0123 456 789" />
        </label>
        <label>
          Email
          <input v-model="form.email" type="email" placeholder="name@gmail.com" />
        </label>
        <label>
          Arrival Date
          <input v-model="formDate" type="date" />
        </label>
        <label>
          Arrival Time
          <input v-model="formClock" type="time" step="900" />
        </label>
        <label>
          Guests
          <input v-model.number="form.guests" type="number" min="1" max="10" />
        </label>

        <div class="selected" v-if="selectedItems.length">
          <p class="mini-title">Pre-order drinks ({{ totalItems }} items)</p>
          <div class="chip-list">
            <div v-for="item in selectedItems" :key="item.drinkId" class="chip">
              <span>{{ item.drink?.name || 'Drink' }}</span>
              <div class="qty">
                <button type="button" @click="updateQty(item.drinkId, -1)">-</button>
                <span>{{ item.qty }}</span>
                <button type="button" @click="updateQty(item.drinkId, 1)">+</button>
              </div>
            </div>
          </div>
        </div>

        <button type="submit" :disabled="!canSubmit || bookingLoading">
          <span v-if="bookingLoading">Processing...</span>
          <span v-else>Book table{{ totalItems ? ' & drinks' : '' }}</span>
        </button>
        <p v-if="bookingOk" class="status success">Đặt bàn thành công! Chúng tôi sẽ liên hệ xác nhận.</p>
        <p v-if="bookingOk && bookingEmailSent" class="status success">Email xác nhận đã gửi tới: {{ form.email }}</p>
        <p v-if="bookingOk && bookingEmailError" class="status error">Đặt bàn thành công nhưng gửi email thất bại: {{ bookingEmailError }}</p>
        <p v-if="bookingError && !bookingOk" class="status error">{{ bookingError }}</p>
      </form>
    </div>

    <div class="panel reco">
      <div class="reco-head">
        <div>
          <p class="eyebrow">Mood-booker</p>
          <h2>Gợi ý đồ uống</h2>
        </div>
        <button class="ghost" type="button" :disabled="recoLoading" @click="fetchReco">
          {{ recoLoading ? 'Đang gợi ý...' : 'Làm mới gợi ý' }}
        </button>
      </div>

      <div class="controls">
        <label>
          Mood
          <select v-model="mood">
            <option value="happy">Happy</option>
            <option value="calm">Calm</option>
            <option value="stressed">Stressed</option>
            <option value="sad">Sad</option>
            <option value="adventurous">Adventurous</option>
          </select>
        </label>
        <template v-if="!isNight">
          <label>
            Caffeine
            <select v-model="caffeinePref">
              <option value="">Any</option>
              <option value="high">High</option>
              <option value="med">Medium</option>
              <option value="low">Low</option>
              <option value="none">None</option>
            </select>
          </label>
          <label>
            Temperature
            <select v-model="tempPref">
              <option value="">Any</option>
              <option value="hot">Hot</option>
              <option value="iced">Iced</option>
              <option value="cold">Cold</option>
            </select>
          </label>
          <label>
            Sweetness: {{ sweetness }}
            <input v-model.number="sweetness" type="range" min="1" max="10" />
          </label>
        </template>
        <template v-else>
          <label>
            Drink type
            <select v-model="nightType">
              <option value="">Any</option>
              <option value="cocktail">Cocktail</option>
              <option value="wine">Wine</option>
              <option value="beer">Beer</option>
              <option value="liqueur">Liqueur</option>
              <option value="coffee">Coffee</option>
            </select>
          </label>
          <label>
            Base
            <select v-model="nightBase">
              <option value="">Any</option>
              <option value="gin">Gin</option>
              <option value="rum">Rum</option>
              <option value="whisky">Whisky</option>
              <option value="wine">Wine</option>
              <option value="beer">Beer</option>
              <option value="liqueur">Liqueur</option>
              <option value="coffee">Coffee</option>
              <option value="signature">Signature</option>
            </select>
          </label>
        </template>
      </div>

      <p v-if="recoError" class="status error">{{ recoError }}</p>
      <div v-else class="reco-list">
        <div v-if="recoLoading" class="status">Đang gợi ý...</div>
        <template v-else>
          <div v-for="drink in reco" :key="drink.drinkId || drink._id" class="card">
            <div class="card-head">
              <div>
                <p class="name">{{ drink.name || 'Drink' }}</p>
                <p class="meta">
                  {{ drink.price ? drink.price.toLocaleString('vi-VN') + ' VND' : '—' }}
                  <span v-if="drink.score !== undefined" class="score">Score: {{ drink.score?.toFixed(2) }}</span>
                </p>
              </div>
              <button type="button" class="mini" @click="addDrink(drink)">Add</button>
            </div>
            <p class="desc">{{ drink.desc || 'Hãy thử ngay thức uống hợp mood của bạn.' }}</p>
          </div>
          <p v-if="!reco.length && !recoLoading" class="status">Chưa có gợi ý. Hãy thử mood khác.</p>
        </template>
      </div>
    </div>
  </section>
</template>

<style scoped>
.booking {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(340px, 1fr));
  gap: 24px;
  color: var(--ink);
}

.panel {
  background: var(--paper);
  padding: clamp(24px, 4vw, 32px);
  border-radius: 16px;
  display: grid;
  gap: 14px;
  border: 1px solid rgba(0, 0, 0, 0.06);
  box-shadow: 0 18px 40px rgba(0, 0, 0, 0.12);
}

.form-fields {
  display: grid;
  gap: 12px;
}

label {
  display: grid;
  gap: 6px;
  font-weight: 700;
}

input,
select {
  border: 1px solid var(--cream-strong);
  padding: 12px 14px;
  border-radius: 10px;
  font-family: inherit;
  background: var(--paper);
  color: var(--ink);
}

button {
  border: 1px solid var(--dark);
  background: var(--dark);
  color: #fff;
  padding: 14px 16px;
  border-radius: 12px;
  cursor: pointer;
  font-weight: 700;
  font-size: 1rem;
}

button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.ghost {
  background: transparent;
  color: var(--ink);
  border: 1px solid rgba(0, 0, 0, 0.12);
}

.mini {
  padding: 8px 12px;
  font-size: 0.95rem;
}

.eyebrow {
  margin: 0;
  letter-spacing: 0.18em;
  text-transform: uppercase;
  font-size: 0.8rem;
}

.lede {
  margin: 0 0 4px;
  color: rgba(0, 0, 0, 0.7);
}

.selected {
  border: 1px dashed rgba(0, 0, 0, 0.12);
  padding: 10px;
  border-radius: 12px;
}

.chip-list {
  display: grid;
  gap: 8px;
}

.chip {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 10px 12px;
  border-radius: 10px;
  background: rgba(0, 0, 0, 0.04);
}

.chip .qty {
  display: flex;
  align-items: center;
  gap: 8px;
}

.chip .qty button {
  padding: 4px 10px;
  border-radius: 8px;
}

.controls {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(160px, 1fr));
  gap: 10px;
}

.reco-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.reco-list {
  display: grid;
  gap: 12px;
}

.card {
  padding: 12px;
  border-radius: 12px;
  background: rgba(0, 0, 0, 0.03);
  border: 1px solid rgba(0, 0, 0, 0.06);
  display: grid;
  gap: 6px;
}

.card-head {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.name {
  margin: 0;
  font-weight: 800;
}

.meta {
  margin: 0;
  color: rgba(0, 0, 0, 0.7);
}

.score {
  margin-left: 8px;
  font-weight: 700;
}

.desc {
  margin: 0;
  color: rgba(0, 0, 0, 0.7);
}

.status {
  margin: 0;
  padding: 10px 12px;
  border-radius: 8px;
  background: rgba(0, 0, 0, 0.04);
}

.status.success {
  background: #e6f5ed;
  color: #156f3d;
}

.status.error {
  background: #fde9ea;
  color: #b00020;
}

.mini-title {
  margin: 0 0 6px;
  font-weight: 800;
}

:global(.theme-night) .panel {
  background: rgba(15, 20, 36, 0.7);
  border-color: rgba(255, 255, 255, 0.08);
  color: #f6efe6;
}

:global(.theme-night) input,
:global(.theme-night) select {
  background: rgba(255, 255, 255, 0.06);
  color: #f6efe6;
  border-color: rgba(255, 255, 255, 0.2);
}

:global(.theme-night) label,
:global(.theme-night) .kicker,
:global(.theme-night) .hero-title,
:global(.theme-night) .hero-copy,
:global(.theme-night) .booking {
  color: #f6efe6;
}

:global(.theme-night) .lede,
:global(.theme-night) .meta,
:global(.theme-night) .desc {
  color: rgba(245, 241, 232, 0.82);
}

:global(.theme-night) .card .name {
  color: #f6efe6;
}

:global(.theme-night) .card {
  background: rgba(255, 255, 255, 0.04);
  border-color: rgba(255, 255, 255, 0.06);
}

:global(.theme-night) .status {
  background: rgba(255, 255, 255, 0.06);
  color: #f6efe6;
}

:global(.theme-night) .status.success {
  background: rgba(62, 146, 98, 0.18);
  color: #a7f3c7;
}

:global(.theme-night) .status.error {
  background: rgba(255, 94, 94, 0.18);
  color: #ffc7c7;
}
</style>
