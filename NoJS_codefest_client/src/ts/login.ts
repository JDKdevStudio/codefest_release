import $ from "jquery";
import Cookies from "js-cookie";



$("#btnStart").on("click", async() => {
    const username = $("#floatingInputData1").val();
    const pass = $("#floatingPassword").val();

    console.log(username);
    console.log(pass);

    let bodyContent = new FormData();
    bodyContent.append("email", username as string);
    bodyContent.append("password", pass as string);

    let response = await fetch("http://25.54.85.248/users/login/?email=zarate2@gmail.com&password=contrase%C3%B1asegura", {
        method: "GET"
    }).then((data) =>{
        const datajson = data.json().then((_data)=>{
            console.log(_data.token);
            localStorage.setItem("token", encodeURIComponent(_data.token));
            console.log(localStorage.getItem("token"))
            window.location.replace("./main.html");
        });    
    }
    );    
});