<script>
  export let todo;
  export let editName;
  export let changeCompleted;
</script>

<tr>
  {#if todo.name !== ""}
    <th> {todo.name}</th>
  {:else}
    <th>
      <input type="text" on:blur={(e) => editName(e?.target.value, todo)} />
    </th>
  {/if}
  <th>
    <input
      type="checkbox"
      checked={todo.completed}
      on:click={() => changeCompleted(todo)}
    />
  </th>
</tr>
