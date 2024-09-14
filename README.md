# todogo - Command Line Application for Managing To-Dos

`todogo` is a simple, yet powerful command line tool to manage your to-do lists efficiently.

## Features

- Manage multiple to-do lists.
- Add, remove, and complete tasks.
- Move tasks between positions.
- Clear all tasks from a list.

## Usage

### Display Application Version

To check the current version of the application, use the following command:

```bash
todogo version
```

### List Management

#### Create a New List

Use the `list add` command to create a new to-do list:

```bash
todogo list create "My List"
```

#### Switch Between Lists

To switch to a specific list, use the `list checkout` command with the list's `ID`:

```bash
todogo list checkout 3
```

#### Remove a List

To remove a list and all its associated to-dos, use the `list rm` command with the list's `ID`:

```bash
todogo list rm 3
```

#### List All Lists

To display all created lists, use the `list ls` command:

```bash
todogo list ls
```

### Todo Management

#### Add a New Todo

To create a new todo item, use the `todo add` command:

```bash
todogo todo add "Learn Golang"
# or the shorthand version:
todogo todo a "Learn Golang"
```

#### Remove a Todo

To remove a todo, use the `todo rm` command with the todo's `ID`:

```bash
todogo todo rm 3
# or the shorthand version:
todogo todo r 3
```

#### Complete a Todo

To mark a todo as completed, use the `todo complete` command with the todo's `ID`:

```bash
todogo todo complete 3
# or the shorthand version:
todogo todo c 3
```

#### List All Todos

To display all todos in the current list, use the `todo ls` command:

```bash
todogo todo ls
```

#### List Completed Todos

To display only completed todos, use the `todo ls --completed` command:

```bash
todogo todo ls --completed
```

#### Move a Todo

To move a todo to a different position within the list, use the `todo mv` command with the todo's `ID` and the target position:

```bash
todogo todo mv 3 --pos 1
```

#### Clear All Todos in a List

To remove all todos from the current list, use the `todo clear` command:

```bash
todogo todo clear
```
