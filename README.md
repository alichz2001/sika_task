
# Sika Task #1

## Description

Simple REST API for user model.

## How to Run

Follow these steps to run the Sika Task #1 project:

1. **Clone the repository:**

    ```bash
    git clone https://github.com/alichz2001/sika_task.git
    ```

2. **Navigate to the project directory:**

    ```bash
    cd sika_task
    ```

3. **Build and run the project using Docker Compose:**

    ```bash
    docker-compose up --build
    ```

   This will set up the necessary environment and launch the Sika Task #1 application. Once the process is complete, you should be able to access the application at [http://localhost:8080](http://localhost:8080).


## Usage

- To fetch users in page:
    ```bash
    curl http://localhsot:8080/users?page=1&count=10
    ```

- To fetch user:
    ```bash
    curl http://localhsot:8080/users/1
    ```
    ```
