<template>
  <div>
    <h3 class="ui header">Menu</h3>

    <div v-if="loading" class="ui active inline loader" aria-label="Đang tải menu"></div>
    <div v-else-if="error" class="ui negative message">
      <div class="header">Không thể tải menu</div>
      <p>{{ error }}</p>
    </div>
    <div v-else class="ui three stackable cards">
      <div class="card" v-for="drink in drinks" :key="drink._id">
        <div class="image">
          <img :src="drink.image || 'https://picsum.photos/400/240?grayscale'" alt="Ảnh đồ uống" />
        </div>
        <div class="content">
          <div class="header">{{ drink.name }}</div>
          <div class="meta">{{ formatCurrency(drink.price) }}</div>
          <div class="description">{{ drink.desc }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getDrinks } from '@/api'

export default {
  name: 'MenuView',
  data: () => ({
    drinks: [],
    loading: false,
    error: null,
  }),
  created() {
    this.fetchDrinks()
  },
  methods: {
    async fetchDrinks() {
      this.loading = true
      this.error = null
      try {
        this.drinks = await getDrinks()
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
  },
}
</script>
