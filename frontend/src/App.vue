<script setup lang="ts">
  import { ref, computed, onMounted } from "vue";
  import { fetchTodos, createTodo, completeTodo, type Todo } from "./api";
  
  const username = ref("");
  const newTodoTitle = ref("");
  const todos = ref<Todo[]>([]);
  const loading = ref(false);
  const error = ref<string | null>(null);
  
  // keep username in localStorage so you don't re-type on refresh
  onMounted(() => {
    const saved = localStorage.getItem("username");
    if (saved) {
      username.value = saved;
      loadTodos();
    }
  });
  
  const hasUsername = computed(() => username.value.trim().length > 0);
  
  async function saveUsername() {
    if (!hasUsername.value) return;
    localStorage.setItem("username", username.value.trim());
    await loadTodos();
  }
  
  async function loadTodos() {
    if (!hasUsername.value) return;
    try {
      loading.value = true;
      error.value = null;
      todos.value = await fetchTodos(username.value.trim());
    } catch (e: any) {
      error.value = e?.message ?? "Failed to load todos";
    } finally {
      loading.value = false;
    }
  }
  
  async function onAddTodo() {
    if (!hasUsername.value || !newTodoTitle.value.trim()) return;
    try {
      loading.value = true;
      error.value = null;
      const todo = await createTodo(username.value.trim(), newTodoTitle.value.trim());
      todos.value.unshift(todo);
      newTodoTitle.value = "";
    } catch (e: any) {
      error.value = e?.message ?? "Failed to add todo";
    } finally {
      loading.value = false;
    }
  }
  
  async function onComplete(todo: Todo) {
    if (todo.is_completed) return;
    try {
      loading.value = true;
      error.value = null;
      await completeTodo(todo.id, username.value.trim());
      // update UI locally
      todo.is_completed = true;
      todo.completed_at = new Date().toISOString();
    } catch (e: any) {
      error.value = e?.message ?? "Failed to complete todo";
    } finally {
      loading.value = false;
    }
  }
  </script>
  
  <template>
    <div class="app">
      <h1>Todo Demo (Vue + Go + ClickHouse)</h1>
  
      <!-- Username setup -->
      <section class="card">
        <h2>1. Choose a username</h2>
        <form @submit.prevent="saveUsername" class="row">
          <input
            v-model="username"
            placeholder="Enter a username (e.g. alice)"
          />
          <button type="submit">Use this name</button>
        </form>
        <p v-if="hasUsername">Current user: <strong>{{ username }}</strong></p>
      </section>
  
      <!-- Error / loading -->
      <p v-if="error" class="error">{{ error }}</p>
      <p v-if="loading">Loading...</p>
  
      <!-- Todo section -->
      <section v-if="hasUsername" class="card">
        <h2>2. Todos for {{ username }}</h2>
  
        <form @submit.prevent="onAddTodo" class="row">
          <input
            v-model="newTodoTitle"
            placeholder="New todo title"
          />
          <button type="submit">Add</button>
        </form>
  
        <ul class="todo-list">
          <li
            v-for="todo in todos"
            :key="todo.id"
            :class="{ completed: todo.is_completed }"
          >
            <span>{{ todo.title }}</span>
            <button
              v-if="!todo.is_completed"
              type="button"
              @click="onComplete(todo)"
            >
              Complete
            </button>
            <span v-else class="badge">Done</span>
          </li>
        </ul>
      </section>
    </div>
  </template>
  
  <style scoped>
  .app {
    max-width: 600px;
    margin: 2rem auto;
    font-family: system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI",
      sans-serif;
    padding: 1rem;
  }
  h1 {
    text-align: center;
    margin-bottom: 1.5rem;
  }
  .card {
    border: 1px solid #ddd;
    border-radius: 8px;
    padding: 1rem;
    margin-bottom: 1rem;
  }
  .row {
    display: flex;
    gap: 0.5rem;
    margin-top: 0.5rem;
  }
  input {
    flex: 1;
    padding: 0.5rem;
  }
  button {
    padding: 0.5rem 0.75rem;
    cursor: pointer;
  }
  .todo-list {
    list-style: none;
    padding: 0;
    margin-top: 1rem;
  }
  .todo-list li {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0.4rem 0;
    border-bottom: 1px solid #eee;
  }
  .todo-list li.completed span:first-child {
    text-decoration: line-through;
    opacity: 0.7;
  }
  .error {
    color: red;
  }
  .badge {
    font-size: 0.8rem;
    padding: 0.1rem 0.4rem;
    border-radius: 999px;
    border: 1px solid #4caf50;
  }
  </style>
  