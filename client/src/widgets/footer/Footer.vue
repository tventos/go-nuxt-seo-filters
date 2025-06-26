<template>
  <USeparator />
  <footer class="flex h-20 items-center">
    <UContainer>
      <div class="flex">
        <div class="w-32 text-gray-400">My shop</div>
        <div>
          <ul class="flex gap-4">
            <li>
              <ULink :to="HOME_PAGE">Home</ULink>
            </li>
            <li>
              <ULink :to="CATALOG_PAGE">Catalog</ULink>
            </li>
          </ul>
        </div>
      </div>
    </UContainer>
  </footer>
</template>
<script setup lang="ts">
import { CATALOG_PAGE, HOME_PAGE } from '~/app/routes.js';
</script>
