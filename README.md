# Task Tracker Challenge

if you want to start this challenge too, here's the link:
https://roadmap.sh/projects/task-tracker

This project is a solution to the **Task Tracker Challenge** from [roadmap.sh](https://roadmap.sh). The goal is to create a task management system that allows users to add, update, complete, and delete tasks efficiently, with the option to exit the program.

## Features

- Add new tasks
- Update existing tasks
- Mark tasks as completed
- Delete tasks
- List all tasks with status
- Handle empty task list with appropriate messaging
- Exit the program using `os.Exit`

## Prerequisites

- I'm using GO 1.24.1 to run this app

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/artomily/task-tracker.git
   ```

2. Navigate to the project directory:

   ```bash
   cd task-tracker
   ```

3. Run the application:

   ```bash
   go run main.go
   ```

## Usage

1. Add tasks by modifying the `tasks` slice in the code.
2. Update and delete tasks using corresponding functions.
3. Mark tasks as completed.
4. The program will display "Task list is empty" if there are no tasks available.
5. Exit the program using the `os.Exit` function.

## Code Structure

- `main.go`: Contains the main logic for listing, updating, deleting tasks, and handling empty task lists.
- `task.go`: (Optional) Can be used for struct definition and additional functionalities.

## Future Improvements

- Implement CRUD operations with persistent storage
- Build a web interface

## Contributing

Contributions are welcome! Feel free to submit a pull request.

## License

This project is licensed under the MIT License.

## Contact

For questions or suggestions, reach out to [vararakya@gmail.com](mailto\:vararakya@gmail.com).

## Upvote This Project
https://roadmap.sh/projects/task-tracker/solutions?u=6628c1794e8cd00d388294d2

