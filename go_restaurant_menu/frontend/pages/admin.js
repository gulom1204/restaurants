import { useState, useEffect } from 'react';
import { motion } from 'framer-motion';
import { fetchMenu, fetchCategories } from '../utils/api';

export default function Admin() {
  const [menuItems, setMenuItems] = useState([]);
  const [categories, setCategories] = useState([]);
  const [activeTab, setActiveTab] = useState('menu');
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
      } catch (err) {
        setError('Failed to load data');
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    loadData();
  }, []);

  if (loading) {
    return (
      <div className="min-h-screen bg-white dark:bg-gray-900">
        <div className="container-custom py-12">
          <div className="text-center">Loading...</div>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="min-h-screen bg-white dark:bg-gray-900">
        <div className="container-custom py-12">
          <div className="text-center text-red-500">{error}</div>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-white dark:bg-gray-900">
      <header className="bg-accent-500 text-white py-4">
        <div className="container-custom">
          <h1 className="text-3xl font-serif font-bold">Admin Panel</h1>
        </div>
      </header>

      <main className="container-custom py-8">
        {/* Admin Navigation */}
        <div className="flex space-x-4 mb-8">
          <button
            onClick={() => setActiveTab('menu')}
            className={`px-6 py-3 rounded-lg transition-all ${
              activeTab === 'menu'
                ? 'bg-accent-500 text-white'
                : 'bg-gray-100 dark:bg-gray-800 hover:bg-gray-200 dark:hover:bg-gray-700'
            }`}
          >
            Menu Management
          </button>
          <button
            onClick={() => setActiveTab('orders')}
            className={`px-6 py-3 rounded-lg transition-all ${
              activeTab === 'orders'
                ? 'bg-accent-500 text-white'
                : 'bg-gray-100 dark:bg-gray-800 hover:bg-gray-200 dark:hover:bg-gray-700'
            }`}
          >
            Orders
          </button>
          <button
            onClick={() => setActiveTab('staff')}
            className={`px-6 py-3 rounded-lg transition-all ${
              activeTab === 'staff'
                ? 'bg-accent-500 text-white'
                : 'bg-gray-100 dark:bg-gray-800 hover:bg-gray-200 dark:hover:bg-gray-700'
            }`}
          >
            Staff Management
          </button>
        </div>

        {/* Content Area */}
        <div className="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-6">
          {activeTab === 'menu' && (
            <div>
              <div className="flex justify-between items-center mb-6">
                <h2 className="text-2xl font-bold">Menu Items</h2>
                <button className="btn btn-primary">Add New Item</button>
              </div>
              <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                {menuItems.map((item) => (
                  <div key={item.id} className="card">
                    {item.image_url && (
                      <img
                        src={item.image_url}
                        alt={item.name}
                        className="w-full h-48 object-cover rounded-lg mb-4"
                      />
                    )}
                    <div className="flex justify-between items-start mb-2">
                      <h3 className="text-xl font-bold">{item.name}</h3>
                      <span className="text-accent-500 font-bold">
                        ${item.price.toFixed(2)}
                      </span>
                    </div>
                    <p className="text-gray-600 dark:text-gray-300 mb-4">
                      {item.description}
                    </p>
                    <div className="flex justify-between items-center">
                      <span
                        className={`px-2 py-1 rounded text-sm ${
                          item.is_available
                            ? 'bg-green-100 text-green-800'
                            : 'bg-red-100 text-red-800'
                        }`}
                      >
                        {item.is_available ? 'Available' : 'Unavailable'}
                      </span>
                      <div className="space-x-2">
                        <button className="btn btn-secondary">Edit</button>
                        <button className="btn btn-primary">Delete</button>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            </div>
          )}

          {activeTab === 'orders' && (
            <div>
              <h2 className="text-2xl font-bold mb-6">Active Orders</h2>
              <div className="space-y-4">
                {/* Здесь будет список заказов */}
                <p className="text-gray-600 dark:text-gray-300">
                  No active orders at the moment.
                </p>
              </div>
            </div>
          )}

          {activeTab === 'staff' && (
            <div>
              <div className="flex justify-between items-center mb-6">
                <h2 className="text-2xl font-bold">Staff Members</h2>
                <button className="btn btn-primary">Add Staff Member</button>
              </div>
              <div className="space-y-4">
                {/* Здесь будет список сотрудников */}
                <p className="text-gray-600 dark:text-gray-300">
                  No staff members added yet.
                </p>
              </div>
            </div>
          )}
        </div>
      </main>
    </div>
  );
} 