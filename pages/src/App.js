import Login from "./Login/login.jsx";
import SignUp from "./SignUp/signup.jsx";
import Navbar from "./Navbar/navbar.jsx";
import Profile from "./Profiles/profile.jsx";
import PublicProfile from "./Profiles/publicprofile.jsx";
import AddProfile from "./Profiles/addprofile.jsx";
import Setting from "./Setting/setting.jsx"
import Home from "./Home/home";
import Help from "./Setting/help";
import Account from "./Setting/account";
import ActivateTag from "./Tag/activateTag";
import NFCTags from "./Tag/NFCtags";

import {
  createBrowserRouter,
  RouterProvider,
  Route,
  Link,
} from "react-router-dom";


function App() {

  /*const currentUser = false;

  const Layout = () => {
    return (
      <div>
        <Navbar />
        <div style={{display: "flex"}}>

        </div>
      </div>
    )
  }
  */
  const router = createBrowserRouter([

    {
      path: "/login.jsx",
      element: <Login/>,
    },
    {
      path: "/signup.jsx",
      element: <SignUp/>,
    },
    {
      path: "/profile.jsx",
      element: <Profile/>,
    },
    {
      path: "publicprofile.jsx",
      element: <PublicProfile/>,
    },
    {
      path: "home.jsx",
      element: <Home/>,
    },
    {
      path: "addprofile.jsx",
      element: <AddProfile/>,
    },
    {
      path: "setting.jsx",
      element: <Setting/>,
    },
    {
      path: "help.jsx",
      element: <Help/>,
    },
    {
      path: "account.jsx",
      element: <Account/>,
    },
    {
      path: "activateTag.jsx",
      element: <ActivateTag/>,
    },
    {
      path: "NFCtags.jsx",
      element:<NFCTags/>,
    }

  ]);

  return (
    <div>
      <RouterProvider router={router} />
    </div>
  );
}

export default App;
