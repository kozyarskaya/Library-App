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

    fetch('http://127.0.0.1:5501/registration', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(input_data)
    })
    .then(response => {
        if (!response.ok) {
            return response.json().then(errorData => {
                throw new Error(`Ошибка: ${errorData.message}`);
            });
        }
        return response.json();
    })
    .then(data => {
        if (data.ok === true) { // Проверка статуса ответа
            alert(`Аккаунт успешно создан!`);
           // reg.classList.add("invisible");
            //alert(`Аккаунт успешно создан!\nВаш ник: ${inputUserName.value}\nВаша почта: ${inputEmail.value}`);
        } else {
            alert('Ошибка при создании аккаунта: ' + data.message);
        }
    })
    .catch(error => {
        alert('Произошла ошибка при регистрации: ' + error.message);
    });
}
