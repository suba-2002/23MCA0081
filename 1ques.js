import axios from 'axios';
import React, { useEffect, useState } from 'react';


const ProductsList = () => {
  const [products, setProducts] = useState([]);

  useEffect(() => {
   
    const apiEndpoint = 'http://localhost:3000/api/products';


    axios.get(apiEndpoint)
      .then(response => {
        setProducts(response.data);
      })
      .catch(error => {
        console.error('Error fetching data: ', error);
      })
  }, []);

  return (
    <div>
      <h1>Top N Products</h1>
      {products.map(product => (
        <div key={product.id}>
          <h2>{product.name}</h2>
          <p>Company: {product.company}</p>
          <p>Category: {product.category}</p>
          <p>Price: {product.price}</p>
          <p>Rating: {product.rating}</p>
          <p>Discount: {product.discount}</p>
          <p>Availability: {product.availability}</p>
        </div>
      ))}
    </div>
  );
};

export default ProductsList;
