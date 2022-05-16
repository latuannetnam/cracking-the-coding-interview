# cracking-the-coding-interview
Solution for Cracking the coding Interview book (6th edition)
## Setup workspace
- https://go.dev/doc/tutorial/workspaces
### Create module directory
- mkdir 01_Arrays_and_Strings
- cd 01_Arrays_and_Strings
- go mod init arrays
### Create workspace and add module directory
- cd ..
- go work init ./01_Arrays_and_Strings

### Add module to workspace
- go work use ./02_Linked_List