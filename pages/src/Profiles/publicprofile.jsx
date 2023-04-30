import React from "react";
import "./publicprofile.css"
import { SocialIcon } from 'react-social-icons';
const PublicProfile = () => {
    return (
        <div className="public-profile">
            <div className="all">
                <div className="profile-container">
                    <button className="backboard"></button>
                    <div className="avatar-container">
                        <img src="https://images.pexels.com/photos/14028501/pexels-photo-14028501.jpeg?auto=compress&cs=tinysrgb&w=1600&lazy=load" alt="avatar" className="profileImage" />
                    </div>
                </div>
                <div className="lists">
                    <div className="name-university">
                        <span className="name">Jason Cromner</span>
                        <span className="university">Missouri University of Science and Technology</span>
                    </div>
                    <button className="social-icon">
                        <SocialIcon url="https://www.facebook.com/my-facebook-page" />
                        <span className="social-id">@fakeID</span>
                    </button>

                    <button className="social-icon">
                        <SocialIcon url="https://www.instagram.com/my-instagram-page" />
                        <span className="social-id">@id</span>
                    </button>

                    <button className="social-icon">
                        <SocialIcon url="https://www.twitter.com/my-twitter-page" />
                        <span className="social-id">@id</span>
                    </button>

                    <button className="social-icon">
                        <SocialIcon url="https://www.snapchat.com/my-snapchat-page" />
                        <span className="social-id">@id</span>
                    </button>
                </div>
            </div>
        </div>
    )
}

export default PublicProfile;