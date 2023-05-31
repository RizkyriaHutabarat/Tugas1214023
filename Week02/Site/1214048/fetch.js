var myHeaders = new Headers();
myHeaders.append("Cookie", "connect.sid=s%3AM6CfLJhCFlu92tNvS7cRegIIcR8rhhUG.AN2Ss3OKnMLlBJEwcDELKykDb293dBuH%2FhX1M3mZI2w");

var requestOptions = {
  method: 'GET',
  headers: myHeaders,
  redirect: 'follow'
};

hasil=""

fetch("https://24pullrequests.com/users.json?page=", requestOptions)
  .then(response => response.text())
  .then(result => tampilkan(result))
  .catch(error => console.log('error', error));

  function tampilkan(result){
    iniJson=JSON.parse(result);
    console.log(iniJson);
    length = iniJson.length;
    console.log(length);

    for(i=0; i<length; i++)
    {
        hasil += "<tr>";
        hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" +iniJson[i].organization.login+"</td>";
        hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" +iniJson[i].organization.avatar_url+"</td>";
        // hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" +iniJson[i].organization[0].link+"</td>";
        
        hasil += "</tr>";
    }
    document.getElementById("inidata").innerHTML = hasil;
  }