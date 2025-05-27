go mod init mypointers

pln

<!-- to run air- -->
echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc

<!-- to run air- -->
echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc

go mod tidy -( it brings in packages which are needed and also removes which are not necessary )

go mod verify -( helps to verify go.sum hashes and verifies imports ) 

go list - return modulename on which your entire application is dependant on. 
 
 go list all - dump up all the modules present in the cache 
 go list -m all - only main & not those which are dependent on them  

go mod graph - it shows which package is depends on which 


go mod vendor

go run -mod=vendor main.go