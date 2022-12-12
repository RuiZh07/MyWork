const ndef = new NFEFReader();

async function writeNFC(){
    try {
        await ndef.write({
            records: [{ recordType: "url", data: "https://themaverse.io/"}]
        });
    } catch {
        const text = document.getElementById('information').innerHTML
        alert("Write failed, Please try again!"+ text )
    };
}


// const loginForm = document.getElementById("login-form");
// const loginButton = document.getElementById("login-form-submit");
// const loginErrorMsg = document.getElementById("login-error-msg");

// loginButton.addEventListener("click", check());

// function check(){
//     const username = loginForm.username.value;
//     const password = loginForm.password.value;
//     alert("show something")
//     if (username === "user" && password === "1234") {
//         alert("You have successfully logged in.")
        
//         } else {
//             alert("Invalid")
//         }
//}

function check(){
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    if (username === "admin" && password === "admin"){
        alert("You have succesfully logged in")
        window.location.replace("./sub.html")
    }else {
        alert("Invalid")
    }
}
// loginButton.addEventListener("click", (e) => {
//     e.preventDefault();
//     const username = loginForm.username.value;
//     const password = loginForm.password.value;
//     alert("test")

//     if (username === "user" && password === "1234") {
//         alert("You have successfully logged in.")
    
//     } else {
//         alert("Invalid")
//         loginErrorMsg.style.opacity = 1;
//     }
// })