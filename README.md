<h3 align="center">
  <div align="center">
    <h1>Currency Bot 💰📈</h1>
  </div>
  <a href="https://github.com/goroutiner/currency_bot">
    <img src="https://github.com/goroutiner/currency_bot/raw/main/images/logo.jpg" width="500" height="500"/>
  </a>
</h3>

**Currency Bot** — это удобный и быстрый Telegram-бот для получения актуальных и исторических данных о курсах валют относительно доллара США. Также бот умеет строить графики изменения курса, что позволяет анализировать тренды.

---

## 📋 Основные функции

1. **Актуальный курс валют**  
   Узнайте текущий курс выбранной валюты по отношению к доллару США (USD).

2. **Исторические данные**  
   Получите информацию о курсах валют за определённый период времени.

3. **Графики изменения курса**  
   Постройте график изменения курса валюты за указанный интервал времени.

---

## 📊 Пример графика

График изменения курса RUB за интервал времени с 2001-01-01 по 2002-01-01:  
<img src="https://github.com/goroutiner/currency_bot/raw/main/images/plot_exemple.jpg" alt="Пример изображения" width="400" height="400">

---

## 📜 Доступные команды

- `/start`  
  Команда запускает телеграм-бота.\
  ![Start bot](https://github.com/goroutiner/currency_bot/raw/main/images/start_exemple.jpg)

- `/change_mode`  
  Команда предлагает изменить режим работы.\
  ![Change mode](https://github.com/goroutiner/currency_bot/raw/main/images/change-mod_exemple.jpg) 
  

- `/change_currency`  
  Команда предлагает изменить валюту.\
  ![Change currency](https://github.com/goroutiner/currency_bot/raw/main/images/change-currency_exemple.jpg)


---

## 🔧 Предварительные настройки бота

1. Необходимо создать бота в Telegram и получить его токен, используя сервис **BotFather**:  [Перейти к BotFather](https://t.me/botfather).
2. Добавьте команды `/start`, `/change_mode` и `/change_currency`, используя **BotFather**, к своему созданному боту.\
![Пример того, как должны выглядеть команды в боте](https://github.com/goroutiner/currency_bot/raw/main/images/commands_exemple.jpg)

3. Зарегистрируйтесь на сервисе [FreecurrencyAPI](https://freecurrencyapi.com/), чтобы получить **API Key**.

---

## 📦 Установка и запуск

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/goroutiner/currency_bot
---

2. Перейдите в рабочую директорию с проектом:
    ```bash
    cd  currency_bot
    ```

3. Открываем в рабочей директории **Dockerfile** записываем свой *токен* и *Api Key* в переменные окружения `BOT_TOKEN` и `API_KEY`
   ```bash
    ENV BOT_TOKEN=""
    ENV API_KEY=""
   ```

4. Собираем Docker-образ с телеграм-ботом:
    ```bash
    docker build -t currency_bot:v1.0 .
    ```

5. Запускаем Docker-образ:
    ```bash
    docker run -it -p 8080:8080 currency_bot:v1.0
    ```

6. Чтобы прекратить работу телеграм-бота, необходимо остановить запущенный контейнер, для этого достаточно нажать сочетание клавиш `Ctrl+C` в открытом терминале.

7. Когда ваш Docker-контейнер запущен, телеграм-бот готов к использованию 💫 Открывайте своего бота в телеграмме и пользуйтесь его функционалом 🤑 

## 🛠️ Технические ресурсы

- **Язык разработки:** Golang  
- **Библиотека для взаимодействия с Telegram API:** [tucnak/telebot](https://github.com/tucnak/telebot?tab=readme-ov-file).
- **API для курсов валют:** [freecurrencyapi.com](https://freecurrencyapi.com/)  
- **Библиотека для графиков:** [gonum/plot](https://github.com/gonum/plot)  
- **Библиотека форматирования валют:** [bojanz/currency](https://github.com/bojanz/currency)

---
