const API_BASE_URL = 'http://localhost:8080/api';

export const fetchMenu = async () => {
  try {
    const response = await fetch(`${API_BASE_URL}/menu`);
    if (!response.ok) throw new Error('Failed to fetch menu');
    return await response.json();
  } catch (error) {
    console.error('Error fetching menu:', error);
    throw error;
  }
};

export const fetchCategories = async () => {
  try {
    const response = await fetch(`${API_BASE_URL}/categories`);
    if (!response.ok) throw new Error('Failed to fetch categories');
    return await response.json();
  } catch (error) {
    console.error('Error fetching categories:', error);
    throw error;
  }
};

export const fetchRestaurants = async () => {
  try {
    const response = await fetch(`${API_BASE_URL}/restaurants`);
    if (!response.ok) throw new Error('Failed to fetch restaurants');
    return await response.json();
  } catch (error) {
    console.error('Error fetching restaurants:', error);
    throw error;
  }
}; 