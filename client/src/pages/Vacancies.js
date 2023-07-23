import vacancyStyle from '../components/styles/vacancy.module.css'
// import Cookies from 'js-cookie';
// import UseChecker from "../components/AuthChecker";
import {useCookies} from "react-cookie";
import {useEffect} from "react";
import {useNavigate} from 'react-router-dom';

const Vacancies = () => {
    const [cookies] = useCookies(['ipw_cookie'])
    // const navigate = useNavigate();
    const isAuthenticated = !!cookies.ipw_cookie
    useEffect(() => {
        if (!isAuthenticated) {
            // navigate('/auth');
            console.log("asd-asd")
        }
    }, [cookies])

    return (
        <div className={vacancyStyle.container}>
            <div className={vacancyStyle.elastic_search}>
                <input className={vacancyStyle.search_bar} placeholder="Поиск" />
                {/*<div className={vacancyStyle.list_levelBlock}>*/}
                {/*    <div className={vacancyStyle.level_block}>Квалификация</div>*/}
                {/*    <div className={vacancyStyle.level_block}>Квалификация</div>*/}
                {/*    <div className={vacancyStyle.level_block}>Квалификация</div>*/}
                {/*</div>*/}
            </div>
            <div className={vacancyStyle.vacancy_container}>
                <div className={vacancyStyle.vacancy_block}>
                    <div className={vacancyStyle.vacancy_block_header}>
                        <div className={vacancyStyle.logo}><img src={require('../components/img/Logo-IPW.jpg')}  alt='logo'/></div>
                        <div className={vacancyStyle.block}>
                            <div className={vacancyStyle.block_name_tag}>
                                <div className={vacancyStyle.company_name}>IT Professionals Work</div>
                                <div className={vacancyStyle.tag}>@ipw</div>
                            </div>
                            <div className={vacancyStyle.emp_position}>Senior Golang Developer</div>
                            <div className={vacancyStyle.city_workTime}>Уфа / Полный рабочий день</div>
                        </div>
                    </div>
                    <div className={vacancyStyle.vacancy_block_desc}>
                        Мы знаем, что такое скорбь, сидим вместе, создаем нашу элиту.
                        Донец сусципит элементум нунц, сэд хендрерит метус конгуэ сэд.
                        Нулла фасилиси. Преддверие и внутренняя оболочка.
                        Ничто не имеет силы, ничто не говорит о том, что есть, конец всему сущему.
                        Сид ид ниси в justo ultrices tempor. Преддверие вариуса сагиттиса маттиса.
                        Сед виверра велит, чтобы я рисковал, а пеллентеск был на месте.
                    </div>
                    <div className={vacancyStyle.vacancy_block_footer}>
                        <div className={vacancyStyle.vacancy_block_skills}>
                            UI/UX дизайнер, Средний (Middle) • Figma Design •
                            UI/UX дизайн • Брендирование • Проектирование интерфейсов •
                            Исследование пользователя • Дизайн продукта • Дизайн иконок •
                            Разработка фирменного стиля • Дизайн логотипов • Прототипирование
                        </div>
                        <div className={vacancyStyle.vacancy_block_btn}>
                            <button className={vacancyStyle.subBtn}>Откликнуться</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export { Vacancies }
