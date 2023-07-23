import React, {useEffect} from "react"
import {useCookies} from 'react-cookie'
import {useNavigate} from 'react-router-dom';
const AuthChecker = () => {
    const [cookies] = useCookies(['ipw_cookie'])
    const navigate = useNavigate();
    useEffect(() => {
        const isAuthenticated = !!cookies.ipw_cookie

        if (!isAuthenticated) {
            navigate('/auth');
        }
    }, [cookies])
}

export default AuthChecker;