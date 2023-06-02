async function read_files(){
    let data = new FormData();
    console.log("reading");
    const input = document.getElementById("file_entry");
    data.append("file", input.files[0]);
    register(data)
}

async function greet(){
    console.log("greeting")
    const data = await fetch("http://localhost:3000/name", {method:"GET", credentials:"include"});
    const data_json = await data.json();
    console.log(data_json);
    
    document.getElementById("greet").innerText = data_json.name;
}

async function register(file){
    console.log("registring")
    const url = "http://localhost:3000/register" 
    await fetch(url, {method:"POST", credentials:"include", body:file});
    window.location.reload();
}

window.onload = greet;

