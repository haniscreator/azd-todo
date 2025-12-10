import { describe, it, expect } from "vitest";
import { summarizeTodos } from "./todoUtils";
import type { Todo } from "./api";

describe("summarizeTodos", () => {
  it("correctly summarizes totals", () => {
    const todos: Todo[] = [
      {
        id: 1,
        username: "alice",
        title: "Buy milk",
        is_completed: false,
        created_at: "",
        completed_at: null,
      },
      {
        id: 2,
        username: "alice",
        title: "Finish project",
        is_completed: true,
        created_at: "",
        completed_at: "",
      },
    ];

    const summary = summarizeTodos(todos);

    expect(summary.total).toBe(2);
    expect(summary.completed).toBe(1);
    expect(summary.pending).toBe(1);
  });
});
