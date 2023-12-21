import React from 'react';
import './login.css';
import { Link } from "react-router-dom";

const Login = () => {
    return (
        <div className="login">
            <div className="bottom">
                <form>
                    {/* Avatar Section */}
                    <div className="avatar">
                        <img src="https://images.pexels.com/photos/14028501/pexels-photo-14028501.jpeg?auto=compress&cs=tinysrgb&w=1600&lazy=load" alt="avatar" className="login-profile-Image" />
                    </div>

                    {/* Email Input Field */}
                    <label for="email">Email</label>
                    <input type="text" size= "35" placeholder="" />

                    {/* Password Input Field */}
                    <label for="password">Password</label>
                    <input type="password" size= "35" placeholder="" />

                    <Link to="/home.jsx">
                        <button>Login</button>
                    </Link>

                    {/* Container for Sign Up */}
                    <div className="container">
                        <span>Don't have an account?</span>
                        {/* Link to Sign Up Page */}
                        <Link to="/signup.jsx" className="sign">Sign Up</Link>
                    </div>
                </form>
            </div>
        </div>
    )
}

export default Login;