import React, { useState } from "react";
import "./addprofile.css"
import {Link} from 'react-router-dom';
import { IoMdArrowRoundBack } from 'react-icons/io';

const AddProfile = () => {
    const [socialMediaList, setSocialMediaList] = useState([
        { name: "", link: "" },
        { name: "", link: "" },
        { name: "", link: "" }
      ]);
    const handleAddSocialMedia = () => {
        setSocialMediaList([...socialMediaList, { name: '', link: '' }]);
       
    };

    const handleInputChange = (index, event) => {
        const { name, value } = event.target;
        const list = [...socialMediaList];
        list[index][name] = value;
        setSocialMediaList(list);
    };
    return (
        <div className="add-profile">
            <div className="top-add-new">
                <div className="add-new-profile">
                <Link to="/profile.jsx">
                    <IoMdArrowRoundBack className="back-icon-add" />
                </Link>
                <span className="title-add"> Add Profile</span>
                </div>
                <div className="bottom-add-profile">
                    <span className="profile-name-container">
                        <label for="profile name">Profile Name</label>
                        <input type="text" size= "25" />
                    </span>
        
                    <div className="social-media-container">
                    {socialMediaList.map((socialMedia, index) => (
                        <div key={index}>
                            <div className="social-media-label">
                                <label for={`Social Media ${index + 1}`}>Social Media {index + 1}</label>
                            </div>
                            <div className="social-media-input">
                            <input type="text" className="input1" 
                                name={`Social Media ${index + 1}`} 
                                onChange={(event) => handleInputChange(index, event)} />
                            <input type="text" name={`link ${index + 1}`} 
                                onChange={(event) => handleInputChange(index, event)} />
                            </div>
                        </div>
                    ))}
                    <button className="button-add-more" onClick={handleAddSocialMedia}>+</button>
                    </div>
                    
                    <Link to="/profile.jsx">
                        <button className="save-profile">Save</button>
                    </Link>
                    
                </div>
            
            </div>
        </div>
    );
};

export default AddProfile;