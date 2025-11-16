<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div id="app" class="ui container">
    <div class="ui secondary pointing menu stackable">
      <RouterLink to="/" class="item" exact-active-class="active">Le’Blanc</RouterLink>
      <RouterLink to="/menu" class="item" exact-active-class="active">Menu</RouterLink>
      <RouterLink to="/mood-booker" class="item" exact-active-class="active">
        Mood-Booker
      </RouterLink>
      <RouterLink to="/booking" class="item" exact-active-class="active">Booking</RouterLink>
      <div class="right menu">
        <RouterLink to="/login" class="item" exact-active-class="active">Login</RouterLink>
        <RouterLink to="/register" class="item" exact-active-class="active">Register</RouterLink>
      </div>
    </div>

    <RouterView />
  </div>
</template>

<style scoped>
.menu {
  margin-bottom: 32px;
}
</style>
