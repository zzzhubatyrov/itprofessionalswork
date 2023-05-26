import { Route, BrowserRouter as Router, Routes } from 'react-router-dom';
import './App.css';

import { Aboutpage } from './pages/Aboutpage';
import { Homepage } from './pages/Homepage';
import { Notfoundpage } from './pages/Notfoundpage';
import { Resume } from './pages/Resume';
import { Vacansies } from './pages/Vacansies';

import { Layout } from './components/Layout';
import { Auth } from './pages/Auth';

function App() {
  return (
    <Router>
        <div className="App">
            <div className="wrapper">
                <Routes>
                    <Route path='/' element={<Layout />}>
                        <Route index element={<Homepage />}/>
                        <Route path='vacansies' element={<Vacansies />}/>
                        <Route path='resume' element={<Resume />}/>
                        <Route path='about' element={<Aboutpage />}/>
                        <Route path='auth' element={<Auth />}/>
                        <Route path='*' element={<Notfoundpage />}/>
                    </Route>
                </Routes>
            </div>        
        </div>        
    </Router>
  );
}

export default App;
