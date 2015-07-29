Param (
    [string]$FirstName,
    [string]$LastName
)

###Templates
$Form = @"
<html>
    <h1>$(Split-Path $MyInvocation.MyCommand.Definition -Leaf)</h1>

    <form method="GET">
        $($MyInvocation.MyCommand.Parameters.Keys | % { "$_ <input type='text' name='-$_'><br>`r`n" })
        <input type="submit" value="Submit">
    </form>
</html>
"@

$Result = @"
    <h1>Hello, $FirstName $LastName</h1>

    <p>Thanks for testing this $FirstName</p>
"@

###Runtime
if ($($PSBoundParameters.Keys)) {
    Write-Output $Result
} else {
    Write-Output $Form
}
