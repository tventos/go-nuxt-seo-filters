<template>
  <Heading>Каталог</Heading>
  <CategoryList />
</template>

<script setup lang="ts">
import Heading from '~/shared/ui/typography/Heading.vue';
import CategoryList from '~/widgets/category-list/ui/CategoryList.vue';

useHead({
  title: 'Каталог',
});
</script>
