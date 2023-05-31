var requestOptions = {
  method: 'GET',
  redirect: 'follow'
};

hasil=""

fetch("https://reqres.in/api/users?page=2", requestOptions)
  .then(response => response.text())
  .then(result => tampilkan(result))
  .catch(error => console.log('error', error));

  function tampilkan(result){
    iniJson=JSON.parse(result);
    console.log(iniJson);
    Length = iniJson.data.length;

    for(i = 0; i<Length; i++)
    {
        hasil += "<tr>";
        hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>"+iniJson.data[i].id+"</td>";
        hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>"+iniJson.data[i].first_name+"</td>";
        hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>"+iniJson.data[i].last_name+"</td>";
        hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'><img src='"+iniJson.data[i].avatar+"'/></td>";
        hasil += "</tr>";
    }



    document.getElementById("inidata").innerHTML = hasil; 
    //console.log(Length);
  }