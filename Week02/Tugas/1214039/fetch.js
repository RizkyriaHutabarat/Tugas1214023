var requestOptions = {
  method: "GET",
  redirect: "follow",
};

hasil = "";

fetch("https://picsum.photos/v2/list?page=1&limit=100", requestOptions)
  .then((response) => response.text())
  .then((result) => tampilkan(result))
  .catch((error) => console.log("error", error));

function tampilkan(result) {
  iniJson = JSON.parse(result);
  console.log(iniJson);
  Lenght = iniJson.length;
  // console.log(Lenght);

  for (i = 0; i < Lenght; i++) {
    hasil += "<tr>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].id + "</td>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].author + "</td>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].width + "</td>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].height + "</td>";
    // hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].url + "</td>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'><img src='" + iniJson[i].download_url + "'></td>";
    hasil += "</tr>";
  }

  document.getElementById("inidata").innerHTML = hasil;
}
