import React, { useState } from "react";
import "./NFCtags.css";
import {Link} from 'react-router-dom';
import { IoMdArrowRoundBack } from 'react-icons/io';


const NFCTags = (props) => {
    const [showModal, setShowModal] = useState(false);

    // Function to open the activation modal.
    const activetag = () => {
        setShowModal(true);
    }

    return (
        <div className="main-tag">
            <div className="account-top">
                <div className="all-profile">
                    <Link to="/home.jsx">
                        <IoMdArrowRoundBack className="back-icon" />
                    </Link>
                    <span className="title-tag">Manage Tags</span>
                </div>

                {/* Tag Information Section */}
                <div className="Tag-info">
                    {/* Activate Tag Button */}
                    <button onClick={activetag}>Activate Tag</button>
                    {/* Order New Tag Button */}
                    <button>Order new tag</button>
                </div>

                {/* Activation Modal */}
                {showModal && (
                    <div className="modal-setting">
                        <div className="modal-content-setting">
                            <h3>Activate Your Tag</h3>
                            <span>Enter Your WaCave Email To Link This Tag</span><br /><br />

                            {/* Hidden Tag Hash Input */}
                            <input type="hidden" name="tagHash"/>

                            {/* Email Input */}
                            <label htmlFor="userEmail">Email</label><br />
                            <input type="email" id="userEmail" name="userEmail" required /><br /><br />

                            {/* Confirm Email Input */}
                            <label htmlFor="confirmEmail">Confirm Email</label><br />
                            <input type="email" id="confirmEmail" name="confirmEmail" required /><br /><br />

                            {/* Modal Buttons for Cancel and Activate */}
                            <div className="modal-buttons-account">
                                <button onClick={() => setShowModal(false)}>Cancel</button>
                                <button onClick={() => setShowModal(false)}>Activate</button>
                            </div>
                        </div>
                    </div>
                )}
                
            </div>
        </div>
    );
};

export default NFCTags;
