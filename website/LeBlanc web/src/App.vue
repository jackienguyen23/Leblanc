<script setup>
import { computed, onBeforeUnmount, onMounted, provide, ref, watchEffect } from 'vue'
import { RouterLink, RouterView, useRoute, useRouter } from 'vue-router'
import darkLogo from '@/assets/dark-logo.png'
import brightLogo from '@/assets/bright-logo.png'

const ADMIN_EMAIL = (import.meta.env.VITE_ADMIN_EMAIL || '').toLowerCase()

const theme = ref(localStorage.getItem('theme') || 'day')
provide('theme', theme)
const router = useRouter()
const route = useRoute()

const clearPersistedUser = () => {
  try {
    localStorage.removeItem('leblancUser')
  } catch (err) {
    console.warn('Could not clear stored user', err)
  }
}

const loadUser = () => {
  try {
    return JSON.parse(localStorage.getItem('leblancUser') || 'null')
  } catch (err) {
    console.warn('Could not parse stored user', err)
    return null
  }
}

clearPersistedUser()
const user = ref(loadUser())
const isAuthed = computed(() => !!user.value)
const isAdmin = computed(() => {
  const email = user.value?.email?.toLowerCase() || ''
  return ADMIN_EMAIL && email === ADMIN_EMAIL
})
const isAdminOnlyShell = computed(() => isAdmin.value && (route.name === 'admin' || route.name === 'account'))
const isPlainLayout = computed(() => route.meta?.layout === 'plain')
const brandTarget = computed(() => (isAdminOnlyShell.value ? '/admin' : '/'))

const showAccountMenu = ref(false)
const accountRef = ref(null)

const handleUserUpdated = (event) => {
  // Event detail can carry the user; fallback to storage for safety.
  user.value = event?.detail ?? loadUser()
  showAccountMenu.value = false
}

const handleStorage = (event) => {
  if (event.key === 'leblancUser') {
    user.value = loadUser()
    showAccountMenu.value = false
  }
}

onMounted(() => {
  window.addEventListener('leblanc-user-updated', handleUserUpdated)
  window.addEventListener('storage', handleStorage)
  document.addEventListener('click', handleClickOutside)
})

onBeforeUnmount(() => {
  window.removeEventListener('leblanc-user-updated', handleUserUpdated)
  window.removeEventListener('storage', handleStorage)
  document.removeEventListener('click', handleClickOutside)
})

watchEffect(() => {
  const body = document.body
  body.classList.remove('theme-day', 'theme-night')
  body.classList.add(theme.value === 'night' ? 'theme-night' : 'theme-day')
  localStorage.setItem('theme', theme.value)
})

const toggleTheme = () => {
  theme.value = theme.value === 'night' ? 'day' : 'night'
}

const logoSrc = computed(() => (theme.value === 'night' ? brightLogo : darkLogo))

const userInitial = computed(() => (user.value?.name?.[0] || 'A').toUpperCase())

const logout = () => {
  localStorage.removeItem('leblancUser')
  user.value = null
  window.dispatchEvent(new CustomEvent('leblanc-user-updated', { detail: null }))
  router.push('/')
}

const toggleAccountMenu = () => {
  showAccountMenu.value = !showAccountMenu.value
}

const handleClickOutside = (event) => {
  if (!accountRef.value) return
  if (!accountRef.value.contains(event.target)) {
    showAccountMenu.value = false
  }
}
</script>

<template>
  <div class="shell">
    <header v-if="!isPlainLayout" class="header">
      <RouterLink :to="brandTarget" class="brand">
        <img :src="logoSrc" alt="Le'Blanc logo" class="brand-logo" />
        <span class="brand-tag">Where every drink tells a story</span>
      </RouterLink>
      <nav class="nav">
        <template v-if="isAdminOnlyShell">
          <RouterLink to="/admin" class="nav-link" exact-active-class="active">Admin</RouterLink>
        </template>
        <template v-else>
          <RouterLink to="/about" class="nav-link" exact-active-class="active">About</RouterLink>
          <RouterLink to="/menu" class="nav-link" exact-active-class="active">Menu</RouterLink>
          <RouterLink to="/booking" class="nav-link" exact-active-class="active">Booking</RouterLink>
          <RouterLink v-if="isAdmin" to="/admin" class="nav-link" exact-active-class="active">Admin</RouterLink>
          <RouterLink v-if="!isAuthed" to="/login" class="nav-link" exact-active-class="active">Login</RouterLink>
          <RouterLink v-if="!isAuthed" to="/register" class="nav-link" exact-active-class="active">Sign up</RouterLink>
        </template>
        <div v-if="isAuthed" ref="accountRef" class="account-wrap">
          <button class="account-pill" type="button" @click.stop="toggleAccountMenu">
            <div class="avatar" aria-hidden="true">{{ userInitial }}</div>
            <span class="account-name">{{ user?.name }}</span>
          </button>
          <div v-if="showAccountMenu" class="account-menu">
            <div class="account-meta">
              <span class="account-name">{{ user?.name }}</span>
              <span class="account-email">{{ user?.email }}</span>
            </div>
            <RouterLink
              v-if="isAdmin"
              to="/admin"
              class="account-link"
              exact-active-class="active"
              @click="showAccountMenu = false"
            >
              Admin dashboard
            </RouterLink>
            <RouterLink to="/account" class="account-link" exact-active-class="active" @click="showAccountMenu = false">
              Account
            </RouterLink>
            <button class="logout" type="button" @click="logout">Log out</button>
          </div>
        </div>
        <button class="theme-toggle" type="button" @click="toggleTheme">
          {{ theme === 'night' ? 'Day Mode' : 'Night Mode' }}
        </button>
      </nav>
    </header>

    <main class="page-frame">
      <RouterView />
    </main>
  </div>
