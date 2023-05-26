import { Link } from 'react-router-dom'

const Notfoundpage = () => {
    return (
        <div>
            <h1>This page doesn't exist. Go <Link to="/">home</Link></h1>
        </div>
        
    )
}

export { Notfoundpage }

