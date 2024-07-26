<script setup>
import { RouterView } from 'vue-router'
</script>

<template>
  <header>
    <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" />

    <div class="wrapper">
      <HelloWorld msg="You did it!" />

      <nav></nav>
    </div>
  </header>

  <RouterView />
</template>

<style scoped></style>
