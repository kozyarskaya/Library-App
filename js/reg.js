// Функция для капитализации первой буквы строки
function capi(string) {
    if (!string) return ""; // Обработка пустой строки
    return string.charAt(0).toUpperCase() + string.slice(1).toLowerCase();
}

// V A L U E S
const inputUserName = document.getElementById("nik"); // Получение элемента ника
const inputEmail = document.getElementById("email"); // Получение элемента электронной почты
const inputPass1 = document.getElementById("password1"); // Получение элемента пароля
const inputPass2 = document.getElementById("password2"); // Получение элемента повторного пароля
const registrationForm = document.getElementById("registration_form"); // Получение формы регистрации
const authorizationForm = document.getElementById("authorization_form"); // Получение формы авторизации
const bookContainer = document.getElementById("book-container"); // Получение контейнера для книг
const inputEmailA = document.getElementById("email-a");
const inputPassA = document.getElementById("password-a");


    // Обработка события отправки формы регистрации
    registrationForm.addEventListener('submit', function(event) {
        event.preventDefault(); // Предотвратить стандартное поведение отправки формы
        // Проверка заполнения всех полей
        if (inputUserName.value.trim() === "" || inputEmail.value.trim() === "" ||
            inputPass1.value.trim() === "" || inputPass2.value.trim() === "") {
            alert(`Пожалуйста заполните все данные!`);
        } else if (inputPass1.value.trim() != inputPass2.value.trim()) {
            alert(`Пароли не совпадают!`);
        } else if (!inputEmail.value.trim().includes('@')) {
            alert('Адрес электронной почты должен содержать @');
        } else {
            signUp(); // Вызов функции отправки данных на сервер
        }
    });

    // Обработка события отправки формы авторизации
    authorizationForm.addEventListener('submit', function(event) {
        event.preventDefault(); // Предотвратить стандартное поведение отправки формы
        // Проверка заполнения всех полей

        if (inputEmailA.value.trim() === "" || inputPassA.value.trim() === "") {
            alert(`Пожалуйста заполните все данные!`);
        } else {
            login(); // Вызов функции отправки данных на сервер
        }
    });

    // Функция переключения на форму авторизации
    function authorization() {
        registrationForm.classList.add("invisible"); // Скрытие формы регистрации
        authorizationForm.classList.remove("invisible"); // Отображение формы авторизации
    }

    // Функция переключения на форму регистрации
    function registrate() {
        registrationForm.classList.remove("invisible"); // Отображение формы регистрации
        authorizationForm.classList.add("invisible"); // Скрытие формы авторизации
    }

    // Функция отправки данных на сервер
    function signUp() {
        const input_data = { 
            username: inputUserName.value.trim(), // Получение значения ника
            email: inputEmail.value.trim(), // Получение значения электронной почты
            password: inputPass1.value.trim() // Получение значения пароля
        };

        // Отправка POST запроса на сервер
        fetch('http://127.0.0.1:5501/api/registration', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' }, // Установка заголовка Content-Type
            body: JSON.stringify(input_data) // Преобразование данных в JSON строку
        })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`); // Обработка ошибок HTTP
            }
            return response.json(); // Парсинг ответа в JSON
        })
        .then(data => {
            console.log(data); // Вывод полученных данных в консоль
            if (data.ok) { // Проверка статуса ответа
                registrationForm.classList.add("invisible"); // Скрытие формы регистрации
                bookContainer.classList.remove("invisible"); // Отображение контейнера для книг

                alert(`Аккаунт создан)\n Ваш ник и почта: ${inputUserName.value} ${inputEmail.value}`); // Вывод сообщения об успешной регистрации
            }
        })
        .catch(error => console.error('Ошибка:', error)); // Обработка ошибок
    }

    // Функция отправки данных на сервер для авторизации
    function login() {
        const input_data = { 
            email: document.getElementById("email-a").value.trim(), // Получение значения электронной почты
            password: document.getElementById("password-a").value.trim() // Получение значения пароля
        };

        // Отправка POST запроса на сервер
        fetch('http://127.0.0.1:5501/api/login', {
            method: 'POST',
            headers: { 
                'Authorization': `Bearer ${token}`, // добавление токина в заголовок
                'Content-Type': 'application/json' }, // Установка заголовка Content-Type
            body: JSON.stringify(input_data) // Преобразование данных в JSON строку
        })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`); // Обработка ошибок HTTP
            }
            return response.json(); // Парсинг ответа в JSON
        })
        .then(data => {
            console.log(data); // Вывод полученных данных в консоль
            if (data.ok) { // Проверка статуса ответа
                authorizationForm.classList.add("invisible"); // Скрытие формы авторизации
                bookContainer.classList.remove("invisible"); // Отображение контейнера для книг

                alert(`С возвращением!\n  ${inputEmailA.value}`); // Вывод сообщения об успешной авторизации
            }
        })
        .catch(error => console.error('Ошибка:', error)); // Обработка ошибок
    }

    // Отправка GET запроса для получения данных о книгах
    if (bookContainer) { 
        fetch('test.json')
            .then(response => {
                if (!response.ok) {
                    throw new Error('HTTP error ' + response.status); 
                }
                return response.json(); 
            })
            .then(data => {
                console.log(data); 
                data.forEach(book => {
                    const bookItem = document.createElement('div'); 
                    bookItem.classList.add('book-item'); 

                    const title = document.createElement('h2'); 
                    title.textContent = book.title; 

                    const date = document.createElement('p'); 
                    date.textContent = `Дата: ${book.date}`; 

                    const text = document.createElement('p'); 
                    text.textContent = book.text; 

                    const link = document.createElement('a'); 
                    link.href = book.link; 
                    link.textContent = 'Читать далее'; 

                    bookItem.appendChild(title); 
                    bookItem.appendChild(date); 
                    bookItem.appendChild(text); 
                    bookItem.appendChild(link); 

                    bookContainer.appendChild(bookItem); 
                });
            })
            .catch(error => console.error('Ошибка:', error)); 
    } else {
        console.log('Элемент с ID "book-container" не найден'); 
    }

