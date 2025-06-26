import type { ProductCategory } from '~/entities/product';

export const useFetchCategoryProducts = async (
  filter = '',
): Promise<ProductCategory[]> => {
  const config = useRuntimeConfig();

  const { data, error } = await useFetch<Product[]>(
    `${config.public.apiBaseUrl}/category/products/${filter}`,
  );

  if (error.value) {
    throw new Error(error.value.message);
  }

  return data.value || [];
};
