import axios from 'axios';
import Cookies from 'js-cookie';
import React from 'react';
import { useNavigate } from 'react-router-dom';
import profileStyles from '../components/styles/profilePage.module.css';

const Profile = () => {
  const navigate = useNavigate();
  
  const Logout = async () => {
    try {
      await axios.post('http://localhost:5000/logout');
      sessionStorage.removeItem('ipwSession');
      Cookies.remove('ipwSession')
      navigate('/auth')
    } catch (error) {
      console.log(error);
    }
  }
  
  return (
    <div className={profileStyles.container}>
      <h1>This is profile</h1>
      <div className={profileStyles.userInfo}>
        <div className={profileStyles.userLogo}>
          <p>userLogo</p>
        </div>
        <div className={profileStyles.userName}>ФИО</div>
        <div className={profileStyles.userNumber}>Номер телефона</div>
        <div className={profileStyles.userCity}>Страна/Город проживания</div>
        <div className={profileStyles.userPosition}>Trainee, Junior, Middle, Senior, Team Lead</div>
        <div className={profileStyles.userWorkType}>Время/Занятость</div>
        <div className={profileStyles.profileInfo}>
          <div className={profileStyles.workExperienceBlock}>Опыт работы/компании</div>
          <div className={profileStyles.hardSkillsBlock}>Скиллы коммуникации/О себе</div>
          <div className={profileStyles.softSkillsBlock}>Опыт/Стек знаний</div>
        </div>
        <button onClick={Logout}>Выйти</button>
      </div>
    </div>
  );
};

export { Profile };

