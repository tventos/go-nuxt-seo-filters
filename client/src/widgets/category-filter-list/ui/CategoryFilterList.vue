<template>
  <UCard>
    <CategoryFilters :filters="filters" :category-slug="props.slug" />
  </UCard>
</template>

<script setup lang="ts">
import CategoryFilters from '~/features/category-filters/ui/CategoryFilters.vue';
import { useFetchCategoryFilters } from '~/entities/category';

interface Props {
  slug: string;
}

const props = defineProps<Props>();
const filters = await useFetchCategoryFilters(props.slug);
</script>
