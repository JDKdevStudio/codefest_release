import $ from "jquery";

$("#btnStart").on("click", async() => {
    const name = $("#floatingnameName1").val() as string;
    const correo = $("#floatingInputemail").val() as string;
    const pass = $("#floatingPassword").val() as string;
    const facultad = $("#a1").val() as string;
    const semestre = $("#a2").val() as string;
    const file = $("#floatingpic") as JQuery<HTMLInputElement>;
    console.log(name);
    console.log(correo);
    console.log(pass);
    console.log(facultad);
    console.log(semestre);

       let bodyContent = new FormData();
       bodyContent.append("name", name);
       bodyContent.append("email", correo);
       bodyContent.append("pwd", pass);
       bodyContent.append("faculty", facultad);
       bodyContent.append("semester", semestre);
       bodyContent.append("avatar", file.prop('files')[0], file.prop('files').name);
       
       let response = await fetch("http://25.54.85.248/users/register/", { 
         method: "POST",
         body: bodyContent,
       });
       
       let data = await response.text();
       console.log(data);
});