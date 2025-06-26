<template>
  <h1 v-if="props.level === 1" class="text-3xl font-medium mb-3"><slot /></h1>
  <h2 v-if="props.level === 2" class="text-xl font-medium mb-3"><slot /></h2>
</template>

<script setup lang="ts">
interface Props {
  level?: 1 | 2;
}
const props = withDefaults(defineProps<Props>(), {
  level: 1,
});
</script>
