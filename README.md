# file-renamer-cli
Go Project For Automatically Renaming Files

This project is a simple command line utility for batch renaming files. It has support for the following actions:

* Trimming left
* Trimming right
* Trimming between two points from left or right
* Altering Letter Case
* Replacing blocks of text
* Inserting text from the left
* Inserting text from the right
* Counting from an arbitrary point with an arbitrary amount of text padding

This project can be compiled using the basic Go compile command: `go build`

Once built, the program uses command line flags to add set arguments and values.

| Command | Use | Argument Count |
|---|---| --- |
| -i | Folder Path Input | 1 |
| -p | Preview | 0 |
| -ie | Include Extensions | 0 |
| -tl | Trim Left | 1 |
| -tr | Trim Right | 1 |
| -tb | Trim Between | 3 |
| -uc | Set Upper Case | 0 |
| -lc | Set Lower Case | 0 |
| -tc | Set Title Case | 0 |
| -r | Replacer | 2 |
| -il | Insert Left | 2 |
| -ir | Insert Right | 2 |
| -c | Counter | 3 |

## Flags

### -i Folder Path Input

This command accepts a single argument, a path to a folder containing the files you want to rename. Whatever path you set, all files in that folder will be renamed when the utility is run.

### -p Preview

This command prints a preview of the file changes that will be made if the command is run.

### -ie Include Extensions

By default the program will exclude extensions (anyting to the left of the last period in the file name, including the last period). If you want to include extensions when making changes, you should include this flag.