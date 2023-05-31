var requestOptions = {
  method: "GET",
  redirect: "follow",
};
hasil = "";

fetch("https://api.reku.id/v2/ads", requestOptions)
  .then((response) => response.text())
  .then((result) => tampilkan(result))
  .catch((error) => console.log("error", error));
function tampilkan(result) {
  iniJson = JSON.parse(result);
  // console.log(iniJson);
  length = iniJson.length;
  // console.log(iniJson);

  for (i = 0; i < length; i++) {
    hasil += "<tr>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].url + "</td>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].alt + "</td>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'><img src='" + iniJson[i].img + "'/></td>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].popup + "</td>";
    hasil += "</tr>";
  }

  document.getElementById("inidata").innerHTML = hasil;
}
