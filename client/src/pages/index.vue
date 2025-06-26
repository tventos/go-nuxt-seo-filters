<template>
  <Heading :level="2">Categories</Heading>

  <CategoryList />
</template>

<script setup lang="ts">
import CategoryList from '~/widgets/category-list/ui/CategoryList.vue';
import Heading from '~/shared/ui/typography/Heading.vue';

useHead({
  title: 'Главная',
});
</script>
