var requestOptions = {
    method: 'GET',
    redirect: 'follow'
  };
  
  hasil = "";

  fetch("https://hp-api.onrender.com/api/characters", requestOptions)
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
          hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].name + "</td>";
          hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].species + "</td>";
          hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].house + "</td>";
          hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'><img src='" + iniJson[i].image + "'></td>";
          hasil += "</tr>";
        }
      
        document.getElementById("inidata").innerHTML = hasil;
      }
      