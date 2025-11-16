<template>
  <div class="ui segment">
    <h3 class="ui header">Mood-Booker</h3>
    <div class="ui form">
      <div class="field">
        <label>Chọn ảnh</label>
        <input type="file" accept="image/*" @change="onFile" />
      </div>

      <div v-if="img" class="ui small image preview">
        <img :src="img" alt="Mood reference" />
      </div>

      <div class="two fields">
        <div class="field">
          <label>Cảm xúc</label>
          <select v-model="emotion" class="ui dropdown">
            <option value="calm">Calm</option>
            <option value="happy">Happy</option>
            <option value="stressed">Stressed</option>
            <option value="sad">Sad</option>
            <option value="adventurous">Adventurous</option>
          </select>
        </div>
        <div class="field">
          <label>Nhiệt độ ưu tiên</label>
          <select v-model="tempPref" class="ui dropdown">
            <option :value="null">Không chọn</option>
            <option value="hot">Nóng</option>
            <option value="iced">Đá</option>
          </select>
        </div>
      </div>

      <div class="field">
        <label>Tone màu từ ảnh</label>
        <div class="ui label">{{ colorLabel }}</div>
      </div>

      <button
        class="ui black button"
        :class="{ loading }"
        @click="recommend"
        :disabled="!img || loading"
      >
        Gợi ý đồ uống
      </button>
    </div>

    <div v-if="error" class="ui negative message">
      <div class="header">Không thể gợi ý đồ uống</div>
      <p>{{ error }}</p>
    </div>

    <div class="ui divider"></div>

    <div class="ui items" v-if="reco.length">
      <div class="item" v-for="drink in reco" :key="drink._id || drink.name">
        <div class="image">
          <img :src="drink.image || 'https://picsum.photos/160/120?blur=1'" alt="Drink" />
        </div>
        <div class="content">
          <div class="header">
            {{ drink.name }}
            <span class="price">— {{ formatCurrency(drink.price) }}</span>
          </div>
          <div class="meta">Điểm hợp mood: {{ formatScore(drink.score) }}</div>
          <div class="description">{{ drink.desc }}</div>
        </div>
      </div>
    </div>
    <div v-else-if="tried && !loading" class="ui info message">
      <div class="header">Chưa có gợi ý</div>
      <p>Hãy thử cảm xúc khác hoặc chọn bức ảnh khác nhé!</p>
    </div>
  </div>
</template>

<script>
import { recoFromFeatures } from '@/api'

export default {
  name: 'MoodBookerView',
  data: () => ({
    img: null,
    emotion: 'calm',
    colorTone: 'neutral',
    tempPref: null,
    loading: false,
    reco: [],
    tried: false,
    error: null,
  }),
  computed: {
    colorLabel() {
      return (this.colorTone || 'neutral').toUpperCase()
    },
  },
  methods: {
    onFile(event) {
      const file = event.target.files?.[0]
      if (!file) {
        return
      }
      if (this.img) {
        URL.revokeObjectURL(this.img)
      }
      const url = URL.createObjectURL(file)
      this.img = url
      this.extractTone(url)
    },
    extractTone(url) {
      const image = new Image()
      image.crossOrigin = 'anonymous'
      image.onload = () => {
        try {
          const canvas = document.createElement('canvas')
          canvas.width = image.width
          canvas.height = image.height
          const ctx = canvas.getContext('2d')
          if (!ctx) {
            throw new Error('Canvas context unavailable')
          }
          ctx.drawImage(image, 0, 0)
          const { data } = ctx.getImageData(0, 0, canvas.width, canvas.height)
          let r = 0
          let g = 0
          let b = 0
          let n = 0
          for (let i = 0; i < data.length; i += 16) {
            r += data[i]
            g += data[i + 1]
            b += data[i + 2]
            n++
          }
          const kelvin = r / n - b / n
          this.colorTone = kelvin > 25 ? 'warm' : kelvin < -25 ? 'cool' : 'neutral'
        } catch (e) {
          this.colorTone = 'neutral'
        }
      }
      image.src = url
    },
    async recommend() {
      if (!this.img) return
      this.loading = true
      this.error = null
      this.tried = true
      try {
        const hour = new Date().getHours()
        const timeOfDay = hour >= 6 && hour < 18 ? 'day' : 'night'
        this.reco = await recoFromFeatures({
          emotion: this.emotion,
          colorTone: this.colorTone,
          context: {
            timeOfDay,
            tempPref: this.tempPref || undefined,
          },
        })
      } catch (err) {
        this.error = err?.message || 'Vui lòng thử lại sau.'
      } finally {
        this.loading = false
      }
    },
    formatCurrency(value) {
      const amount = Number(value) || 0
      return `${amount.toLocaleString('vi-VN')}đ`
    },
    formatScore(value) {
      if (value === undefined || value === null) return '—'
      return Number(value).toFixed(2)
    },
  },
  beforeUnmount() {
    if (this.img) {
      URL.revokeObjectURL(this.img)
    }
  },
}
</script>

<style scoped>
.preview {
  margin-bottom: 16px;
}

.price {
  font-weight: 400;
  color: #5c5c5c;
}
</style>
