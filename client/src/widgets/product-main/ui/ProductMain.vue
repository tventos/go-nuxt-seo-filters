<template>
  <div class="flex gap-10">
    <div class="w-3/12">
      <Placeholder class="h-60" />
    </div>
    <div class="w-3/4">
      <Heading>{{ product.name }}</Heading>
      <div class="w-full mt-5 text-gray-600">{{ product.description }}</div>
    </div>
  </div>

  <div class="w-full mt-10">
    <Heading :level="2">Characteristics</Heading>
    <ProductCharacteristics
      :data="product.characteristics"
      :category-slug="product.category.slug"
    />
  </div>
</template>

<script setup lang="ts">
import { useFetchProduct } from '~/entities/product/api/useFetchProduct';
import Heading from '~/shared/ui/typography/Heading.vue';
import Placeholder from '~/shared/ui/placeholder/Placeholder.vue';
import ProductCharacteristics from '~/widgets/product-characteristics/ui/ProductCharacteristics.vue';

const route = useRoute<{ slug: string }>();
const product = await useFetchProduct(route.params.slug);

useHead({
  title: product.name,
  description: product.description,
});
</script>
