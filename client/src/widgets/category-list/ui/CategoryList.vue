<template>
  <div class="flex">
    <UCard class="w-3/12" v-for="category in categories" :key="category.id">
      <ULink
        class="flex text-xl font-medium mb-2 w-full"
        :to="`/category/${category.slug}`"
        >{{ category.name }}</ULink
      >
      <CategoryListChildren :categories="category.children"
    /></UCard>
  </div>
</template>

<script setup lang="ts">
import { useCategoryStore } from '~/entities/category';
import CategoryListChildren from './CategoryListChildren.vue';

const categoryStore = useCategoryStore();
await categoryStore.init();

const categories = ref(categoryStore.categories);
</script>
