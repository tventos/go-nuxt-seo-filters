<template>
  <Heading>{{ categoryInfo.h1 }}</Heading>

  <div class="flex gap-6">
    <div class="w-3/12">
      <CategoryFilterList :slug="categorySlug" />
    </div>
    <div class="w-3/4">
      <div class="flex gap-6">
        <CategoryProductList :filter="filter" />
      </div>

      <div
        v-if="categoryInfo.content"
        v-html="categoryInfo.content"
        class="mt-12"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import Heading from '~/shared/ui/typography/Heading.vue';
import CategoryProductList from '~/widgets/category-product-list/ui/CategoryProductList.vue';
import { buildCategoryFilters } from '~/shared/lib/buildCategoryUrl';
import { useFetchCategoryInfo } from '~/entities/category';
import CategoryFilterList from '~/widgets/category-filter-list/ui/CategoryFilterList.vue';
import { useCategoryFiltersStore } from '~/features/category-filters/model/useCategoryFiltersStore';

const route = useRoute();
const filter = buildCategoryFilters(route.params.slug, route.query);
const categoryInfo = await useFetchCategoryInfo(filter);
const categoryFiltersStore = useCategoryFiltersStore();
const categorySlug = route.params.slug[0];

useHead({
  title: categoryInfo.title,
  description: categoryInfo.description,
});

if (categoryInfo.filter_slug && categoryInfo.filter_value_id) {
  categoryFiltersStore.setInitialFiltersData({
    [categoryInfo.filter_slug]: [categoryInfo.filter_value_id],
  });
}
</script>
