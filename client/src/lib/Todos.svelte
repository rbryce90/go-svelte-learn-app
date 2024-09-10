<script lang="ts">
  import Todo from "./Todo.svelte";
  type Todo = {
    name: string;
    completed: boolean;
  };

  let todos: Todo[] = [
    { name: "Do this app", completed: false },
    { name: "Do something else", completed: true },
  ];

  const changeCompleted = (todo: Todo) => {
    todo.completed = todo.completed ? false : true;
  };
  const addTodo = () => {
    console.log("addTodo ");
    const newTodos = [...todos];
    const newTodo: Todo = {
      name: "",
      completed: false,
    };
    newTodos.push(newTodo);
    todos = newTodos;
  };

  const editName = (input: string, todo: Todo) => {
    todo.name = input;
  };

  const saveChanges = () => {
    console.log("save: ", todos);
  };
</script>

<div class="todos-div">
  <h1>Todos</h1>
  <table>
    <tr>
      <th class="table-header"> Name </th>
      <th class="table-header"> Completed </th>
    </tr>
    {#each todos as todo}
      <Todo {todo} {editName} {changeCompleted} />
    {/each}
    <tr>
      <th>
        <button on:click={() => addTodo()}>Add</button>
      </th>
      <th>
        <button on:click={() => saveChanges()}>Save</button>
      </th>
    </tr>
  </table>
</div>

<style>
  .table-header {
    font-size: 1.5em;
    font-weight: 900;
    font-family: sans-serif;
    padding: 20px;
  }
  .todos-div button {
    width: 100%;
    height: 3em;
    font-size: larger;
    font-weight: bold;
    border-radius: 5px;
    background-color: rgb(107, 107, 205);
    cursor: pointer;
    border: none;
    color: white;
  }
  .todos-div {
    font-family: sans-serif;
    width: 50%;
    display: flex;
    justify-content: center;
    flex-direction: column;
    align-items: center;
  }
  .todos-div table {
    width: 100%;
  }
  .todos-div h1 {
    display: flex;
    justify-content: center;
  }

  .todos-div tr {
    width: 100%;
  }
  .todos-div th {
    width: 50%;
    padding: 10px;
  }
</style>
