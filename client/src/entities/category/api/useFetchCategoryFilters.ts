import type { CategoryFilter } from '~/entities/category';

export const useFetchCategoryFilters = async (
  categorySlug,
): Promise<CategoryFilter[]> => {
  const config = useRuntimeConfig();

  const { data, error } = await useFetch<CategoryFilter[]>(
    `${config.public.apiBaseUrl}/category/filters/list/${categorySlug}`,
  );

  if (error.value) {
    throw new Error(error.value.message);
  }

  return data.value || [];
};
