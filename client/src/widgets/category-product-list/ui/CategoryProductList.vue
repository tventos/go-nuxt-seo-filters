<template>
  <div class="w-3/12" v-for="product of products" :key="product.id">
    <ProductCard :product="product" />
  </div>

  <div class="flex align-middle flex-wrap w-full" v-if="products.length === 0">
    <UIcon
      class="w-full mb-5 text-gray-500"
      name="i-lucide-search-x"
      size="70"
    />
    <div class="text-center text-2xl w-full">Ничего не найдено.</div>
  </div>
</template>

<script setup lang="ts">
import ProductCard from '~/entities/product/ui/ProductCard.vue';
import type { ProductCategory } from '~/entities/product';
import { useFetchCategoryProducts } from '~/entities/category';

interface Props {
  filter?: string;
}

const props = defineProps<Props>();

const products = await useFetchCategoryProducts(props.filter);
</script>
