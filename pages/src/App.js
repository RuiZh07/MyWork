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
  Routes,
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
      path: "/",
      element: <Login/>,
    },
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
    },
    {
      path: "navbar.jsx",
      element: <Navbar/>,
    }

  ]);

  return (

    <div>
      <RouterProvider router={router}>
        <Routes>
          <Route path="/" element={<Login />} />
          <Route path="/login" element={<Login />} />
          <Route path="/signup" element={<SignUp />} />
          <Route path="/profile" element={<Profile />} />
          <Route path="/publicprofile" element={<PublicProfile />} />
          <Route path="/home" element={<Home />} />
          <Route path="/addprofile" element={<AddProfile />} />
          <Route path="/setting" element={<Setting />} />
          <Route path="/help" element={<Help />} />
          <Route path="/account" element={<Account />} />
          <Route path="/activateTag" element={<ActivateTag />} />
          <Route path="/NFCtags" element={<NFCTags />} />
          <Route path="/navbar" element={<Navbar />} />
        </Routes>
      </RouterProvider>
    </div>
  );
}

export default App;
