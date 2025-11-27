<script setup>
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { getUsers } from '@/api'

const router = useRouter()
const users = ref([])
const loading = ref(true)
const error = ref('')
const ADMIN_EMAIL = (import.meta.env.VITE_ADMIN_EMAIL || '').toLowerCase()

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
        <p class="eyebrow">Admin</p>
        <h1>User directory</h1>
        <p class="lede">List of all registered users.</p>
      </div>
      <button class="btn" type="button" @click="fetchUsers" :disabled="loading">
        {{ loading ? 'Refreshing…' : 'Refresh' }}
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
  gap: 14px;
}

.head {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.eyebrow {
  margin: 0;
  letter-spacing: 0.2em;
  text-transform: uppercase;
  font-size: 0.8rem;
  color: var(--tan);
}

h1 {
  margin: 4px 0;
}

.lede {
  margin: 0;
  color: rgba(0, 0, 0, 0.65);
}

.btn {
  padding: 10px 14px;
  border-radius: 10px;
  border: 1px solid rgba(0, 0, 0, 0.2);
  background: #0f1424;
  color: #f6efe6;
  font-weight: 800;
  cursor: pointer;
}

.banner {
  padding: 12px 14px;
  border-radius: 10px;
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
  gap: 6px;
}

.row {
  display: grid;
  grid-template-columns: 1.2fr 1.8fr 1fr;
  gap: 10px;
  padding: 12px 14px;
  border-radius: 12px;
  background: rgba(0, 0, 0, 0.02);
}

.row.header {
  font-weight: 900;
  background: transparent;
}

.name {
  font-weight: 800;
}

.email {
  word-break: break-all;
}

.date {
  color: rgba(0, 0, 0, 0.65);
}

@media (max-width: 720px) {
  .row {
    grid-template-columns: 1fr;
  }
}
</style>
