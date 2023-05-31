var myHeaders = new Headers();
myHeaders.append("Cookie", "connect.sid=s%3AM6CfLJhCFlu92tNvS7cRegIIcR8rhhUG.AN2Ss3OKnMLlBJEwcDELKykDb293dBuH%2FhX1M3mZI2w");

var requestOptions = {
  method: "GET",
  headers: myHeaders,
  redirect: "follow",
};

hasil = "";

fetch("https://bobsburgers-api.herokuapp.com/characters/?limit=9&skip=25", requestOptions)
  .then((response) => response.text())
  .then((result) => tampilkan(result))
  .catch((error) => console.log("error", error));

function tampilkan(result) {
  let iniJson = JSON.parse(result);
  console.log(iniJson);
  Lenght = iniJson.length;
  // console.log(Lenght);

  for (i = 0; i < Lenght; i++) {
    hasil += "<tr>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].id + "</td>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].name + "</td>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'><img src='" + iniJson[i].image + "'/></td>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].gender + "</td>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].hairColor + "</td>";
    hasil += "</tr>";
  }

  document.getElementById("inidata").innerHTML = hasil;
}
