# new-project

Practice repository for setting up Git with a `development` branch (PROM-42164).

## Set up the repository yourself

These steps create a local repository, add a `development` branch, merge it into `main` and publish the result to GitHub.

If the repository already exists on GitHub, run `git clone <repository-url>` and `cd new-project` instead of steps 1 and 2.

1. Create the project directory and go into it:

   ```sh
   mkdir new-project
   cd new-project
   ```

2. Initialize a Git repository with `main` as the default branch:

   ```sh
   git init -b main
   ```

3. Create `README.md` with some initial text:

   ```sh
   echo "# new-project" > README.md
   ```

4. Stage the file for commit:

   ```sh
   git add README.md
   ```

5. Commit it with the message `init`:

   ```sh
   git commit -m "init"
   ```

6. Create the `development` branch and switch to it:

   ```sh
   git switch -c development
   ```

   On Git older than 2.23, use `git checkout -b development`.

7. Add your instructions to `README.md` in any editor, then stage the change:

   ```sh
   git add README.md
   ```

8. Commit using the Jira Smart Commit format, `<ISSUE-KEY> #<command> <text>`:

   ```sh
   git commit -m "PROM-42164 #comment Add step-by-step setup instructions to README"
   ```

   The commands you can use:

   - `#comment <text>` adds a comment to the issue.
   - `#time <value>` logs work time, for example `#time 1h 30m`.
   - `#<transition>` moves the issue to a workflow status, for example `#done`.

   Jira only processes Smart Commits when the repository is connected to Jira. See [Process issues with smart commits](https://support.atlassian.com/jira-software-cloud/docs/process-issues-with-smart-commits/#Smart-Commit-commands).

9. Switch back to `main` and merge `development` into it:

   ```sh
   git switch main
   git merge --no-ff development
   ```

   `--no-ff` always creates a merge commit, so the history shows where the branch was merged.

10. Check that everything is committed and the history looks right:

    ```sh
    git status
    git log --oneline --graph --all
    ```

    `git status` should say `nothing to commit, working tree clean`.

11. Publish both branches to a public GitHub repository. With the GitHub CLI:

    ```sh
    gh repo create new-project --public --source=. --push
    git push origin development
    ```

    Or, after creating an empty public repository on github.com:

    ```sh
    git remote add origin https://github.com/<your-user>/new-project.git
    git push -u origin main development
    ```
