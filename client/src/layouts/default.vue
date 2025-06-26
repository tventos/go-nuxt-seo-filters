<template>
  <UApp>
    <Header />

    <UContainer class="pt-3.5 pb-3.5">
      <slot />
    </UContainer>

    <Footer />
  </UApp>
</template>

<script setup lang="ts">
import Header from '~/widgets/header/Header.vue';
import Footer from '~/widgets/footer/Footer.vue';
</script>
