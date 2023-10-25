import $ from "jquery";



$("#reg-project").on("click",async()=>{
    var titulo = $("#floatingname").val() as string;
    var descripcion = $("textarea").val();
    var archivo = $("#floatingpic").val();
    var urlVideo = $("#floatingurl").val();


    let headersList = {
        "Accept": "*/*",
        "User-Agent": "Thunder Client (https://www.thunderclient.com)",
        "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0LCJuYW1lIjoiWmFyYXRlIiwiZW1haWwiOiJ6YXJhdGUyQGdtYWlsLmNvbSIsImF2YXRhcl9pbWFnZSI6IjE2OTgyNTEzNTEwNDg2MDQxMDBfNTEuanBlZyIsImZhY3VsdHkiOiJzaXN0ZW1hcyIsInNlbWVzdGVyIjoxLCJzdGF0dXMiOnRydWUsInVzZXJfdHlwZSI6MiwidG9rZW4iOiIiLCJleHAiOjE2OTgzNjAwNTJ9.nOUH9GjuL5ay8JUD8i95q8Ky-D27a406sSGswez_T08"
       }
       
       let bodyContent = new FormData();
       bodyContent.append("title", "The Nexus Battles II");
       bodyContent.append("desc", "Juego de cartas y ecommerce en línea");
       bodyContent.append("url_video", "https://www.youtube.com/watch?v=83bc-086rPU");
       bodyContent.append("banner", "c:\Users\josed\Downloads\image.png");
       
       let response = await fetch("http://localhost/projects/", { 
         method: "POST",
         body: bodyContent,
         headers: headersList
       });
       
       let data = await response.text();
       console.log(data);
       
})