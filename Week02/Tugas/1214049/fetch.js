var myHeaders = new Headers();
myHeaders.append("Cookie", "connect.sid=s%3AB-5o56pbb1ziXq9i1jszvrJChy4go_Hc.cu3rvBI4vJlppszC6I8MzXVV00Nbf2bQYVooG1Lf8SI");

var requestOptions = {
  method: 'GET',
  headers: myHeaders,
  redirect: 'follow'
};

hasil=""

fetch("https://api.pokemontcg.io/v2/cards?q=name:gardevoir", requestOptions)
  .then(response => response.text())
  .then(result => tampilkan(result))
  .catch(error => console.log('error', error));

  function tampilkan(result){
    iniJson = JSON.parse(result);
    console.log(iniJson);
    Length = iniJson.data.length;
    console.log(Length);

    for(i=0; i<Length; i++)
    {
      hasil += "<tr>";
      hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" +iniJson.data[i].id+" </td>"
      hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" +iniJson.data[i].name+" </td>"
      hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" +iniJson.data[i].supertype+" </td>"
      hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'><img src='" +iniJson.data[i].images.small+"'/> </td>"
      hasil += "</td>";
    }

    document.getElementById("inidata").innerHTML=hasil;
  }