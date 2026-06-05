import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import '../node_modules/@uswds/uswds/dist/css/uswds.min.css';
import './index.css';
import App from './App';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App />
  </StrictMode>
);
