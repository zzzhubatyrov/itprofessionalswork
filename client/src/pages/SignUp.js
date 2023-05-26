import axios from 'axios'
import React, { useState } from 'react'
import formSignUp from '../components/styles/auth.module.css'

const SignUpPage = () => {
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    
    const SingUp = async (e) => {
        e.preventDefault()
        try {
            const response = await axios.post('http://localhost:5000/login', {email: email, password: password})
            console.log(response)
        } catch (error) {
            console.log(error)
        }
        console.log({email, password})
    }
        
    return (
        <>
            <form className={formSignUp.SignUpForm} onSubmit={SingUp}>
                <input type="email" placeholder="Email" value={email} onChange={e => setEmail(e.target.value)} />
                <input type="password" placeholder="Password" value={password} onChange={e => setPassword(e.target.value)} />
                <input type="submit" />
            </form>            
        </>
    )
}

export { SignUpPage }

