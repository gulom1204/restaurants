import React from 'react';

const Home = () => {
  const categories = ['Appetizers', 'Main Course', 'Desserts', 'Beverages'];

  return (
    <div className="container">
      <h1>Welcome to the Restaurant Menu</h1>
      <p>This is the frontend for the restaurant menu project.</p>
      <div className="menu">
        <h2>Menu Categories</h2>
        <ul>
          {categories.map((category, index) => (
            <li key={index} className="category-item">{category}</li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default Home; 