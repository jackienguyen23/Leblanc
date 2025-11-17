<script setup>
import { ref, computed } from 'vue'
import { createBooking } from '@/api'

const form = ref({
  name: '',
  phone: '',
  email: '',
  time: '',
  guests: 2,
})

const submitting = ref(false)
const ok = ref(false)
const error = ref('')

const canSubmit = computed(() => form.value.name && form.value.phone && form.value.time)

const book = async () => {
  if (!canSubmit.value || submitting.value) return
  submitting.value = true
  error.value = ''
  ok.value = false
  try {
    const payload = {
      ...form.value,
      items: [],
      channel: 'web',
    }
    const res = await createBooking(payload)
    ok.value = Boolean(res?.ok || res?._id)
  } catch (err) {
    error.value = err?.message || 'Không thể đặt bàn lúc này.'
  } finally {
    submitting.value = false
  }
}
</script>

<template>
  <section class="booking">
    <div class="panel intro">
      <h1>Contact & Make Reservation</h1>
      <p>
        Visit us at Le'Blanc and experience coffee that heals a story. We're here to brighten your day with every cup.
      </p>
      <div class="info">
        <div>
          <strong>Location</strong>
          <p>Ho Chi Minh City, Vietnam</p>
        </div>
        <div>
          <strong>Hours</strong>
          <p>Mon - Sun: 7:00 AM - 02:00 PM</p>
        </div>
      </div>
    </div>

    <form class="panel form" @submit.prevent="book">
      <label>
        Name
        <input v-model="form.name" placeholder="Nguyễn Văn A" />
      </label>
      <label>
        Phone
        <input v-model="form.phone" placeholder="0123 456 789" />
      </label>
      <label>
        Gmail
        <input v-model="form.email" type="email" placeholder="name@gmail.com" />
      </label>
      <label>
        Arrival Time
        <input v-model="form.time" type="datetime-local" />
      </label>
      <label>
        Guests
        <input v-model.number="form.guests" type="number" min="1" max="10" />
      </label>
      <button type="submit" :disabled="!canSubmit || submitting">
        <span v-if="submitting">Processing...</span>
        <span v-else>Request Booking</span>
      </button>
      <p v-if="ok" class="status success">Đặt bàn thành công! Chúng tôi sẽ liên hệ xác nhận.</p>
      <p v-else-if="error" class="status error">{{ error }}</p>
    </form>
  </section>
</template>

<style scoped>
.booking {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 28px;
  color: #1f1208;
}

.panel {
  border: none;
  background: transparent;
  padding: clamp(32px, 5vw, 72px) 0;
  display: flex;
  flex-direction: column;
  gap: 18px;
}

.info {
  display: grid;
  gap: 12px;
}

.form label {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

input {
  border: 1px solid #d8d2ca;
  padding: 14px 16px;
  border-radius: 8px;
  font-family: inherit;
}

button {
  border: 1px solid #1f1208;
  background: #1f1208;
  color: #fff;
  padding: 16px 24px;
  border-radius: 999px;
  cursor: pointer;
  font-weight: 600;
  font-size: 1rem;
}

button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.status {
  margin: 0;
  padding: 10px 12px;
  border-radius: 4px;
}

.status.success {
  border: none;
  background: #e6f5ed;
  color: #156f3d;
}

.status.error {
  border: none;
  background: #fde9ea;
  color: #b00020;
}
</style>
