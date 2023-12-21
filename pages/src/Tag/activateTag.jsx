import React from "react";
import "./activateTag.css"
const ActivateTag = () => {
    return (
        <div>
          {/* Title and Instruction */}  
          <h1>Activate Your Tag</h1>
          <span>Enter Your WaCave Email To Link This Tag</span><br /><br />

          {/* Activation Form */}
          <form action="/activateTag" method="post">
            <input type="hidden" name="tagHash"/>
            <label htmlFor="userEmail">Email</label><br />

            <input type="email" id="userEmail" name="userEmail" required /><br /><br />
            <label htmlFor="confirmEmail">Confirm Email</label><br />
            
            <input type="email" id="confirmEmail" name="confirmEmail" required /><br /><br />
            <button type="submit">Activate</button>
          </form>

          <p>Don't have an account yet? <a href="/auth/signup">Sign up</a></p>
        </div>
    );
};


export default ActivateTag;
