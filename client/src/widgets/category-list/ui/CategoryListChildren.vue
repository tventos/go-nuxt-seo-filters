<template>
  <ul class="category-list-children">
    <li v-for="category of props.categories" :key="category.id">
      <ULink :to="`/category/${category.slug}`">{{ category.name }}</ULink>
      <CategoryListChildren
        v-if="category.children"
        :categories="category.children"
      />
    </li>
  </ul>
</template>

<script setup lang="ts">
import type { Category } from '~/entities/category';

const props = defineProps({
  categories: Array as Category[],
});
</script>

<style scoped>
@reference "tailwindcss";
.category-list-children {
  li > ul {
    @apply pl-3;
  }
}
</style>
