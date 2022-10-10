import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import SignInPage from './pages/signin/SignInPage';
import Layout from './components/layout/Layout';
import InventoryPage from './pages/inventory/InventoryPage';
import TransactionPage from './pages/transaction/TransactionPage';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/signin" element={<SignInPage />} />
        <Route element={<Layout />}>
          <Route path="/inventory" element={<InventoryPage />} />
          <Route path="/transaction" element={<TransactionPage />} />
        </Route>
      </Routes>
    </Router>
  );
}

export default App;
