import axios from "axios";

const api = axios.create({
  baseURL: "/api", // goes through Vite proxy → http://localhost:8080/api
});

export type Todo = {
  id: number;
  username: string;
  title: string;
  is_completed: boolean;
  created_at: string;
  completed_at?: string | null;
};

export async function fetchTodos(username: string): Promise<Todo[]> {
  const res = await api.get<Todo[]>("/todos", {
    params: { username },
  });
  return res.data;
}

export async function createTodo(username: string, title: string): Promise<Todo> {
  const res = await api.post<Todo>("/todos", { username, title });
  return res.data;
}

export async function completeTodo(id: number, username: string): Promise<void> {
  await api.patch(`/todos/${id}/complete`, { username });
}
