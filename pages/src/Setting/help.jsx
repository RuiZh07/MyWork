import React from "react";
import "./help.css"
import {Link} from 'react-router-dom';
import AddCircleOutlinedIcon from '@mui/icons-material/AddCircleOutlined';
import { IoMdArrowRoundBack } from 'react-icons/io';
import { IoIosArrowForward } from 'react-icons/io';
const Help = () => {
    return (
        <div className="main-help">
            <div className="top">
                <div className="all-help">
                    <Link to="/setting.jsx">
                        <IoMdArrowRoundBack className="back-icon" />
                    </Link>
                    <span className="title-help">Help & Feedback</span>
                </div>

                <div className="question">
                    <h4>1. What is Wecave?</h4>
                    <p>It is a digital business card. It allows you to transfers all of the information 
                        on you Wecave profile with a single tap onto someone else’s phone.</p>

                    <h4>2. Is there a website for Wecave?</h4>
                    <p>No. While we are working on the website, you can use the app first. </p>
                    <div className="different-color">
                    <h4>3. Does Wecave require a subscription to use?</h4>
                    <p>No, Wecave doesn’t require.</p>

                    <h4>4. How do I connect my business information to Wecave product or QR code?</h4>
                    <p>Once you create your digital business card, 
                        you can use you in-app QR code to start sharing immediately. 
                        If you purchase a Wecave product, 
                        as soon as you receive it you will activate the device to your digital business card and all of you info will connect.</p>
                    <h3><center>Send Feeback</center></h3>
                    <center><p>Email to: zhaor0724@gmail.com</p></center>
                    </div>
                </div>
                
                
            </div>
        </div>
    );
};

export default Help;