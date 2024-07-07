import './App.css';
import Register from './components/Register/Register';
import Login from './components/Login/Login'
import { Routes, Route } from 'react-router-dom';
import Layout from './components/Applications/Layout';
import Home from './components/Home/Home';
import RequireAuth from './components/Applications/RequireAuth';
import Qrmark from './components/Qrmark/Qrmark';
import About from './components/About/About';
import NotFound from './components/NotFound';
import SchoolDetail from './components/SchoolDetail/SchoolDetail'
import Admin from './components/Admin/Admin';
import Verify from './components/Verify/Verify';

function App() {
  return (
    <Routes>
      <Route path="/" element={<Layout/ >}>
        <Route path="register" element={<Register/>}></Route>
        <Route path="login" element={<Login/>}></Route>
        <Route path="about" element={<About/>}></Route>
        <Route path="verify/:token" element={<Verify/>}></Route>

        <Route exact path="/" element={<RequireAuth allowedRoles={[]} />}>
          <Route exact path="" element={<Home />} />
          <Route exact path="qrmark" element={<Qrmark/>}/>
          <Route exact path="school/:id" element={<SchoolDetail/>}/>
        </Route>

        <Route exact path="/" element={<RequireAuth allowedRoles={["admin"]} />}>
          <Route exact path="admin" element={<Admin />} />
        </Route>

        <Route path="*" element={<NotFound/>}></Route>
      </Route>
    </Routes>
  );
}

export default App;