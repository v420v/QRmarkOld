import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { CurrentUserProvider } from './context/AuthProvider';
import { SessionProvider } from './context/SessionProvider';
import { ChakraProvider } from '@chakra-ui/react'

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <BrowserRouter>
    <ChakraProvider>
    <SessionProvider>
    <CurrentUserProvider>
      <Routes>
        <Route path="/*" element={<App/>} />
      </Routes>
    </CurrentUserProvider>
    </SessionProvider>
    </ChakraProvider>
  </BrowserRouter>
);
