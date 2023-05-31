var requestOptions = {
  method: 'GET',
  redirect: 'follow',
};

hasil = '';

fetch('https://api.reku.id/v2/bidask', requestOptions)
  .then((response) => response.text())
  .then((result) => tampilkan(result))
  .catch((error) => console.log('error', error));

function tampilkan(result) {
  iniJson = JSON.parse(result);
  console.log(iniJson);
  Lenght = iniJson.length;
  // console.log(Lenght);

  for (i = 0; i < Lenght; i++) {
    hasil += '<tr>';
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].accid + '</td>';
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].code + '</td>';
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].changepct + '</td>';
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].bid + '</td>';
    // hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].url + "</td>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].ask + '</td>';
    hasil += '</tr>';
  }

  document.getElementById('inidata').innerHTML = hasil;
}
