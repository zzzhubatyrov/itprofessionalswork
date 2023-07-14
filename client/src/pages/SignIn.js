import axios from 'axios'
import React, { useState } from 'react'
import formSignIn from '../components/styles/auth.module.css'

const SignInPage = () => {
    const [name, setName] = useState("")
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [confirmPassword, setConfirmPassword] = useState("")

    const SingIn = async (e) => {
        e.preventDefault()
        try {
            const response = await axios.post('http://localhost:5000/register', {name: name, email: email, password: password})
            console.log(response)
        } catch (error) {
            console.log(error)
        }
        console.log({name, email, password})
    }
        
    return (
        <>
            <form className={formSignIn.SignInForm} onSubmit={SingIn}>
                <input type="text" placeholder="Имя / Фамилия" value={name} onChange={e => setName(e.target.value)} />
                <input type="email" placeholder="E-mail" value={email} onChange={e => setEmail(e.target.value)} />
                <input type="password" placeholder="Пароль" value={password} onChange={e => setPassword(e.target.value)} />
                <input type="password" placeholder="Подтвердите пароль" value={confirmPassword} onChange={e => setConfirmPassword(e.target.value)} />
                <input type="submit" value="Зарегистрироваться" />
            </form>            
        </>
    )
}

export { SignInPage }

