import type { LocationQuery } from '#vue-router';

export const buildCategoryFilters = (path: string[], query: LocationQuery) => {
  let filter = path.join('/');
  filter += '?' + new URLSearchParams(query).toString();

  return filter;
};
