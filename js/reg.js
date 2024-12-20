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
const bookContainerG = document.getElementById("book-containerG"); // Получение контейнера для книг
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
                bookContainerG.classList.remove("invisible"); // Отображение контейнера для книг

                alert(`С возвращением!\n  ${inputEmailA.value}`); // Вывод сообщения об успешной авторизации
            }
        })
        .catch(error => console.error('Ошибка:', error)); // Обработка ошибок
    }
    
    // Выбор элементов по их ID
const bookContainer = document.getElementById('book-container');
const addBookBtn = document.getElementById('add-book-btn');
const addBookForm = document.getElementById('add-book-form');
const addBookFormContent = document.getElementById('add-book-form-content');

// Загрузка книг с сервера используя fetch API
if (bookContainer) { 
fetch('test.json')
    .then(response => response.json())
    .then(data => {
        // Перебор массива книг и создание элементов для каждой книги
        data.forEach(book => {
            const bookItem = document.createElement('div');
            bookItem.classList.add('book-item');
            bookItem.dataset.id = book._id; // Добавление атрибута data-id для идентификации книги

            
            // Создание элементов для отображения информации о книге
        

            const title = document.createElement('h2');
            title.textContent = book.title;

            const author = document.createElement('p');
            author.textContent = `Автор: ${book.author}`;

            const date = document.createElement('p');
            date.textContent = `Дата: ${book.date}`;

            const preview = document.createElement('p');
            preview.textContent = book.preview;

            const text = document.createElement('p');
            text.textContent = book.text;

            // Создание кнопок для редактирования и удаления книги
            const editBtn = document.createElement('button');
            editBtn.textContent = 'Редактировать';
            editBtn.onclick = () => editBook(book._id); // Вызов функции редактирования книги
            editBtn.classList.add("btnleft"); // Добавление класса для стилизации кнопки

            const deleteBtn = document.createElement('button');
            deleteBtn.textContent = 'Удалить';
            deleteBtn.onclick = () => deleteBook(book._id); // Вызов функции удаления книги
            deleteBtn.classList.add("btnleft"); // Добавление класса для стилизации кнопки

            // Добавление элементов в контейнер книги
            bookItem.appendChild(title);
            bookItem.appendChild(author);
            bookItem.appendChild(date);
            bookItem.appendChild(preview);
            //bookItem.appendChild(text);
            bookItem.appendChild(editBtn);
            bookItem.appendChild(deleteBtn);

            // Добавление контейнера книги в общий контейнер книг
            bookContainer.appendChild(bookItem);
        });
    })
    .catch(error => console.error('Ошибка:', error)); // Обработка ошибок при загрузке данных

    // Обработка клика по кнопке добавления новой книги
addBookBtn.onclick = () => {
    addBookForm.classList.remove('invisible'); // Показать форму добавления книги
};

// Обработка отправки формы добавления новой книги
addBookFormContent.addEventListener('submit', (e) => {
    e.preventDefault(); // Предотвратить стандартное поведение отправки формы

    // Получение значений полей формы
    const title = document.getElementById('title').value;
    const author = document.getElementById('author').value;
    const date = document.getElementById('date').value;
    const preview = document.getElementById('preview').value;
    //const text = document.getElementById('text').value; // Комментировано, но можно раскомментировать если нужно

    // Отправка POST запроса на сервер для добавления новой книги
    fetch('http://127.0.0.1:5501/api/books/', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ title, author, date, preview, text: '' }) // Преобразование данных в JSON строку
    })
    .then(response => response.json())
    .then(data => {
        // Создание нового элемента книги после успешного добавления на сервер
        const bookItem = document.createElement('div');
        bookItem.classList.add('book-item');
        bookItem.dataset.id = data._id; // Добавление атрибута data-id

        const titleElement = document.createElement('h2');
        titleElement.textContent = data.title;

        const authorElement = document.createElement('p');
        authorElement.textContent = `Автор: ${data.author}`;

        const dateElement = document.createElement('p');
        dateElement.textContent = `Дата: ${data.date}`;

        const previewElement = document.createElement('p');
        previewElement.textContent = data.preview;

       // const textElement = document.createElement('p');
        //textElement.textContent = data.text;

        const editBtn = document.createElement('button');
        editBtn.textContent = 'Редактировать';
        editBtn.onclick = () => editBook(data._id);

        const deleteBtn = document.createElement('button');
        deleteBtn.textContent = 'Удалить';
        deleteBtn.onclick = () => deleteBook(data._id);

        bookItem.appendChild(titleElement);
        bookItem.appendChild(authorElement);
        bookItem.appendChild(dateElement);
        bookItem.appendChild(previewElement);
       // bookItem.appendChild(textElement);
        bookItem.appendChild(editBtn);
        bookItem.appendChild(deleteBtn);

        // Добавление нового элемента книги в контейнер книг
        bookContainer.appendChild(bookItem);

        // Скрыть форму добавления книги и сбросить ее поля
        addBookForm.classList.add('invisible');
        addBookFormContent.reset();
    })
    .catch(error => console.error('Ошибка:', error)); // Обработка ошибок при добавлении книги
});
    // Функция редактирования книги
