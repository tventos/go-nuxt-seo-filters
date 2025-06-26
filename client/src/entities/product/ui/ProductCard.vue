<template>
  <UCard>
    <template #header>
      <Placeholder class="h-32" />
    </template>

    <div class="h-7">
      <ULink
        class="font-medium"
        :to="`${PRODUCT_PAGE}/${props.product.slug}`"
        >{{ props.product.name }}</ULink
      >
    </div>

    <div class="flex gap-2 h-5">
      <UBadge color="neutral" v-for="filter of props.product.filter_values">{{
        filter.value
      }}</UBadge>
    </div>

    <template #footer>
      <div class="flex">
        <div>
          <UButton as="a" :to="`${PRODUCT_PAGE}/${props.product.slug}`">
            Buy
          </UButton>
        </div>
      </div>
    </template>
  </UCard>
</template>

<script setup lang="ts">
import type { ProductCategory } from '~/entities/product';
import { PRODUCT_PAGE } from '~/app/routes';
import Placeholder from '~/shared/ui/placeholder/Placeholder.vue';

interface Props {
  product: ProductCategory;
}

const props = defineProps<Props>();
</script>
