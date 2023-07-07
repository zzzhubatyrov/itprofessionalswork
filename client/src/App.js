import { Route, BrowserRouter as Router, Routes } from 'react-router-dom';
import './App.css';

import { Aboutpage } from './pages/Aboutpage';
import { Homepage } from './pages/Homepage';
import { Notfoundpage } from './pages/Notfoundpage';
import { Resumes } from './pages/Resumes';
import { Vacancies } from './pages/Vacancies';

import { Layout } from './components/Layout';
import { Auth } from './pages/Auth';
import { Profile } from './pages/Profile';
import { Employers } from './pages/Employers';

function App() {
  return (
    <Router>
        <div className="App">
            <div className="wrapper">
                <Routes>
                    <Route path='/' element={<Layout />}>
                        <Route index element={<Homepage />}/>
                        <Route path='vacancies' element={<Vacancies />}/>
                        <Route path='resumes' element={<Resumes />}/>
                        <Route path='about' element={<Aboutpage />}/>
                        <Route path='auth' element={<Auth />}/>
                        <Route path='*' element={<Notfoundpage />}/>

                        <Route path='profile' element={<Profile />}/>
                        // TODO Warning test link
                        <Route path='employers' element={<Employers />}/>
                    </Route>
                </Routes>
            </div>        
        </div>        
    </Router>
  );
}

export default App;
