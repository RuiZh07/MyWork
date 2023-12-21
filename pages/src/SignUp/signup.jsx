import React from 'react';
import './signup.css';
import { Link } from "react-router-dom";

const SignUp = () => {
    return (
        <div className="signup">
            <div className="bottom">
                {/* Sign Up Form */}
                <form>
                    <p> Create an Account</p>

                    <label for="name">Name</label>
                    <input type="text" size= "35" placeholder="" />
                    
                    <label for="university">University</label>
                    <input type="text" size= "35" placeholder="" />
   
                    <label for="email">Email</label>
                    <input type="text" size= "35" placeholder="" />
                                     
                    <label for="password">Password</label>
                    <input type="password" size= "35" placeholder="" />

                    <label for="password">Re-enter Password</label>
                    <input type="password" size= "35" placeholder="" />

                    <div className="container">
                        <span>Already have an account?</span>
                        <Link to="/login.jsx">Log in</Link>
                    </div>

                    {/* Sign Up Button */}
                    <Link to="/login.jsx">
                        <button>Sign up</button>
                    </Link>
                </form>
            </div>
        </div>
    )
}

export default SignUp;