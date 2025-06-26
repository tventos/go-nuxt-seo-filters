export interface Category {
  id: number;
  name: string;
  slug: string;
  children?: Category[];
}

export interface CategoryInfo {
  title: string;
  description: string;
  h1: string;
  content: string;
  filter_value_id?: number;
  filter_slug?: string;
}

export interface CategoryFilterValues {
  id: number;
  value: string;
}

export interface CategoryFilter {
  id: number;
  name: string;
  slug: string;
  values: CategoryFilterValues[];
}
