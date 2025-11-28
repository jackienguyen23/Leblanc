<script setup>
import { inject, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { getUsers } from '@/api'

const router = useRouter()
const users = ref([])
const loading = ref(true)
const error = ref('')
const ADMIN_EMAIL = (import.meta.env.VITE_ADMIN_EMAIL || '').toLowerCase()
const theme = inject('theme', ref('day'))

const isAdminUser = (user) => {
  if (!user) return false
  const email = user.email?.toLowerCase() || ''
  return ADMIN_EMAIL && email === ADMIN_EMAIL
}

const ensureAdmin = () => {
  let stored = null
  try {
    stored = JSON.parse(localStorage.getItem('leblancUser') || 'null')
  } catch (err) {
    stored = null
  }
  if (!isAdminUser(stored)) {
    router.replace('/')
    return null
  }
  return stored
}

const fetchUsers = async () => {
  loading.value = true
  error.value = ''
  try {
    users.value = await getUsers()
  } catch (err) {
    error.value = err?.response?.data?.error || 'Could not load users.'
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  const u = ensureAdmin()
  if (u) fetchUsers()
})

const formatDate = (iso) => {
  try {
    return new Date(iso).toLocaleString()
  } catch (err) {
    return iso
  }
}
</script>

<template>
  <section class="admin">
    <header class="head">
      <div>
        <p class="eyebrow">ADMIN PAGE</p>
        <h1>User directory</h1>
        <p class="lede">List of all registered users.</p>
      </div>
      <button
        class="btn"
        :class="{ 'is-night': theme === 'night' }"
        type="button"
        @click="fetchUsers"
        :disabled="loading"
      >
        {{ loading ? 'Refreshing...' : 'Refresh' }}
      </button>
    </header>

    <div v-if="error" class="banner error">{{ error }}</div>
    <div v-else-if="loading" class="banner">Loading users...</div>
    <div v-else class="table">
      <div class="row header">
        <span>Name</span>
        <span>Email</span>
        <span>Joined</span>
      </div>
      <div v-for="u in users" :key="u._id || u.email" class="row">
        <span class="name">{{ u.name }}</span>
        <span class="email">{{ u.email }}</span>
        <span class="date">{{ formatDate(u.createdAt) }}</span>
      </div>
    </div>
  </section>
</template>

<style scoped>
.admin {
  display: grid;
  gap: 18px;
}

.head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 22px;
}

.eyebrow {
  margin: 0;
  letter-spacing: 0.26em;
  text-transform: uppercase;
  font-size: 1.04rem;
  color: var(--tan);
}

h1 {
  margin: 6px 0;
  font-size: 2.6rem;
  line-height: 1.1;
}

.lede {
  margin: 0;
  font-size: 1.12rem;
  color: var(--ink);
  opacity: 0.82;
}

.btn {
  padding: 13px 18px;
  border-radius: 13px;
  border: 1px solid rgba(0, 0, 0, 0.2);
  background: #0f1424;
  color: #f6efe6;
  font-weight: 800;
  font-size: 1.05rem;
  cursor: pointer;
  transition: transform 0.12s ease, background 0.12s ease, color 0.12s ease;
}

.btn.is-night {
  background: var(--tan);
  color: #0f1424;
  border-color: rgba(0, 0, 0, 0.25);
}

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.banner {
  padding: 16px 18px;
  border-radius: 13px;
  font-size: 1.05rem;
  background: rgba(0, 0, 0, 0.04);
  border: 1px solid rgba(0, 0, 0, 0.08);
}

.banner.error {
  background: rgba(220, 90, 36, 0.12);
  border-color: rgba(220, 90, 36, 0.3);
  color: #a53a0b;
}

.table {
  display: grid;
  gap: 8px;
}

.row {
  display: grid;
  grid-template-columns: 1.2fr 1.8fr 1fr;
  gap: 13px;
  padding: 16px 18px;
  border-radius: 16px;
  background: rgba(255, 255, 255, 0.04);
  font-size: 1.05rem;
}

.row.header {
  font-weight: 900;
  background: transparent;
}

.name {
  font-weight: 800;
  font-size: 1.06rem;
}

.email {
  word-break: break-all;
}

.date {
  color: var(--ink);
  opacity: 0.82;
}

@media (max-width: 720px) {
  .row {
    grid-template-columns: 1fr;
  }
}
</style>
