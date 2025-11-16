<template>
  <div class="ui segment">
    <h3 class="ui header">Đặt bàn nhanh</h3>
    <div class="ui form">
      <div class="two fields">
        <div class="field">
          <label>Họ tên</label>
          <input v-model="name" placeholder="Nguyễn Văn A" />
        </div>
        <div class="field">
          <label>SĐT</label>
          <input v-model="phone" placeholder="0123 456 789" />
        </div>
        <div class="field">
          <label>Email</label>
          <input type="email" v-model="email">
        </div>
      </div>
      <div class="field">
        <label>Thời gian</label>
        <input type="datetime-local" v-model="time" />
      </div>
      <button
        class="ui black button"
        :class="{ loading: submitting }"
        @click="book"
        :disabled="submitting || !canSubmit"
      >
        Đặt
      </button>
    </div>

    <div v-if="ok" class="ui positive message">
      <div class="header">Đặt bàn thành công!</div>
      <p>Chúng tôi sẽ liên hệ xác nhận trong thời gian sớm nhất.</p>
    </div>
    <div v-else-if="error" class="ui negative message">
      <div class="header">Có lỗi xảy ra</div>
      <p>{{ error }}</p>
    </div>
  </div>
</template>

<script>
import { createBooking } from '@/api'

export default {
  name: 'BookingView',
  data: () => ({
    name: '',
    phone: '',
    time: '',
    ok: false,
    error: null,
    submitting: false,
  }),
  computed: {
    canSubmit() {
      return this.name && this.phone && this.time
    },
  },
  methods: {
    async book() {
      if (!this.canSubmit || this.submitting) return
      this.submitting = true
      this.error = null
      this.ok = false
      try {
        const payload = {
          name: this.name,
          phone: this.phone,
          time: this.time,
          items: [],
        }
        const res = await createBooking(payload)
        this.ok = Boolean(res?._id || res?.ok)
      } catch (err) {
        this.error = err?.message || 'Vui lòng thử lại.'
      } finally {
        this.submitting = false
      }
    },
  },
}
</script>
