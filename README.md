pogo
====
PowerShell as a Service


###Request Contexts
* __/command/__ - Executes named PowerShell command.
* __/script/__ - Executes named script.
* __/exit/__ - Terminates the service instance.

###Running Commands
Commands are handled by the /command/ context. Most common commands are readily supported and will return any structured data in JSON format. Parameters are passed via url query parameters. Named values will be broken out into key/value pairs and added to the commandstring.


Run __[Get-Date](https://technet.microsoft.com/en-us/library/hh849887.aspx)__ command.
    http://127.0.0.1:8080/command/Get-Date
```json
{
    "value":  "\/Date(1436237104231)\/",
    "DisplayHint":  2,
    "DateTime":  "Monday, July 6, 2015 9:45:04 PM"
}
```


Run __[Write-Host](https://technet.microsoft.com/en-us/library/ee177031.aspx)__ command.
    http://127.0.0.1:8080/command/Write-Host?"Hello,%20World!"
```json
Hello, World!
```


Run __[Get-Item](https://technet.microsoft.com/en-us/library/hh849788.aspx)__ command.
    http://127.0.0.1:8080/command/Get-Item?-Path="pogo.exe"
```json
{
    "Name":  "pogo.exe",
    "Length":  9076224,
    "DirectoryName":  "C:\\Users\\onyxhat\\Documents\\GitHub\\pogo",
    "Directory":  {
                      "Name":  "pogo",
                      "Parent":  {
                                     "Name":  "GitHub",
                                     "Parent":  "Documents",
                                     "Exists":  true,
                                     "Root":  "C:\\",
                                     "FullName":  "C:\\Users\\onyxhat\\Documents\\GitHub",
                                     "Extension":  "",
                                     "CreationTime":  "\/Date(1436236701491)\/",
                                     "CreationTimeUtc":  "\/Date(1436236701491)\/",
                                     "LastAccessTime":  "\/Date(1436236829685)\/",
                                     "LastAccessTimeUtc":  "\/Date(1436236829685)\/",
                                     "LastWriteTime":  "\/Date(1436236829685)\/",
                                     "LastWriteTimeUtc":  "\/Date(1436236829685)\/",
                                     "Attributes":  16
                                 },
                      "Exists":  true,
...
}
```


###TODO
* Add ___Authentication___
* Implement POST method of form values as parameters (takes precedence over query params)
* Command restrictions
* Additional configuration values
* Remote configuration
* Configurable script directory
