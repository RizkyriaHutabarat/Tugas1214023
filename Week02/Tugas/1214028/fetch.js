var requestOptions = {
  method: 'GET',
  redirect: 'follow'
};

hasil=""

fetch("https://api.disneyapi.dev/characters", requestOptions)
  .then(response => response.text())
  .then(result => tampilkan(result))  .catch(error => console.log('error', error));

  function tampilkan(result){
    iniJson=JSON.parse(result);
    console.log(iniJson);
    Length = iniJson.data.length;
    //console.log(Length);

    for(i=0; i<Length; i++)
    {
        hasil += "<tr>";
        hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900 text-center'>"+iniJson.data[i]._id +"</td>";
        hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900 text-center'>"+iniJson.data[i].name +"</td>";
        hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900 text-center'><img src='"+iniJson.data[i].imageUrl +"'/></td>";
        hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900 text-center'>"+iniJson.data[i].films +"</td>";
        hasil += "<tr>";
    }
    document.getElementById("inidata").innerHTML = hasil;
}