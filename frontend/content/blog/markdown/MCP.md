# Understanding MCP: Model Context Protocol (with Code Sample)

As AI systems evolve from simple prompt-response interactions to rich, stateful applications, the need for managing context becomes crucial. This is where the **Model Context Protocol (MCP)** enters the scene. MCP is an emerging standard designed to structure and maintain the contextual memory of AI systems, allowing models to function as collaborative agents over time.

## What is MCP?

**MCP** provides a formalized structure for how context is passed between users, tools, and language models. Think of it as the "conversation memory" layer that sits between your prompts and the model's response. It helps coordinate sequences of inputs, outputs, tool calls, and state updates.

Unlike monolithic chat logs, MCP treats *context* as a structured, versioned, and inspectable data model. It’s especially helpful for:

- Managing large conversations across multiple sessions
- Integrating external tools or APIs
- Maintaining memory or user preferences
- Coordinating multiple agents in a system

## Core Concepts

MCP structures interactions around these components:

- **Messages**: Typed interactions (e.g., `user`, `assistant`, `tool_call`) with metadata.
- **Turns**: Logical groupings of actions triggered by a single input.
- **Context**: The evolving record of state over time.
- **Updates**: Explicit deltas to the context, often triggered by events or decisions.

## Example: MCP in Practice

Here is a simplified code sample that shows how an MCP-compliant system might handle a conversation with a tool call:

```go
 fileServer := http.FileServer(http.Dir(config.Paths.AssetFiles))
 mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

 // Routes
 mux.HandleFunc("/", handler.ServeHomepage)
 mux.HandleFunc("/home", handler.ServeHomepage)
 mux.HandleFunc("/resume", handler.ServeResume)
 mux.HandleFunc("/project", handler.ServeProjectsList)
 mux.HandleFunc("/blog", handler.ServeBlogList)
 mux.HandleFunc("/blog/", func(w http.ResponseWriter, r *http.Request) {
  // Check if the request is for the table of contents
  if strings.Contains(r.URL.Path, "/table-of-contents") {
   handler.ServeBlogTableOfContents(w, r)
   return
  }
  // Otherwise, serve the regular blog content
  handler.ServeBlogContent(w, r)
 })

 // Apply middleware chain
 var handler http.Handler = mux

 if !config.Features.CacheEnabled {
  handler = middleware.NoCacheMiddleware(handler)
 }
```

```json
{
  "context": {
    "messages": [
      {
        "id": "msg-001",
        "role": "user",
        "type": "text",
        "content": "What's the weather in Montreal?",
        "timestamp": "2025-04-14T10:00:00Z"
      },
      {
        "id": "msg-002",
        "role": "assistant",
        "type": "tool_call",
        "tool": "weather_api",
        "parameters": { "city": "Montreal" },
        "timestamp": "2025-04-14T10:00:01Z"
      },
      {
        "id": "msg-003",
        "role": "tool",
        "type": "tool_result",
        "tool": "weather_api",
        "output": "Sunny, 12°C",
        "timestamp": "2025-04-14T10:00:02Z"
      },
      {
        "id": "msg-004",
        "role": "assistant",
        "type": "text",
        "content": "It's sunny in Montreal right now, with a temperature of 12°C.",
        "timestamp": "2025-04-14T10:00:03Z"
      }
    ],
    "turns": [
      {
        "id": "turn-001",
        "messages": ["msg-001", "msg-002", "msg-003", "msg-004"]
      }
    ]
  }
}
```

## Sample Header for testing

MCP structures interactions around these components:

- **Messages**: Typed interactions (e.g., `user`, `assistant`, `tool_call`) with metadata.
- **Turns**: Logical groupings of actions triggered by a single input.
- **Context**: The evolving record of state over time.
- **Updates**: Explicit deltas to the context, often triggered by events or decisions.

```python
from datetime import datetime
import uuid

def timestamp():
    return datetime.utcnow().isoformat() + "Z"

def new_message(role, content=None, type="text", **kwargs):
    return {
        "id": str(uuid.uuid4()),
        "role": role,
        "type": type,
        "content": content,
        "timestamp": timestamp(),
        **kwargs
    }

# Simulate a conversation context
context = {
    "messages": [],
    "turns": []
}

# Start a new turn
turn_id = str(uuid.uuid4())
turn_messages = []

# User asks a question
msg_user = new_message("user", "What's the weather in Montreal?")
turn_messages.append(msg_user)

# Assistant calls the weather tool
msg_tool_call = new_message("assistant", type="tool_call", tool="weather_api", parameters={"city": "Montreal"})
turn_messages.append(msg_tool_call)

# Simulated tool response
msg_tool_result = new_message("tool", type="tool_result", tool="weather_api", output="Sunny, 12°C")
turn_messages.append(msg_tool_result)

# Assistant responds with final answer
msg_assistant = new_message("assistant", "It's sunny in Montreal right now, with a temperature of 12°C.")
turn_messages.append(msg_assistant)

# Add messages to context
context["messages"].extend(turn_messages)
context["turns"].append({
    "id": turn_id,
    "messages": [m["id"] for m in turn_messages]
})

# Output context
import json
print(json.dumps(context, indent=2))
```

## The Ultimate Guide to Boosting Your Productivity in 2025

In the age of constant notifications and remote work, staying productive is more challenging than ever. Here’s a practical guide to help you optimize your time and energy.

## Set Clear Goals

Start each day with 2–3 major goals. Focus on results, not just tasks.

## Use the 80/20 Rule

Identify the 20% of tasks that generate 80% of your results. Prioritize those.

## Create a Morning Routine

A solid morning routine sets the tone for the rest of the day. Include exercise, planning, and quiet time.

## Minimize Distractions

Silence your phone, close unnecessary tabs, and use apps like Focus@Will or Freedom to stay on track.

## Batch Similar Tasks

Group similar tasks like emails, calls, or admin work. This reduces context switching and saves time.

## Take Regular Breaks

Use techniques like Pomodoro (25 min focus, 5 min break) to maintain energy throughout the day.

## Leverage Automation Tools

Automate repetitive tasks with tools like Zapier, IFTTT, or Notion workflows.

## Weekly

Spend 15 minutes each week reviewing what worked and what didn’t. Adjust accordingly.

## Declutter Your Digital Space

A messy desktop or overflowing inbox drains mental energy. Clear it weekly.

## Protect Your Time

Learn to say no. Block off focus hours and treat them like meetings with yourself.
