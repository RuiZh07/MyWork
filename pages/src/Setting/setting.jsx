import React, { useState } from "react";
import "./setting.css"
import {Link} from 'react-router-dom';
import AddCircleOutlinedIcon from '@mui/icons-material/AddCircleOutlined';
import { IoMdArrowRoundBack } from 'react-icons/io';
import { IoIosArrowForward } from 'react-icons/io';
const Setting = () => {
    const deleteAccount = (email) => {
        fetch('/delete-account', {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ email })
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Error deleting account');
            }
            // Redirect to login or other page
            window.location.href = '/Login/login.jsx';
        })
        .catch(error => {
            console.error(error);
            alert('Failed to delete account');
        });
    }

    const [showModal, setShowModal] = useState(false);
    const [email, setEmail] = useState("");
  
    const handleDeleteAccount = () => {
      setShowModal(true);
    }
  
    const confirmDeleteAccount = () => {
      // Call a function to delete the account with the provided email address
      deleteAccount(email);
      setShowModal(false);
    }

    return (
        <div className="main-profile">
            <div className="top">
                <div className="all-setting">
                    <Link to="/home.jsx">
                        <IoMdArrowRoundBack className="back-icon" />
                    </Link>
                    <span className="setting-title-profile">Setting</span>
                </div>
                
                <div className="avatar-setting">
                    <img src="https://images.pexels.com/photos/14028501/pexels-photo-14028501.jpeg?auto=compress&cs=tinysrgb&w=1600&lazy=load" alt="avatar" className="setting-profile-Image" />
                </div>

                <div className="name-university-setting">
                    <span className="name-settings">Jason J. Cromner</span>
                    <span className="university-settings">Missouri University of Science and Technology</span>
                    
                </div>

                <div className="settings-portion">
                    <div className="account-setting">
                        <Link to="/account.jsx">
                            <button>
                                <span>Account Security</span>
                                <IoIosArrowForward className="back-icon" />   
                        </button>
                        </Link>
                    </div>

                    <div className="subscription">
                        <Link to="/home.jsx">
                            <button>
                                <span>Subscription</span>
                                <IoIosArrowForward className="back-icon" />
                            </button>
                        </Link>
                    </div>

                    <div className="help-feedback">
                        <Link to="/help.jsx">
                            <button>
                                <span>Help & Feedback</span>
                                <IoIosArrowForward className="back-icon" />  
                        </button>
                        </Link>
                    </div>
                    <div className="delete-account">
                            <button onClick={handleDeleteAccount}>
                                <span>Delete Account</span>
                        </button>
                    </div>
                </div>
            </div>

            {showModal && (
                <div className="modal-setting">
                    <div className="modal-content-setting">
                        <h3>Are you sure you want to delete your account?</h3>
                        <p>Please enter your email address to confirm the action:</p>
                        <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} />
                        <div className="modal-buttons-account">
                            <button onClick={() => setShowModal(false)}>Cancel</button>
                            <button onClick={confirmDeleteAccount}>Delete Account</button>
                        </div>
                    </div>
                </div>
            )}   
        </div>
    );
};

export default Setting;