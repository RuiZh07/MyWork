import React from 'react';
import './home.css';
import { Link } from "react-router-dom";

const Home = () => {
    return (
        <div className="home">
            <div className="bottom">
                
                {/* Avatar Section */}
                <div className="avatar">
                    <img src="https://images.pexels.com/photos/14028501/pexels-photo-14028501.jpeg?auto=compress&cs=tinysrgb&w=1600&lazy=load" alt="avatar" className="home-profileImage" />
                </div>

                {/* User Info Section */}
                <div className="list">

                    <div className="name-university-home">
                        <span className="name-home">Jason J. Cromner</span>
                        <span className="university-home">Missouri University of Science and Technology</span>
                        <span className="link-home">Profile Link: 127.0.0.1:8000/JJCromer</span>
                    </div>

                    {/* Navigation Buttons */}
                    <div className="home-button">

                        {/* Link to Profile Page */}
                        <Link to="/profile.jsx">
                            <button><p>Profile</p></button>
                        </Link>

                        {/* Link to NFC Tags Page */}
                        <Link to="/NFCtags.jsx">
                            <button><p>Manage Tag</p></button>
                        </Link>

                        {/* Link to Setting Page */}
                        <Link to="/setting.jsx">
                            <button><p>Setting</p></button>
                        </Link>

                        {/* Link to Logout Page */}
                        <Link to="/login.jsx">
                            <button><p>Log out</p></button>
                        </Link>
                    </div>

                </div>

            </div>
        </div>

    )
}

export default Home;