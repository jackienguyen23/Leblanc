<script setup>
import { reactive, ref } from 'vue'
import { useRouter, RouterLink } from 'vue-router'
import { registerUser } from '@/api'

const router = useRouter()
const form = reactive({
  name: '',
  email: '',
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

  if (!form.name || !form.email || !form.password) {
    error.value = 'Please fill out your name, Gmail and password.'
    return
  }

  loading.value = true
  try {
    const res = await registerUser({
      name: form.name,
      email: form.email,
      password: form.password,
    })
    persistUser(res.user)
    message.value = 'Account created! Redirecting you to log in...'
    setTimeout(() => router.push('/login'), 1100)
  } catch (err) {
    error.value = err?.response?.data?.error || 'Could not create your account. Please try again.'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="ui padded segment auth-card">
    <h2 class="ui header">Create your LeBlanc account</h2>
    <p>Share your name, Gmail, and choose a password to get started.</p>

    <form class="ui form" @submit.prevent="handleSubmit">
      <div class="field">
        <label>Name</label>
        <input v-model="form.name" type="text" placeholder="E.g. Jackie Nguyen" autocomplete="name" />
      </div>

      <div class="field">
        <label>Gmail</label>
        <input
          v-model="form.email"
          type="email"
          placeholder="name@gmail.com"
          autocomplete="email"
        />
      </div>

      <div class="field">
        <label>Password</label>
        <input
          v-model="form.password"
          type="password"
          placeholder="Create a password"
          autocomplete="new-password"
        />
      </div>

      <button class="ui primary button" type="submit" :class="{ loading }" :disabled="loading">
        Create account
      </button>
      <RouterLink to="/login" class="ui button basic">Already have an account? Sign in</RouterLink>
    </form>

    <div v-if="error" class="ui negative message">
      <div class="header">Something went wrong</div>
      <p>{{ error }}</p>
    </div>

    <div v-if="message" class="ui positive message">
      <div class="header">Almost there</div>
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
