var myHeaders = new Headers();
myHeaders.append("Cookie", "__cf_bm=zoH_8G_fOQpHWMdM5YOc4nukFdJcsu0SlnyqS4elU_U-1677568638-0-AT2HLk6cnGHyfKb9GMvATLTvSY7OB2jk53qHBUrEOmq4i2QZiDq4Wg130GCrf7CWbqGW1jaOcMW63l1JsUNIkQ8=");

var requestOptions = {
  method: "GET",
  headers: myHeaders,
  redirect: "follow",
};

hasil = "";

fetch("https://indodax.com/api/pairs", requestOptions)
  .then((response) => response.text())
  .then((result) => tampilkan(result))
  .catch((error) => console.log("error", error));

function tampilkan(result) {
  iniJson = JSON.parse(result);
  console.log(iniJson);
  Length = iniJson.length;
  // console.log(Length);

  for (i = 0; i < Length; i++) {
    hasil += "<tr>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].id + "</td>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].coingecko_id + "</td>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].traded_currency + "</td>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900 text-center'><img src='" + iniJson[i].url_logo_png + "'/></td>";
    hasil += "</tr>";
  }

  document.getElementById("inidata").innerHTML = hasil;
}
