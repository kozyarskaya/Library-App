function capi(string) {
    if (!string) return ""; // Обработка пустой строки
    return string.charAt(0).toUpperCase() + string.slice(1).toLowerCase();
}

// V A L U E S
const inputUserName = document.getElementById("nik")
const inputEmail = document.getElementById("email")
const inputPass1 = document.getElementById("password1")
const inputPass2 = document.getElementById("password2")
const reg = document.getElementById("registration_form")
//const msg = document.getElementById(".msg")


function registrate() {
    if (inputUserName.value.trim() === "" || inputEmail.value.trim() === "" ||
    inputPass1.value.trim() === "" || inputPass2.value.trim() === "") {
        alert(`Пожалуйста заполните все данные!`);
    } else if (inputPass1.value.trim() != inputPass2.value.trim()) {
        alert(`Пароли не совпадают!`);
    } else if (!inputEmail.value.trim().includes('@')) {
        alert('Адрес электронной почты должен содержать @');
    }else{
        resp();
    }
}

function resp() {
    const input_data = { 
        username: inputUserName.value.trim(),
        email: inputEmail.value.trim(),
        password: inputPass1.value.trim()
    };
    fetch('http://127.0.0.1:5501/api', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(input_data)
        })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            console.log(data);
            if (data.ok) { // Проверка статуса ответа
                reg.classList.add("invisible");
                alert(`Аккаунт создан)\n Ваш ник и почта: ${inputUserName.value} ${inputEmail.value}`);
            }
        })
        .catch(error => console.error('Ошибка:', error));

    /*if (data.ok) {
        reg.classList.add("invisible");
        alert(`Аккаунт создан)\n Ваш ник и почта: ${inputUserName.value} ${inputEmail.value}`);
    }*/
}