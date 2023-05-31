var myHeaders = new Headers();
myHeaders.append("Cookie", "connect.sid=s%3AGkR2_95H8CudJl645iDbAYQ5Gaz3Jc_v.nFTFo5yg319rSYpVHBs6DXnCaOJDOHQd92wh4E9WYV8");

var raw = "";

var requestOptions = {
  method: 'GET',
  headers: myHeaders,
  redirect: 'follow'
};

let hasil = ""

fetch("https://jakpost.vercel.app/api/category/indonesia", requestOptions)
  .then(response => response.text())
  .then(result => tampilkan(result))
  .catch(error => console.log('error', error));

let tampilkan = (result) => {
  // console.log(result)
  let inijson=JSON.parse(result)
  console.log(inijson)
  length = inijson.posts.length
  console.log(length)
  for (let i = 0; i < length; i++) {
    const element = inijson.posts[i];
    hasil += "<tr>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>"+element.title+"</td>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'><img src='"+element.image+"' /></td>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>"+element.headline+"</td>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>"+element.category+"</td>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>"+element.pusblised_at+"</td>";
    hasil += "</tr>";
  }
  document.getElementById("inidata").innerHTML=hasil
}
