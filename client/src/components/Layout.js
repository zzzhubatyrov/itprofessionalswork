import { Link, Outlet } from 'react-router-dom';

const Layout = () => {
    return (
        <>
            <header className="header">
                <Link className="logo" to="/">IPW</Link>
                <nav className="navbar">
                    <Link className="navItems vacansies" to="/vacancies">Вакансии</Link> {/* Для компаний*/}
                    {/*<Link className="navItems resume" to="/resume">Резюме</Link>*/}
                    <Link className="navItems contacts" to="/about">О нас</Link>
                </nav>
            </header>
            <div className="container">
                <Outlet />
            </div>

            {/* <footer>
                <div className="copyright">Zhubatyrov©</div>
                <div className="social"></div>
            </footer> */}
        </>
    )
}

export { Layout };