// Редактирование книги
function editBook(id) {
    try {
        const bookItem = bookContainer.querySelector(`.book-item[data-id="${id}"]`);
        if (!bookItem) throw new Error('Элемент не найден');

        const title = bookItem.querySelector('h2').textContent;
        const author = bookItem.querySelector('p:nth-child(2)').textContent.replace('Автор: ', '');
        const date = bookItem.querySelector('p:nth-child(3)').textContent.replace('Дата: ', '');
        const preview = bookItem.querySelector('p:nth-child(4)').textContent;

        // Создание формы для редактирования книги
        const editForm = document.createElement('form');
        editForm.innerHTML = `
            <input class="input" type="text" id="title" value="${title}">
            <input class="input" type="text" id="author" value="${author}">
            <input class="input" type="date" id="date" value="${date}">
            <input class="input" type="text" id="preview" value="${preview}">
            <button class="btn" type="submit">Сохранить</button>
        `;

        // Очистить контент элемента книги и добавить форму редактирования
        bookItem.innerHTML = '';
        bookItem.appendChild(editForm);

        // Обработка отправки формы редактирования книги
        editForm.addEventListener('submit', (e) => {
            e.preventDefault();

            // Получить новые значения полей формы
            const newTitle = document.getElementById('title').value;
            const newAuthor = document.getElementById('author').value;
            const newDate = document.getElementById('date').value;
            const newPreview = document.getElementById('preview').value;

            // Отправка PUT запроса на сервер для обновления книги
            fetch(`http://127.0.0.1:5501/api/books/${id}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ title: newTitle, author: newAuthor, date: newDate, preview: newPreview })
            })
            .then(response => response.json())
            .then(data => {
                // Создать новые элементы для отображения обновленной информации о книге
                const titleElement = document.createElement('h2');
                titleElement.textContent = data.title;

                const authorElement = document.createElement('p');
                authorElement.textContent = `Автор: ${data.author}`;

                const dateElement = document.createElement('p');
                dateElement.textContent = `Дата: ${data.date}`;

                const previewElement = document.createElement('p');
                previewElement.textContent = data.preview;

                const editBtn = document.createElement('button');
                editBtn.textContent = 'Редактировать';
                editBtn.onclick = () => editBook(data._id);

                const deleteBtn = document.createElement('button');
                deleteBtn.textContent = 'Удалить';
                deleteBtn.onclick = () => deleteBook(data._id);

                // Очистить контент элемента книги и добавить новые элементы
                bookItem.innerHTML = '';
                bookItem.appendChild(titleElement);
                bookItem.appendChild(authorElement);
                bookItem.appendChild(dateElement);
                bookItem.appendChild(previewElement);
                bookItem.appendChild(editBtn);
                bookItem.appendChild(deleteBtn);
            })
            .catch(error => console.error('Ошибка:', error));
        });
    } catch (error) {
        console.error('Ошибка:', error);
    }
}



// Функция удаления книги
function deleteBook(id) {
    // Отправка DELETE запроса на сервер для удаления книги
    fetch(`http://localhost:3000/api/books/${id}`, {
        method: 'DELETE'
    })
    .then(() => {
        // Найти элемент книги по ID и удалить его из DOM
        const bookItem = bookContainer.querySelector(`.book-item[data-id="${id}"]`);
        if (bookItem) bookItem.remove();
    })
    .catch(error => console.error('Ошибка:', error)); // Обработка ошибок при удалении книги
}
} else {
    console.log('Элемент с ID "book-container" не найден'); 
}



