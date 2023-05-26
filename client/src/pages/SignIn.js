import axios from 'axios'
import React, { useState } from 'react'
import formSignIn from '../components/styles/auth.module.css'

const SignInPage = () => {
    const [surname, setSurname] = useState("")
    const [name, setName] = useState("")
    const [lastname, setLastname] = useState("")
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    
    const SingIn = async (e) => {
        e.preventDefault()
        try {
            const response = await axios.post('http://localhost:5000/register', {surname: surname, name: name, lastname: lastname, email: email, password: password})
            console.log(response)
        } catch (error) {
            console.log(error)
        }
        console.log({surname, name, lastname, email, password})
    }
        
    return (
        <>
            <form className={formSignIn.SignInForm} onSubmit={SingIn}>
                <input type="text" placeholder="Фамилия" value={surname} onChange={e => setSurname(e.target.value)} />
                <input type="text" placeholder="Имя" value={name} onChange={e => setName(e.target.value)} />
                <input type="text" placeholder="Отчество" value={lastname} onChange={e => setLastname(e.target.value)} />
                <input type="email" placeholder="Email" value={email} onChange={e => setEmail(e.target.value)} />
                <input type="password" placeholder="Password" value={password} onChange={e => setPassword(e.target.value)} />
                <input type="submit" />
            </form>            
        </>
    )
}

export { SignInPage }

