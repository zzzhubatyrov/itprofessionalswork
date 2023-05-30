import React, { useEffect } from 'react';
import { NavLink } from 'react-router-dom';
import aboutPage from '../components/styles/aboutPage.module.css';

const setActive = (isActive) => isActive ? aboutPage.socialLink : ''

const Aboutpage = () => {
    useEffect(() => {
        document.title = 'О нас';
    })
    return (
        <div className={aboutPage.container}>
            <div className={aboutPage.developers}>
                <div className={aboutPage.devBlock}>
                    <div className={aboutPage.name}>Родион Жубатыров</div>
                    <hr />
                    <div className={aboutPage.social}>
                        <div className={`${aboutPage.linkBlock}`}>
                            <h3>VK </h3>
                            <NavLink className={setActive} to='https://vk.com/zzzhubatyrov'>zzzhubatyrov</NavLink>
                        </div>
                        <div className={`${aboutPage.linkBlock}`}>
                            <h3>GitHub </h3>
                            <NavLink className={setActive} to='https://github.com/zzzhubatyrov'>zzzhubatyrov</NavLink>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export { Aboutpage };

