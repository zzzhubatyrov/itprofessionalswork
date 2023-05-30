import Cookies from 'js-cookie';
import React from 'react';
import { Link } from 'react-router-dom';

const Homepage = () => {

    const isAuthenticated = !!Cookies.get('ipwSession');
    if (isAuthenticated) {
        return (
            <>
                <div className="centName">Здесь вы сможете создать своё Резюме!</div>
                <div className="name">IT Professionals Work</div>
                <Link className="btn" to="/profile">Профиль</Link>
            </>
        )
    }

    return (
        <>
            <div className="centName">Здесь вы сможете создать своё Резюме!</div>
            <div className="name">IT Professionals Work</div>
            <Link className="btn" to="/auth">Регистрация / Авторизация</Link>
        </>
    )
}

export { Homepage };

