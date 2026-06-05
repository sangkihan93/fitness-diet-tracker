import './App.css';
import { AppLayout } from './components/layout/AppLayout';
import { DashboardPage } from './features/dashboard/DashboardPage';

function App() {
  return (
    <AppLayout>
      <DashboardPage />
    </AppLayout>
  );
}

export default App;
