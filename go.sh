#!/bin/bash

echo " ----- Go template ------"

echo "enter folder name: "
read folderName
mkdir $folderName;
cd $folderName;

echo "Enter module name"
read modName
echo "creating a module ... "
go mod init example.com/$modName


echo "Enter file name: "
read fileName
echo "creating a go file "
touch $fileName.go

echo "--------------------"
echo "package main" >> $fileName.go
echo "" >> $fileName.go
echo "import \"fmt\" " >> $fileName.go
echo "" >> $fileName.go
echo "func main(){ " >> $fileName.go
echo "" >> $fileName.go
echo "  fmt.Println(\"Hello World!\") " >> $fileName.go
echo "}" >> $fileName.go
echo " "


# echo " "
# echo "func getUserInput( prompt string ) string {" >> $fileName.go
# echo " "
# echo "    fmt.Print( prompt ) " >> $fileName.go
# echo "    var value string" >> $fileName.go
# echo "    fmt.Scanln( &value )" >> $fileName.go
# echo ""
# echo "    return value" >> $fileName.go
# echo "}" >> $fileName.go
# echo " " >> $fileName.go






echo "running $fileName.go ..."
go run $fileName.go
echo ""

