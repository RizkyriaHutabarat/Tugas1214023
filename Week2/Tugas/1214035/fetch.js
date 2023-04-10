var requestOptions = {
    method: 'GET',
    redirect: 'follow'
  };
  
  hasil = "";
  fetch("https://thronesapi.com/api/v2/Characters", requestOptions)
    .then(response => response.text())
    .then(result => tampilkan (result))
    .catch(error => console.log('error', error));

    function tampilkan(result){
    iniJson = JSON.parse(result);
    console.log(iniJson);
    length = iniJson.length;
    

    for(i=0; i<length; i++)
        {
            hasil += "<tr>";
            hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>"+iniJson[i].fullName+"</td>";
            hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>"+iniJson[i].title+"</td>";
            hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>"+iniJson[i].family+"</td>";
            hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900 text-center'><img width='100' height='100' src='"+iniJson[i].imageUrl +"'/></td>";
            hasil += "</tr>";
        }


        document.getElementById("iniJson").innerHTML = hasil;
  }