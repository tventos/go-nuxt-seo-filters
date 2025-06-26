<template>
  <div
    v-for="item of props.data"
    :key="item.filter_id"
    class="flex flex-wrap border-b border-gray-200 pb-2 mb-2"
  >
    <div class="w-6/12">{{ item.filter_name }}</div>
    <div v-if="!item.has_seo" class="w-6/12">{{ item.value }}</div>
    <div v-if="item.has_seo" class="w-6/12">
      <ULink
        :to="`${CATALOG_PAGE}/${categorySlug}/${item.seo_slug}`"
        class="text-green-600 border-b"
        >{{ item.value }}</ULink
      >
    </div>
  </div>
</template>

<script setup lang="ts">
import type { ProductCharacteristic } from '~/entities/product';
import { CATALOG_PAGE } from '~/app/routes';

interface Props {
  categorySlug: string;
  data: ProductCharacteristic[];
}
const props = defineProps<Props>();
</script>
