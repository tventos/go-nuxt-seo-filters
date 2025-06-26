import type { Category } from '~/entities/category';

export interface ProductCharacteristic {
  filter_id: number;
  filter_name: string;
  value_id: number;
  value: string;
  seo_slug: string;
  has_seo: boolean;
}

export interface Product {
  id: number;
  name: string;
  description: string;
  slug: string;
  category: Omit<Category, 'children'>;
  characteristics: ProductCharacteristic[];
}

export interface ProductCategory {
  id: number;
  name: string;
  description: string;
  slug: string;
  filter_values: Array<{ value: string }>;
}
