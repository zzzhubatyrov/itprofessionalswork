import { Link } from 'react-router-dom'

const Homepage = () => {
    return (
        <>
            <div className="centName">Здесь вы сможете создать своё Резюме!</div>
            <div className="name">IT Professionals Work</div>
            <Link className="btn" to="/auth">Регистрация / Авторизация</Link>
        </>
    )
}

export { Homepage }

