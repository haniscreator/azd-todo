import type { Todo } from "./api";

export function summarizeTodos(todos: Todo[]) {
  const total = todos.length;
  const completed = todos.filter((t) => t.is_completed).length;
  const pending = total - completed;

  return { total, completed, pending };
}
