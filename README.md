# Roadmap

### **Phase 1: Basic CLI to Fetch and Display RSS Feeds**

**Session 1: Setup CLI and Basic Structure**

- **Goal**: Initialize the project and set up a basic CLI framework.
- **Tasks**:
    [x] Initialize a new Go module.
    [ ] Choose a CLI framework (e.g., [Cobra](https://github.com/spf13/cobra) or Go’s built-in `flag` package).
    [ ] Set up basic commands like `check` and `list`.
    [ ] Research and select a Go RSS parser library (e.g., `github.com/mmcdole/gofeed`).

**Session 2: RSS Feed Parsing**

- **Goal**: Fetch and parse RSS feeds.
- **Tasks**:
    [ ] Integrate the RSS parser.
   [ ] Create a function to fetch RSS feeds from a list of URLs.
    [ ]. Parse the feed to display titles, publication dates, and URLs.
    [ ] Save parsed articles into a local JSON or SQLite file to track read/unread status.

### **Phase 2: Notification System**

**Session 3: Track Unread Articles**

- **Goal**: Add functionality to track unread articles.
- **Tasks**:
    [ ] Implement a function to differentiate between read and unread articles.
    [ ] Store article status in the local storage (e.g., a simple JSON file or SQLite).
    [ ] Test the functionality by marking articles as read/unread.

**Session 4: Terminal Notifications**

- **Goal**: Display a notification when there are unread articles.
- **Tasks**:
    [ ] Research terminal notification tools (e.g., `notify-send` for Linux, `osascript` for macOS).
   [ ] Implement a check at the start of the CLI for unread articles.
    [ ] Show a notification or alert in the terminal with the number of unread articles.

### **Phase 3: View and Open Articles**

**Session 5: Command to List Articles**

- **Goal**: List all unread articles in the terminal.
- **Tasks**:
    [ ] Implement a `list` command to show all unread articles with their titles and publication dates.
    [ ] Add options to mark articles as read directly from the list.

**Session 6: Command to Open Articles**

- **Goal**: Add functionality to open articles from the terminal.
- **Tasks**:
    [ ] Implement an `open` command to open an article in the browser.
    [ ] Research and integrate system-agnostic ways to open URLs (`exec.Command` or libraries like `github.com/skratchdot/open-golang`).

### **Phase 4: AI Summaries (Advanced)**

**Session 7: Integrate AI for Summaries**

- **Goal**: Implement basic AI summaries for articles (requires third-party API).
- **Tasks**:
    [ ] Research AI summary services (OpenAI, Hugging Face, etc.).
    [ ] Set up API integration to fetch summaries of articles.
    [ ] Display summaries in the terminal when requested.

**Session 8: Display Summaries in CLI**

- **Goal**: Add command to generate and show AI summaries.
- **Tasks**:
    [ ] Implement a `summarize` command to fetch and display AI-generated summaries for unread articles.
    [ ] Add the option to read the full article if the summary seems interesting.

### **Optional Cool Features to Add Later**

1. **Search Command**: Add a `search` feature to filter articles by keywords or categories.
2. **Category-based Notifications**: Customize notifications for specific RSS categories (e.g., videos, articles).
3. **Neofetch Integration**: Display the number of unread articles when running `neofetch` or similar terminal startup utilities.
4. **Offline Mode**: Cache articles for offline reading.
5. **Multiple RSS Feeds Management**: Add, remove, and organize feeds within the CLI.