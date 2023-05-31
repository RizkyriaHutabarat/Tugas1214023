var myHeaders = new Headers();
myHeaders.append("Cookie", "connect.sid=s%3AGkR2_95H8CudJl645iDbAYQ5Gaz3Jc_v.nFTFo5yg319rSYpVHBs6DXnCaOJDOHQd92wh4E9WYV8");

var raw = "";

var requestOptions = {
  method: 'GET',
  headers: myHeaders,
  redirect: 'follow'
};

let hasil = ""

fetch("https://cat-fact.herokuapp.com/facts", requestOptions)
  .then(response => response.text())
  .then(result => tampilkan(result))
  .catch(error => console.log('error', error));

let tampilkan = (result) => {
  // console.log(result)
  let inijson=JSON.parse(result)
  console.log(inijson)
  length = inijson.length
  console.log(length)
  for (let i = 0; i < length; i++) {
    const element = inijson[i];
    hasil += "<tr>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>"+element.text+"</td>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>"+element.updatedAt+"</td>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>"+element.status.verified+"</td>";
    hasil += "</tr>";
  }
  document.getElementById("inidata").innerHTML=hasil
}
