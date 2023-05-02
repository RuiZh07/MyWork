import React from "react";
import "./profile.css"
import {Link} from 'react-router-dom';
import AddCircleOutlinedIcon from '@mui/icons-material/AddCircleOutlined';
import { IoMdArrowRoundBack } from 'react-icons/io';

const Profile = () => {
    return (
        <div className="main-profile">
            <div className="top">
                <div className="all-profile">
                    <Link to="/home.jsx">
                        <IoMdArrowRoundBack className="back-icon" />
                    </Link>
                    <span className="title-profile">Profile</span>
                </div>
                
                <div className="all">
                    <Link to="/publicprofile.jsx">
                        <button>For party</button>
                    </Link>

                    <Link to="/signup.jsx">
                        <button>For career</button>
                    </Link>
                </div>

                <div className="add">
                    <Link to="/addprofile.jsx">
                        <AddCircleOutlinedIcon clasName="plus" />
                        <p className="add-new">Add new Profile</p>
                    </Link>
                    
                </div>
                
            </div>
        </div>
    );
};

export default Profile;