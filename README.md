pogo
====
PowerShell as a Service


###Request Contexts
* __/command/__ - Executes named PowerShell command.
* __/script/__ - Executes named script.
* __/exit/__ - Terminates the service instance.

###Example
    http://127.0.0.1:8080/command/__Get-Date__
Returns PowerShell Data in JSON  
```json
{
    "value":  "\/Date(1436237104231)\/",
    "DisplayHint":  2,
    "DateTime":  "Monday, July 6, 2015 9:45:04 PM"
}
```
