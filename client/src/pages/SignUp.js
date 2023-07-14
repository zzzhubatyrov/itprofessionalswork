import axios from 'axios';
import Cookies from 'js-cookie';
import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import formSignUp from '../components/styles/auth.module.css';


const SignUpPage = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const isAuthenticated = !!Cookies.get('ipw_cookie');
  const navigate = useNavigate();

  useEffect(() => {
    if (isAuthenticated) {
      navigate('/profile')
    }
  }, [isAuthenticated, navigate])

  const SignUp = async (e) => {
    e.preventDefault()
    try {
      const response = await axios.post('http://localhost:5000/login', {
        email: email,
        password: password,
      });
      const token = response.data.cookieValue;
      Cookies.set('ipw_cookie', token, { expires: 1 });
      sessionStorage.setItem('ipw_cookie', token);
      navigate('/profile')
    } catch (error) {
      console.log(error);
    }

    console.log({ email, password });
  };

  return (
    <>
      <form className={formSignUp.SignUpForm} onSubmit={SignUp}>
        <input
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <input type="submit" value="Авторизация" />
      </form>
    </>
  );
};

export { SignUpPage };

