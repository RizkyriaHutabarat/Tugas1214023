var requestOptions = {
    method: 'GET',
    redirect: 'follow'
  };

  hasil=""

  fetch("https://www.cheapshark.com/api/1.0/deals?storeID=1&upperPrice=15", requestOptions)
    .then(response => response.text())
    .then(result => tampilkan(result))
    .catch(error => console.log('error', error));

    function tampilkan(result) {
        iniJson = JSON.parse(result);
        console.log(iniJson);
        Lenght = iniJson.length;
        // console.log(Lenght);
      
        for (i = 0; i < Lenght; i++) {
          hasil += "<tr>";
          hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].title + "</td>";
          hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].gameID + "</td>";
          hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].salePrice + "</td>";
          hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].normalPrice + "</td>";
          // hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].url + "</td>";
          hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'><img src='" + iniJson[i].thumb + "'></td>";
          hasil += "</tr>";
        }
      
        document.getElementById("inidata").innerHTML = hasil;
      }