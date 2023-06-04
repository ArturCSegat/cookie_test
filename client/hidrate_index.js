async function read_files(){
    let data = new FormData();
    console.log("reading");
    const input = document.getElementById("file_entry");
    data.append("rede", input.files[0]);
    register(data)
}

async function greet(){
    console.log("greeting")
    const data = await fetch("http://localhost:1337/all-nodes/", {method:"GET", credentials:"include"});
    const data_json = await data.json();
    console.log(data_json);
    
    for (const node of data_json.nodes){
        document.getElementById("greet").innerText += string_node_builder(node);
    }
}

async function register(file){
    console.log("registring")
    const url = "http://localhost:1337/upload_csv/" 
    await fetch(url, {method:"POST", credentials:"include", body:file});
    window.location.reload();
}


function string_node_builder(node){
    let string_node = "";
    string_node += "id: ";
    string_node += node.id.toString();
    string_node += " | ";

    string_node += "lat: ";
    string_node += node.lat.toString();
    string_node += " | ";

    string_node += "lng: ";
    string_node += node.lng.toString();
    string_node += " | ";

    string_node += "neighbours: ";
    string_node += node.neighbours.toString();
    string_node += "\n";

    return string_node;
}

window.onload = greet;

