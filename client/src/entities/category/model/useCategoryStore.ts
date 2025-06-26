import { defineStore } from 'pinia';
import { type Category, useFetchCategories } from '~/entities/category';

export const useCategoryStore = defineStore('category', {
  state: () => ({
    categories: [] as Category[],
    isLoading: false as boolean,
    error: null as string | null,
    isInitialDataLoaded: false as boolean,
  }),

  actions: {
    async init() {
      if (this.isInitialDataLoaded) {
        return;
      }

      await this.fetchCategories();
    },
    async fetchCategories() {
      this.isLoading = true;
      this.error = null;

      try {
        this.categories = await useFetchCategories();
        this.isInitialDataLoaded = true;
      } catch (e: any) {
        this.error = e.message;
      } finally {
        this.isLoading = false;
      }
    },
  },
});
