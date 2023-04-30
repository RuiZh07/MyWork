import React from "react";
import "./NFCtags.css";
import {Link} from 'react-router-dom';
import { IoMdArrowRoundBack } from 'react-icons/io';

const NFCTags = (props) => {
    return (
        <div className="main-account-setting">
            <div className="account-top">
                <div className="all-profile">
                    <Link to="/home.jsx">
                        <IoMdArrowRoundBack className="back-icon" />
                    </Link>
                    <span className="title-tag">Manage Tags</span>
                </div>
                <div className="Tag-info">
                    
                </div>

            </div>

        </div>
    );
};

export default NFCTags;

