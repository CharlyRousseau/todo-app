<script lang="ts">
    import { onMount } from "svelte";
    import { writable } from "svelte/store";
    import { Trash2 } from "lucide-svelte";

    interface Todo {
        id: number;
        content: string;
        completed: boolean;
    }

    let todos = writable<Todo[]>([]);
    let newTodo = "";

    async function fetchTodos() {
        const res = await fetch("http://localhost:8080/todos");
        const data: Todo[] = await res.json();
        data.sort((a, b) =>
            a.completed === b.completed ? 0 : a.completed ? 1 : -1,
        );
        todos.set(data);
    }

    async function addTodo() {
        if (newTodo.trim()) {
            await fetch("http://localhost:8080/todos", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ content: newTodo }),
            });
            newTodo = "";
            fetchTodos();
        }
    }

    async function toggleTodoCompletion(todo: Todo) {
        const updatedTodo = { ...todo, completed: !todo.completed };
        const res = await fetch(`http://localhost:8080/todos/${todo.id}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(updatedTodo),
        });

        if (res.ok) {
            // Update the todo in the store
            todos.update((todos) =>
                todos.map((t) => (t.id === todo.id ? updatedTodo : t)),
            );
            // Fetch the todos after the server has responded to the update request
            await fetchTodos();
        }
    }

    async function deleteTodo(todo: Todo) {
        const res = await fetch(`http://localhost:8080/todos/${todo.id}`, {
            method: "DELETE",
        });

        if (res.ok) {
            // Remove the todo from the store
            todos.update((todos) => todos.filter((t) => t.id !== todo.id));
        }
    }
    onMount(fetchTodos);
</script>

<main>
    <h1>Todo List</h1>
    <div class="input-container">
        <input bind:value={newTodo} placeholder="Enter a new todo" />
        <button on:click={addTodo}>Add Todo</button>
    </div>
    <ul>
        {#each $todos as todo (todo.id)}
            <li class="todo-item" class:completed={todo.completed}>
                <div>
                    <input
                        type="checkbox"
                        bind:checked={todo.completed}
                        on:click={() => toggleTodoCompletion(todo)}
                    />
                    <span>{todo.content}</span>
                </div>
                <!-- Use the Trash component from lucide-svelte -->
                <button class="delete-button" on:click={() => deleteTodo(todo)}>
                    <Trash2 size="16" />
                </button>
            </li>
        {/each}
    </ul>
</main>

<style>
    main {
        font-family: Arial, sans-serif;
        padding: 1rem;
        max-width: 800px;
        margin: 0 auto;
    }

    .input-container {
        margin-bottom: 1rem;
    }

    input,
    button {
        padding: 0.5rem;
        font-size: 1rem;
        border: none;
        border-radius: 8px;
        margin-right: 0.5rem;
    }

    input[type="checkbox"] {
        margin-right: 0.5rem;
    }

    ul {
        list-style-type: none;
        padding: 0;
    }

    li {
        margin-bottom: 0.5rem;
        display: flex;
        align-items: center;
    }

    li.completed {
        text-decoration: line-through;
    }

    .todo-item {
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    .delete-button {
        background: none;
        border: none;
        cursor: pointer;
        color: black;
        transition: color 0.3s ease;
    }

    .delete-button:hover {
        color: red;
    }
</style>
