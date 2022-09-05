# golang-exercises
## section 1: getting started

### Installing Go
To install go in arch linux \
```$ sudo pacman -S go ```\
Test if the installation is succesful by running\
```$ go version```

### Setting up Go in Neovim
- Install the latest stable version of Go's language server (gopls)
```
$ go install golang.org/x/tools/gopls@latest
```
- Setup gopls in nvim-lsp config
``` 
require('lspconfig')['gopls'].setup{
  on_attach = on_attach,
  flags = lsp_flags
} 
```
- Install vim-go plugin via packer
``` 
local use = require('packer').use 

require('packer').startup({function() 
  use "fatih/vim-go" 
})
```
