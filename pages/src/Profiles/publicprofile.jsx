import React from "react";
import {Link} from 'react-router-dom';
import "./publicprofile.css"
import { SocialIcon } from 'react-social-icons';
import { IoMdArrowRoundBack } from 'react-icons/io';
const PublicProfile = () => {
    return (
        <div className="main-profile">
            <div className="top">
                <div className="all-profile">
                    <Link to="/profile.jsx">
                        <IoMdArrowRoundBack className="back-icon" />
                    </Link>
                    <span className="title-profile">Party Profile</span>
                </div>
            
            <div className="all">
                
                <div className="profile-container">
                    <button className="backboard"></button>
                
                <div className="avatar-container">
                    <img src="https://images.pexels.com/photos/14028501/pexels-photo-14028501.jpeg?auto=compress&cs=tinysrgb&w=1600&lazy=load" alt="avatar" className="profileImage" />
                </div>
                </div>
                <div className="lists">
                    <div className="name-university-public">
                        <span className="name-public">Jason Cromner</span>
                        <span className="university-public">Missouri University of Science and Technology</span>
                    </div>
                    <button className="social-icon">
                        <SocialIcon url="https://www.facebook.com/my-facebook-page" style={{ height: 35, width: 35 }}/>
                        <span className="social-id">@ID</span>
                    </button>

                    <button className="social-icon">
                        <SocialIcon url="https://www.instagram.com/my-instagram-page" style={{ height: 35, width: 35 }}/>
                        <span className="social-id">@ID</span>
                    </button>

                    <button className="social-icon">
                        <SocialIcon url="https://www.twitter.com/my-twitter-page" style={{ height: 35, width: 35 }}/>
                        <span className="social-id">@ID</span>
                    </button>

                    <button className="social-icon">
                        <SocialIcon url="https://www.snapchat.com/my-snapchat-page" style={{ height: 35, width: 35 }}/>
                        <span className="social-id">@ID</span>
                    </button>
                </div>
            </div>
            </div>
        </div>
    )
}

export default PublicProfile;