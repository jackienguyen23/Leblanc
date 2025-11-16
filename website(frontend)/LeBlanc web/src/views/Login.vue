<script setup>
import { reactive, ref } from 'vue'
import { useRouter, RouterLink } from 'vue-router'
import { loginUser } from '@/api'

const router = useRouter()
const form = reactive({
  name: '',
  password: '',
})

const loading = ref(false)
const error = ref('')
const message = ref('')

const persistUser = (user) => {
  try {
    localStorage.setItem('leblancUser', JSON.stringify(user))
  } catch (err) {
    console.warn('Could not persist user', err)
  }
}

const handleSubmit = async () => {
  error.value = ''
  message.value = ''

  if (!form.name || !form.password) {
    error.value = 'Please provide both your name and password.'
    return
  }

  loading.value = true
  try {
    const res = await loginUser({
      name: form.name,
      password: form.password,
    })
    persistUser(res.user)
    message.value = `Welcome back, ${res.user.name}! Taking you home...`
    setTimeout(() => router.push('/'), 900)
  } catch (err) {
    error.value = err?.response?.data?.error || 'Unable to sign you in. Please try again.'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="ui padded segment auth-card">
    <h2 class="ui header">Sign in to LeBlanc</h2>
    <p>Use the name you registered with and your password.</p>

    <form class="ui form" @submit.prevent="handleSubmit">
      <div class="field">
        <label>Name</label>
        <input v-model="form.name" type="text" placeholder="E.g. Jackie Nguyen" autocomplete="username" />
      </div>

      <div class="field">
        <label>Password</label>
        <input
          v-model="form.password"
          type="password"
          placeholder="Enter your password"
          autocomplete="current-password"
        />
      </div>

      <button class="ui primary button" type="submit" :class="{ loading }" :disabled="loading">
        Log in
      </button>
      <RouterLink to="/register" class="ui button basic">First visit? Create an account</RouterLink>
    </form>

    <div v-if="error" class="ui negative message">
      <div class="header">Login failed</div>
      <p>{{ error }}</p>
    </div>

    <div v-if="message" class="ui positive message">
      <div class="header">Success</div>
      <p>{{ message }}</p>
    </div>
  </div>
</template>

<style scoped>
.auth-card {
  max-width: 520px;
  margin: 0 auto;
}

.ui.form {
  margin-top: 16px;
}

.ui.button + .ui.button {
  margin-left: 12px;
}
</style>
