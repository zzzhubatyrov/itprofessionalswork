import { SignInPage } from './SignIn'
import { SignUpPage } from './SignUp'
import React, {useState} from 'react';
import UseChecker from '../components/AuthChecker';
import styles from '../components/styles/auth.module.css'
// import {Link} from "react-router-dom";

const Auth = () => {
    const [isRegistration, setIsRegistration] = useState(true);
    const toggleForm = () => {
        setIsRegistration(!isRegistration)
    }

    return (
        <div className={styles.container}>
            <h1 className={styles.header_form}>
                {isRegistration ? (
                    <>
                        <button className={`${styles.header_link} ${styles.btn}`}>
                            Авторизация
                        </button>
                        /
                        <button className={`${styles.header_link} ${styles.btn}`} onClick={toggleForm}>Регистрация</button>
                    </>
                ) : (
                    <>
                        <button className={`${styles.header_link} ${styles.btn}`}>Регистрация</button>
                        /
                        <button className={`${styles.header_link} ${styles.btn}`} onClick={toggleForm}>
                            Авторизация
                        </button>
                    </>
                )}
            </h1>
            {isRegistration ? <SignUpPage /> : <SignInPage />}
            <UseChecker />
        </div>
    )
}

export { Auth }

