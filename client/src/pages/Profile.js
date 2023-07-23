import axios from 'axios';
import {useNavigate} from 'react-router-dom';
import styles from '../components/styles/profilePage.module.css';
import UseChecker from "../components/AuthChecker";
import {useEffect, useState} from "react";
import Cookies from 'js-cookie';
import ModalWindow from "../components/ModalWindows";

const Profile = () => {
  const navigate = useNavigate();
  const [userData, setUserData] = useState(null);
  // const [userDataRole, setUserDataRole] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const token = Cookies.get("ipw_cookie"); // Получаем токен из local storage или куки
        if (token) {
          const response = await axios.get('http://localhost:5000/get-user', {
            withCredentials: true,
            headers: {
              Authorization: `Bearer ${token}`, // Включаем токен в заголовок запроса
            },
          });
          console.log(response.data)
          console.log(response.data.role)
          setUserData(response.data);
          // setUserDataRole(response.data.role)
        }
      } catch (error) {
        console.error(error);
      }
    };

    fetchData();
  }, []);

  const logout = async () => {
    try {
      await axios.post('http://localhost:5000/logout');
      sessionStorage.removeItem('ipw_cookie');
      Cookies.remove('ipw_cookie');
      navigate('/auth');
    } catch (error) {
      console.log(error);
    }
  };

  const [isModalOpen, setIsModalOpen] = useState(false);

  const openModal = () => {
    setIsModalOpen(true);
  };

  const closeModal = () => {
    setIsModalOpen(false);
  };


  return (
      <div className={styles.container}>
        <div className={styles.profileBlock}>
          {userData && (
              <>
                <div className={styles.profileInfoBlock}>
                  <div className={styles.profileInfoBlock_container}>
                    <div className={styles.photo}>
                      <img src={userData.photo} alt='logo' />
                      {/*<img src={require('../components/img/zhubatyrov.jpg')} alt='logo' />*/}
                    </div>
                    <div className={styles.userInfoBlock_left}>
                      <div className={styles.user_nameAgeCity}>
                        <div className={styles.user_nameTag}>
                          <div className={styles.user_name}>{userData.name}</div>
                          <div className={styles.user_role}>{userData.role.name}</div>
                          <div className={styles.user_tag}>{userData.tag}</div>
                        </div>
                        <div className={styles.user_ageCity}>
                          <div className={styles.user_age}><span>Мужчина, {userData.age} лет</span></div>
                          <div className={styles.user_city}>Город: {userData.location}</div>
                        </div>
                        {/*<span>не готов к переезду<br/>не готов к командировкам</span>*/}
                      </div>
                      <div className={styles.user_contacts}>
                        <div className={styles.user_email}>{userData.email}</div>
                        <div className={styles.user_number}>{userData.number}</div>
                      </div>
                    </div>
                    <button className={`${styles.cv_btn} ${styles.non_active}`}>Скачать CV</button>
                    <button className={`${styles.for_hr}`} onClick={openModal}>Работодателям</button>
                    <ModalWindow isOpen={isModalOpen} onClose={closeModal} />
                    {/*${styles.non_active}*/}
                  </div>
                  <div className={styles.brnd}>
                    <span><button onClick={logout}>Выйти</button></span>
                    <span>itprofessionalswork</span>
                  </div>
                </div>
                <div className={styles.userInfoBlock_right}>
                  <div className={styles.aboutUser}>
                    <div className={styles.aboutMe_head}>Обо мне:</div>
                    <div className={styles.aboutMe_description}>{userData.description}</div>
                  </div>
                  <div className={styles.educationBlock}>
                    <div className={styles.education_head}>Среднее специальное образование:</div>
                    <div className={styles.education_description}>
                      <div className={styles.year_education}>2024</div>
                      <div className={styles.education_nameBlock}>
                        <div className={styles.education_name}>Нефтекамский Машиностроительный Колледж</div>
                        <div className={styles.education_specialize}>Информационные системы и программирование, Программист</div>
                      </div>
                    </div>
                  </div>
                  <div className={styles.resume}></div>
                </div>
              </>
          )}
        </div>
        <UseChecker />
      </div>
  );
};

export { Profile };