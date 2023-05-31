async function greet(){
    const data = await fetch("http://localhost:3000/name", {method:"GET", credentials:"include"});
    const data_json = await data.json();
    console.log(data_json);
    
    document.getElementById("greet").innerText = data_json.name;
}

async function register(value){
    console.log(value);
    const url = "http://localhost:3000/register/" + value 
    await fetch(url, {method:"GET", credentials:"include"});
    window.location.reload();
}

window.onload = greet;

