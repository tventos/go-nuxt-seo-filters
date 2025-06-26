import type { Product } from '~/entities/product';

export const useFetchProduct = async (slug: string): Promise<Product> => {
  const config = useRuntimeConfig();

  const { data, error } = await useFetch<Product>(
    `${config.public.apiBaseUrl}/product/${slug}`,
  );

  if (error.value) {
    throw new Error(error.value.message);
  }

  return data.value || [];
};
