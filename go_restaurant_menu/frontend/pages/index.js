import { useState, useEffect } from 'react';
import { motion } from 'framer-motion';
import { fetchMenu, fetchCategories } from '../utils/api';
import Navbar from '../components/Navbar';

const PLACEHOLDER_IMG = "https://via.placeholder.com/400x300?text=No+Image";

export default function Home() {
  const [menuItems, setMenuItems] = useState([]);
  const [categories, setCategories] = useState([]);
  const [activeCategory, setActiveCategory] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const loadData = async () => {
      try {
        setLoading(true);
        const [menuData, categoriesData] = await Promise.all([
          fetchMenu(),
          fetchCategories()
        ]);
        setMenuItems(menuData);
        setCategories(categoriesData);
        if (categoriesData.length > 0) {
          setActiveCategory(categoriesData[0].id);
        }
      } catch (err) {
        setError('Failed to load menu');
        console.error(err);
      } finally {
        setLoading(false);
      }
    };
    loadData();
  }, []);

  const filteredItems = activeCategory
    ? menuItems.filter(item => item.category_id === activeCategory)
    : [];

  if (loading) {
    return (
      <div className="min-h-screen bg-white dark:bg-gray-900">
        <Navbar />
        <main className="container-custom py-12">
          <div className="text-center" role="status" aria-live="polite">Loading menu...</div>
        </main>
      </div>
    );
  }

  if (error) {
    return (
      <div className="min-h-screen bg-white dark:bg-gray-900">
        <Navbar />
        <main className="container-custom py-12">
          <div className="text-center text-red-500" role="alert">{error}</div>
        </main>
      </div>
    );
  }

  const activeCategoryObj = categories.find(cat => cat.id === activeCategory);

  return (
    <div className="min-h-screen bg-white dark:bg-gray-900">
      <Navbar />
      <main className="container-custom py-8">
        {/* Категории */}
        <nav aria-label="Категории меню" className="mb-8">
          <ul className="flex flex-wrap gap-4 justify-center">
            {categories.map((category) => (
              <li key={category.id}>
                <button
                  onClick={() => setActiveCategory(category.id)}
                  className={`px-6 py-2 rounded-full font-semibold transition-all text-base md:text-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-accent-500
                    ${activeCategory === category.id
                      ? 'bg-accent-500 text-white scale-105'
                      : 'bg-gray-100 dark:bg-gray-800 text-gray-700 dark:text-gray-200 hover:bg-accent-100 dark:hover:bg-accent-900'}
                  `}
                  aria-current={activeCategory === category.id ? 'page' : undefined}
                >
                  {category.icon && <span className="mr-2" aria-hidden="true">{category.icon}</span>}
                  {category.name}
                </button>
              </li>
            ))}
          </ul>
        </nav>

        {/* Заголовок категории */}
        {activeCategoryObj && (
          <motion.h2
            initial={{ opacity: 0, y: -10 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.4 }}
            className="text-3xl md:text-4xl font-serif font-bold text-center mb-10 text-accent-600 dark:text-accent-400"
          >
            {activeCategoryObj.name}
          </motion.h2>
        )}

        {/* Сетка блюд */}
        <section
          className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8"
          aria-label="Блюда категории"
        >
          {filteredItems.length === 0 && (
            <div className="col-span-full text-center text-gray-500 dark:text-gray-400 text-lg">Нет блюд в этой категории.</div>
          )}
          {filteredItems.map((item) => (
            <motion.article
              key={item.id}
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.5 }}
              className="bg-white dark:bg-gray-800 rounded-2xl shadow-lg p-6 flex flex-col items-center hover:shadow-2xl transition-shadow relative"
              tabIndex="0"
            >
              <div className="mb-4 w-full aspect-[4/3] rounded-xl overflow-hidden bg-gray-100 dark:bg-gray-700 flex items-center justify-center">
                <img
                  src={item.image_url && item.image_url.trim() !== '' ? item.image_url : PLACEHOLDER_IMG}
                  alt={item.name}
                  className="object-cover w-full h-full"
                  loading="lazy"
                  onError={e => { e.target.onerror = null; e.target.src = PLACEHOLDER_IMG; }}
                />
              </div>
              <h3 className="text-xl font-bold text-center mb-2">{item.name}</h3>
              <p className="text-gray-600 dark:text-gray-300 text-center mb-4 line-clamp-2">{item.description}</p>
              <div className="flex items-center justify-between w-full mt-auto">
                <span className="text-accent-500 font-bold text-lg">{item.price.toLocaleString()} сум</span>
                <button
                  className="ml-4 px-4 py-2 rounded-full bg-accent-500 text-white font-semibold shadow hover:bg-accent-600 focus:outline-none focus:ring-2 focus:ring-accent-500 transition-all"
                  type="button"
                >
                  Добавить
                </button>
              </div>
              {!item.is_available && (
                <div className="absolute top-4 right-4 bg-red-500 text-white text-xs px-3 py-1 rounded-full">Нет в наличии</div>
              )}
            </motion.article>
          ))}
        </section>
      </main>
      <footer className="bg-gray-100 dark:bg-gray-800 py-6 mt-12">
        <div className="container-custom text-center text-gray-600 dark:text-gray-300">
          <p>© 2024 Restaurant Menu. All rights reserved.</p>
        </div>
      </footer>
    </div>
  );
} 