</template>

<style scoped>
.shell {
  min-height: 100vh;
  display: grid;
  gap: 12px;
}

.header {
  width: 100%;
  margin: 0;
  padding: 22px 28px 8px;
  display: flex;
  justify-content: space-between;
  gap: 18px;
  align-items: center;
  position: sticky;
  top: 0;
  z-index: 10;
  background: var(--paper);
}

.brand {
  display: grid;
  gap: 6px;
  align-items: center;
  text-decoration: none;
  color: inherit;
}

.brand-logo {
  height: 96px;
  width: auto;
  display: block;
}

.brand-tag {
  font-size: 0.8rem;
  letter-spacing: 0.12em;
  text-transform: uppercase;
  opacity: 0.8;
}

.nav {
  display: flex;
  gap: 12px;
  align-items: center;
  flex-wrap: wrap;
}

.account-pill {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 8px 12px;
  border-radius: 999px;
  background: rgba(0, 0, 0, 0.04);
  border: 1px solid rgba(0, 0, 0, 0.08);
}

.account-wrap {
  position: relative;
}

.account-pill {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  padding: 8px 12px;
  border-radius: 999px;
  border: 1px solid rgba(0, 0, 0, 0.08);
  background: rgba(0, 0, 0, 0.04);
  cursor: pointer;
  font-weight: 800;
}

.avatar {
  height: 34px;
  width: 34px;
  border-radius: 50%;
  background: linear-gradient(145deg, #b88443, #e1c58d);
  color: #0b0b0b;
  display: grid;
  place-items: center;
  font-weight: 900;
  letter-spacing: 0.02em;
}

.account-name {
  font-weight: 800;
  color: var(--ink);
}

.account-menu {
  position: absolute;
  right: 0;
  top: calc(100% + 6px);
  min-width: 220px;
  display: grid;
  gap: 10px;
  padding: 12px;
  border-radius: 14px;
  background: var(--paper);
  border: 1px solid rgba(0, 0, 0, 0.08);
  box-shadow: 0 18px 36px rgba(0, 0, 0, 0.14);
  z-index: 20;
}

.account-meta {
  display: grid;
  line-height: 1.2;
}

.account-email {
  font-size: 0.9rem;
  color: rgba(0, 0, 0, 0.65);
}

.account-link {
  padding: 8px 10px;
  border-radius: 10px;
  border: 1px solid rgba(0, 0, 0, 0.08);
  text-decoration: none;
  color: var(--ink);
  font-weight: 800;
}

.logout {
  border: none;
  background: rgba(0, 0, 0, 0.08);
  padding: 10px 10px;
  border-radius: 10px;
  cursor: pointer;
  font-weight: 800;
}

.nav-link {
  text-decoration: none;
  color: var(--ink);
  font-weight: 800;
  letter-spacing: 0.02em;
  padding: 10px 0;
  font-size: 1.05rem;
}

.nav-link.active {
  color: var(--tan);
}

.theme-toggle {
  padding: 10px 14px;
  border-radius: 12px;
  border: 1px solid var(--dark);
  background: transparent;
  color: var(--ink);
  font-weight: 800;
  letter-spacing: 0.02em;
  cursor: pointer;
}

.theme-toggle:hover {
  background: rgba(0, 0, 0, 0.08);
}

.cta {
  padding: 10px 14px;
  background: var(--dark);
  color: #fff;
  border-radius: 10px;
  text-decoration: none;
  font-weight: 900;
  letter-spacing: 0.04em;
}

.page-frame {
  max-width: min(1600px, 98vw);
  margin: 0 auto;
  width: 100%;
  padding: 0 24px 64px;
  box-sizing: border-box;
}

@media (max-width: 720px) {
  .brand-word {
    font-size: 1.2rem;
  }
}
</style>
