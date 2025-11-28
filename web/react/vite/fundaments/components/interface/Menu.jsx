export default function Menu() {
    return (
        <nav>
            <div className="logo">
                <a href="/">
                    <img src="logo.png" alt="Logo" />
                </a>
            </div>
            <ul className="menu">
                <li><a href="/">Home</a></li>
                <li><a href="/about">About</a></li>
                <li><a href="/contact">Contact</a></li>
            </ul>
            <div className="actions">
                <button>Login</button>
                <button>Register</button>
            </div>
        </nav>
    )
}