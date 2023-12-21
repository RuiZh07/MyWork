import React from "react";
import "./navbar.css";
import AccountCircleOutlinedIcon from '@mui/icons-material/AccountCircleOutlined';
import SettingsOutlinedIcon from '@mui/icons-material/SettingsOutlined';
import ContactPageOutlinedIcon from '@mui/icons-material/ContactPageOutlined';
const Navbar = () => {
    return (
        <div className='background'>
            <div className="navbar">
                {/* User Account Icon */}
                <AccountCircleOutlinedIcon className="icons" />

                {/* Contact Page Icon */}
                <ContactPageOutlinedIcon className="icons" />

                {/* Settings Icon */}
                <SettingsOutlinedIcon className="icons"/>
            </div>    
        </div>
    )
   
}

export default Navbar;