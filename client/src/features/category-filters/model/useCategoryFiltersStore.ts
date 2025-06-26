import { defineStore } from 'pinia';
import { useSubmitCategoryApplyFilters } from '~/entities/category';

export const useCategoryFiltersStore = defineStore('categoryFilters', {
  state: () => ({
    categorySlug: '' as string,
    values: {} as Record<string, number[]>,
  }),

  actions: {
    setCategorySlug(categorySlug: string) {
      this.categorySlug = categorySlug;
    },

    setInitialFiltersData(filterValues: Record<string, number[]>) {
      this.values = filterValues;
    },

    selectFilter(key: string, value: number) {
      if (!this.values[key]) {
        this.values[key] = [value];
      } else {
        const index = this.values[key].indexOf(value);
        if (index === -1) {
          this.values[key].push(value);
        } else {
          this.values[key].splice(index, 1);
          if (this.values[key].length === 0) {
            delete this.values[key];
          }
        }
      }
    },

    isSelected(key: string, value: number): boolean {
      return !!this.values[key]?.includes(value);
    },

    async applyFilters() {
      return await useSubmitCategoryApplyFilters(
        this.categorySlug,
        this.values,
      );
    },
  },
});
