# Architecture Documentation

## High-Level Overview
GitMen is designed as a modular, scalable CLI tool that integrates core Git operations, terminal UI components, and AI-powered assistance.

### Modules
1. **CLI Commands (`cmd/`)**:
   - Handles user input via Cobra.
   - Routes commands to appropriate business logic (e.g., `gitMen stage`, `gitMen resolve`).

2. **Terminal UI (`tui/`)**:
   - Built with Bubble Tea for interactive workflows (e.g., staging files, resolving conflicts).
   - Includes drag-and-drop file import functionality.

3. **Core Logic (`core/`)**:
   - Implements backend operations like Git commands, error handling, and project navigation.
   - Provides robust error search via Google integration.

4. **Configuration Management (`config/`)**:
   - Manages user preferences stored in TOML files.
   - Supports dynamic theme switching and auto battery-based adjustments.

5. **AI Models (`ai/`)**:
   - Integrates AI services via API keys (e.g., OpenAI GPT, Google Bard).
   - Provides contextual assistance for Git workflows.

6. **Themes (`themes/`)**:
   - Stores built-in and user-defined themes.
   - Marketplace integration for community contributions.

7. **Plugins (`plugins/`)**:
   - Enables extensibility for additional functionalities (e.g., Docker, Kubernetes).
   - Designed with a lightweight plugin API.

8. **Assets (`assets/`)**:
   - Contains icons, animations, and fonts for enhanced visuals.

9. **Testing (`tests/`)**:
   - Includes unit and integration tests for core logic and TUI components.

## Data Flow
### Git Commands
1. User enters a Git-related command via CLI.
2. Command is routed to `core/git.go` for processing.
3. Results are displayed in the terminal or TUI.

### AI Assistance
1. User triggers AI-powered assistance (e.g., `gitMen resolve`).
2. AI model is called via `ai/ai.go`.
3. Suggestions are displayed in the terminal.

### Drag-and-Drop
1. Files are dragged into the terminal.
2. `tui/drag_drop.go` handles the event.
3. File paths are added to the staging area.

## Design Patterns
1. **Factory Pattern**:
   - Used for initializing AI models.
2. **Modular Design**:
   - Keeps components decoupled for scalability.
3. **Observer Pattern**:
   - Enables dynamic updates in TUI components.

## Future Scalability
The architecture supports:
- Additional AI integrations.
- Expanded plugin system.
- Community-driven themes and workflows.