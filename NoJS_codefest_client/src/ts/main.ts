import $ from "jquery";

const insertNewMessage = (message: any) => {
    $.get("../templates/itemDisplay.html", function (data: string) {
        let response = data;
        let placeholders: Record<string, string> = {
            "{{title}}": message.title,
            "{{description}}": message.desc,
            "{{img}}": "http://25.54.85.248/assets/" + message.url_banner
        };
        //Replace in template
        for (let placeholder in placeholders) {
            if (placeholders.hasOwnProperty(placeholder)) {
                response = response.replace(placeholder, placeholders[placeholder]);
            }
        }



        $("#insert-star").append(response);
    });
}


$(async () => {
    let headersList = {
        "Authorization": "Bearer " + localStorage.getItem("access_token")
    }

    let response = await fetch("http://25.54.85.248/projects/", {
        method: "GET",
        headers: headersList
    });
    if (response.status == 200) {
        console.log(response)
        response.json().then((mydata) => {
            for (let a = 0; a < mydata.length; a++) {
                console.log(mydata[a]);
                insertNewMessage(mydata[a])
            }
        });
    }
});