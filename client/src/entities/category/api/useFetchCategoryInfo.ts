import type { CategoryInfo } from '~/entities/category';

export const useFetchCategoryInfo = async (filter): Promise<CategoryInfo> => {
  const config = useRuntimeConfig();

  const { data, error } = await useFetch<CategoryInfo>(
    `${config.public.apiBaseUrl}/category/info/${filter}`,
  );

  if (error.value) {
    throw new Error(error.value.message);
  }

  return data.value || [];
};
