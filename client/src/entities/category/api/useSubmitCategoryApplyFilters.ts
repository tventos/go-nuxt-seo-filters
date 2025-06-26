export const useSubmitCategoryApplyFilters = async (
  categorySlug: string,
  filters: { [key: string]: number[] },
): Promise<{ link: string }> => {
  const config = useRuntimeConfig();

  try {
    return await $fetch<{ link: string }>(
      `${config.public.apiBaseUrl}/category/filters/apply`,
      {
        method: 'POST',
        body: {
          category_slug: categorySlug,
          filters,
        },
      },
    );
  } catch (e) {
    throw new Error(e);
  }
};
