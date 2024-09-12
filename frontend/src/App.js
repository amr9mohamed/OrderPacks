import React from 'react';
import OrderForm from './components/OrderForm';
import PackageSizeForm from './components/PackageSizeForm';

function App() {
  return (
      <div className="App">
        <h1>Package Management System</h1>
        <OrderForm />
        <PackageSizeForm />
      </div>
  );
}

export default App;
