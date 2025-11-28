export default function Menu() {
    return (
        <nav>
            <div className="logo">
                <Link to="/">
                    <img src="logo.png" alt="Logo" />
                </Link>
            </div>
            <ul className="menu">
                <li><Link to="/">Home</Link></li>
                <li><Link to="/about">About</Link></li>
                <li><Link to="/contact">Contact</Link></li>
            </ul>
            <div className="actions">
                <button>Login</button>
                <button>Register</button>
            </div>
        </nav>
    )
}