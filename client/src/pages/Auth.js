import { SignInPage } from './SignIn'
import { SignUpPage } from './SignUp'

const Auth = () => {
    return (
        <>
            <h1>SignUp and SignIn page</h1>
            <div className='forms'>
                <SignUpPage />
                <SignInPage />
            </div>
        </>
    )
}

export { Auth }

