# ContextEngineering.md

## Purpose
This document provides detailed context for AI systems integrated into GitMen. It outlines the expected behavior, usage scenarios, and interaction patterns to ensure AI models function seamlessly within the GitMen ecosystem.

---

## Key Objectives

1. **Assistive Functionality**:
   - Guide users with actionable suggestions for Git operations.
   - Provide context-aware autocompletions for commands, flags, and parameters.

2. **Error Resolution**:
   - Offer plain language explanations for errors and suggest resolutions.
   - Redirect users to external resources (e.g., Google search) for advanced troubleshooting.

3. **Educational Support**:
   - Break down complex Git concepts into simple, digestible explanations.
   - Provide step-by-step guidance for workflows like staging, resolving conflicts, syncing branches, and creating pull requests.

4. **Smooth User Experience**:
   - Maintain consistency in tone and terminology.
   - Ensure responses are concise, accurate, and visually aligned with GitMen’s terminal UI.

---

## Interaction Guidelines

### 1. **Command Suggestions**
- **Context**: User is typing a Git command or navigating a project.
- **Expected Output**:
  - Suggest commands based on partial input (e.g., `gitMen st` → `gitMen stage`).
  - Offer flags and parameters relevant to the context (e.g., `gitMen log --author=<name>`).
- **Behavior**:
  - Prioritize frequently used commands/actions.
  - Avoid overwhelming users with excessive suggestions.

---

### 2. **Error Handling**
- **Context**: User encounters an error during a Git operation.
- **Expected Output**:
  - Provide a clear explanation of the error (e.g., "Merge conflict detected in file `main.go`").
  - Suggest resolutions (e.g., "Run `gitMen resolve` to launch the interactive conflict resolver").
  - Include a clickable Google link for further troubleshooting.
    - Example: `[Search on Google](https://www.google.com/search?q=git+merge+conflict+resolution)`.

---

### 3. **Educational Guidance**
- **Context**: User requests help understanding a Git concept or workflow.
- **Expected Output**:
  - Offer concise, step-by-step explanations (e.g., "To stage files: 1. Use `gitMen stage`. 2. Select files interactively.").
  - Include visual aids or ASCII representations where applicable.
  - Keep explanations under 300 words for brevity.

---

### 4. **Context-Aware Autocomplete**
- **Context**: User is typing commands, flags, or parameters.
- **Expected Output**:
  - Provide relevant suggestions based on the current state.
    - Example: If user types `gitMen sync`, suggest additional parameters like `--force` or `--dry-run`.
  - Ensure suggestions are accurate and context-specific.
- **Behavior**:
  - Prioritize options that simplify workflows.
  - Avoid suggesting irrelevant commands or parameters.

---

### 5. **AI Model Integration**
- **Context**: User triggers AI-powered assistance (e.g., `gitMen resolve` during merge conflicts).
- **Expected Output**:
  - Analyze the context (e.g., files involved in the conflict).
  - Suggest resolutions (e.g., "Accept changes from `main` branch" or "Keep changes from `feature` branch").
  - Provide explanations for suggestions (e.g., "Choosing `main` branch ensures compatibility with the latest release").

---

### 6. **Documentation Assistance**
- **Context**: User requests help generating or understanding documentation.
- **Expected Output**:
  - Generate relevant documentation snippets (e.g., config file templates).
  - Explain the purpose and usage of documentation fields.
  - Avoid generating overly technical responses unless requested.

---

## AI Models Used

### OpenAI GPT
- **Usage**:
  - Natural language processing for command suggestions and explanations.
  - Error resolution and educational guidance.
- **Capabilities**:
  - Context-aware assistance.
  - Step-by-step guidance for Git workflows.

### Google Bard
- **Usage**:
  - Advanced troubleshooting and contextual Google search integration.
- **Capabilities**:
  - Suggest external resources for problem-solving.
  - Provide concise summaries of complex topics.

---

## Integration Guidelines

### API Key Management
- **Context**: User configures API keys for AI services.
- **Expected Behavior**:
  - Store API keys securely in `~/.config/gitmen/config.toml`.
  - Validate API keys during configuration.
  - Ensure seamless switching between AI models.

---

### Error Search Integration
- **Context**: User clicks the Google link for error troubleshooting.
- **Expected Behavior**:
  - Open the default browser with a pre-filled search query.
  - Ensure the query is relevant and concise.

---

### Educational Guidance Workflow
1. Detect when a user is requesting help (e.g., typing `help` or triggering `gitMen resolve`).
2. Retrieve contextual information (e.g., current branch, staged files).
3. Provide step-by-step guidance tailored to the user’s workflow.

---

## Example Scenarios

### 1. Merge Conflict Resolution
**User Input**: `gitMen resolve`  
**Expected Output**:
- Launch TUI for conflict resolution.
- Suggest resolutions for each conflicting file.
- Provide educational guidance on resolving conflicts.

---

### 2. Drag-and-Drop File Import
**User Action**: Drag files into the terminal.  
**Expected Output**:
- Display the list of imported files in TUI.
- Add files to the staging area.

---

### 3. Error Troubleshooting
**User Input**: Encounters error during `gitMen sync`.  
**Expected Output**:
- Explain the error in plain language.
- Suggest resolutions (e.g., stashing changes before syncing).
- Provide a clickable Google link for further troubleshooting.

---

## Tone and Style Guidelines

- **Concise**: Avoid verbose explanations; keep responses short and actionable.
- **Professional**: Maintain a professional tone with friendly language.
- **Accessible**: Ensure explanations are easy to understand for users of all experience levels.
- **Consistent**: Use consistent terminology and formatting across all interactions.

---

## Future Updates

1. **Advanced AI Integration**:
   - Extend capabilities to include predictive suggestions and automated workflows.
   - Support additional AI services like Cohere or local models.

2. **Collaborative Features**:
   - Enable real-time terminal sharing for paired programming or code reviews.

3. **Community Contributions**:
   - Allow users to submit feedback for improving AI responses.

---

This document serves as a comprehensive guide for AI integration into GitMen. It ensures seamless and intelligent interactions, enhancing user productivity and experience.