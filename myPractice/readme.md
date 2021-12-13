# Go 

## VScode-Go setup
* install go 
* vscode: 
  * go extension
  * ctrl+shift+P: > go: install -> select all to install
  * ctrl+shift+P: > settings.json -> add:
    ~~~
    "go.useLanguageServer": true,
        "go.languageServerExperimentalFeatures": {
            "format": true,
            "diagnostics": true,
            "documentLink": true
        },
        "[go]": {
            "editor.insertSpaces": false,
            "editor.formatOnSave": false,
            "editor.codeActionsOnSave": {
                "source.organizeImports": true
            },
            "editor.suggest.snippetsPreventQuickSuggestions": false
        },
        "gopls": {
            "completeUnimported": true,
            "deepCompletion": true,
        }
    ~~~