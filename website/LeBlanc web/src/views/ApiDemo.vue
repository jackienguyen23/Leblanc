<template>
  <div class="api-demo">
    <h2>API Integration Demo</h2>
    
    <div class="config">
      <label>
        <input type="checkbox" v-model="useGraphQL" />
        Use GraphQL (currently using {{ useGraphQL ? 'GraphQL' : 'REST' }})
      </label>
    </div>

    <div class="section">
      <h3>Drinks</h3>
      <button @click="fetchDrinks" :disabled="loading">
        {{ loading ? 'Loading...' : 'Fetch Drinks' }}
      </button>
      <div v-if="drinks.length" class="results">
        <p>Found {{ drinks.length }} drinks:</p>
        <ul>
          <li v-for="drink in drinks.slice(0, 5)" :key="drink._id">
            {{ drink.name }} - ${{ drink.price }}
          </li>
        </ul>
      </div>
    </div>

    <div class="section">
      <h3>Recommendations</h3>
      <button @click="fetchRecommendations" :disabled="loading">
        Get Recommendations
      </button>
      <div v-if="recommendations.length" class="results">
        <p>Top recommendations:</p>
        <ul>
          <li v-for="reco in recommendations" :key="reco.drinkId">
            Drink ID: {{ reco.drinkId }} - Score: {{ reco.score }}
          </li>
        </ul>
      </div>
    </div>

    <div class="section">
      <h3>Authentication Demo</h3>
      <div class="form">
        <input v-model="username" placeholder="Username" />
        <input v-model="email" placeholder="Email" type="email" />
        <input v-model="password" placeholder="Password" type="password" />
        <button @click="handleRegister" :disabled="loading">Register</button>
        <button @click="handleLogin" :disabled="loading">Login</button>
      </div>
      <div v-if="authResult" class="results">
        <p>{{ authResult }}</p>
      </div>
    </div>

    <div v-if="error" class="error">
      Error: {{ error }}
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import {
  getDrinks,
  getDrinksGraphQL,
  getDrinksREST,
  recoFromFeatures,
  recoFromFeaturesGraphQL,
  recoFromFeaturesREST,
  registerUser,
  registerUserGraphQL,
  registerUserREST,
  loginUser,
  loginUserGraphQL,
  loginUserREST,
} from '@/api'

const useGraphQL = ref(false)
const loading = ref(false)
const error = ref(null)
const drinks = ref([])
const recommendations = ref([])
const authResult = ref(null)

const username = ref('')
const email = ref('')
const password = ref('')

const fetchDrinks = async () => {
  loading.value = true
  error.value = null
  try {
    // Use the unified API or specific implementation
    const result = useGraphQL.value 
      ? await getDrinksGraphQL() 
      : await getDrinksREST()
    drinks.value = result
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

const fetchRecommendations = async () => {
  loading.value = true
  error.value = null
  try {
    const emotionFit = {
      calm: 0.8,
      happy: 0.6,
      stressed: 0.2,
      sad: 0.1,
      adventurous: 0.5,
    }

    const result = useGraphQL.value
      ? await recoFromFeaturesGraphQL(emotionFit, 'med', 'iced', 7)
      : await recoFromFeaturesREST({
          emotion: 'calm',
          colorTone: 'cool',
          context: {
            timeOfDay: 'day',
            tempPref: 'iced',
          },
        })
    
    recommendations.value = result
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

const handleRegister = async () => {
  if (!username.value || !email.value || !password.value) {
    error.value = 'Please fill in all fields'
    return
  }

  loading.value = true
  error.value = null
  authResult.value = null
  try {
    const payload = {
      name: username.value,
      email: email.value,
      password: password.value,
    }

    const result = useGraphQL.value
      ? await registerUserGraphQL({ name: payload.name, email: payload.email, password: payload.password })
      : await registerUserREST(payload)
    
    authResult.value = `Registration successful! Welcome, ${result.user?.name || username.value}`
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

const handleLogin = async () => {
  if (!username.value || !password.value) {
    error.value = 'Please fill in username and password'
    return
  }

  loading.value = true
  error.value = null
  authResult.value = null
  try {
    const payload = {
      name: username.value,
      password: password.value,
    }

    const result = useGraphQL.value
      ? await loginUserGraphQL(payload)
      : await loginUserREST(payload)
    
    authResult.value = `Login successful! Welcome back, ${result.user?.name || username.value}`
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.api-demo {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

.config {
  background: #f5f5f5;
  padding: 15px;
  border-radius: 8px;
  margin-bottom: 20px;
}

.section {
  margin: 30px 0;
  padding: 20px;
  border: 1px solid #ddd;
  border-radius: 8px;
}

.section h3 {
  margin-top: 0;
}

button {
  background: #4CAF50;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 4px;
  cursor: pointer;
  margin: 5px;
}

button:hover:not(:disabled) {
  background: #45a049;
}

button:disabled {
  background: #ccc;
  cursor: not-allowed;
}

.results {
  margin-top: 15px;
  padding: 15px;
  background: #e8f5e9;
  border-radius: 4px;
}

.results ul {
  list-style: none;
  padding: 0;
}

.results li {
  padding: 5px 0;
  border-bottom: 1px solid #c8e6c9;
}

.error {
  background: #ffebee;
  color: #c62828;
  padding: 15px;
  border-radius: 4px;
  margin-top: 20px;
}

.form {
  display: flex;
  flex-direction: column;
  gap: 10px;
  max-width: 400px;
}

.form input {
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.form button {
  align-self: flex-start;
}
</style>
