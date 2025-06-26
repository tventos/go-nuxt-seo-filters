<template>
  <div
    class="mb-4 pb-4 border-b border-gray-200"
    v-for="filter of props.filters"
    :key="`filter-${filter.slug}-${filter.id}`"
  >
    <div class="text-lg font-bold mb-2">
      {{ filter.name }}
    </div>

    <div>
      <UCheckbox
        v-for="filterValue of filter.values"
        :key="filterValue.id"
        :label="filterValue.value"
        @change="() => selectFilter(filter.slug, filterValue.id)"
        :model-value="
          categoryFiltersStore.isSelected(filter.slug, filterValue.id)
        "
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { type CategoryFilter } from '~/entities/category';
import { useCategoryFiltersStore } from '~/features/category-filters/model/useCategoryFiltersStore';

interface Props {
  categorySlug: string;
  filters: CategoryFilter[];
}

const props = defineProps<Props>();
const categoryFiltersStore = useCategoryFiltersStore();
const router = useRouter();

categoryFiltersStore.setCategorySlug(props.categorySlug);

const selectFilter = async (key, value) => {
  categoryFiltersStore.selectFilter(key, value);

  const response = await categoryFiltersStore.applyFilters();

  await router.push(response.link);
};
</script>
