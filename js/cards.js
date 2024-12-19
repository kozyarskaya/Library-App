document.addEventListener('DOMContentLoaded', function() {
    const bookContainer = document.getElementById("book-container");
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
});


    /*fetch('http://localhost:8080/api/user')
            .then(response => response.json())
            .then(data => {
                console.log(data);
                document.getElementById("user-data").innerHTML = `ID: ${data.id}, Username: ${data.username}, Email: ${data.email}`;
            })
            .catch(error => console.error('Ошибка:', error));*/