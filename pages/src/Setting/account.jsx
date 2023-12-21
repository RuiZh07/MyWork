import React, { useState } from "react";

import "./account.css"
import {Link} from 'react-router-dom';
import { IoMdArrowRoundBack } from 'react-icons/io';
import { IoIosArrowForward } from 'react-icons/io';
import Modal from 'react-modal';

const Account = () => {
    // State to manage the modal's open/closed state.
    const [isModalOpen, setIsModalOpen] = useState(false);

    // Function to open the modal.
    const openModal = () => {
        setIsModalOpen(true);
    }

    // Function to close the modal.
    const closeModal = () => {
        setIsModalOpen(false);
    }

    return (
        <div className="main-account-setting">
            <div className="account-top">
                <div className="all-profile">
                    <Link to="/setting.jsx">
                        <IoMdArrowRoundBack className="back-icon" />
                    </Link>
                    <span className="title-account">Account Security</span>
                </div>

                {/* Account Information Portion */}
                <div className="account-portion">
                    {/* Name Section */}
                    <div className="account-name">
                        <button>
                            <span>Name</span>
                            <span>Jason Cromner</span>
                        </button>
                    </div>

                    {/* University Section */}
                    <div className="account-university">
                        <button>
                            <span>University</span>
                            <span>MST</span> 
                        </button>
                    </div>

                    {/* Username Section */}
                    <div className="account-username">
                        <button>
                            <span>Username</span>
                            <span>JasonC </span>   
                        </button>
                    </div>

                    {/* Change Password Section */}
                    <div className="account-change-password">
                        <button onClick={openModal}>
                            <span>Change Password</span>
                            <span>*****</span>
                        </button>
                    </div>
                </div>
            </div>

            {/* Modal for Change Password */}
            <Modal isOpen={isModalOpen} onRequestClose={closeModal} contentLabel="Change Password" className="modal" overlayClassName="modal-overlay">
                <div className="modal-content-account">
                    <h2>Change Password</h2>
                    <form>
                        <label className="modal-label" htmlFor="current-password">Current Password</label>
                        <input type="password" id="current-password" name="current-password" />

                        <label className="modal-label" htmlFor="new-password">New Password</label>
                        <input type="password" id="new-password" name="new-password" />

                        <label className="modal-label" htmlFor="confirm-password">Confirm Password</label>
                        <input type="password" id="confirm-password" name="confirm-password" />

                        <div className="modal-buttons-account">
                            <button onClick={closeModal}>Cancel</button>
                            <button type="submit">Save</button>
                        </div>
                    </form>
                </div>
            </Modal>
        </div>
    );
};

export default Account;