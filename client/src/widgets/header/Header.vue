<template>
  <header class="flex w-full h-20 items-center">
    <UContainer>
      <div class="flex align-center w-full">
        <div class="flex w-32">
          <ULink
            :to="HOME_PAGE"
            active
            class="flex items-center justify-center font-bold"
            >My shop</ULink
          >
        </div>
        <div>
          <UNavigationMenu :items="items" class="w-full justify-center" />
        </div>
      </div>
    </UContainer>
  </header>
</template>

<script setup type="ts">
import { useCategoryStore } from "~/entities/category";
import { CATALOG_PAGE, HOME_PAGE } from '~/app/routes';

const categoryStore = useCategoryStore();
await categoryStore.init();

const catalogItems = categoryStore.categories.map((category) => ({
  label: category.name,
  to: `${CATALOG_PAGE}/${category.slug}`,
}));

const items = ref([
  {
    label: 'Home',
    icon: 'i-lucide-house',
    to: HOME_PAGE,
  },
  {
    label: 'Catalog',
    icon: 'i-lucide-shopping-cart',
    to: CATALOG_PAGE,
    children: catalogItems
  }
  ]);
</script>
