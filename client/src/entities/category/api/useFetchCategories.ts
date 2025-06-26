import type { Category } from '~/entities/category';

export const useFetchCategories = async (): Promise<Category[]> => {
  const config = useRuntimeConfig();

  const { data, error } = await useFetch<Category[]>(
    `${config.public.apiBaseUrl}/category`,
  );

  if (error.value) {
    throw new Error(error.value.message);
  }

  return data.value || [];
};
