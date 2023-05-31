function PostSignUp(student_id, student_name, course_name, email, telephone) {
  var myHeaders = new Headers();
  myHeaders.append("create", "chapter03");
  myHeaders.append("Content-Type", "application/json");

  var raw = JSON.stringify({
    student_id: student_id,
    student_name: student_name,
    course_name: course_name,
    email: email,
    telephone: telephone,
  });

  var requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: raw,
    redirect: "follow",
  };

  fetch("https://eo1hyy84q4xffbg.m.pipedream.net", requestOptions)
    .then((response) => response.text())
    .then((result) => console.log(result))
    .catch((error) => console.log("error", error));
}
function PushButton() {
  student_id = document.getElementById("student_id").value;
  student_name = document.getElementById("student_name").value;
  course_name = document.getElementById("course_name").value;
  email = document.getElementById("email").value;
  telephone = document.getElementById("telephone").value;
  PostSignUp(student_id, student_name, course_name, email, telephone);
}
